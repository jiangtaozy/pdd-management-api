/*
 * Maintained by jemo from 2021.1.20 to now
 * Created by jemo on 2021.1.20 11:36:41
 * Sync Pdd Goods Flow Data
 * 同步拼多多商品流量数据
 */

package handle

import (
  "log"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncPddGoodsFlowData(requestBody map[string]interface{}, responseBody map[string]interface{}) {
  startDate := requestBody["startDate"]
  endDate := requestBody["endDate"]
  if startDate != endDate {
    return
  }
  var result = responseBody["result"].(map[string]interface{})
  goodsDetailList := result["goodsDetailList"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddGoodsFlowData (
      adStrategy,
      cate3AvgGoodsVcr,
      cate3IsPgvAbove,
      cate3PctGoodsVcr,
      cfmOrdrCnt,
      cfmOrdrCntPpr,
      cfmOrdrCntPprIsPercent,
      cfmOrdrCntYtd,
      cfmOrdrGoodsQty,
      cfmOrdrGoodsQtyPpr,
      cfmOrdrGoodsQtyPprIsPercent,
      cfmOrdrGoodsQtyYtd,
      cfmOrdrRtoPpr,
      cfmOrdrRtoPprIsPercent,
      goodsFavCnt,
      goodsFavCntPpr,
      goodsFavCntPprIsPercent,
      goodsFavCntYtd,
      goodsId,
      goodsLabel,
      goodsName,
      goodsPv,
      goodsPvPpr,
      goodsPvPprIsPercent,
      goodsPvYtd,
      goodsUv,
      goodsUvPpr,
      goodsUvPprIsPercent,
      goodsUvYtd,
      goodsVcr,
      goodsVcrPpr,
      goodsVcrPprIsPercent,
      goodsVcrYtd,
      hdThumbUrl,
      imprUsrCnt,
      imprUsrCntPpr,
      imprUsrCntPprIsPercent,
      imprUsrCntYtd,
      isCreated1m,
      isNewstyle,
      payOrdrAmt,
      payOrdrAmtPpr,
      payOrdrAmtPprIsPercent,
      payOrdrAmtYtd,
      payOrdrCnt,
      payOrdrCntPpr,
      payOrdrCntPprIsPercent,
      payOrdrCntYtd,
      payOrdrGoodsQty,
      payOrdrGoodsQtyPpr,
      payOrdrGoodsQtyPprIsPercent,
      payOrdrGoodsQtyYtd,
      payOrdrUsrCnt,
      payOrdrUsrCntPpr,
      payOrdrUsrCntPprIsPercent,
      payOrdrUsrCntYtd,
      pctGoodsVcr,
      pctGoodsVcrYtd,
      statDate,
      url
    ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("sync-pdd-goods-flow-data-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      pddGoodsFlowData
    SET
      adStrategy = ?,
      cate3AvgGoodsVcr = ?,
      cate3IsPgvAbove = ?,
      cate3PctGoodsVcr = ?,
      cfmOrdrCnt = ?,
      cfmOrdrCntPpr = ?,
      cfmOrdrCntPprIsPercent = ?,
      cfmOrdrCntYtd = ?,
      cfmOrdrGoodsQty = ?,
      cfmOrdrGoodsQtyPpr = ?,
      cfmOrdrGoodsQtyPprIsPercent = ?,
      cfmOrdrGoodsQtyYtd = ?,
      cfmOrdrRtoPpr = ?,
      cfmOrdrRtoPprIsPercent = ?,
      goodsFavCnt = ?,
      goodsFavCntPpr = ?,
      goodsFavCntPprIsPercent = ?,
      goodsFavCntYtd = ?,
      goodsLabel = ?,
      goodsName = ?,
      goodsPv = ?,
      goodsPvPpr = ?,
      goodsPvPprIsPercent = ?,
      goodsPvYtd = ?,
      goodsUv = ?,
      goodsUvPpr = ?,
      goodsUvPprIsPercent = ?,
      goodsUvYtd = ?,
      goodsVcr = ?,
      goodsVcrPpr = ?,
      goodsVcrPprIsPercent = ?,
      goodsVcrYtd = ?,
      hdThumbUrl = ?,
      imprUsrCnt = ?,
      imprUsrCntPpr = ?,
      imprUsrCntPprIsPercent = ?,
      imprUsrCntYtd = ?,
      isCreated1m = ?,
      isNewstyle = ?,
      payOrdrAmt = ?,
      payOrdrAmtPpr = ?,
      payOrdrAmtPprIsPercent = ?,
      payOrdrAmtYtd = ?,
      payOrdrCnt = ?,
      payOrdrCntPpr = ?,
      payOrdrCntPprIsPercent = ?,
      payOrdrCntYtd = ?,
      payOrdrGoodsQty = ?,
      payOrdrGoodsQtyPpr = ?,
      payOrdrGoodsQtyPprIsPercent = ?,
      payOrdrGoodsQtyYtd = ?,
      payOrdrUsrCnt = ?,
      payOrdrUsrCntPpr = ?,
      payOrdrUsrCntPprIsPercent = ?,
      payOrdrUsrCntYtd = ?,
      pctGoodsVcr = ?,
      pctGoodsVcrYtd = ?,
      url = ?
    WHERE
      goodsId = ?
    AND
      statDate = ?
  `)
  if err != nil {
    log.Println("sync-pdd-goods-flow-data-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(goodsDetailList); i++ {
    data := goodsDetailList[i].(map[string]interface{})
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        pddGoodsFlowData
      WHERE
        goodsId = ?
      AND
        statDate = ?
    `, data["goodsId"], data["statDate"]).Scan(&count)
    if err != nil {
      log.Println("sync-pdd-goods-flow-data-count-error: ", err)
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        data["adStrategy"],
        data["cate3AvgGoodsVcr"],
        data["cate3IsPgvAbove"],
        data["cate3PctGoodsVcr"],
        data["cfmOrdrCnt"],
        data["cfmOrdrCntPpr"],
        data["cfmOrdrCntPprIsPercent"],
        data["cfmOrdrCntYtd"],
        data["cfmOrdrGoodsQty"],
        data["cfmOrdrGoodsQtyPpr"],
        data["cfmOrdrGoodsQtyPprIsPercent"],
        data["cfmOrdrGoodsQtyYtd"],
        data["cfmOrdrRtoPpr"],
        data["cfmOrdrRtoPprIsPercent"],
        data["goodsFavCnt"],
        data["goodsFavCntPpr"],
        data["goodsFavCntPprIsPercent"],
        data["goodsFavCntYtd"],
        data["goodsId"],
        data["goodsLabel"],
        data["goodsName"],
        data["goodsPv"],
        data["goodsPvPpr"],
        data["goodsPvPprIsPercent"],
        data["goodsPvYtd"],
        data["goodsUv"],
        data["goodsUvPpr"],
        data["goodsUvPprIsPercent"],
        data["goodsUvYtd"],
        data["goodsVcr"],
        data["goodsVcrPpr"],
        data["goodsVcrPprIsPercent"],
        data["goodsVcrYtd"],
        data["hdThumbUrl"],
        data["imprUsrCnt"],
        data["imprUsrCntPpr"],
        data["imprUsrCntPprIsPercent"],
        data["imprUsrCntYtd"],
        data["isCreated1m"],
        data["isNewstyle"],
        data["payOrdrAmt"],
        data["payOrdrAmtPpr"],
        data["payOrdrAmtPprIsPercent"],
        data["payOrdrAmtYtd"],
        data["payOrdrCnt"],
        data["payOrdrCntPpr"],
        data["payOrdrCntPprIsPercent"],
        data["payOrdrCntYtd"],
        data["payOrdrGoodsQty"],
        data["payOrdrGoodsQtyPpr"],
        data["payOrdrGoodsQtyPprIsPercent"],
        data["payOrdrGoodsQtyYtd"],
        data["payOrdrUsrCnt"],
        data["payOrdrUsrCntPpr"],
        data["payOrdrUsrCntPprIsPercent"],
        data["payOrdrUsrCntYtd"],
        data["pctGoodsVcr"],
        data["pctGoodsVcrYtd"],
        data["statDate"],
        data["url"],
      )
      if err != nil {
        log.Println("sync-pdd-goods-flow-data-insert-exec-error: ", err)
        log.Println("data: ", data)
      }
    } else {
      _, err = stmtUpdate.Exec(
        data["adStrategy"],
        data["cate3AvgGoodsVcr"],
        data["cate3IsPgvAbove"],
        data["cate3PctGoodsVcr"],
        data["cfmOrdrCnt"],
        data["cfmOrdrCntPpr"],
        data["cfmOrdrCntPprIsPercent"],
        data["cfmOrdrCntYtd"],
        data["cfmOrdrGoodsQty"],
        data["cfmOrdrGoodsQtyPpr"],
        data["cfmOrdrGoodsQtyPprIsPercent"],
        data["cfmOrdrGoodsQtyYtd"],
        data["cfmOrdrRtoPpr"],
        data["cfmOrdrRtoPprIsPercent"],
        data["goodsFavCnt"],
        data["goodsFavCntPpr"],
        data["goodsFavCntPprIsPercent"],
        data["goodsFavCntYtd"],
        data["goodsLabel"],
        data["goodsName"],
        data["goodsPv"],
        data["goodsPvPpr"],
        data["goodsPvPprIsPercent"],
        data["goodsPvYtd"],
        data["goodsUv"],
        data["goodsUvPpr"],
        data["goodsUvPprIsPercent"],
        data["goodsUvYtd"],
        data["goodsVcr"],
        data["goodsVcrPpr"],
        data["goodsVcrPprIsPercent"],
        data["goodsVcrYtd"],
        data["hdThumbUrl"],
        data["imprUsrCnt"],
        data["imprUsrCntPpr"],
        data["imprUsrCntPprIsPercent"],
        data["imprUsrCntYtd"],
        data["isCreated1m"],
        data["isNewstyle"],
        data["payOrdrAmt"],
        data["payOrdrAmtPpr"],
        data["payOrdrAmtPprIsPercent"],
        data["payOrdrAmtYtd"],
        data["payOrdrCnt"],
        data["payOrdrCntPpr"],
        data["payOrdrCntPprIsPercent"],
        data["payOrdrCntYtd"],
        data["payOrdrGoodsQty"],
        data["payOrdrGoodsQtyPpr"],
        data["payOrdrGoodsQtyPprIsPercent"],
        data["payOrdrGoodsQtyYtd"],
        data["payOrdrUsrCnt"],
        data["payOrdrUsrCntPpr"],
        data["payOrdrUsrCntPprIsPercent"],
        data["payOrdrUsrCntYtd"],
        data["pctGoodsVcr"],
        data["pctGoodsVcrYtd"],
        data["url"],
        data["goodsId"],
        data["statDate"],
      )
      if err != nil {
        log.Println("sync-pdd-goods-flow-data-update-exec-error: ", err)
      }
    }
  }
}
