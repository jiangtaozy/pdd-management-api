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
  var id = body["id"]
  var name = body["name"].(string)
  db := database.DB
  stmtUpdate, err := db.Prepare("UPDATE searchItem SET name = ? WHERE id = ?")
  if err != nil {
    log.Println("update-search-title-prepare-error: ", err)
  }
  _, err = stmtUpdate.Exec(name, id)
  if err != nil {
    log.Println("update-search-title-exec-error: ", err)
  }
  stmtUpdateItem, err := db.Prepare(`
    UPDATE
      item
    SET
      womenProductId = ?
    WHERE
      searchId = ?
  `)
  if err != nil {
    log.Println("update-search-title-prepare-update-item-error: ", err)
  }
  womenProductId := body["womenProductId"]
  _, err = stmtUpdateItem.Exec(womenProductId, id)
  if err != nil {
    log.Println("update-search-title-update-item-exec-error: ", err)
  }
  io.WriteString(w, "ok")
}
