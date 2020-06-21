/*
 * Maintained by jemo from 2020.5.14 to now
 * Created by jemo on 2020.5.14 11:01:01
 * Search Title by Id
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SearchTitleById(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  idArray := query["id"]
  id := idArray[0]
  db := database.DB
  var title string
  err := db.QueryRow("SELECT name FROM searchItem WHERE id = ?", id).Scan(&title)
  if err != nil {
    log.Println("search-title-by-id-query-error: ", err)
  }
  json.NewEncoder(w).Encode(map[string]interface{}{
    "id": id,
    "name": title,
  })
}
