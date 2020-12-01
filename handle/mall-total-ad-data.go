/*
 * Maintained by jemo from 2020.11.28 to now
 * Created by jemo on 2020.11.28 11:16:58
 * Mall Total Ad Data
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func MallTotalAdData(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      SUM(pddAdUnitDailyData.impression) impression,
      SUM(pddAdUnitDailyData.click) click,
      SUM(pddAdUnitDailyData.spend) spend,
      SUM(pddAdUnitDailyData.orderNum) orderNum,
      SUM(pddAdUnitDailyData.gmv) gmv,
      SUM(pddAdUnitDailyData.mallFavNum) mallFavNum,
      SUM(pddAdUnitDailyData.goodsFavNum) goodsFavNum,
      pddAdUnit.mallId,
      pddAdUnit.scenesType
    FROM
      pddAdUnitDailyData
    LEFT JOIN pddAdUnit
      ON pddAdUnitDailyData.adId = pddAdUnit.adId
    GROUP BY
      pddAdUnit.mallId,
      pddAdUnit.scenesType
  `)
  if err != nil {
    log.Println("mall-total-ad-data-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      impression int64
      click int64
      spend int64
      orderNum int64
      gmv int64
      mallFavNum int64
      goodsFavNum int64
      mallId int64
      scenesType int64
    )
    err := rows.Scan(
      &impression,
      &click,
      &spend,
      &orderNum,
      &gmv,
      &mallFavNum,
      &goodsFavNum,
      &mallId,
      &scenesType,
    )
    if err != nil {
      log.Println("mall-total-ad-data-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    data := map[string]interface{}{
      "impression": impression,
      "click": click,
      "spend": spend,
      "orderNum": orderNum,
      "gmv": gmv,
      "mallFavNum": mallFavNum,
      "goodsFavNum": goodsFavNum,
      "mallId": mallId,
      "scenesType": scenesType,
    }
    list = append(list, data)
  }
  err = json.NewEncoder(w).Encode(list)
  if err != nil {
    log.Println("mall-total-ad-data-encode-err: ", err)
    http.Error(w, err.Error(), 500)
  }
}
