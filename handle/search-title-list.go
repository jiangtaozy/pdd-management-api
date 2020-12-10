/*
 * Maintained by jemo from 2020.5.9 to now
 * Created by jemo on 2020.5.9 17:43:53
 * Search Title List
 */

package handle

import (
  "log"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SearchTitleList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      searchItem.id,
      searchItem.name,
      item.price,
      item.imgUrl,
      item.detailUrl,
      item.womenProductId
    FROM searchItem
    LEFT JOIN item
      ON searchItem.id = item.searchId
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
      imgUrl sql.NullString
      price sql.NullFloat64
      detailUrl sql.NullString
      womenProductId sql.NullInt64
    )
    if err := rows.Scan(
      &id,
      &name,
      &price,
      &imgUrl,
      &detailUrl,
      &womenProductId,
    ); err != nil {
      log.Println("search-title-list-scan-error: ", err)
    }
    title := map[string]interface{}{
      "id": id,
      "name": name,
      "price": price.Float64,
      "imgUrl": imgUrl.String,
      "detailUrl": detailUrl.String,
      "womenProductId": womenProductId.Int64,
    }
    titleList = append(titleList, title)
  }
  json.NewEncoder(w).Encode(titleList)
}
