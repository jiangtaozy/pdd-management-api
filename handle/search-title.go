/*
 * Maintained by jemo from 2020.5.9 to now
 * Created by jemo on 2020.5.9 17:20:09
 * Search Title
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SearchTitle(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("search-data-error: ", err)
  }
  var searchTitle = body["searchTitle"].(string)
  db := database.DB
  var count int
  err = db.QueryRow("SELECT COUNT(*) FROM searchItem WHERE name = ?", searchTitle).Scan(&count)
  if err != nil {
    log.Println("search-data-count-error: ", err)
  }
  if count == 0 {
    stmtInsertItem, err := db.Prepare("INSERT INTO searchItem (name) VALUES(?)")
    if err != nil {
      log.Println("search-title-insert-item-prepare-error: ", err)
    }
    _, err = stmtInsertItem.Exec(searchTitle)
    if err != nil {
      log.Println("search-title-insert-item-exec-error: ", err)
    }
  }
  io.WriteString(w, "ok")
}
