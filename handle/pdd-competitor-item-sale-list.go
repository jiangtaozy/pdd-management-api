/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 17:16:54
 * Pdd Competitor Item Sale List
 * 拼多多竞争对手商品销量列表
 */

package handle

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddCompetitorItemSaleList(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  itemId := query["itemId"][0]
  db := database.DB
  rows, err := db.Query(`
    SELECT
      goodsId,
      date,
      sale
    FROM
      pddCompetitorItemSale
    WHERE
      goodsId = ?
  `, itemId)
  if err != nil {
    log.Println("pdd-competitor-item-sale-list-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      goodsId string
      date string
      sale int64
    )
    if err := rows.Scan(
      &goodsId,
      &date,
      &sale,
    ); err != nil {
      log.Println("pdd-competitor-item-sale-list-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    pddCompetitorItemSale := map[string]interface{}{
      "goodsId": goodsId,
      "date": date,
      "sale": sale,
    }
    list = append(list, pddCompetitorItemSale)
  }
  json.NewEncoder(w).Encode(list)
}
