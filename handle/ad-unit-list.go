/*
 * Maintained by jemo from 2020.5.27 to now
 * Created by jemo on 2020.5.27 11:06:07
 * Ad Unit List
 */

package handle

import (
  "encoding/json"
  //"database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdUnitList(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  planIdArray := query["planId"]
  planId := planIdArray[0]
  db := database.DB
  rows, err := db.Query("SELECT mallId, planId, adId, adName, goodsId, goodsName, scenesType FROM pddAdUnit WHERE planId = ?", planId)
  if err != nil {
    log.Println("ad-unit-list-query-error: ", err)
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
    )
    if err := rows.Scan(&mallId, &planId, &adId, &adName, &goodsId, &goodsName, &scenesType); err != nil {
      log.Println("ad-unit-list-scan-error: ", err)
    }
    adUnit := map[string]interface{}{
      "mallId": mallId,
      "planId": planId,
      "adId": adId,
      "adName": adName,
      "goodsId": goodsId,
      "goodsName": goodsName,
      "scenesType": scenesType,
    }
    adUnitList = append(adUnitList, adUnit)
  }
  json.NewEncoder(w).Encode(adUnitList)
}
