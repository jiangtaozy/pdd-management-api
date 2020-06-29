/*
 * Maintained by jemo from 2020.6.27 to now
 * Created by jemo on 2020.6.27 21:30:18
 * Ad Day Data
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdDayData(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query("SELECT SUM(impression) impression, SUM(click) click, SUM(spend) spend, SUM(orderNum) orderNum, SUM(gmv) gmv, SUM(mallFavNum) mallFavNum, SUM(goodsFavNum) goodsFavNum, date from pddAdUnitDailyData GROUP BY date ORDER BY date ASC")
  if err != nil {
    log.Println("ad-day-data-query-error: ", err)
  }
  defer rows.Close()
  var adDayDataList []interface{}
  for rows.Next() {
    var (
      impression int64
      click int64
      spend int64
      orderNum int64
      gmv int64
      mallFavNum int64
      goodsFavNum int64
      date string
    )
    err := rows.Scan(&impression, &click, &spend, &orderNum, &gmv, &mallFavNum, &goodsFavNum, &date)
    if err != nil {
      log.Println("ad-day-data-scan-error: ", err)
    }
    adDayData := map[string]interface{}{
      "impression": impression,
      "click": click,
      "date": date,
      "spend": spend,
      "orderNum": orderNum,
      "gmv": gmv,
      "mallFavNum": mallFavNum,
      "goodsFavNum": goodsFavNum,
    }
    adDayDataList = append(adDayDataList, adDayData)
  }
  err = json.NewEncoder(w).Encode(adDayDataList)
  if err != nil {
    log.Println("ad-day-data-encode-err: ", err)
  }
}

