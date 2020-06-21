/*
 * Maintained by jemo from 2020.5.25 to now
 * Created by jemo on 2020.5.25 9:43:53
 * Pdd Item List
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddItemList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query("SELECT id, quantity, skuGroupPriceMin, skuGroupPriceMax, pddId, goodsName, displayPriority, thumbUrl, isOnsale, soldQuantity, outGoodsSn, soldQuantityForThirtyDays, favCnt, ifNewGoods, goodsInfoScr, createdAt FROM pddItem ORDER BY createdAt DESC")
  if err != nil {
    log.Println("pdd-item-list-query-error: ", err)
  }
  defer rows.Close()
  var itemList []interface{}
  for rows.Next() {
    var (
      id int64
      quantity int64
      skuGroupPriceMin int64
      skuGroupPriceMax int64
      pddId int64
      goodsName string
      displayPriority string
      thumbUrl string
      isOnsale bool
      soldQuantity int64
      outGoodsSn string
      soldQuantityForThirtyDays int64
      favCnt int64
      ifNewGoods bool
      goodsInfoScr string
      createdAt string
    )
    if err := rows.Scan(&id, &quantity, &skuGroupPriceMin, &skuGroupPriceMax, &pddId, &goodsName, &displayPriority, &thumbUrl, &isOnsale, &soldQuantity, &outGoodsSn, &soldQuantityForThirtyDays, &favCnt, &ifNewGoods, &goodsInfoScr, &createdAt); err != nil {
      log.Println("pdd-item-list-scan-error: ", err)
    }
    item := map[string]interface{}{
      "id": id,
      "quantity": quantity,
      "skuGroupPriceMin": skuGroupPriceMin,
      "skuGroupPriceMax": skuGroupPriceMax,
      "pddId": pddId,
      "goodsName": goodsName,
      "displayPriority": displayPriority,
      "thumbUrl": thumbUrl,
      "isOnsale": isOnsale,
      "soldQuantity": soldQuantity,
      "outGoodsSn": outGoodsSn,
      "soldQuantityForThirtyDays": soldQuantityForThirtyDays,
      "favCnt": favCnt,
      "ifNewGoods": ifNewGoods,
      "goodsInfoScr": goodsInfoScr,
      "createdAt": createdAt,
    }
    itemList = append(itemList, item)
  }
  json.NewEncoder(w).Encode(itemList)
}
