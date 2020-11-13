/*
 * Maintained by jemo from 2020.11.13 to now
 * Created by jemo on 2020.11.13 11:30:31
 * Save unit/listPage
 * 保存推广单元数据
 */

package handle

import (
  "log"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SaveUnitListPage(requestBody map[string]interface{}, responseBody map[string]interface{}) {
  beginDate := requestBody["beginDate"].(string)
  endDate := requestBody["endDate"].(string)
  if beginDate != endDate {
    return
  }
  date := beginDate
  scenesType := requestBody["scenesType"].(float64)
  list := responseBody["result"].(map[string]interface{})["result"].([]interface{})
  db := database.DB
  // ad daily data
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddAdUnitDailyData (
      adId,
      impression,
      click,
      ctr,
      transactionCost,
      spend,
      roi,
      orderNum,
      cpc,
      cvr,
      gmv,
      cpm,
      mallFavNum,
      goodsFavNum,
      inquiryNum,
      uniqueView,
      rankAverage,
      rankMedian,
      avgPayAmount,
      appActivateNum,
      costPerAppActivate,
      appActivateRate,
      appRegisterNum,
      costPerAppRegister,
      appPayNum,
      costPerAppPay,
      date
    )
    VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("save-ad-unit-data-of-date-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE pddAdUnitDailyData
    SET
      impression = ?,
      click = ?,
      ctr = ?,
      transactionCost = ?,
      spend = ?,
      roi = ?,
      orderNum = ?,
      cpc = ?,
      cvr = ?,
      gmv = ?,
      cpm = ?,
      mallFavNum = ?,
      goodsFavNum = ?,
      inquiryNum = ?,
      uniqueView = ?,
      rankAverage = ?,
      rankMedian = ?,
      avgPayAmount = ?,
      appActivateNum = ?,
      costPerAppActivate = ?,
      appActivateRate = ?,
      appRegisterNum = ?,
      costPerAppRegister = ?,
      appPayNum = ?,
      costPerAppPay = ?
    WHERE
      adId = ? AND
      date = ?
  `)
  if err != nil {
    log.Println("save-ad-unit-data-of-date-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  // ad plan
  stmtInsertPlan, err := db.Prepare(`
    INSERT INTO pddAdPlan (
      mallId,
      planId,
      planName,
      scenesType
    )
    VALUES(?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("save-ad-unit-data-of-date-insert-plan-prepare-error: ", err)
  }
  defer stmtInsertPlan.Close()
  stmtUpdatePlan, err := db.Prepare(`
    UPDATE pddAdPlan
    SET
      planName = ?,
      scenesType = ?
    WHERE
      mallId = ? AND
      planId = ?
  `)
  if err != nil {
    log.Println("save-ad-unit-data-of-date-update-plan-prepare-error: ", err)
  }
  defer stmtUpdatePlan.Close()
  // ad unit
  stmtInsertUnit, err := db.Prepare(`
    INSERT INTO pddAdUnit (
      mallId,
      planId,
      adId,
      adName,
      goodsId,
      goodsName,
      scenesType
    )
    VALUES(?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("save-ad-unit-data-of-date-insert-unit-prepare-error: ", err)
  }
  defer stmtInsertUnit.Close()
  stmtUpdateUnit, err := db.Prepare(`
    UPDATE pddAdUnit
    SET
      adName = ?,
      goodsId = ?,
      goodsName = ?,
      scenesType = ?
    WHERE
      mallId = ? AND
      planId = ? AND
      adId = ?
  `)
  if err != nil {
    log.Println("save-ad-unit-data-of-date-update-unit-prepare-error: ", err)
  }
  defer stmtUpdateUnit.Close()
  for i := 0; i < len(list); i++ {
    unit := list[i].(map[string]interface{})
    // insert ad data
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        pddAdUnitDailyData
      WHERE
        adId = ? AND
        date = ?
    `, unit["adId"], date).Scan(&count)
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
      if unit["impression"] != nil {
        _, err = stmtInsert.Exec(
          unit["adId"],
          unit["impression"],
          unit["click"],
          unit["ctr"],
          unit["transactionCost"],
          unit["spend"],
          unit["roi"],
          unit["orderNum"],
          unit["cpc"],
          unit["cvr"],
          unit["gmv"],
          unit["cpm"],
          unit["mallFavNum"],
          unit["goodsFavNum"],
          unit["inquiryNum"],
          unit["uniqueView"],
          unit["rankAverage"],
          unit["rankMedian"],
          unit["avgPayAmount"],
          unit["appActivateNum"],
          unit["costPerAppActivate"],
          unit["appActivateRate"],
          unit["appRegisterNum"],
          unit["costPerAppRegister"],
          unit["appPayNum"],
          unit["costPerAppPay"],
          date,
        )
        if err != nil {
          log.Println("save-ad-unit-data-of-date-insert-exec-error: ", err)
        }
      }
    } else {
      _, err = stmtUpdate.Exec(
        unit["impression"],
        unit["click"],
        unit["ctr"],
        unit["transactionCost"],
        unit["spend"],
        unit["roi"],
        unit["orderNum"],
        unit["cpc"],
        unit["cvr"],
        unit["gmv"],
        unit["cpm"],
        unit["mallFavNum"],
        unit["goodsFavNum"],
        unit["inquiryNum"],
        unit["uniqueView"],
        unit["rankAverage"],
        unit["rankMedian"],
        unit["avgPayAmount"],
        unit["appActivateNum"],
        unit["costPerAppActivate"],
        unit["appActivateRate"],
        unit["appRegisterNum"],
        unit["costPerAppRegister"],
        unit["appPayNum"],
        unit["costPerAppPay"],
        unit["adId"],
        date,
      )
      if err != nil {
        log.Println("save-ad-unit-data-of-date-update-exec-error: ", err)
      }
    }
    // insert plan
    var planCount int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        pddAdPlan
      WHERE
        mallId = ? AND
        planId = ?
    `, unit["mallId"], unit["planId"]).Scan(&planCount)
    if err != nil {
      log.Println("save-ad-unit-data-of-date-plan-count-error: ", err)
    }
    if planCount == 0 {
      _, err = stmtInsertPlan.Exec(
        unit["mallId"],
        unit["planId"],
        unit["planName"],
        scenesType,
      )
      if err != nil {
        log.Println("save-ad-unit-data-of-date-insert-plan-exec-error: ", err)
      }
    } else if planCount == 1 {
      _, err = stmtUpdatePlan.Exec(
        unit["planName"],
        scenesType,
        unit["mallId"],
        unit["planId"],
      )
      if err != nil {
        log.Println("save-ad-unit-data-of-date-update-plan-exec-error: ", err)
      }
    } else {
      stmtDeletePlan, err := db.Prepare(`
        DELETE
        FROM
          pddAdPlan
        WHERE
          mallId = ? AND
          planId = ?
      `)
      if err != nil {
        log.Println("save-ad-unit-data-of-date-prepare-delete-plan-error: ", err)
      }
      _, err = stmtDeletePlan.Exec(
        unit["mallId"],
        unit["planId"],
      )
      if err != nil {
        log.Println("save-ad-unit-data-of-date-exec-delete-plan-error: ", err)
      }
    }
    // insert ad unit
    var unitCount int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        pddAdUnit
      WHERE
        mallId = ? AND
        planId = ? AND
        adId = ?
    `, unit["mallId"], unit["planId"], unit["adId"]).Scan(&unitCount)
    if err != nil {
      log.Println("save-ad-unit-data-of-date-unit-count-error: ", err)
    }
    if unitCount == 0 {
      _, err = stmtInsertUnit.Exec(
        unit["mallId"],
        unit["planId"],
        unit["adId"],
        unit["adName"],
        unit["goodsId"],
        unit["goodsName"],
        scenesType,
      )
      if err != nil {
        log.Println("save-ad-unit-data-of-date-insert-unit-exec-error: ", err)
      }
    } else {
      _, err = stmtUpdateUnit.Exec(
        unit["adName"],
        unit["goodsId"],
        unit["goodsName"],
        scenesType,
        unit["mallId"],
        unit["planId"],
        unit["adId"],
      )
      if err != nil {
        log.Println("save-ad-unit-data-of-date-update-unit-exec-error: ", err)
      }
    }
  }
}
