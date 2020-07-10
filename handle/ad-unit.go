/*
 * Maintained by jemo from 2020.5.28 to now
 * Created by jemo on 2020.5.28 09:09:11
 * Ad Unit
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdUnit(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  adIdArray := query["adId"]
  adId := adIdArray[0]
  db := database.DB
  var (
    adName string
    goodsId int64
    goodsName string
    thumbUrl sql.NullString
    skuGroupPriceMax int64
    shippingPrice float64
    suitPrice float64
    siteType int32
  )
  err := db.QueryRow("SELECT pddAdUnit.adId, pddAdUnit.adName, pddAdUnit.goodsId, pddAdUnit.goodsName, pddItem.thumbUrl, pddItem.skuGroupPriceMax, item.shippingPrice, item.suitPrice, item.siteType FROM pddAdUnit LEFT JOIN pddItem ON pddAdUnit.goodsId = pddItem.pddId LEFT JOIN item ON pddItem.outGoodsSn = item.searchId WHERE (item.forSell = true OR item.forSell IS NULL) AND pddAdUnit.adId = ?", adId).Scan(&adId, &adName, &goodsId, &goodsName, &thumbUrl, &skuGroupPriceMax, &shippingPrice, &suitPrice, &siteType)
  if err != nil {
    log.Println("ad-unit-query-error: ", err)
  }
  json.NewEncoder(w).Encode(map[string]interface{}{
    "adId": adId,
    "adName": adName,
    "goodsId": goodsId,
    "goodsName": goodsName,
    "thumbUrl": thumbUrl.String,
    "skuGroupPriceMax": skuGroupPriceMax,
    "shippingPrice": shippingPrice,
    "suitPrice": suitPrice,
    "siteType": siteType,
  })
}
