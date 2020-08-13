/*
 * Maintained by jemo from 2020.8.12 to now
 * Created by jemo on 2020.8.12 20:08:18
 * Update Item Order
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func ItemOrderUpdate(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("item-order-update-decode-error: ", err)
  }
  var id = body["id"]
  var outerOrderId = body["outerOrderId"]
  db := database.DB
  stmtUpdate, err := db.Prepare("UPDATE itemOrder SET outerOrderId = ? WHERE id = ?")
  if err != nil {
    log.Println("item-order-update-prepare-error: ", err)
  }
  _, err = stmtUpdate.Exec(outerOrderId, id)
  if err != nil {
    log.Println("item-order-update-exec-error: ", err)
  }
  io.WriteString(w, "ok")
}
