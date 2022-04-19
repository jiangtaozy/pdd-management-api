/*
 * Maintained by jemo from 2022.4.16 to now
 * Created by jemo on 2022.4.16 15:58:46
 * Item Sku Update
 */

package handle

import (
  "io"
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func ItemSkuUpdate(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("item-sku-update-decoder-error: ", err)
  }
  var id = body["id"]
  var shortSkuName = body["shortSkuName"]
  var shortSkuNum = body["shortSkuNum"]
  db := database.DB
  update, err := db.Prepare("update itemSku set shortSkuName = ? where id = ?")
  if err != nil {
    log.Println("item-sku-update-prepare-error: ", err)
  }
  _, err = update.Exec(shortSkuName, id)
  if err != nil {
    log.Println("item-sku-update-exec-error: ", err)
  }
  var count int
  err = db.QueryRow("select count(*) from itemSkuNum where shortSkuName = ?", shortSkuName).Scan(&count)
  if err != nil {
    log.Println("item-sku-update-count-error: ", err)
  }
  if count == 0 {
    insert, err := db.Prepare("insert into itemSkuNum (shortSkuName, shortSkuNum) values(?, ?)")
    if err != nil {
      log.Println("item-sku-update-insert-prepare-error: ", err)
    }
    _, err = insert.Exec(shortSkuName, shortSkuNum)
    if err != nil {
      log.Println("item-sku-update-insert-exec-error: ", err)
    }
  } else {
    updateSku, err := db.Prepare("update itemSkuNum set shortSkuNum = ? where shortSkuName = ?")
    if err != nil {
      log.Println("item-sku-update-sku-prepare-error: ", err)
    }
    _, err = updateSku.Exec(shortSkuNum, shortSkuName)
    if err != nil {
      log.Println("item-sku-update-sku-exec-error: ", err)
    }
  }
  io.WriteString(w, "ok")
}
