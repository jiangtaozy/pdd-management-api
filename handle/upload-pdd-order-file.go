/*
 * Maintained by jemo from 2020.5.18 to now
 * Created by jemo on 2020.5.18 14:44:07
 * Upload Pinduoduo Order File
 */

package handle

import (
  "io"
  "encoding/csv"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func UploadPddOrderFile(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20) // 32M
  file, _, err := r.FormFile("file")
  if err != nil {
    log.Println("upload-pdd-order-file-form-file-err: ", err)
  }
  defer file.Close()
  reader := csv.NewReader(file)
  reader.FieldsPerRecord = -1
  data, err := reader.ReadAll()
  if err != nil {
    log.Println("upload-pdd-order-file-read-all-err: ", err)
  }
  db := database.ConnectDB()
  stmtInsert, err := db.Prepare("INSERT INTO itemOrder (productName, orderId, orderStatus, productTotalPrice, storeDiscount, platformDiscount, postage, onsiteInstallationFee, homeDeliveryFee, homeDeliveryAndInstallationFee, userPaidAmount, merchantReceivedAmount, numberOfProducts, idCardName, identificationNumber, receiver, phone, whetherUnderReview, province, city, district, street, paymentTime, joinSuccessTime, orderConfirmationTime, commitmentDeliveryTime, deliveryTime, confirmDeliveryTime, productId, productSku, userBuyPhone, skuId, merchantCodeSkuDimension, merchantCodeProductDimension, trackingNumber, courierCompany, haitaoCustomsOrderNumber, paymentId, paymentMethod, whetherDrawOrZeroYuanTry, whetherShunfengAddPrice, merchantNotes, afterSaleStatus, buyerMessage, relatedGoodsCode, goodsName, goodsType, goodsChild, warehouseName, warehouseAddress, whetherStoreMention, storeName, storeCustomCode, travelInformation, consumerInformation) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Println("upload-pdd-order-file-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare("UPDATE itemOrder SET orderStatus = ?, deliveryTime = ?, confirmDeliveryTime = ?, trackingNumber = ?, courierCompany = ? WHERE orderId = ?")
  if err != nil {
    log.Println("upload-pdd-order-file-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(data); i++ {
    if i > 0 {
      if data[i][26] == "\t" {
        data[i][26] = ""
      }
      if data[i][27] == "\t" {
        data[i][27] = ""
      }
      var orderCount int
      err = db.QueryRow("SELECT COUNT(*) FROM itemOrder WHERE orderId = ?", data[i][1]).Scan(&orderCount)
      if err != nil {
        log.Println("upload-pdd-order-file-count-error: ", err)
      }
      if orderCount == 0 {
        _, err = stmtInsert.Exec(data[i][0], data[i][1], data[i][2], data[i][3], data[i][4], data[i][5], data[i][6], data[i][7], data[i][8], data[i][9], data[i][10], data[i][11], data[i][12], data[i][13], data[i][14], data[i][15], data[i][16], data[i][17], data[i][18], data[i][19], data[i][20], data[i][21], data[i][22], data[i][23], data[i][24], data[i][25], NewNullString(data[i][26]), NewNullString(data[i][27]), data[i][28], data[i][29], data[i][30], data[i][31], data[i][32], data[i][33], data[i][34], data[i][35], data[i][36], data[i][37], data[i][38], data[i][39], data[i][40], data[i][41], data[i][42], data[i][43], data[i][44], data[i][45], data[i][46], data[i][47], data[i][48], data[i][49], data[i][50], data[i][51], data[i][52], data[i][53], data[i][54])
        if err != nil {
          log.Println("upload-pdd-order-file-insert-exec-error: ", err)
        }
      } else {
  //stmtUpdate, err := db.Prepare("UPDATE itemOrder SET orderStatus = ?, deliveryTime = ?, confirmDeliveryTime = ?, trackingNumber = ?, courierCompany = ? WHERE orderId = ?")
        _, err = stmtUpdate.Exec(data[i][2], data[i][26], NewNullString(data[i][27]), NewNullString(data[i][34]), data[i][35], data[i][1])
        if err != nil {
          log.Println("upload-pdd-order-file-update-exec-error: ", err)
        }
      }
      /*
productName,0 orderId,1 orderStatus,2 productTotalPrice,3 storeDiscount,4 platformDiscount,5 postage, 6 onsiteInstallationFee, 7 homeDeliveryFee, 8 homeDeliveryAndInstallationFee, 9 userPaidAmount, 10 merchantReceivedAmount,11 numberOfProducts, 12 idCardName, 13 identificationNumber, 14 receiver, 15 phone, 16 whetherUnderReview, 17 province, 18 city, 19 district, 20 street, 21 paymentTime, 22 joinSuccessTime,23 orderConfirmationTime, 24 commitmentDeliveryTime, 25 deliveryTime, 26 confirmDeliveryTime,27 productId,28 productSku,19 userBuyPhone,30 skuId,31 merchantCodeSkuDimension,32 merchantCodeProductDimension,33 trackingNumber,34 courierCompany,35 haitaoCustomsOrderNumber, paymentId, paymentMethod, whetherDrawOrZeroYuanTry, whetherShunfengAddPrice, merchantNotes, afterSaleStatus, buyerMessage, relatedGoodsCode, goodsName, goodsType, goodsChild, warehouseName, warehouseAddress, whetherStoreMention, storeName, storeCustomCode, travelInformation, consumerInformation
      */
    }
  }
  io.WriteString(w, "ok")
}

func NewNullString(s string) sql.NullString {
  if len(s) == 0 {
    return sql.NullString{}
  }
  return sql.NullString{
    String: s,
    Valid: true,
  }
}
