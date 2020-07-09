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
  "database/sql"
)

func PddItemList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query("SELECT pddItem.id, pddItem.quantity, pddItem.skuGroupPriceMin, pddItem.skuGroupPriceMax, pddItem.pddId, pddItem.goodsName, pddItem.displayPriority, pddItem.thumbUrl, pddItem.isOnsale, pddItem.soldQuantity, pddItem.outGoodsSn, pddItem.soldQuantityForThirtyDays, pddItem.favCnt, pddItem.ifNewGoods, pddItem.goodsInfoScr, pddItem.createdAt, item.name, item.shippingPrice, item.suitPrice, item.siteType FROM pddItem LEFT JOIN item ON pddItem.outGoodsSn = item.searchId WHERE item.forSell = true OR item.forSell IS NULL ORDER BY pddItem.createdAt DESC")
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
      goodsInfoScr sql.NullString
      createdAt string
      name sql.NullString
      shippingPrice sql.NullFloat64
      suitPrice sql.NullFloat64
      siteType sql.NullInt32
    )
    if err := rows.Scan(&id, &quantity, &skuGroupPriceMin, &skuGroupPriceMax, &pddId, &goodsName, &displayPriority, &thumbUrl, &isOnsale, &soldQuantity, &outGoodsSn, &soldQuantityForThirtyDays, &favCnt, &ifNewGoods, &goodsInfoScr, &createdAt, &name, &shippingPrice, &suitPrice, &siteType); err != nil {
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
      "goodsInfoScr": goodsInfoScr.String,
      "createdAt": createdAt,
      "name": name.String,
      "shippingPrice": shippingPrice.Float64,
      "suitPrice": suitPrice.Float64,
      "siteType": siteType.Int32,
    }
    itemList = append(itemList, item)
  }
  json.NewEncoder(w).Encode(itemList)
}
