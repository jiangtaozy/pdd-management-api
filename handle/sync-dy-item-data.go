/*
 * Maintained by jemo from 2020.11.25 to now
 * Created by jemo on 2020.11.25 16:39:39
 * Sync Douyin Item Data
 * 同步抖音商品数据
 */

package handle

import (
  "log"
  "strconv"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/openapi-fxg"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncDyItemData(w http.ResponseWriter, r *http.Request) {
  shopId := "973906"
  accessToken := GetAccessToken(shopId)
  page := 0
  data, err := getAndSaveProductListByPage(page, accessToken)
  if err != nil {
    log.Println("sync-dy-item-data-get-product-list-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  allPages := int(data["all_pages"].(float64))
  for i := 1; i < allPages; i++ {
    _, err = getAndSaveProductListByPage(i, accessToken)
    if err != nil {
      log.Println("sync-dy-item-data-get-all-product-list-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  }
  json.NewEncoder(w).Encode(data)
}

func getAndSaveProductListByPage(page int, accessToken string) (map[string]interface{}, error) {
  param := map[string]interface{}{
    "page": strconv.Itoa(page),
    "size": "10",
  }
  data, err := openapiFxg.ProductList(
    appId,
    appSecret,
    accessToken,
    param,
  )
  if err != nil {
    log.Println("sync-dy-item-data-error: ", err)
    return nil, err
  }
  list := data["data"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO dyItem (
      checkStatus,
      cosRatio,
      createTime,
      description,
      discountPrice,
      extra,
      firstCid,
      img,
      marketPrice,
      mobile,
      name,
      outProductId,
      payType,
      productId,
      productIdStr,
      recommendRemark,
      secondCid,
      settlementPrice,
      specId,
      status,
      thirdCid,
      updateTime,
      usp
    ) VALUES(
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    )
  `)
  if err != nil {
    log.Println("sync-dy-item-data-insert-prepare-error: ", err)
    return nil, err
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      dyItem
    SET
      checkStatus = ?,
      cosRatio = ?,
      createTime = ?,
      description = ?,
      discountPrice = ?,
      extra = ?,
      firstCid = ?,
      img = ?,
      marketPrice = ?,
      mobile = ?,
      name = ?,
      outProductId = ?,
      payType = ?,
      productId = ?,
      recommendRemark = ?,
      secondCid = ?,
      settlementPrice = ?,
      specId = ?,
      status = ?,
      thirdCid = ?,
      updateTime = ?,
      usp = ?
    WHERE
      productIdStr = ?
  `)
  if err != nil {
    log.Println("sync-dy-item-data-update-prepare-error: ", err)
    return nil, err
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    item := list[i].(map[string]interface{})
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        dyItem
      WHERE
        productIdStr = ?
    `, item["product_id_str"]).Scan(&count)
    if err != nil {
      log.Println("sync-dy-item-data-count-error: ", err)
      return nil, err
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        item["check_status"],
        item["cos_ratio"],
        item["create_time"],
        item["description"],
        item["discount_price"],
        item["extra"],
        item["first_cid"],
        item["img"],
        item["market_price"],
        item["mobile"],
        item["name"],
        item["out_product_id"],
        item["pay_type"],
        item["product_id"],
        item["product_id_str"],
        item["recommend_remark"],
        item["second_cid"],
        item["settlement_price"],
        item["spec_id"],
        item["status"],
        item["third_cid"],
        item["update_time"],
        item["usp"],
      )
      if err != nil {
        log.Println("sync-dy-item-data-insert-exec-error: ", err)
        log.Println("item: ", item)
        return nil, err
      }
    } else {
      _, err = stmtUpdate.Exec(
        item["check_status"],
        item["cos_ratio"],
        item["create_time"],
        item["description"],
        item["discount_price"],
        item["extra"],
        item["first_cid"],
        item["img"],
        item["market_price"],
        item["mobile"],
        item["name"],
        item["out_product_id"],
        item["pay_type"],
        item["product_id"],
        item["recommend_remark"],
        item["second_cid"],
        item["settlement_price"],
        item["spec_id"],
        item["status"],
        item["third_cid"],
        item["update_time"],
        item["usp"],
        item["product_id_str"],
      )
      if err != nil {
        log.Println("sync-dy-item-data-update-exec-error: ", err)
        log.Println("item: ", item)
        return nil, err
      }
    }
  }
  return data, nil
}
