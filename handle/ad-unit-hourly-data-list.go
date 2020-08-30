/*
 * Maintained by jemo from 2020.8.29 to now
 * Created by jemo on 2020.8.29 16:52:22
 * Ad Unit Hourly Data List
 */

package handle

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdUnitHourlyDataList(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  idArray := query["id"]
  id := idArray[0]
  db := database.DB
  rows, err := db.Query(`
    SELECT
      hour,
      SUM(impression) impression,
      SUM(click) click,
      SUM(spend) spend,
      SUM(orderNum) orderNum,
      SUM(gmv) gmv,
      SUM(mallFavNum) mallFavNum,
      SUM(goodsFavNum) goodsFavNum
    FROM pddAdUnitHourlyData
    WHERE adId = ?
    GROUP BY hour
  `, id)
  if err != nil {
    log.Println("ad-unit-hourly-data-list-query-error: ", err)
  }
  defer rows.Close()
  var hourlyData []interface{}
  for rows.Next() {
    var (
      hour int64
      impression int64
      click int64
      spend int64
      orderNum int64
      gmv int64
      mallFavNum int64
      goodsFavNum int64
    )
    err := rows.Scan(
      &hour,
      &impression,
      &click,
      &spend,
      &orderNum,
      &gmv,
      &mallFavNum,
      &goodsFavNum,
    )
    if err != nil {
      log.Println("ad-unit-hourly-data-list-scan-error: ", err)
    }
    hourly := map[string]interface{}{
      "hour": hour,
      "impression": impression,
      "click": click,
      "spend": spend,
      "orderNum": orderNum,
      "gmv": gmv,
      "mallFavNum": mallFavNum,
      "goodsFavNum": goodsFavNum,
    }
    hourlyData = append(hourlyData, hourly)
  }
  json.NewEncoder(w).Encode(hourlyData)
}
