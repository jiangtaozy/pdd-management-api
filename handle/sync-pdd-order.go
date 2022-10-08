/*
 * Maintained by jemo from 2021.1.14 to now
 * Created by jemo on 2021.1.14 17:10:36
 * Sync Pdd Order
 */

package handle

import (
  "log"
  "time"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncPddOrder(requestBody map[string]interface{}, responseBody map[string]interface{}) {
  var result = responseBody["result"].(map[string]interface{})
  pageItems := result["pageItems"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO itemOrder (
      mallId,
      productName,
      orderId,
      orderStatus,
      orderStatusStr,
      productTotalPrice,
      storeDiscount,
      platformDiscount,
      postage,
      serviceAmount,
      onsiteInstallationFee,
      homeDeliveryFee,
      homeDeliveryAndInstallationFee,
      userPaidAmount,
      receiver,
      phone,
      province,
      city,
      district,
      street,
      paymentTime,
      joinSuccessTime,
      orderConfirmationTime,
      commitmentDeliveryTime,
      deliveryTime,
      confirmDeliveryTime,
      productId,
      productSku,
      numberOfProducts,
      skuId,
      merchantCodeSkuDimension,
      merchantCodeProductDimension,
      trackingNumber,
      afterSaleStatus,
      buyerMessage,
      goodsName,
      goodsType,
      merchantNotes
    ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("save-pdd-order-data-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      itemOrder
    SET
      mallId = ?,
      orderStatus = ?,
      orderStatusStr = ?,
      paymentTime = ?,
      joinSuccessTime = ?,
      orderConfirmationTime = ?,
      commitmentDeliveryTime = ?,
      deliveryTime = ?,
      confirmDeliveryTime = ?,
      trackingNumber = ?,
      afterSaleStatus = ?,
      productTotalPrice = ?,
      storeDiscount = ?,
      platformDiscount = ?,
      postage = ?,
      serviceAmount = ?,
      onsiteInstallationFee = ?,
      homeDeliveryFee = ?,
      homeDeliveryAndInstallationFee = ?,
      userPaidAmount = ?,
      merchantNotes = ?
    WHERE
      orderId = ?
  `)
  if err != nil {
    log.Println("save-pdd-order-data-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(pageItems); i++ {
    order := pageItems[i].(map[string]interface{})
    var orderCount int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        itemOrder
      WHERE
        orderId = ?
    `, order["order_sn"]).Scan(&orderCount)
    if err != nil {
      log.Println("save-pdd-order-data-count-error: ", err)
    }
    payTime := time.Unix(int64(order["confirm_time"].(float64)), 0)
    groupTime := payTime
    confirmTime := payTime
    var promiseShippingTime time.Time
    if order["promise_shipping_time"] != nil {
      promiseShippingTime = time.Unix(int64(order["promise_shipping_time"].(float64)), 0)
    }
    shippingTime := time.Unix(int64(order["shipping_time"].(float64)), 0)
    receiveTime := shippingTime
    if orderCount == 0 {
      _, err = stmtInsert.Exec(
        order["mall_id"],
        order["goods_name"],
        order["order_sn"],
        order["order_status"],
        order["order_status_str"],
        order["goods_amount"],
        order["merchant_discount"],
        order["platform_discount"],
        order["shipping_amount"],
        order["service_amount"],
        order["home_install_value"],
        order["delivery_home_value"],
        order["delivery_install_value"],
        order["order_amount"],
        order["receive_name"],
        order["mobile"],
        order["province_name"],
        order["city_name"],
        order["district_name"],
        order["shipping_address"],
        payTime,
        groupTime,
        confirmTime,
        promiseShippingTime,
        shippingTime,
        receiveTime,
        order["goods_id"],
        order["spec"],
        order["goods_number"],
        order["sku_id"],
        order["out_sku_sn"],
        order["out_goods_sn"],
        order["tracking_number"],
        order["after_sales_status"],
        order["buyer_memo"],
        order["goods_name"],
        order["goods_type"],
        order["mall_remark"],
      )
      if err != nil {
        log.Println("save-pdd-order-data-insert-exec-error: ", err)
      }
    } else {
      _, err = stmtUpdate.Exec(
        order["mall_id"],
        order["order_status"],
        order["order_status_str"],
        payTime,
        groupTime,
        confirmTime,
        promiseShippingTime,
        shippingTime,
        receiveTime,
        order["tracking_number"],
        order["after_sales_status"],
        order["goods_amount"],
        order["merchant_discount"],
        order["platform_discount"],
        order["shipping_amount"],
        order["service_amount"],
        order["home_install_value"],
        order["delivery_home_value"],
        order["delivery_install_value"],
        order["order_amount"],
        order["mall_remark"],
        order["order_sn"],
      )
      if err != nil {
        log.Println("save-pdd-order-data-update-exec-error: ", err)
      }
    }
  }
}
