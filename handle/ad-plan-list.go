/*
 * Maintained by jemo from 2020.5.27 to now
 * Created by jemo on 2020.5.27 10:26:23
 * Ad Plan List
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdPlanList(w http.ResponseWriter, r *http.Request) {
  db := database.ConnectDB()
  rows, err := db.Query("SELECT mallId, planId, planName, stickTime, isStick, scenesType FROM pddAdPlan")
  if err != nil {
    log.Println("ad-plan-list-query-error: ", err)
  }
  defer rows.Close()
  var adPlanList []interface{}
  for rows.Next() {
    var (
      mallId int64
      planId int64
      planName string
      stickTime sql.NullString
      isStick bool
      scenesType int64
    )
    if err := rows.Scan(&mallId, &planId, &planName, &stickTime, &isStick, &scenesType); err != nil {
      log.Println("adPlan-list-scan-error: ", err)
    }
    adPlan := map[string]interface{}{
      "mallId": mallId,
      "planId": planId,
      "planName": planName,
      "stickTime": stickTime.String,
      "isStick": isStick,
      "scenesType": scenesType,
    }
    adPlanList = append(adPlanList, adPlan)
  }
  json.NewEncoder(w).Encode(adPlanList)
}
