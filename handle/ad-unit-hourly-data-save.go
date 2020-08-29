/*
 * Maintained by jemo from 2020.8.29 to now
 * Created by jemo on 2020.8.29 07:26:08
 * Save ad unit hourly data
 */

package handle

import (
  "io"
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdUnitHourlyDataSave(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("ad-unit-hourly-data-save-json-decode-error: ", err)
  }
  result := body["result"].([]interface{})
  data := result[0].(map[string]interface{})
  entityId := data["entityId"]
  date := data["date"]
  hourlyReportList := data["hourlyReportList"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddAdUnitHourlyData (
      adId,
      date,
      hour,
      impression,
      click,
      spend,
      orderNum,
      gmv,
      mallFavNum,
      goodsFavNum,
      inquiryNum,
      uniqueView,
      rankAverage,
      rankMedian
    )
    VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("ad-unit-hourly-data-save-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE pddAdUnitHourlyData
    SET
      impression = ?,
      click = ?,
      spend = ?,
      orderNum = ?,
      gmv = ?,
      mallFavNum = ?,
      goodsFavNum = ?,
      inquiryNum = ?,
      uniqueView = ?,
      rankAverage = ?,
      rankMedian = ?
    WHERE
      adId = ? AND
      date = ? AND
      hour = ?
  `)
  if err != nil {
    log.Println("ad-unit-hourly-data-save-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(hourlyReportList); i++ {
    hourlyReport := hourlyReportList[i].(map[string]interface{})
    var count int
    err = db.QueryRow(`
      SELECT COUNT(*)
      FROM pddAdUnitHourlyData
      WHERE adId = ?
        AND date = ?
        AND hour = ?
    `, entityId, date, hourlyReport["hour"]).Scan(&count)
    if err != nil {
      log.Println("ad-unit-hourly-data-save-count-error: ", err)
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        entityId,
        date,
        hourlyReport["hour"],
        hourlyReport["impression"],
        hourlyReport["click"],
        hourlyReport["spend"],
        hourlyReport["orderNum"],
        hourlyReport["gmv"],
        hourlyReport["mallFavNum"],
        hourlyReport["goodsFavNum"],
        hourlyReport["inquiryNum"],
        hourlyReport["uniqueView"],
        hourlyReport["rankAverage"],
        hourlyReport["rankMedian"],
      )
      if err != nil {
        log.Println("ad-unit-hourly-data-save-insert-error: ", err)
      }
    } else {
      _, err = stmtUpdate.Exec(
        hourlyReport["impression"],
        hourlyReport["click"],
        hourlyReport["spend"],
        hourlyReport["orderNum"],
        hourlyReport["gmv"],
        hourlyReport["mallFavNum"],
        hourlyReport["goodsFavNum"],
        hourlyReport["inquiryNum"],
        hourlyReport["uniqueView"],
        hourlyReport["rankAverage"],
        hourlyReport["rankMedian"],
        entityId,
        date,
        hourlyReport["hour"],
      )
      if err != nil {
        log.Println("ad-unit-hourly-data-save-update-error: ", err)
      }
    }
  }
  io.WriteString(w, "ok")
}
