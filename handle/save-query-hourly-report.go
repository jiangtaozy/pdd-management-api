/*
 * Maintained by jemo from 2020.11.19 to now
 * Created by jemo on 2020.11.19 14:54:27
 * Save Query Hourly Report
 */

package handle

import (
  "log"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SaveQueryHourlyReport(requestBody map[string]interface{}, responseBody map[string]interface{}) {
  date := requestBody["date"]
  entityId := requestBody["entityId"]
  mallId := requestBody["mallId"]
  if mallId == "" {
    mallId = 0
  }
  planId := requestBody["planId"]
  if planId == "" {
    planId = 0
  }
  unitId := requestBody["unitId"]
  if unitId == "" {
    unitId = 0
  }
  scenesType := requestBody["scenesType"]
  list := responseBody["result"].(map[string]interface{})["hourlyReportList"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddAdHourlyData (
      entityId,
      date,
      hour,
      mallId,
      planId,
      unitId,
      scenesType,
      appActivateNum,
      appActivateRate,
      appPayNum,
      appRegisterNum,
      avgPayAmount,
      click,
      costPerAppActivate,
      costPerAppPay,
      costPerAppRegister,
      costPerFormSubmit,
      costPerLeadOrder,
      costPerMallFav,
      cpc,
      cpm,
      ctr,
      cvr,
      formSubmitNum,
      gmv,
      goodsFavNum,
      impression,
      inquiryNum,
      leadOrderNum,
      liveClickNum,
      liveCommentNum,
      liveShareNum,
      mallFavNum,
      mallFavRate,
      nonCommercialCVR,
      nonFanClickNum,
      nonFanSpend,
      orderNum,
      rankAverage,
      rankMedian,
      roi,
      spend,
      spendWithSdk,
      transactionCost,
      uniqueView
    ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("save-query-hourly-report-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE pddAdHourlyData
    SET
      mallId = ?,
      planId = ?,
      unitId = ?,
      scenesType = ?,
      appActivateNum = ?,
      appActivateRate = ?,
      appPayNum = ?,
      appRegisterNum = ?,
      avgPayAmount = ?,
      click = ?,
      costPerAppActivate = ?,
      costPerAppPay = ?,
      costPerAppRegister = ?,
      costPerFormSubmit = ?,
      costPerLeadOrder = ?,
      costPerMallFav = ?,
      cpc = ?,
      cpm = ?,
      ctr = ?,
      cvr = ?,
      formSubmitNum = ?,
      gmv = ?,
      goodsFavNum = ?,
      impression = ?,
      inquiryNum = ?,
      leadOrderNum = ?,
      liveClickNum = ?,
      liveCommentNum = ?,
      liveShareNum = ?,
      mallFavNum = ?,
      mallFavRate = ?,
      nonCommercialCVR = ?,
      nonFanClickNum = ?,
      nonFanSpend = ?,
      orderNum = ?,
      rankAverage = ?,
      rankMedian = ?,
      roi = ?,
      spend = ?,
      spendWithSdk = ?,
      transactionCost = ?,
      uniqueView = ?
    WHERE
      entityId = ? AND
      date = ? AND
      hour = ?
  `)
  if err != nil {
    log.Println("save-query-hourly-report-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    data := list[i].(map[string]interface{})
    if data["rankMedian"] == nil {
      data["rankMedian"] = 0
    }
    var count int
    err = db.QueryRow(`
      SELECT COUNT(*)
      FROM
        pddAdHourlyData
      WHERE
        entityId = ? AND
        date = ? AND
        hour = ?
    `, entityId, date, data["hour"]).Scan(&count)
    if err != nil {
      log.Println("save-query-hourly-report-count-error: ", err)
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        entityId,
        date,
        data["hour"],
        mallId,
        planId,
        unitId,
        scenesType,
        data["appActivateNum"],
        data["appActivateRate"],
        data["appPayNum"],
        data["appRegisterNum"],
        data["avgPayAmount"],
        data["click"],
        data["costPerAppActivate"],
        data["costPerAppPay"],
        data["costPerAppRegister"],
        data["costPerFormSubmit"],
        data["costPerLeadOrder"],
        data["costPerMallFav"],
        data["cpc"],
        data["cpm"],
        data["ctr"],
        data["cvr"],
        data["formSubmitNum"],
        data["gmv"],
        data["goodsFavNum"],
        data["impression"],
        data["inquiryNum"],
        data["leadOrderNum"],
        data["liveClickNum"],
        data["liveCommentNum"],
        data["liveShareNum"],
        data["mallFavNum"],
        data["mallFavRate"],
        data["nonCommercialCVR"],
        data["nonFanClickNum"],
        data["nonFanSpend"],
        data["orderNum"],
        data["rankAverage"],
        data["rankMedian"],
        data["roi"],
        data["spend"],
        data["spendWithSdk"],
        data["transactionCost"],
        data["uniqueView"],
      )
      if err != nil {
        log.Println("save-query-hourly-report-insert-error: ", err)
        log.Println("data: ", data)
      }
    } else {
      _, err = stmtUpdate.Exec(
        mallId,
        planId,
        unitId,
        scenesType,
        data["appActivateNum"],
        data["appActivateRate"],
        data["appPayNum"],
        data["appRegisterNum"],
        data["avgPayAmount"],
        data["click"],
        data["costPerAppActivate"],
        data["costPerAppPay"],
        data["costPerAppRegister"],
        data["costPerFormSubmit"],
        data["costPerLeadOrder"],
        data["costPerMallFav"],
        data["cpc"],
        data["cpm"],
        data["ctr"],
        data["cvr"],
        data["formSubmitNum"],
        data["gmv"],
        data["goodsFavNum"],
        data["impression"],
        data["inquiryNum"],
        data["leadOrderNum"],
        data["liveClickNum"],
        data["liveCommentNum"],
        data["liveShareNum"],
        data["mallFavNum"],
        data["mallFavRate"],
        data["nonCommercialCVR"],
        data["nonFanClickNum"],
        data["nonFanSpend"],
        data["orderNum"],
        data["rankAverage"],
        data["rankMedian"],
        data["roi"],
        data["spend"],
        data["spendWithSdk"],
        data["transactionCost"],
        data["uniqueView"],
        entityId,
        date,
        data["hour"],
      )
      if err != nil {
        log.Println("save-query-hourly-report-update-error: ", err)
      }
    }
  }
}

