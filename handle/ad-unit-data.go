/*
 * Maintained by jemo from 2020.5.27 to now
 * Created by jemo on 2020.5.27 16:07:25
 * Ad Unit Data
 */

package handle

import (
  "io"
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdUnitData(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("ad-unit-data-decode-err: ", err)
  }
  adId := body["adId"]
  unitDataString := body["unitData"].(string)
  unitDataMap := make(map[string]interface{})
  err = json.Unmarshal([]byte(unitDataString), &unitDataMap)
  if err != nil {
    log.Println("ad-unit-data-json-unmarshal-error: ", err)
  }
  result := unitDataMap["result"].(map[string]interface{})
  list := result["result"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare("INSERT INTO pddAdUnitDailyData (adId, impression, click, ctr, transactionCost, spend, roi, orderNum, cpc, cvr, gmv, cpm, mallFavNum, goodsFavNum, inquiryNum, uniqueView, rankAverage, rankMedian, avgPayAmount, appActivateNum, costPerAppActivate, appActivateRate, appRegisterNum, costPerAppRegister, appPayNum, costPerAppPay, date, entityId, dimensionType) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Println("ad-unit-data-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare("UPDATE pddAdUnitDailyData SET impression = ?, click = ?, ctr = ?, transactionCost = ?, spend = ?, roi = ?, orderNum = ?, cpc = ?, cvr = ?, gmv = ?, cpm = ?, mallFavNum = ?, goodsFavNum = ?, inquiryNum = ?, uniqueView = ?, rankAverage = ?, rankMedian = ?, avgPayAmount = ?, appActivateNum = ?, costPerAppActivate = ?, appActivateRate = ?, appRegisterNum = ?, costPerAppRegister = ?, appPayNum = ?, costPerAppPay = ?, entityId = ?, dimensionType = ? WHERE adId = ? AND date = ?")
  if err != nil {
    log.Println("ad-unit-data-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    unit := list[i].(map[string]interface{})
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM pddAdUnitDailyData WHERE adId = ? AND date = ?", adId, unit["date"]).Scan(&count)
    if err != nil {
      log.Println("ad-unit-data-count-error: ", err)
    }
    if count == 0 {
      _, err = stmtInsert.Exec(adId, unit["impression"], unit["click"], unit["ctr"], unit["transactionCost"], unit["spend"], unit["roi"], unit["orderNum"], unit["cpc"], unit["cvr"], unit["gmv"], unit["cpm"], unit["mallFavNum"], unit["goodsFavNum"], unit["inquiryNum"], unit["uniqueView"], unit["rankAverage"], unit["rankMedian"], unit["avgPayAmount"], unit["appActivateNum"], unit["costPerAppActivate"], unit["appActivateRate"], unit["appRegisterNum"], unit["costPerAppRegister"], unit["appPayNum"], unit["costPerAppPay"], unit["date"], unit["entityId"], unit["dimensionType"])
      if err != nil {
        log.Println("ad-unit-data-insert-exec-error: ", err)
      }
    } else {
      _, err = stmtUpdate.Exec(unit["impression"], unit["click"], unit["ctr"], unit["transactionCost"], unit["spend"], unit["roi"], unit["orderNum"], unit["cpc"], unit["cvr"], unit["gmv"], unit["cpm"], unit["mallFavNum"], unit["goodsFavNum"], unit["inquiryNum"], unit["uniqueView"], unit["rankAverage"], unit["rankMedian"], unit["avgPayAmount"], unit["appActivateNum"], unit["costPerAppActivate"], unit["appActivateRate"], unit["appRegisterNum"], unit["costPerAppRegister"], unit["appPayNum"], unit["costPerAppPay"], unit["entityId"], unit["dimensionType"], adId, unit["date"])
      if err != nil {
        log.Println("ad-unit-data-update-exec-error: ", err)
      }
    }
  }
  io.WriteString(w, "ok")
}
