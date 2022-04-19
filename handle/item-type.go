/*
 * Maintained by jemo from 2022.4.18 to now
 * Created by jemo on 2022.4.18 10:32:55
 * 商品类型修改
 */

package handle

import (
  "io"
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func ItemType(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("item-type-decode-body-error: ", err)
  }
  db := database.DB
  var count int
  err = db.QueryRow("SELECT COUNT(*) FROM itemType WHERE id = ?", body["id"]).Scan(&count)
  if err != nil {
    log.Println("item-type-count-error: ", err)
  }
  if count == 0 {
    stmtInsert, err := db.Prepare("INSERT INTO itemType (typeName, typeNum) VALUES(?, ?)")
    if err != nil {
      log.Println("item-type-insert-prepare-error: ", err)
    }
    _, err = stmtInsert.Exec(body["typeName"], body["typeNum"])
    if err != nil {
      log.Println("item-type-insert-exec-error: ", err)
    }
  } else {
    stmtUpdate, err := db.Prepare("UPDATE itemType SET typeName = ?, typeNum = ? WHERE id = ?")
    if err != nil {
      log.Println("item-type-update-prepare-error: ", err)
    }
    _, err = stmtUpdate.Exec(body["typeName"], body["typeNum"], body["id"])
    if err != nil {
      log.Println("item-type-update-exec-error", err)
    }
  }

  io.WriteString(w, "ok")
}
