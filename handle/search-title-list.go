/*
 * Maintained by jemo from 2020.5.9 to now
 * Created by jemo on 2020.5.9 17:43:53
 * Search Title List
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SearchTitleList(w http.ResponseWriter, r *http.Request) {
  db := database.ConnectDB()
  rows, err := db.Query("SELECT * FROM searchItem")
  if err != nil {
    log.Println("search-title-list-query-error: ", err)
  }
  defer rows.Close()
  var titleList []interface{}
  for rows.Next() {
    var (
      id int64
      name string
    )
    if err := rows.Scan(&id, &name); err != nil {
      log.Println("search-title-list-scan-error: ", err)
    }
    title := map[string]interface{}{
      "id": id,
      "name": name,
    }
    titleList = append(titleList, title)
  }
  json.NewEncoder(w).Encode(titleList)
}
