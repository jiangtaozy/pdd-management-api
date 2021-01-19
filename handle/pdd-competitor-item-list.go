/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 15:44:10
 * Pdd Competitor Item List
 * 拼多多竞争对手商品列表
 */

package handle

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddCompetitorItemList(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  itemId := query["itemId"][0]
  db := database.DB
  rows, err := db.Query(`
    SELECT
      id,
      name,
      price,
      goodsId,
      competitorId,
      relatedItemId
    FROM
      pddCompetitorItem
    WHERE
      relatedItemId = ?
  `, itemId)
  if err != nil {
    log.Println("pdd-competitor-item-list-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      id int64
      name string
      price int64
      goodsId string
      competitorId int64
      relatedItemId int64
    )
    if err := rows.Scan(
      &id,
      &name,
      &price,
      &goodsId,
      &competitorId,
      &relatedItemId,
    ); err != nil {
      log.Println("pdd-competitor-item-list-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    pddCompetitorItem := map[string]interface{}{
      "id": id,
      "name": name,
      "price": price,
      "goodsId": goodsId,
      "competitorId": competitorId,
      "relatedItemId": relatedItemId,
    }
    list = append(list, pddCompetitorItem)
  }
  json.NewEncoder(w).Encode(list)
}
