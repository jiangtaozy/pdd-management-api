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
      item.id,
      item.searchId,
      item.name,
      item.price,
      item.imgUrl,
      item.detailUrl,
      item.womenProductId,
      pddItem.goodsName
    FROM item
    LEFT JOIN pddItem
      ON item.searchId = pddItem.outGoodsSn
    ORDER BY searchId DESC
  `)
  if err != nil {
    log.Println("search-title-list-query-error: ", err)
  }
  defer rows.Close()
  var titleList []interface{}
  for rows.Next() {
    var (
      id int64
      searchId int64
      name string
      imgUrl sql.NullString
      price sql.NullFloat64
      detailUrl sql.NullString
      womenProductId sql.NullInt64
      goodsName sql.NullString
    )
    if err := rows.Scan(
      &id,
      &searchId,
      &name,
      &price,
      &imgUrl,
      &detailUrl,
      &womenProductId,
      &goodsName,
    ); err != nil {
      log.Println("search-title-list-scan-error: ", err)
    }
    title := map[string]interface{}{
      "id": id,
      "searchId": searchId,
      "name": name,
      "price": price.Float64,
      "imgUrl": imgUrl.String,
      "detailUrl": detailUrl.String,
      "womenProductId": womenProductId.Int64,
      "goodsName": goodsName.String,
    }
    titleList = append(titleList, title)
  }
  json.NewEncoder(w).Encode(titleList)
}
