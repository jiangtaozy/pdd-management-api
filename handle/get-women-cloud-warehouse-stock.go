/*
 * Maintained by jemo from 2021.2.10 to now
 * Created by jemo on 2021.2.10 16:58:16
 * Get Women Cloud Warehouse Stock
 * 获取女装网云仓库存
 */

package handle

import (
  "io"
  "log"
  "time"
  "strconv"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func GetWomenCloudWarehouseStock(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  id := body["id"].(float64)
  womenProductId := strconv.Itoa(int(body["womenProductId"].(float64)))
  if err != nil {
    log.Println("get-women-detail-data-decode-body-err: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  timeStamp := strconv.Itoa(int(time.Now().UnixNano() / 1000000))
  client := &http.Client{}
  req, err := http.NewRequest(
    "GET",
    "https://www.hznzcn.com/product/productContentOutOffList",
    nil,
  )
  if err != nil {
    log.Println("get-women-cloud-warehouse-stock-new-request-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  query := req.URL.Query()
  query.Add("Option", "OutOffList")
  query.Add("productId", womenProductId)
  query.Add("ycProductTag", "1")
  query.Add("t", timeStamp)
  req.URL.RawQuery = query.Encode()
  resp, err := client.Do(req)
  if err != nil {
    log.Println("get-women-cloud-warehouse-stock-client-do-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  var result []interface{}
  json.NewDecoder(resp.Body).Decode(&result)
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO womenItemCloudWarehouseSku (
      searchId,
      productId,
      skuDesc,
      ycAvailNum,
      ycStockTips
    ) VALUES (?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("get-women-cloud-warehouse-stock-insert-prepare-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      womenItemCloudWarehouseSku
    SET
      ycAvailNum = ?,
      ycStockTips = ?
    WHERE
      searchId = ?
    AND
      skuDesc = ?
  `)
  if err != nil {
    log.Println("get-women-cloud-warehouse-stock-update-prepare-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(result); i++ {
    sku := result[i].(map[string]interface{})
    skuDesc := sku["Spec"]
    ycAvailNum := sku["ycAvailNum"]
    ycStockTips := sku["ycStockTips"]
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        womenItemCloudWarehouseSku
      WHERE
        searchId = ?
      AND
        skuDesc = ?
    `, id, skuDesc).Scan(&count)
    if err != nil {
      log.Println("get-women-cloud-warehouse-stock-count-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        id,
        womenProductId,
        skuDesc,
        ycAvailNum,
        ycStockTips,
      )
      if err != nil {
        log.Println("get-women-cloud-warehouse-stock-insert-exec-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    } else {
      _, err = stmtUpdate.Exec(
        ycAvailNum,
        ycStockTips,
        id,
        skuDesc,
      )
      if err != nil {
        log.Println("get-women-cloud-warehouse-stock-update-exec-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    }
  }
  io.WriteString(w, "ok")
}
