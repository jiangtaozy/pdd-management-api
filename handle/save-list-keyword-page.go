/*
 * Maintained by jemo from 2020.11.11 to now
 * Created by jemo on 2020.11.11 15:36:31
 * Save listKeywordPage
 */

package handle

import (
  "log"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SaveListKeywordPage(keywordList string) {
  keywordMap := make(map[string]interface{})
  err := json.Unmarshal([]byte(keywordList), &keywordMap)
  if err != nil {
    log.Println("save-list-keyword-page-json-unmarshal-error: ", err)
  }
  result := keywordMap["result"].(map[string]interface{})
  list := result["keywordList"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO adUnitKeyword (
      adId,
      mallId,
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
      date,
      entityId,
      dimensionType,
      bid,
      bidPremium,
      bidPremiumValue,
      keyword,
      keywordAdIdx,
      qualityScore,
      keywordAdIdxOri,
      keywordId,
      keywordType,
      planStrategy,
      dataOperateStatus
    ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("save-list-keyword-page-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
  UPDATE adUnitKeyword
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
    costPerAppPay = ?,
    entityId = ?,
    dimensionType = ?,
    bid = ?,
    bidPremium = ?,
    bidPremiumValue = ?,
    keyword = ?,
    keywordAdIdx = ?,
    qualityScore = ?,
    keywordAdIdxOri = ?,
    keywordType = ?,
    planStrategy = ?,
    dataOperateStatus = ?
  WHERE
    mallId = ? AND
    adId = ? AND
    keywordId = ? AND
    date = ?
  `)
  if err != nil {
    log.Println("save-list-keyword-page-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    keyword := list[i].(map[string]interface{})
    var date string
    if keyword["beginDate"].(string) == keyword["endDate"].(string) {
      date = keyword["beginDate"].(string)
    } else {
      continue
    }
    var count int
    err = db.QueryRow(`
      SELECT COUNT(*)
      FROM
        adUnitKeyword
      WHERE
        mallId = ? AND
        adId = ? AND
        keywordId = ? AND
        date = ?
    `, keyword["mallId"], keyword["adId"], keyword["keywordId"], date).Scan(&count)
    if err != nil {
      log.Println("save-list-keyword-page-count-error: ", err)
    }
    if keyword["uniqueView"] == nil {
      keyword["uniqueView"] = 0
    }
    if keyword["rankAverage"] == nil {
      keyword["rankAverage"] = 0
    }
    if keyword["rankMedian"] == nil {
      keyword["rankMedian"] = 0
    }
    if keyword["avgPayAmount"] == nil {
      keyword["avgPayAmount"] = 0
    }
    if keyword["appActivateNum"] == nil {
      keyword["appActivateNum"] = 0
    }
    if keyword["costPerAppActivate"] == nil {
      keyword["costPerAppActivate"] = 0
    }
    if keyword["appActivateRate"] == nil {
      keyword["appActivateRate"] = 0
    }
    if keyword["appRegisterNum"] == nil {
      keyword["appRegisterNum"] = 0
    }
    if keyword["costPerAppRegister"] == nil {
      keyword["costPerAppRegister"] = 0
    }
    if keyword["appPayNum"] == nil {
      keyword["appPayNum"] = 0
    }
    if keyword["costPerAppPay"] == nil {
      keyword["costPerAppPay"] = 0
    }
    if count == 0 {
      if keyword["impression"] != nil {
        _, err = stmtInsert.Exec(
          keyword["adId"],
          keyword["mallId"],
          keyword["impression"],
          keyword["click"],
          keyword["ctr"],
          keyword["transactionCost"],
          keyword["spend"],
          keyword["roi"],
          keyword["orderNum"],
          keyword["cpc"],
          keyword["cvr"],
          keyword["gmv"],
          keyword["cpm"],
          keyword["mallFavNum"],
          keyword["goodsFavNum"],
          keyword["inquiryNum"],
          keyword["uniqueView"],
          keyword["rankAverage"],
          keyword["rankMedian"],
          keyword["avgPayAmount"],
          keyword["appActivateNum"],
          keyword["costPerAppActivate"],
          keyword["appActivateRate"],
          keyword["appRegisterNum"],
          keyword["costPerAppRegister"],
          keyword["appPayNum"],
          keyword["costPerAppPay"],
          date,
          keyword["entityId"],
          keyword["dimensionType"],
          keyword["bid"],
          keyword["bidPremium"],
          keyword["bidPremiumValue"],
          keyword["keyword"],
          keyword["keywordAdIdx"],
          keyword["qualityScore"],
          keyword["keywordAdIdxOri"],
          keyword["keywordId"],
          keyword["keywordType"],
          keyword["planStrategy"],
          keyword["dataOperateStatus"],
        )
        if err != nil {
          log.Println("save-list-keyword-page-insert-exec-error: ", err)
          log.Println("keyword: ", keyword)
        }
      }
    } else {
      _, err = stmtUpdate.Exec(
        keyword["impression"],
        keyword["click"],
        keyword["ctr"],
        keyword["transactionCost"],
        keyword["spend"],
        keyword["roi"],
        keyword["orderNum"],
        keyword["cpc"],
        keyword["cvr"],
        keyword["gmv"],
        keyword["cpm"],
        keyword["mallFavNum"],
        keyword["goodsFavNum"],
        keyword["inquiryNum"],
        keyword["uniqueView"],
        keyword["rankAverage"],
        keyword["rankMedian"],
        keyword["avgPayAmount"],
        keyword["appActivateNum"],
        keyword["costPerAppActivate"],
        keyword["appActivateRate"],
        keyword["appRegisterNum"],
        keyword["costPerAppRegister"],
        keyword["appPayNum"],
        keyword["costPerAppPay"],
        keyword["entityId"],
        keyword["dimensionType"],
        keyword["bid"],
        keyword["bidPremium"],
        keyword["bidPremiumValue"],
        keyword["keyword"],
        keyword["keywordAdIdx"],
        keyword["qualityScore"],
        keyword["keywordAdIdxOri"],
        keyword["keywordType"],
        keyword["planStrategy"],
        keyword["dataOperateStatus"],
        keyword["mallId"],
        keyword["adId"],
        keyword["keywordId"],
        date,
      )
      if err != nil {
        log.Println("save-list-keyword-page-update-exec-error: ", err)
      }
    }
  }
}
