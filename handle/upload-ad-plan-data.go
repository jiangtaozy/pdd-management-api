/*
 * Maintained by jemo from 2020.5.26 to now
 * Created by jemo on 2020.5.26 17:12:30
 * Upload Ad Plan Data
 */

package handle

import (
  "io"
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func UploadAdPlanData(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("upload-ad-plan-data-decode-err: ", err)
  }
  var adPlanData = body["adPlanData"].(string)
  var scenesType = body["scenesType"]
  adPlanMap := make(map[string]interface{})
  err = json.Unmarshal([]byte(adPlanData), &adPlanMap)
  if err != nil {
    log.Println("upload-ad-plan-data-unmarshal-error: ", err)
  }
  var result = adPlanMap["result"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare("INSERT INTO pddAdPlan (mallId, planId, planName, stickTime, isStick, scenesType) VALUES(?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Println("upload-ad-plan-data-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare("UPDATE pddAdPlan SET planName = ?, stickTime = ?, isStick = ?, scenesType = ? WHERE mallId = ? AND planId = ?")
  if err != nil {
    log.Println("upload-ad-plan-data-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  stmtInsertUnit, err := db.Prepare("INSERT INTO pddAdUnit (mallId, planId, adId, adName, goodsId, goodsName, scenesType) VALUES(?, ?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Println("upload-ad-plan-data-insert-unit-prepare-error: ", err)
  }
  defer stmtInsertUnit.Close()
  stmtUpdateUnit, err := db.Prepare("UPDATE pddAdUnit SET adName = ?, goodsId = ?, goodsName = ?, scenesType = ? WHERE mallId = ? AND planId = ? AND adId = ?")
  if err != nil {
    log.Println("upload-ad-plan-data-update-unit-prepare-error: ", err)
  }
  defer stmtUpdateUnit.Close()
  for i := 0; i < len(result); i++ {
    var adPlan = result[i].(map[string]interface{})
    // insert ad plan
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM pddAdPlan WHERE mallId = ? AND planId = ?", adPlan["mallId"], adPlan["planId"]).Scan(&count)
    if err != nil {
      log.Println("upload-ad-plan-data-count-error: ", err)
    }
    if count == 0 {
      _, err = stmtInsert.Exec(adPlan["mallId"], adPlan["planId"], adPlan["planName"], adPlan["stickTime"], adPlan["isStick"], scenesType)
      if err != nil {
        log.Println("upload-ad-plan-data-insert-exec-error: ", err)
      }
    } else {
      _, err = stmtUpdate.Exec(adPlan["planName"], adPlan["stickTime"], adPlan["isStick"], scenesType, adPlan["mallId"], adPlan["planId"])
      if err != nil {
        log.Println("upload-ad-plan-data-update-exec-error: ", err)
      }
    }
    // insert ad unit
    var adUnitBaseMapList = adPlan["adUnitBaseMapList"].([]interface{})
    for j := 0; j < len(adUnitBaseMapList); j++ {
      var adUnit = adUnitBaseMapList[j].(map[string]interface{})
      var unitCount int
      err = db.QueryRow("SELECT COUNT(*) FROM pddAdUnit WHERE mallId = ? AND planId = ? AND adId = ?", adUnit["mallId"], adUnit["planId"], adUnit["adId"]).Scan(&unitCount)
      if err != nil {
        log.Println("upload-ad-plan-data-unit-count-error: ", err)
      }
      if unitCount == 0 {
        _, err = stmtInsertUnit.Exec(adUnit["mallId"], adUnit["planId"], adUnit["adId"], adUnit["adName"], adUnit["goodsId"], adUnit["goodsName"], scenesType)
        if err != nil {
          log.Println("upload-ad-plan-data-insert-unit-exec-error: ", err)
        }
      } else {
        _, err = stmtUpdateUnit.Exec(adUnit["adName"], adUnit["goodsId"], adUnit["goodsName"], scenesType, adUnit["mallId"], adUnit["planId"], adUnit["adId"], )
        if err != nil {
          log.Println("upload-ad-plan-data-update-unit-exec-error: ", err)
        }
      }
    }
  }
  io.WriteString(w, "ok")
}
