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
  "database/sql"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SearchTitleList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      searchItem.id,
      searchItem.name,
      item.detailUrl
    FROM searchItem
    LEFT JOIN item
      ON searchItem.id = item.searchId
    WHERE
      item.forSell = true OR
      item.forSell IS NULL
    ORDER BY id DESC
  `)
  if err != nil {
    log.Println("search-title-list-query-error: ", err)
  }
  defer rows.Close()
  var titleList []interface{}
  for rows.Next() {
    var (
      id int64
      name string
      detailUrl sql.NullString
    )
    if err := rows.Scan(&id, &name, &detailUrl); err != nil {
      log.Println("search-title-list-scan-error: ", err)
    }
    title := map[string]interface{}{
      "id": id,
      "name": name,
      "detailUrl": detailUrl.String,
    }
    titleList = append(titleList, title)
  }
  json.NewEncoder(w).Encode(titleList)
}
