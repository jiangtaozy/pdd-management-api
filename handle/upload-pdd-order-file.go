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
  db := database.DB
  //stmtInsert, err := db.Prepare("INSERT INTO itemOrder (0 productName, 1 orderId, 2 orderStatus, 3 productTotalPrice, 4 storeDiscount, 5 platformDiscount, 6 postage, 7 onsiteInstallationFee, 8 homeDeliveryFee, 9 homeDeliveryAndInstallationFee, 10 userPaidAmount, 11 merchantReceivedAmount, 12 numberOfProducts, 13 whetherUnderReview, 14 paymentTime, 15 joinSuccessTime, 16 orderConfirmationTime, 17 commitmentDeliveryTime, 18 deliveryTime, 19 confirmDeliveryTime, 20 productId, 21 productSku, 22 skuId, 23 merchantCodeSkuDimension, 24 merchantCodeProductDimension, 25 merchantNotes, 26 buyerMessage, 27 afterSaleStatus, 28 consumerInformation, 29 trackingNumber, 30 courierCompany, 30 receiver, 32 phone, 33 province, 34 city, 35 district, 36 street, 37 whetherDrawOrZeroYuanTry, 38 whetherShunfengAddPrice, 39 userBuyPhone, 40 travelInformation, 42 whetherStoreMention, 43 storeName, 44 storeCustomCode, 45 relatedGoodsCode, 46 goodsName, 47 goodsType, 48 goodsChild, 49 warehouseName, 50 warehouseAddress, 51 idCardName, 52 identificationNumber, 53 haitaoCustomsOrderNumber, 54 paymentId, 55 paymentMethod) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
  stmtInsert, err := db.Prepare("INSERT INTO itemOrder (productName, orderId, orderStatus, productTotalPrice, storeDiscount, platformDiscount, postage, onsiteInstallationFee, homeDeliveryFee, homeDeliveryAndInstallationFee, userPaidAmount, merchantReceivedAmount, numberOfProducts, whetherUnderReview, paymentTime, joinSuccessTime, orderConfirmationTime, commitmentDeliveryTime, deliveryTime, confirmDeliveryTime, productId, productSku, skuId, merchantCodeSkuDimension, merchantCodeProductDimension, merchantNotes, buyerMessage, afterSaleStatus, consumerInformation, trackingNumber, courierCompany, receiver, phone, province, city, district, street, whetherDrawOrZeroYuanTry, whetherShunfengAddPrice, userBuyPhone, travelInformation, whetherStoreMention, storeName, storeCustomCode, relatedGoodsCode, goodsName, goodsType, goodsChild, warehouseName, warehouseAddress, idCardName, identificationNumber, haitaoCustomsOrderNumber, paymentId, paymentMethod) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
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
      if data[i][18] == "\t" {
        data[i][18] = ""
      }
      if data[i][19] == "\t" {
        data[i][19] = ""
      }
      if data[i][20] == "\t" {
        data[i][20] = ""
      }
      var orderCount int
      err = db.QueryRow("SELECT COUNT(*) FROM itemOrder WHERE orderId = ?", data[i][1]).Scan(&orderCount)
      if err != nil {
        log.Println("upload-pdd-order-file-count-error: ", err)
      }
      if orderCount == 0 {
        _, err = stmtInsert.Exec(data[i][0], data[i][1], data[i][2], data[i][3], data[i][4], data[i][5], data[i][6], data[i][7], data[i][8], data[i][9], data[i][10], data[i][11], data[i][12], data[i][13], data[i][14], data[i][15], data[i][16], data[i][17], NewNullString(data[i][18]), NewNullString(data[i][19]), NewNullString(data[i][20]), data[i][21], data[i][22], data[i][23], data[i][24], data[i][25], data[i][26], data[i][27], data[i][28], data[i][29], data[i][30], data[i][31], data[i][32], data[i][33], data[i][34], data[i][35], data[i][36], data[i][37], data[i][38], data[i][39], data[i][40], data[i][42], data[i][43], data[i][44], data[i][45], data[i][46], data[i][47], data[i][48], data[i][49], data[i][50], data[i][51], data[i][52], data[i][53], data[i][54], data[i][55])
        if err != nil {
          log.Println("upload-pdd-order-file-insert-exec-error: ", err)
        }
      } else {
        // stmtUpdate, err := db.Prepare("UPDATE itemOrder SET orderStatus = ?, deliveryTime = ?, confirmDeliveryTime = ?, trackingNumber = ?, courierCompany = ? WHERE orderId = ?")
        _, err = stmtUpdate.Exec(data[i][2], data[i][18], NewNullString(data[i][19]), NewNullString(data[i][29]), data[i][30], data[i][1])
        if err != nil {
          log.Println("upload-pdd-order-file-update-exec-error: ", err)
        }
      }
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
