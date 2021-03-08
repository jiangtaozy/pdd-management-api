/*
 * Maintained by jemo from 2021.3.8 to now
 * Created by jemo on 2021.3.8 14:20:23
 * Download Limit Time Data
 * 下载限时折扣数据
 */

package handle

import (
  "io"
  "log"
  "math"
  "net/http"
  "github.com/tealeg/xlsx/v3"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func DownloadLimitTime(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  planId := query["planId"][0]
  db := database.DB
  rows, err := db.Query(`
    SELECT
      pddAdUnit.goodsId,
      pddItem.skuGroupPriceMin,
      womenItem.price,
      min(pddItemPriceHistory.skuGroupPriceMin) AS skuGroupPriceMin
    FROM
      pddAdUnit
    LEFT JOIN
      pddItem
    ON
      pddAdUnit.goodsId = pddItem.pddId
    LEFT JOIN
      womenItem
    ON
      pddItem.outGoodsSn = womenItem.searchId
    LEFT JOIN
      pddItemPriceHistory
    ON
      pddAdUnit.goodsId = pddItemPriceHistory.pddId
    WHERE
      pddAdUnit.planId = ? AND
      pddItem.isOnsale = 1
    GROUP BY
      pddAdUnit.goodsId
  `, planId)
  if err != nil {
    log.Println("download-limit-time.go-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  wb, err := xlsx.OpenFile("/home/jemo/Downloads/data.xlsx")
  if err != nil {
    log.Println("download-limit-time.go-open-file-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  sh, ok := wb.Sheet["Sheet1"]
  if !ok {
    log.Println("download-limit-time.go-add-sheet-error: ", ok)
    http.Error(w, err.Error(), 500)
    return
  }
  i := 1
  for rows.Next() {
    var (
      goodsId int64
      skuGroupPriceMin int64
      price int64
      historySkuGroupPriceMin int64
    )
    err = rows.Scan(
      &goodsId,
      &skuGroupPriceMin,
      &price,
      &historySkuGroupPriceMin,
    )
    if err != nil {
      log.Println("download-limit-time.go-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    costPrice := price + (5.5 + 6) * 100
    discount := math.Round(float64(costPrice) / (float64(costPrice) / float64(historySkuGroupPriceMin) + 0.05) / float64(skuGroupPriceMin) * 100) * 10
    row, err := sh.Row(i)
    if err != nil {
      log.Println("download-limit-time.go-row-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    cell := row.GetCell(0)
    cell.SetInt64(goodsId)
    cell = row.GetCell(2)
    cell.SetInt64(int64(discount))
    i++
  }
  err = wb.Save("data.xlsx")
  if err != nil {
    log.Println("download-limit-time.go-save-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  io.WriteString(w, "ok")
}
