/*
 * Maintained by jemo from 2020.5.9 to now
 * Created by jemo on 2020.5.9 17:43:53
 * Search Title List
 */

package handle

import (
  "log"
  "time"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SearchTitleList(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  db := database.DB
  query := r.URL.Query()
  keyword := query["keyword"][0]
  sqlStr := `
    SELECT
      item.searchId,
      item.name,
      item.price,
      item.imgUrl,
      item.detailUrl,
      item.womenProductId,
      item.keyName,
      item.itemTypeKey,
      item.itemNum,
      pddItem.goodsName,
      pddItem.pddId,
      pddItem.isOnsale
    FROM item
    LEFT JOIN pddItem
      ON item.searchId = pddItem.outGoodsSn
  `
  var args []interface{}
  if len(keyword) > 0 {
    sqlStr += `
    where
      item.keyName LIKE ?
    `
    keyword = "%" + keyword + "%"
    args = append(args, keyword)
  } else {
    sqlStr += `
    where
      pddItem.isOnsale = 1
      order by
        pddItem.createdAt desc
    `
  }
  rows, err := db.Query(sqlStr, args...)
  if err != nil {
    log.Println("search-title-list-query-keyword-error: ", err)
  }
  defer rows.Close()
  var titleList []interface{}
  for rows.Next() {
    var (
      searchId int64
      name string
      imgUrl sql.NullString
      price sql.NullFloat64
      detailUrl sql.NullString
      womenProductId sql.NullInt64
      keyName sql.NullString
      itemTypeKey sql.NullInt64
      itemNum sql.NullString
      goodsName sql.NullString
      pddId sql.NullInt64
      isOnsale sql.NullBool
    )
    if err := rows.Scan(
      &searchId,
      &name,
      &price,
      &imgUrl,
      &detailUrl,
      &womenProductId,
      &keyName,
      &itemTypeKey,
      &itemNum,
      &goodsName,
      &pddId,
      &isOnsale,
    ); err != nil {
      log.Println("search-title-list-scan-error: ", err)
    }
    title := map[string]interface{}{
      "searchId": searchId,
      "name": name,
      "price": price.Float64,
      "imgUrl": imgUrl.String,
      "detailUrl": detailUrl.String,
      "womenProductId": womenProductId.Int64,
      "keyName": keyName.String,
      "itemTypeKey": itemTypeKey.Int64,
      "itemNum": itemNum.String,
      "goodsName": goodsName.String,
      "pddId": pddId.Int64,
      "isOnsale": isOnsale.Bool,
    }
    titleList = append(titleList, title)
  }
  now := time.Now()
  diff := now.Sub(start)
  log.Println("time: ", diff)
  json.NewEncoder(w).Encode(titleList)
}
