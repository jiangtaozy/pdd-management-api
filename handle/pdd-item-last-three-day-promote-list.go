/*
 * Maintained by jemo from 2020.12.11 to now
 * Created by jemo on 2020.12.11 17:01:10
 * Pinduoduo Item Last Three Day Promote List
 * 过去三天推广商品列表
 */

package handle

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddItemLastThreeDayPromoteList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      daily.adId,
      unit.goodsId,
      unit.goodsName
    FROM
      pddAdUnitDailyData AS daily
    LEFT JOIN
      pddAdUnit AS unit
    ON
      daily.adId = unit.adId
    WHERE
      daily.date >= ( CURDATE() - INTERVAL 3 DAY )
    GROUP BY
      daily.adId
  `)
  if err != nil {
    log.Println("pdd-item-last-three-day-promote-list-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      adId int64
      goodsId int64
      goodsName string
    )
    err = rows.Scan(
      &adId,
      &goodsId,
      &goodsName,
    )
    if err != nil {
      log.Println("pdd-item-last-three-day-promote-list-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    item := map[string]interface{}{
      "adId": adId,
      "goodsId": goodsId,
      "goodsName": goodsName,
    }
    list = append(list, item)
  }
  json.NewEncoder(w).Encode(list)
}
