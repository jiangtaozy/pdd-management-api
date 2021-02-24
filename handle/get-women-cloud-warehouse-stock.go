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
  "strings"
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
    log.Println("get-women-cloud-warehouse-stock-decode-body-err: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  err = FetchWomenCloudWarehouseStock(id, womenProductId)
  if err != nil {
    log.Println("get-women-cloud-warehouse-stock-fetch-err: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  io.WriteString(w, "ok")
}

func FetchWomenCloudWarehouseStock(id float64, womenProductId string) error {
  timeStamp := strconv.Itoa(int(time.Now().UnixNano() / 1000000))
  client := &http.Client{}
  req, err := http.NewRequest(
    "GET",
    "https://www.hznzcn.com/product/productContentOutOffList",
    nil,
  )
  if err != nil {
    log.Println("get-women-cloud-warehouse-stock-new-request-error: ", err)
    return err
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
    return err
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
      ycStockTips,
      skuColor,
      skuSize
    ) VALUES (?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("get-women-cloud-warehouse-stock-insert-prepare-error: ", err)
    return err
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      womenItemCloudWarehouseSku
    SET
      ycAvailNum = ?,
      ycStockTips = ?,
      skuColor = ?,
      skuSize = ?
    WHERE
      searchId = ?
    AND
      skuDesc = ?
  `)
  if err != nil {
    log.Println("get-women-cloud-warehouse-stock-update-prepare-error: ", err)
    return err
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(result); i++ {
    sku := result[i].(map[string]interface{})
    skuDesc := sku["Spec"].(string)
    skuDescList := strings.Split(skuDesc, ",")
    skuColor := skuDescList[0]
    skuSize := skuDescList[1]
    if skuSize == "XXL" {
      skuSize = "2XL"
    } else if skuSize == "XXXL" {
      skuSize = "3XL"
    } else if skuSize == "XXXXL" {
      skuSize = "4XL"
    } else if skuSize == "L(120-135斤)" {
      skuSize = "L"
    } else if skuSize == "M（110-125斤）" {
      skuSize = "M"
    } else if skuSize == "S（100-115斤）" {
      skuSize = "S"
    } else if skuSize == "XL（130-160斤）" {
      skuSize = "XL"
    } else if skuSize == "XS(100斤以内)" {
      skuSize = "XS"
    }
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
      return err
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        id,
        womenProductId,
        skuDesc,
        ycAvailNum,
        ycStockTips,
        skuColor,
        skuSize,
      )
      if err != nil {
        log.Println("get-women-cloud-warehouse-stock-insert-exec-error: ", err)
        return err
      }
    } else {
      _, err = stmtUpdate.Exec(
        ycAvailNum,
        ycStockTips,
        skuColor,
        skuSize,
        id,
        skuDesc,
      )
      if err != nil {
        log.Println("get-women-cloud-warehouse-stock-update-exec-error: ", err)
        return err
      }
    }
  }
  return nil
}
