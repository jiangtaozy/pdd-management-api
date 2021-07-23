/*
 * Maintained by jemo from 2021.7.20 to now
 * Created by jemo on 2021.7.20 16:05:22
 * 同步女装网订单
 */

package handle

import (
  "log"
  "strconv"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncWomenOrder(requestBody map[string]interface{}, responseBody map[string]interface{}) {
  orders := responseBody["orders"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO order1688 (
      orderId,
      shippingFare,
      actualPayment,
      orderStatus,
      orderStatusStr,
      orderCreatedTime,
      receiver,
      shippingAddress,
      courierCompany,
      trackingNumber,
      orderType,
      deliveryTime,
      productStatus
    ) VALUES(
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    )
  `)
  if err != nil {
    log.Println("sync-women-order-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      order1688
    SET
      orderStatus = ?,
      orderStatusStr = ?,
      courierCompany = ?,
      trackingNumber = ?,
      productStatus = ?
    WHERE
      orderId = ?
  `)
  if err != nil {
    log.Println("sync-women-order-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  orderType := 1
  for i := 0; i < len(orders); i++ {
    order := orders[i].(map[string]interface{})
    orderId := order["orderCode"]
    shippingFare, err := strconv.ParseFloat(order["expFee"].(string), 64)
    if err != nil {
      log.Println("sync-women-order-parse-exp-fee-error: ", err)
    }
    actualPayment, err := strconv.ParseFloat(order["totalFee"].(string), 64)
    if err != nil {
      log.Println("sync-women-order-parse-total-fee-error: ", err)
    }
    orderStatus := order["orderStatus"]
    orderStatusStr := order["statusName"]
    orderCreatedTime := order["created"]
    receiver := order["consigneeRealName"]
    shippingAddress := order["consigneeAddress"]
    courierCompany := order["expCom"]
    trackingNumber := order["expNum"]
    deliveryTime := order["sendDate"]
    if deliveryTime == "" {
      deliveryTime = nil
    }
    productStatus := order["statusName"]
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        order1688
      WHERE
        orderId = ?
    `, orderId).Scan(&count)
    if err != nil {
      log.Println("sync-women-order-count-error: ", err)
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        orderId,
        shippingFare,
        actualPayment,
        orderStatus,
        orderStatusStr,
        orderCreatedTime,
        receiver,
        shippingAddress,
        courierCompany,
        trackingNumber,
        orderType,
        deliveryTime,
        productStatus,
      )
      if err != nil {
        log.Println("sync-women-order-insert-exec-error: ", err)
      }
    } else {
      _, err = stmtUpdate.Exec(
        orderStatus,
        orderStatusStr,
        courierCompany,
        trackingNumber,
        productStatus,
        orderId,
      )
      if err != nil {
        log.Println("sync-women-order-update-exec-error: ", err)
      }
    }
  }
}
