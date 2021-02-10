/*
 * Maintained by jemo from 2020.7.6 to now
 * Created by jemo on 2020.7.6 10:59:45
 * Upload hznzcn Order File
 */

package handle

import (
  "io"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
  "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func UploadHznzcnOrderFile(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20) // 32M
  file, _, err := r.FormFile("file")
  if err != nil {
    log.Println("upload-hznzcn-order-file-form-file-err: ", err)
  }
  defer file.Close()
  xlFile, err := excelize.OpenReader(file)
  if err != nil {
    log.Println("upload-hznzcn-order-file-excelize-open-reader-error: ", err)
  }
  rows, err := xlFile.GetRows("订单导出")
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO order1688 (
      orderId,
      totalPrice,
      shippingFare,
      discount,
      actualPayment,
      orderStatus,
      orderCreatedTime,
      orderPaymentTime,
      receiver,
      shippingAddress,
      phone,
      productTitle,
      price,
      amount,
      courierCompany,
      trackingNumber,
      orderType,
      productSku,
      orderTotalPrice,
      agentDeliveryFee,
      paymentMethod,
      outerOrderId,
      deliveryTime,
      productStatus,
      distributionAmount,
      deliveryAmount
    ) VALUES(
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    )
  `)
  if err != nil {
    log.Println("upload-hznzcn-order-file-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare("UPDATE order1688 SET orderStatus = ?, courierCompany = ?, trackingNumber = ?, productStatus = ? WHERE orderId = ?")
  if err != nil {
    log.Println("upload-hznzcn-order-file-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 1; i < len(rows); i++ {
    row := rows[i]
    orderId := getXlsxCell(rows[0], row, "订单编号")
    if orderId == "" {
      continue
    }
    totalPrice := getXlsxCell(rows[0], row, "商品总价")
    shippingFare := getXlsxCell(rows[0], row, "运费")
    discount := getXlsxCell(rows[0], row, "优惠金额")
    actualPayment := getXlsxCell(rows[0], row, "付款金额")
    orderStatusStr := getXlsxCell(rows[0], row, "订单状态").(string)
    orderStatusMap := map[string]interface{}{
      "已付款": 1,
      "已发货": 2,
      "已完成": 4,
      "已退换货": 5, // 已换货
      "已退款": 5, // 已退货
      "已作废": 6,
    }
    orderStatus := orderStatusMap[orderStatusStr]
    orderCreatedTime := getXlsxCell(rows[0], row, "下单日期")
    orderPaymentTime := getXlsxCell(rows[0], row, "付款时间")
    if orderPaymentTime == "" {
      orderPaymentTime = nil
    }
    receiver := getXlsxCell(rows[0], row, "收货人姓名")
    shippingAddress := getXlsxCell(rows[0], row, "收货人地址")
    phone := getXlsxCell(rows[0], row, "收货人手机")
    productTitle := getXlsxCell(rows[0], row, "商品名称")
    price := getXlsxCell(rows[0], row, "商品单价")
    amount := getXlsxCell(rows[0], row, "购买数量")
    courierCompany := getXlsxCell(rows[0], row, "配送方式")
    trackingNumber := getXlsxCell(rows[0], row, "运单号")
    productSku := getXlsxCell(rows[0], row, "商品规格")
    orderTotalPrice := getXlsxCell(rows[0], row, "订单总金额")
    agentDeliveryFee := getXlsxCell(rows[0], row, "代发费")
    paymentMethod := getXlsxCell(rows[0], row, "付款方式")
    outerOrderId := getXlsxCell(rows[0], row, "外部订单号")
    deliveryTime := getXlsxCell(rows[0], row, "发货时间")
    if deliveryTime == "" {
      deliveryTime = nil
    }
    productStatus := getXlsxCell(rows[0], row, "货物状态")
    distributionAmount := getXlsxCell(rows[0], row, "已配数量")
    deliveryAmount := getXlsxCell(rows[0], row, "已发数量")
    orderType := 1
    var orderCount int
    err = db.QueryRow("SELECT COUNT(*) FROM order1688 WHERE orderId = ?", orderId).Scan(&orderCount)
    if err != nil {
      log.Println("upload-hznzcn-order-file-count-error: ", err)
    }
    if orderCount == 0 {
      _, err = stmtInsert.Exec(
        orderId,
        totalPrice,
        shippingFare,
        discount,
        actualPayment,
        orderStatus,
        orderCreatedTime,
        orderPaymentTime,
        receiver,
        shippingAddress,
        phone,
        productTitle,
        price,
        amount,
        courierCompany,
        trackingNumber,
        orderType,
        productSku,
        orderTotalPrice,
        agentDeliveryFee,
        paymentMethod,
        outerOrderId,
        deliveryTime,
        productStatus,
        distributionAmount,
        deliveryAmount,
      )
      if err != nil {
        log.Println("upload-hznzcn-order-file-insert-exec-error: ", err)
        log.Println("orderStatusStr: ", orderStatusStr)
        log.Println("orderPaymentTime: ", orderPaymentTime)
      }
    } else {
      _, err = stmtUpdate.Exec(orderStatus, courierCompany, trackingNumber, productStatus, orderId)
      if err != nil {
        log.Println("upload-hznzcn-order-file-update-exec-error: ", err)
      }
    }
  }
  io.WriteString(w, "ok")
}

func getXlsxCell(heading []string, row []string, key string) interface{} {
  for i := 0; i < len(heading); i++ {
    if heading[i] == key {
      return row[i]
    }
  }
  return nil
}
