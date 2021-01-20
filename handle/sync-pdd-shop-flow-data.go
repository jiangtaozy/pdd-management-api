/*
 * Maintained by jemo from 2021.1.20 to now
 * Created by jemo on 2021.1.20 14:22:20
 * Sync Pdd Shop Flow Data
 * 同步拼多多店铺流量数据
 */

package handle

import (
  "log"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncPddShopFlowData(requestBody map[string]interface{}, responseBody map[string]interface{}) {
  date := requestBody["queryDate"]
  if date == "" {
    return
  }
  var data = responseBody["result"].(map[string]interface{})
  if data["gpv"] == nil {
    return
  }
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddShopFlowData (
      cfmOrdrCnt,
      cfmOrdrCntIsPercent,
      cfmOrdrCntPct,
      cfmOrdrUsrCnt,
      cfmOrdrUsrCntIsPercent,
      cfmOrdrUsrCntPct,
      goodsFavCnt,
      goodsFavCntIsPercent,
      goodsFavCntPct,
      gpv,
      gpvIsPercent,
      gpvPct,
      guv,
      guvIsPercent,
      guvPct,
      payOrdrAmt,
      payOrdrAmtIsPercent,
      payOrdrAmtPct,
      payOrdrCnt,
      payOrdrCntIsPercent,
      payOrdrCntPct,
      payOrdrUsrCnt,
      payOrdrUsrCntIsPercent,
      payOrdrUsrCntPct,
      payUvRto,
      payUvRtoIsPercent,
      payUvRtoPct,
      vstGoodsCnt,
      vstGoodsCntIsPercent,
      vstGoodsCntPct,
      date
    ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("sync-pdd-shop-flow-data-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      pddShopFlowData
    SET
      cfmOrdrCnt = ?,
      cfmOrdrCntIsPercent = ?,
      cfmOrdrCntPct = ?,
      cfmOrdrUsrCnt = ?,
      cfmOrdrUsrCntIsPercent = ?,
      cfmOrdrUsrCntPct = ?,
      goodsFavCnt = ?,
      goodsFavCntIsPercent = ?,
      goodsFavCntPct = ?,
      gpv = ?,
      gpvIsPercent = ?,
      gpvPct = ?,
      guv = ?,
      guvIsPercent = ?,
      guvPct = ?,
      payOrdrAmt = ?,
      payOrdrAmtIsPercent = ?,
      payOrdrAmtPct = ?,
      payOrdrCnt = ?,
      payOrdrCntIsPercent = ?,
      payOrdrCntPct = ?,
      payOrdrUsrCnt = ?,
      payOrdrUsrCntIsPercent = ?,
      payOrdrUsrCntPct = ?,
      payUvRto = ?,
      payUvRtoIsPercent = ?,
      payUvRtoPct = ?,
      vstGoodsCnt = ?,
      vstGoodsCntIsPercent = ?,
      vstGoodsCntPct = ?
    WHERE
      date = ?
  `)
  if err != nil {
    log.Println("sync-pdd-shop-flow-data-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      pddShopFlowData
    WHERE
      date = ?
  `, date).Scan(&count)
  if err != nil {
    log.Println("sync-pdd-shop-flow-data-count-error: ", err)
  }
  if count == 0 {
    _, err = stmtInsert.Exec(
      data["cfmOrdrCnt"],
      data["cfmOrdrCntIsPercent"],
      data["cfmOrdrCntPct"],
      data["cfmOrdrUsrCnt"],
      data["cfmOrdrUsrCntIsPercent"],
      data["cfmOrdrUsrCntPct"],
      data["goodsFavCnt"],
      data["goodsFavCntIsPercent"],
      data["goodsFavCntPct"],
      data["gpv"],
      data["gpvIsPercent"],
      data["gpvPct"],
      data["guv"],
      data["guvIsPercent"],
      data["guvPct"],
      data["payOrdrAmt"],
      data["payOrdrAmtIsPercent"],
      data["payOrdrAmtPct"],
      data["payOrdrCnt"],
      data["payOrdrCntIsPercent"],
      data["payOrdrCntPct"],
      data["payOrdrUsrCnt"],
      data["payOrdrUsrCntIsPercent"],
      data["payOrdrUsrCntPct"],
      data["payUvRto"],
      data["payUvRtoIsPercent"],
      data["payUvRtoPct"],
      data["vstGoodsCnt"],
      data["vstGoodsCntIsPercent"],
      data["vstGoodsCntPct"],
      date,
    )
    if err != nil {
      log.Println("sync-pdd-shop-flow-data-insert-exec-error: ", err)
      log.Println("data: ", data)
    }
  } else {
    _, err = stmtUpdate.Exec(
      data["cfmOrdrCnt"],
      data["cfmOrdrCntIsPercent"],
      data["cfmOrdrCntPct"],
      data["cfmOrdrUsrCnt"],
      data["cfmOrdrUsrCntIsPercent"],
      data["cfmOrdrUsrCntPct"],
      data["goodsFavCnt"],
      data["goodsFavCntIsPercent"],
      data["goodsFavCntPct"],
      data["gpv"],
      data["gpvIsPercent"],
      data["gpvPct"],
      data["guv"],
      data["guvIsPercent"],
      data["guvPct"],
      data["payOrdrAmt"],
      data["payOrdrAmtIsPercent"],
      data["payOrdrAmtPct"],
      data["payOrdrCnt"],
      data["payOrdrCntIsPercent"],
      data["payOrdrCntPct"],
      data["payOrdrUsrCnt"],
      data["payOrdrUsrCntIsPercent"],
      data["payOrdrUsrCntPct"],
      data["payUvRto"],
      data["payUvRtoIsPercent"],
      data["payUvRtoPct"],
      data["vstGoodsCnt"],
      data["vstGoodsCntIsPercent"],
      data["vstGoodsCntPct"],
      date,
    )
    if err != nil {
      log.Println("sync-pdd-shop-flow-data-update-exec-error: ", err)
    }
  }
}
