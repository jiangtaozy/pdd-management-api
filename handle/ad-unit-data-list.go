/*
 * Maintained by jemo from 2020.5.28 to now
 * Created by jemo on 2020.5.28 07:55:27
 * Ad Unit Data List
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdUnitDataList(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  adIdArray := query["adId"]
  adId := adIdArray[0]
  db := database.ConnectDB()
  rows, err := db.Query("SELECT adId, impression, click, ctr, transactionCost, spend, roi, orderNum, cpc, cvr, gmv, cpm, mallFavNum, goodsFavNum, inquiryNum, uniqueView, rankAverage, rankMedian, avgPayAmount, appActivateNum, costPerAppActivate, appActivateRate, appRegisterNum, costPerAppRegister, appPayNum, costPerAppPay, date FROM pddAdUnitDailyData WHERE adId = ? ORDER BY date ASC", adId)
  if err != nil {
    log.Println("ad-unit-data-list-query-error: ", err)
  }
  defer rows.Close()
  var unitDataList []interface{}
  for rows.Next() {
    var (
      adId int64
      impression int64
      click int64
      ctr float64
      transactionCost int64
      spend int64
      roi float64
      orderNum int64
      cpc float64
      cvr float64
      gmv int64
      cpm float64
      mallFavNum int64
      goodsFavNum int64
      inquiryNum int64
      uniqueView int64
      rankAverage int64
      rankMedian int64
      avgPayAmount int64
      appActivateNum int64
      costPerAppActivate int64
      appActivateRate int64
      appRegisterNum float64
      costPerAppRegister int64
      appPayNum int64
      costPerAppPay int64
      date string
    )
    if err := rows.Scan(&adId, &impression, &click, &ctr, &transactionCost, &spend, &roi, &orderNum, &cpc, &cvr, &gmv, &cpm, &mallFavNum, &goodsFavNum, &inquiryNum, &uniqueView, &rankAverage, &rankMedian, &avgPayAmount, &appActivateNum, &costPerAppActivate, &appActivateRate, &appRegisterNum, &costPerAppRegister, &appPayNum, &costPerAppPay, &date); err != nil {
      log.Println("ad-unit-data-list-scan-error: ", err)
    }
    unitData := map[string]interface{}{
      "adId": adId,
      "impression": impression,
      "click": click,
      "ctr": ctr,
      "transactionCost": transactionCost,
      "spend": spend,
      "roi": roi,
      "orderNum": orderNum,
      "cpc": cpc,
      "cvr": cvr,
      "gmv": gmv,
      "cpm": cpm,
      "mallFavNum": mallFavNum,
      "goodsFavNum": goodsFavNum,
      "inquiryNum": inquiryNum,
      "uniqueView": uniqueView,
      "rankAverage": rankAverage,
      "rankMedian": rankMedian,
      "avgPayAmount": avgPayAmount,
      "appActivateNum": appActivateNum,
      "costPerAppActivate": costPerAppActivate,
      "appActivateRate": appActivateRate,
      "appRegisterNum": appRegisterNum,
      "costPerAppRegister": costPerAppRegister,
      "appPayNum": appPayNum,
      "costPerAppPay": costPerAppPay,
      "date": date,
    }
    unitDataList = append(unitDataList, unitData)
  }
  json.NewEncoder(w).Encode(unitDataList)
}

