/*
 * Maintained by jemo from 2020.12.3 to now
 * Created by jemo on 2020.12.3 15:03:11
 * Sync Douyin Order Data
 * 同步抖音订单数据
 */

package handle

import (
  "log"
  "time"
  "strconv"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/openapi-fxg"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncDyOrderData(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  startTime := query["startTime"][0]
  endTime := query["endTime"][0]
  shopId := "973906"
  accessToken := GetAccessToken(shopId)
  page := 0
  size := 10
  data, err := getAndSaveOrderListByPage(
    page,
    size,
    accessToken,
    startTime,
    endTime,
  )
  if err != nil {
    log.Println("sync-dy-order-data-get-by-page-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  total := int(data["total"].(float64))
  allPages := total / size + 1
  for i := 1; i < allPages; i++ {
    _, err = getAndSaveOrderListByPage(
      i,
      size,
      accessToken,
      startTime,
      endTime,
    )
    if err != nil {
      log.Println("sync-dy-order-data-get-all-order-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  }
  json.NewEncoder(w).Encode(data)
}

func getAndSaveOrderListByPage(
  page int,
  size int,
  accessToken string,
  startTime string,
  endTime string,
) (
  map[string]interface{},
  error,
) {
  param := map[string]interface{}{
    "start_time": startTime,
    "end_time": endTime,
    "order_by": "create_time",
    "page": strconv.Itoa(page),
    "size": strconv.Itoa(size),
  }
  data, err := openapiFxg.OrderList(
    appId,
    appSecret,
    accessToken,
    param,
  )
  if err != nil {
    log.Println("sync-dy-order-data-get-and-save-order-list-by-page-error: ", err)
    return nil, err
  }
  if data["list"] == nil {
    return data, nil
  }
  list := data["list"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO dyOrder (
      orderId,
      shopId,
      openId,
      orderStatus,
      orderType,
      orderTag,
      childNum,
      postAddrProvinceId,
      postAddrProvinceName,
      postAddrCityId,
      postAddrCityName,
      postAddrTownId,
      postAddrTownName,
      postAddrDetail,
      postCode,
      postReceiver,
      postTel,
      buyerWords,
      sellerWords,
      logisticsId,
      logisticsCode,
      logisticsTime,
      receiptTime,
      createTime,
      updateTime,
      expShipTime,
      cancelReason,
      payType,
      payTime,
      postAmount,
      couponAmount,
      shopCouponAmount,
      couponMetaId,
      orderTotalAmount,
      isComment,
      urgeCnt,
      bType,
      subBType,
      cBiz,
      isInsurance,
      cType,
      cosRatio,
      userName,
      finalStatus,
      shippedNum
    ) VALUES (
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    )
  `)
  if err != nil {
    log.Println("sync-dy-order-data-insert-prepare-error: ", err)
    return nil, err
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      dyOrder
    SET
      shopId = ?,
      openId = ?,
      orderStatus = ?,
      orderType = ?,
      orderTag = ?,
      childNum = ?,
      postAddrProvinceId = ?,
      postAddrProvinceName = ?,
      postAddrCityId = ?,
      postAddrCityName = ?,
      postAddrTownId = ?,
      postAddrTownName = ?,
      postAddrDetail = ?,
      postCode = ?,
      postReceiver = ?,
      postTel = ?,
      buyerWords = ?,
      sellerWords = ?,
      logisticsId = ?,
      logisticsCode = ?,
      logisticsTime = ?,
      receiptTime = ?,
      createTime = ?,
      updateTime = ?,
      expShipTime = ?,
      cancelReason = ?,
      payType = ?,
      payTime = ?,
      postAmount = ?,
      couponAmount = ?,
      shopCouponAmount = ?,
      couponMetaId = ?,
      orderTotalAmount = ?,
      isComment = ?,
      urgeCnt = ?,
      bType = ?,
      subBType = ?,
      cBiz = ?,
      isInsurance = ?,
      cType = ?,
      cosRatio = ?,
      userName = ?,
      finalStatus = ?,
      shippedNum = ?
    WHERE
      orderId = ?
  `)
  if err != nil {
    log.Println("sync-dy-order-data-update-prepare-error: ", err)
    return nil, err
  }
  defer stmtUpdate.Close()
  for i:= 0; i < len(list); i++ {
    order := list[i].(map[string]interface{})
    postAddr := order["post_addr"].(map[string]interface{})
    province := postAddr["province"].(map[string]interface{})
    city := postAddr["city"].(map[string]interface{})
    town := postAddr["town"].(map[string]interface{})
    postAddrProvinceId := province["id"]
    postAddrProvinceName := province["name"]
    postAddrCityId := city["id"]
    postAddrCityName := city["name"]
    postAddrTownId := town["id"]
    postAddrTownName := town["name"]
    postAddrDetail := postAddr["detail"]
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        dyOrder
      WHERE
        orderId = ?
    `, order["order_id"]).Scan(&count)
    if err != nil {
      log.Println("sync-dy-order-data-count-error: ", err)
      return nil, err
    }
    logisticsTime, err := getDateTime(order["logistics_time"].(string))
    if err != nil {
      log.Println("sync-dy-order-data-logistics-time-error: ", err)
      return nil, err
    }
    receiptTime, err := getDateTime(order["receipt_time"].(string))
    if err != nil {
      log.Println("sync-dy-order-data-receipt-time-error: ", err)
      return nil, err
    }
    createTime, err := getDateTime(order["create_time"].(string))
    if err != nil {
      log.Println("sync-dy-order-data-create-time-error: ", err)
      return nil, err
    }
    updateTime, err := time.Unix(int64(order["update_time"].(float64)), 0), nil
    if err != nil {
      log.Println("sync-dy-order-data-update-time-error: ", err)
      return nil, err
    }
    expShipTime, err := time.Unix(int64(order["exp_ship_time"].(float64)), 0), nil
    if err != nil {
      log.Println("sync-dy-order-data-exp-ship-time-error: ", err)
      return nil, err
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        order["order_id"],
        order["shop_id"],
        order["open_id"],
        order["order_status"],
        order["order_type"],
        order["order_tag"],
        order["child_num"],
        postAddrProvinceId,
        postAddrProvinceName,
        postAddrCityId,
        postAddrCityName,
        postAddrTownId,
        postAddrTownName,
        postAddrDetail,
        order["post_code"],
        order["post_receiver"],
        order["post_tel"],
        order["buyer_words"],
        order["seller_words"],
        order["logistics_id"],
        order["logistics_code"],
        logisticsTime,
        receiptTime,
        createTime,
        updateTime,
        expShipTime,
        order["cancel_reason"],
        order["pay_type"],
        order["pay_time"],
        order["post_amount"],
        order["coupon_amount"],
        order["shop_coupon_amount"],
        order["coupon_meta_id"],
        order["order_total_amount"],
        order["is_comment"],
        order["urge_cnt"],
        order["b_type"],
        order["sub_b_type"],
        order["c_biz"],
        order["is_insurance"],
        order["c_type"],
        order["cos_ratio"],
        order["user_name"],
        order["final_status"],
        order["shipped_num"],
      )
      if err != nil {
        log.Println("sync-dy-order-data-insert-exec-error: ", err)
        log.Println("order: ", order)
        return nil, err
      }
    } else {
      _, err = stmtUpdate.Exec(
        order["shop_id"],
        order["open_id"],
        order["order_status"],
        order["order_type"],
        order["order_tag"],
        order["child_num"],
        postAddrProvinceId,
        postAddrProvinceName,
        postAddrCityId,
        postAddrCityName,
        postAddrTownId,
        postAddrTownName,
        postAddrDetail,
        order["post_code"],
        order["post_receiver"],
        order["post_tel"],
        order["buyer_words"],
        order["seller_words"],
        order["logistics_id"],
        order["logistics_code"],
        logisticsTime,
        receiptTime,
        createTime,
        updateTime,
        expShipTime,
        order["cancel_reason"],
        order["pay_type"],
        order["pay_time"],
        order["post_amount"],
        order["coupon_amount"],
        order["shop_coupon_amount"],
        order["coupon_meta_id"],
        order["order_total_amount"],
        order["is_comment"],
        order["urge_cnt"],
        order["b_type"],
        order["sub_b_type"],
        order["c_biz"],
        order["is_insurance"],
        order["c_type"],
        order["cos_ratio"],
        order["user_name"],
        order["final_status"],
        order["shipped_num"],
        order["order_id"],
      )
      if err != nil {
        log.Println("sync-dy-order-data-update-exec-error: ", err)
        log.Println("order: ", order)
        return nil, err
      }
    }
    child := order["child"].([]interface{})
    err = saveChildOrder(child)
    if err != nil {
      log.Println("sync-dy-order-data-save-child-error: ", err)
      return nil, err
    }
  }
  return data, nil
}

func getDateTime(str string) (time.Time, error) {
    timeInt, err := strconv.ParseInt(str, 10, 64)
    if err != nil {
      log.Println("sync-dy-order-data-get-date-time-error: ", err)
      return time.Time{}, err
    }
    return time.Unix(timeInt, 0), nil
}

func saveChildOrder(list []interface{}) error {
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO dyChildOrder (
      orderId,
      shopId,
      openId,
      orderStatus,
      orderType,
      orderTag,
      postAddrProvinceId,
      postAddrProvinceName,
      postAddrCityId,
      postAddrCityName,
      postAddrTownId,
      postAddrTownName,
      postAddrDetail,
      postCode,
      postReceiver,
      postTel,
      buyerWords,
      sellerWords,
      logisticsId,
      logisticsCode,
      logisticsTime,
      receiptTime,
      createTime,
      updateTime,
      expShipTime,
      cancelReason,
      payType,
      payTime,
      postAmount,
      couponAmount,
      shopCouponAmount,
      couponMetaId,
      totalAmount,
      isComment,
      urgeCnt,
      bType,
      subBType,
      isInsurance,
      cType,
      cosRatio,
      userName,
      finalStatus,
      shippedNum,
      code,
      comboAmount,
      comboId,
      comboNum,
      itemIds,
      outProductId,
      outSkuId,
      pid,
      platformFullAmount,
      productId,
      productName,
      productPic,
      specDesc
    ) VALUES (
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    )
  `)
  if err != nil {
    log.Println("sync-dy-order-data-child-order-insert-prepare-error: ", err)
    return err
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      dyChildOrder
    SET
      shopId = ?,
      openId = ?,
      orderStatus = ?,
      orderType = ?,
      orderTag = ?,
      postAddrProvinceId = ?,
      postAddrProvinceName = ?,
      postAddrCityId = ?,
      postAddrCityName = ?,
      postAddrTownId = ?,
      postAddrTownName = ?,
      postAddrDetail = ?,
      postCode = ?,
      postReceiver = ?,
      postTel = ?,
      buyerWords = ?,
      sellerWords = ?,
      logisticsId = ?,
      logisticsCode = ?,
      logisticsTime = ?,
      receiptTime = ?,
      createTime = ?,
      updateTime = ?,
      expShipTime = ?,
      cancelReason = ?,
      payType = ?,
      payTime = ?,
      postAmount = ?,
      couponAmount = ?,
      shopCouponAmount = ?,
      couponMetaId = ?,
      totalAmount = ?,
      isComment = ?,
      urgeCnt = ?,
      bType = ?,
      subBType = ?,
      isInsurance = ?,
      cType = ?,
      cosRatio = ?,
      userName = ?,
      finalStatus = ?,
      shippedNum = ?,
      code = ?,
      comboAmount = ?,
      comboId = ?,
      comboNum = ?,
      itemIds = ?,
      outProductId = ?,
      outSkuId = ?,
      pid = ?,
      platformFullAmount = ?,
      productId = ?,
      productName = ?,
      productPic = ?,
      specDesc = ?
    WHERE
      orderId = ?
  `)
  if err != nil {
    log.Println("sync-dy-order-data-child-order-update-prepare-error: ", err)
    return err
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    order := list[i].(map[string]interface{})
    postAddr := order["post_addr"].(map[string]interface{})
    province := postAddr["province"].(map[string]interface{})
    city := postAddr["city"].(map[string]interface{})
    town := postAddr["town"].(map[string]interface{})
    postAddrProvinceId := province["id"]
    postAddrProvinceName := province["name"]
    postAddrCityId := city["id"]
    postAddrCityName := city["name"]
    postAddrTownId := town["id"]
    postAddrTownName := town["name"]
    postAddrDetail := postAddr["detail"]
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        dyChildOrder
      WHERE
        orderId = ?
    `, order["order_id"]).Scan(&count)
    if err != nil {
      log.Println("sync-dy-order-data-child-order-count-error: ", err)
      return err
    }
    logisticsTime, err := getDateTime(order["logistics_time"].(string))
    if err != nil {
      log.Println("sync-dy-order-data-logistics-time-error: ", err)
      return err
    }
    receiptTime, err := getDateTime(order["receipt_time"].(string))
    if err != nil {
      log.Println("sync-dy-order-data-receipt-time-error: ", err)
      return err
    }
    createTime, err := getDateTime(order["create_time"].(string))
    if err != nil {
      log.Println("sync-dy-order-data-create-time-error: ", err)
      return err
    }
    updateTime, err := time.Unix(int64(order["update_time"].(float64)), 0), nil
    if err != nil {
      log.Println("sync-dy-order-data-update-time-error: ", err)
      return err
    }
    expShipTime, err := time.Unix(int64(order["exp_ship_time"].(float64)), 0), nil
    if err != nil {
      log.Println("sync-dy-order-data-exp-ship-time-error: ", err)
      return err
    }
    itemIdsByte, err := json.Marshal(order["item_ids"])
    if err != nil {
      log.Println("sync-dy-order-data-marshal-item-ids-error: ", err)
      return err
    }
    itemIds := string(itemIdsByte)
    specDescByte, err := json.Marshal(order["spec_desc"])
    if err != nil {
      log.Println("sync-dy-order-data-marshal-spec-desc-error: ", err)
      return err
    }
    specDesc := string(specDescByte)
    if count == 0 {
      _, err = stmtInsert.Exec(
        order["order_id"],
        order["shop_id"],
        order["open_id"],
        order["order_status"],
        order["order_type"],
        order["order_tag"],
        postAddrProvinceId,
        postAddrProvinceName,
        postAddrCityId,
        postAddrCityName,
        postAddrTownId,
        postAddrTownName,
        postAddrDetail,
        order["post_code"],
        order["post_receiver"],
        order["post_tel"],
        order["buyer_words"],
        order["seller_words"],
        order["logistics_id"],
        order["logistics_code"],
        logisticsTime,
        receiptTime,
        createTime,
        updateTime,
        expShipTime,
        order["cancel_reason"],
        order["pay_type"],
        order["pay_time"],
        order["post_amount"],
        order["coupon_amount"],
        order["shop_coupon_amount"],
        order["coupon_meta_id"],
        order["total_amount"],
        order["is_comment"],
        order["urge_cnt"],
        order["b_type"],
        order["sub_b_type"],
        order["is_insurance"],
        order["c_type"],
        order["cos_ratio"],
        order["user_name"],
        order["final_status"],
        order["shipped_num"],
        order["code"],
        order["combo_amount"],
        order["combo_id"],
        order["combo_num"],
        itemIds,
        order["out_product_id"],
        order["out_sku_id"],
        order["pid"],
        order["platform_full_amount"],
        order["product_id"],
        order["product_name"],
        order["product_pic"],
        specDesc,
      )
      if err != nil {
        log.Println("sync-dy-order-data-child-order-insert-exec-error: ", err)
        log.Println("order: ", order)
        return err
      }
    } else {
      _, err = stmtUpdate.Exec(
        order["shop_id"],
        order["open_id"],
        order["order_status"],
        order["order_type"],
        order["order_tag"],
        postAddrProvinceId,
        postAddrProvinceName,
        postAddrCityId,
        postAddrCityName,
        postAddrTownId,
        postAddrTownName,
        postAddrDetail,
        order["post_code"],
        order["post_receiver"],
        order["post_tel"],
        order["buyer_words"],
        order["seller_words"],
        order["logistics_id"],
        order["logistics_code"],
        logisticsTime,
        receiptTime,
        createTime,
        updateTime,
        expShipTime,
        order["cancel_reason"],
        order["pay_type"],
        order["pay_time"],
        order["post_amount"],
        order["coupon_amount"],
        order["shop_coupon_amount"],
        order["coupon_meta_id"],
        order["total_amount"],
        order["is_comment"],
        order["urge_cnt"],
        order["b_type"],
        order["sub_b_type"],
        order["is_insurance"],
        order["c_type"],
        order["cos_ratio"],
        order["user_name"],
        order["final_status"],
        order["shipped_num"],
        order["code"],
        order["combo_amount"],
        order["combo_id"],
        order["combo_num"],
        itemIds,
        order["out_product_id"],
        order["out_sku_id"],
        order["pid"],
        order["platform_full_amount"],
        order["product_id"],
        order["product_name"],
        order["product_pic"],
        specDesc,
        order["order_id"],
      )
      if err != nil {
        log.Println("sync-dy-order-data-child-order-update-exec-error: ", err)
        log.Println("order: ", order)
        return err
      }
    }
  }
  return nil
}
