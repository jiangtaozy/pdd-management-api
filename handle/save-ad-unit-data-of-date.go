/*
 * Maintained by jemo from 2020.6.29 to now
 * Created by jemo on 2020.6.29 12:18:19
 * Save Ad Unit Data Of Date
 */

package handle

import (
  "io"
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SaveAdUnitDataOfDate(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("save-ad-unit-data-of-date-decode-err: ", err)
  }
  date := body["date"]
  unitListDataString := body["unitListData"].(string)
  unitListDataMap := make(map[string]interface{})
  err = json.Unmarshal([]byte(unitListDataString), &unitListDataMap)
  if err != nil {
    log.Println("save-ad-unit-data-of-date-json-unmarshal-error: ", err)
  }
  list := unitListDataMap["result"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare("INSERT INTO pddAdUnitDailyData (adId, impression, click, ctr, transactionCost, spend, roi, orderNum, cpc, cvr, gmv, cpm, mallFavNum, goodsFavNum, inquiryNum, uniqueView, rankAverage, rankMedian, avgPayAmount, appActivateNum, costPerAppActivate, appActivateRate, appRegisterNum, costPerAppRegister, appPayNum, costPerAppPay, date) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Println("save-ad-unit-data-of-date-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare("UPDATE pddAdUnitDailyData SET impression = ?, click = ?, ctr = ?, transactionCost = ?, spend = ?, roi = ?, orderNum = ?, cpc = ?, cvr = ?, gmv = ?, cpm = ?, mallFavNum = ?, goodsFavNum = ?, inquiryNum = ?, uniqueView = ?, rankAverage = ?, rankMedian = ?, avgPayAmount = ?, appActivateNum = ?, costPerAppActivate = ?, appActivateRate = ?, appRegisterNum = ?, costPerAppRegister = ?, appPayNum = ?, costPerAppPay = ? WHERE adId = ? AND date = ?")
  if err != nil {
    log.Println("save-ad-unit-data-of-date-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    unit := list[i].(map[string]interface{})
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM pddAdUnitDailyData WHERE adId = ? AND date = ?", unit["adId"], date).Scan(&count)
    if err != nil {
      log.Println("save-ad-unit-data-of-date-count-error: ", err)
    }
    if unit["uniqueView"] == nil {
      unit["uniqueView"] = 0
    }
    if unit["rankAverage"] == nil {
      unit["rankAverage"] = 0
    }
    if unit["rankMedian"] == nil {
      unit["rankMedian"] = 0
    }
    if unit["avgPayAmount"] == nil {
      unit["avgPayAmount"] = 0
    }
    if unit["appActivateNum"] == nil {
      unit["appActivateNum"] = 0
    }
    if unit["costPerAppActivate"] == nil {
      unit["costPerAppActivate"] = 0
    }
    if unit["appActivateRate"] == nil {
      unit["appActivateRate"] = 0
    }
    if unit["appRegisterNum"] == nil {
      unit["appRegisterNum"] = 0
    }
    if unit["costPerAppRegister"] == nil {
      unit["costPerAppRegister"] = 0
    }
    if unit["appPayNum"] == nil {
      unit["appPayNum"] = 0
    }
    if unit["costPerAppPay"] == nil {
      unit["costPerAppPay"] = 0
    }
    if count == 0 {
      _, err = stmtInsert.Exec(unit["adId"], unit["impression"], unit["click"], unit["ctr"], unit["transactionCost"], unit["spend"], unit["roi"], unit["orderNum"], unit["cpc"], unit["cvr"], unit["gmv"], unit["cpm"], unit["mallFavNum"], unit["goodsFavNum"], unit["inquiryNum"], unit["uniqueView"], unit["rankAverage"], unit["rankMedian"], unit["avgPayAmount"], unit["appActivateNum"], unit["costPerAppActivate"], unit["appActivateRate"], unit["appRegisterNum"], unit["costPerAppRegister"], unit["appPayNum"], unit["costPerAppPay"], date)
      if err != nil {
        log.Println("save-ad-unit-data-of-date-insert-exec-error: ", err)
      }
    } else {
      _, err = stmtUpdate.Exec(unit["impression"], unit["click"], unit["ctr"], unit["transactionCost"], unit["spend"], unit["roi"], unit["orderNum"], unit["cpc"], unit["cvr"], unit["gmv"], unit["cpm"], unit["mallFavNum"], unit["goodsFavNum"], unit["inquiryNum"], unit["uniqueView"], unit["rankAverage"], unit["rankMedian"], unit["avgPayAmount"], unit["appActivateNum"], unit["costPerAppActivate"], unit["appActivateRate"], unit["appRegisterNum"], unit["costPerAppRegister"], unit["appPayNum"], unit["costPerAppPay"], unit["adId"], date)
      if err != nil {
        log.Println("save-ad-unit-data-of-date-update-exec-error: ", err)
      }
    }
  }
  io.WriteString(w, "ok")
}
