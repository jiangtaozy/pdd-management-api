/*
 * Maintained by jemo from 2020.6.13 to now
 * Created by jemo on 2020.6.13 16:07:23
 * Ad Unit List All
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdUnitListAll(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query("SELECT adUnit.mallId, adUnit.planId, adUnit.adId, adUnit.adName, adUnit.goodsId, adUnit.goodsName, adUnit.scenesType, adData.impression, adData.click, adData.spend, adData.orderNum, adData.gmv, adData.mallFavNum, adData.goodsFavNum, adPlan.planName FROM pddAdUnit AS adUnit LEFT JOIN (SELECT adId, SUM(impression) impression, SUM(click) click, SUM(spend) spend, SUM(orderNum) orderNum, SUM(gmv) gmv, SUM(mallFavNum) mallFavNum, SUM(goodsFavNum) goodsFavNum FROM pddAdUnitDailyData GROUP BY adId) AS adData ON adUnit.adId = adData.adId LEFT JOIN pddAdPlan AS adPlan ON adUnit.planId = adPlan.planId ORDER BY adData.orderNum DESC")
  if err != nil {
    log.Println("ad-unit-list-all-query-error: ", err)
  }
  defer rows.Close()
  var adUnitList []interface{}
  for rows.Next() {
    var (
      mallId int64
      planId int64
      adId int64
      adName string
      goodsId int64
      goodsName string
      scenesType int64
      impression sql.NullInt64
      click sql.NullInt64
      spend sql.NullInt64
      orderNum sql.NullInt64
      gmv sql.NullInt64
      mallFavNum sql.NullInt64
      goodsFavNum sql.NullInt64
      planName string
    )
    if err := rows.Scan(&mallId, &planId, &adId, &adName, &goodsId, &goodsName, &scenesType, &impression, &click, &spend, &orderNum, &gmv, &mallFavNum, &goodsFavNum, &planName); err != nil {
      log.Println("ad-unit-list-all-scan-error: ", err)
    }
    adUnit := map[string]interface{}{
      "mallId": mallId,
      "planId": planId,
      "adId": adId,
      "adName": adName,
      "goodsId": goodsId,
      "goodsName": goodsName,
      "scenesType": scenesType,
      "impression": impression.Int64,
      "click": click.Int64,
      "spend": spend.Int64,
      "orderNum": orderNum.Int64,
      "gmv": gmv.Int64,
      "mallFavNum": mallFavNum.Int64,
      "goodsFavNum": goodsFavNum.Int64,
      "planName": planName,
    }
    if impression.Valid && click.Valid && impression.Int64 != 0 {
      adUnit["ctr"] = float64(click.Int64) / float64(impression.Int64)
    }
    if click.Valid && orderNum.Valid && click.Int64 != 0 {
      adUnit["cvr"] = float64(orderNum.Int64) / float64(click.Int64)
    }
    if click.Valid && mallFavNum.Valid && click.Int64 != 0 {
      adUnit["cmfr"] = float64(mallFavNum.Int64) / float64(click.Int64)
    }
    if click.Valid && goodsFavNum.Valid && click.Int64 != 0 {
      adUnit["cgfr"] = float64(goodsFavNum.Int64) / float64(click.Int64)
    }
    adUnitList = append(adUnitList, adUnit)
  }
  err = json.NewEncoder(w).Encode(adUnitList)
  if err != nil {
    log.Println("ad-unit-list-all-encode-error: ", err)
  }
}
