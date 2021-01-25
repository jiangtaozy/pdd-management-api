/*
 * Maintained by jemo from 2021.1.24 to now
 * Created by jemo on 2021.1.24 17:55:32
 * Pdd Item Price History Data
 * 拼多多商品价格历史数据
 */

package handle

import (
  "log"
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddItemPriceHistoryData(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      pddItem.pddId,
      pddItem.goodsName,
      priceHistory.skuGroupPriceMin,
      priceHistory.date
    FROM
      pddItem
    LEFT JOIN
      pddItemPriceHistory AS priceHistory
    ON
      pddItem.pddId = priceHistory.pddId
  `)
  if err != nil {
    log.Println("pdd-item-price-history-data-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      pddId int64
      goodsName string
      skuGroupPriceMin sql.NullInt64
      date sql.NullString
    )
    if err := rows.Scan(
      &pddId,
      &goodsName,
      &skuGroupPriceMin,
      &date,
    ); err != nil {
      log.Println("pdd-item-price-history-data-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    history := map[string]interface{}{
      "skuGroupPriceMin": skuGroupPriceMin.Int64,
      "date": date.String,
    }
    hasItemInItemList := false
    for i := 0; i < len(list); i++ {
      item := list[i].(map[string]interface{})
      if item["pddId"] == pddId {
        hasItemInItemList = true
        priceHistory := item["priceHistory"].([]interface{})
        priceHistoryLength := len(priceHistory)
        if priceHistoryLength > 0 {
          if priceHistory[priceHistoryLength - 1].(map[string]interface{})["skuGroupPriceMin"].(int64) != skuGroupPriceMin.Int64 {
            item["priceHistory"] = append(item["priceHistory"].([]interface{}), history)
          }
        } else {
          item["priceHistory"] = append(item["priceHistory"].([]interface{}), history)
        }
        break
      }
    }
    if !hasItemInItemList {
      var priceHistory []interface{}
      priceHistory = append(priceHistory, history)
      item := map[string]interface{}{
        "pddId": pddId,
        "goodsName": goodsName,
        "priceHistory": priceHistory,
      }
      list = append(list, item)
    }
  }
  for i := 0; i < len(list); i++ {
    item := list[i].(map[string]interface{})
    pddId := item["pddId"].(int64)
    priceHistory := item["priceHistory"].([]interface{})
    if len(priceHistory) > 0 {
      price := priceHistory[len(priceHistory) - 1].(map[string]interface{})
      date := price["date"].(string)
      var (
        goodsPv sql.NullInt64
        payOrdrCnt sql.NullInt64
      )
      if date != "" {
        err = db.QueryRow(`
          SELECT
            SUM(goodsPv) goodsPv,
            SUM(payOrdrCnt) payOrdrCnt
          FROM
            pddGoodsFlowData
          WHERE
            goodsId = ?
          AND
            statDate > ?
          GROUP BY
            goodsId
        `, pddId, date).Scan(
          &goodsPv,
          &payOrdrCnt,
        )
        if err != nil && err != sql.ErrNoRows {
          log.Println("pdd-item-price-history-data-query-flow-data-scan-error: ", err)
          http.Error(w, err.Error(), 500)
          return
        }
        item["afterChangePriceGoodsPv"] = goodsPv.Int64
        item["payOrdrCnt"] = payOrdrCnt.Int64
      }
    }
  }
  json.NewEncoder(w).Encode(list)
}
