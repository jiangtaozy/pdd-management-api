/*
 * Maintained by jemo from 2020.5.13 to now
 * Created by jemo on 2020.5.13 16:37:26
 * Update Search Title
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func UpdateSearchTitle(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("update-search-data-error: ", err)
  }
  var searchId = body["searchId"]
  var itemTypeKey = body["itemTypeKey"]
  var itemNum = body["itemNum"]
  db := database.DB
  stmtUpdate, err := db.Prepare(`
    update
      item
    set
      itemTypeKey = ?,
      itemNum = ?
    where searchId = ?
  `)
  if err != nil {
    log.Println("update-search-title-prepare-error: ", err)
  }
  _, err = stmtUpdate.Exec(itemTypeKey, itemNum, searchId)
  if err != nil {
    log.Println("update-search-title-exec-error: ", err)
  }
  io.WriteString(w, "ok")
}
