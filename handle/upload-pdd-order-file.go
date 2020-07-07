/*
 * Maintained by jemo from 2020.5.18 to now
 * Created by jemo on 2020.5.18 14:44:07
 * Upload Pinduoduo Order File
 */

package handle

import (
  "io"
  "encoding/csv"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
  "strings"
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
  stmtUpdate, err := db.Prepare("UPDATE itemOrder SET trackingNumber = ?, courierCompany = ? WHERE orderId = ?")
  if err != nil {
    log.Println("upload-pdd-order-file-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 1; i < len(data); i++ {
    orderId := getXlsxCell(data[0], data[i], "订单号")
    trackingNumber := strings.TrimSpace(getXlsxCell(data[0], data[i], "快递单号").(string))
    courierCompany := strings.TrimSpace(getXlsxCell(data[0], data[i], "快递公司").(string))
    var orderCount int
    err = db.QueryRow("SELECT COUNT(*) FROM itemOrder WHERE orderId = ?", data[i][1]).Scan(&orderCount)
    if err != nil {
      log.Println("upload-pdd-order-file-count-error: ", err)
    }
    if orderCount != 0 {
      _, err = stmtUpdate.Exec(trackingNumber, courierCompany, orderId)
      if err != nil {
        log.Println("upload-pdd-order-file-update-exec-error: ", err)
      }
    } else {
      log.Println("upload-pdd-order-file-new-order-error: have new order to save")
    }
  }
  io.WriteString(w, "ok")
}
