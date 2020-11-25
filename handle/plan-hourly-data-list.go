/*
 * Maintained by jemo from 2020.11.22 to now
 * Created by jemo on 2020.11.22 17:46:31
 * Plan Hourly Data List
 */

package handle

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PlanHourlyDataList(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  id := query["planId"][0]
  db := database.DB
  rows, err := db.Query(`
    SELECT
      entityId,
      date,
      hour,
      impression,
      click,
      spend,
      orderNum,
      gmv,
      mallFavNum,
      goodsFavNum
    FROM
      pddAdHourlyData
    WHERE
      entityId = ?
  `, id)
  if err != nil {
    log.Println("plan-hourly-data-list-query-error: ", err)
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      entityId int64
      date string
      hour int64
      impression int64
      click int64
      spend int64
      orderNum int64
      gmv int64
      mallFavNum int64
      goodsFavNum int64
    )
    err = rows.Scan(
      &entityId,
      &date,
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
      log.Println("plan-hourly-data-list-scan-error: ", err)
    }
    obj := map[string]interface{}{
      "entityId": entityId,
      "date": date,
      "hour": hour,
      "impression": impression,
      "click": click,
      "spend": spend,
      "orderNum": orderNum,
      "gmv": gmv,
      "mallFavNum": mallFavNum,
      "goodsFavNum": goodsFavNum,
    }
    list = append(list, obj)
  }
  json.NewEncoder(w).Encode(list)
}
