/*
 * Maintained by jemo from 2020.7.3 to now
 * Created by jemo on 2020.7.3 17:26:49
 * Upload 1688 Order File
 */

package handle

import (
  "io"
  "github.com/extrame/xls"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
  "strconv"
  "strings"
)

func Upload1688OrderFile(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20) // 32M
  file, _, err := r.FormFile("file")
  if err != nil {
    log.Println("upload-1688-order-file-form-file-err: ", err)
  }
  defer file.Close()
  xlFile, err := xls.OpenReader(file, "utf-8")
  if err != nil {
    log.Println("upload-1688-order-file-xls-open-reader-error: ", err)
  }
  sheet := xlFile.GetSheet(0)
  lines := int(sheet.MaxRow)
  db := database.DB
  stmtInsert, err := db.Prepare("INSERT INTO order1688 (orderId, sellerCompany, totalPrice, shippingFare, discount, actualPayment, orderStatus, orderCreatedTime, orderPaymentTime, receiver, shippingAddress, postcode, phone, productTitle, price, amount, courierCompany, trackingNumber, orderType) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Println("upload-1688-order-file-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare("UPDATE order1688 SET orderStatus = ?, courierCompany = ?, trackingNumber = ? WHERE orderId = ?")
  if err != nil {
    log.Println("upload-1688-order-file-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 1; i <= lines; i++ {
    row := sheet.Row(i)
    orderId := getCell(sheet.Row(0), row, "订单编号")
    if orderId == "" {
      continue
    }
    sellerCompany := getCell(sheet.Row(0), row, "卖家公司名")
    totalPrice := getCell(sheet.Row(0), row, "货品总价(元)")
    shippingFare := getCell(sheet.Row(0), row, "运费(元)")
    discount := getCell(sheet.Row(0), row, "涨价或折扣(元)")
    actualPayment := getCell(sheet.Row(0), row, "实付款(元)")
    orderStatusStr := getCell(sheet.Row(0), row, "订单状态").(string)
    orderStatusMap := map[string]interface{}{
      "等待卖家发货": 1,
      "等待买家确认收货": 2,
      "已收货": 3,
      "交易成功": 4,
      "交易关闭": 6,
      "已发货": 2,
      "退款中": 5,
    }
    orderStatus := orderStatusMap[orderStatusStr]
    orderCreatedTimeString := getCell(sheet.Row(0), row, "订单创建时间")
    orderCreatedTimeFloat, _ := strconv.ParseFloat(orderCreatedTimeString.(string), 64)
    orderCreatedTime := timeFromExcelTime(orderCreatedTimeFloat, false)
    orderPaymentTimeString := getCell(sheet.Row(0), row, "订单付款时间")
    orderPaymentTimeFloat, _ := strconv.ParseFloat(orderPaymentTimeString.(string), 64)
    orderPaymentTime := timeFromExcelTime(orderPaymentTimeFloat, false)
    receiver := getCell(sheet.Row(0), row, "收货人姓名")
    shippingAddress := getCell(sheet.Row(0), row, "收货地址")
    postcode := getCell(sheet.Row(0), row, "邮编")
    phone := getCell(sheet.Row(0), row, "联系手机")
    productTitle := getCell(sheet.Row(0), row, "货品标题")
    price := getCell(sheet.Row(0), row, "单价(元)")
    amount := getCell(sheet.Row(0), row, "数量")
    var courierCompany string
    var trackingNumber string
    companyNumber := getCell(sheet.Row(0), row, "物流公司运单号")
    companyNumberArray := strings.Split(companyNumber.(string), ":")
    if len(companyNumberArray) > 1 {
      courierCompany = companyNumberArray[0]
      trackingNumber = companyNumberArray[1]
    }
    orderType := 0
    var orderCount int
    err = db.QueryRow("SELECT COUNT(*) FROM order1688 WHERE orderId = ?", orderId).Scan(&orderCount)
    if err != nil {
      log.Println("upload-1688-order-file-count-error: ", err)
    }
    if orderCount == 0 {
      _, err = stmtInsert.Exec(orderId, sellerCompany, totalPrice, shippingFare, discount, actualPayment, orderStatus, orderCreatedTime, orderPaymentTime, receiver, shippingAddress, postcode, phone, productTitle, price, amount, courierCompany, trackingNumber, orderType)
      if err != nil {
        log.Println("upload-1688-order-file-insert-exec-error: ", err)
        log.Println("orderStatusStr: ", orderStatusStr)
        log.Println("orderId: ", orderId)
      }
    } else {
      _, err = stmtUpdate.Exec(orderStatus, courierCompany, trackingNumber, orderId)
      if err != nil {
        log.Println("upload-1688-order-file-update-exec-error: ", err)
        log.Println("orderStatusStr: ", orderStatusStr)
      }
    }
  }
  io.WriteString(w, "ok")
}

func getCell(heading *xls.Row, row *xls.Row, key string) interface{} {
  for i := 0; i < heading.LastCol(); i++ {
    if heading.Col(i) == key {
      return row.Col(i)
    }
  }
  return nil
}
