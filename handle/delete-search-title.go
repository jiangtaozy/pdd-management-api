/*
 * Maintained by jemo from 2020.5.13 to now
 * Created by jemo on 2020.5.13 17:36:27
 * Delete Search Title
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func DeleteSearchTitle(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("delete-search-data-error: ", err)
  }
  var id = body["id"]
  log.Println("id: ", id)
  db := database.DB
  stmtDelete, err := db.Prepare("DELETE FROM searchItem WHERE id = ?")
  if err != nil {
    log.Println("delete-search-title-prepare-error: ", err)
  }
  _, err = stmtDelete.Exec(id)
  if err != nil {
    log.Println("delete-search-title-exec-error: ", err)
  }
  io.WriteString(w, "ok")
}
