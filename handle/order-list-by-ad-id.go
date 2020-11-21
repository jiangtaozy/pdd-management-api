/*
 * Maintained by jemo from 2020.11.17 to now
 * Created by jemo on 2020.11.17 17:46:07
 * Order List By Ad Id
 */

package handle

import (
  "log"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func OrderListByAdId(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  id := query["adId"][0]
  db := database.DB
  rows, err := db.Query(`
    SELECT
      itemOrder.orderId,
      itemOrder.orderStatus,
      itemOrder.orderStatusStr,
      itemOrder.userPaidAmount,
      itemOrder.platformDiscount,
      itemOrder.afterSaleStatus,
      itemOrder.paymentTime,
      order1688.actualPayment,
      order1688.orderStatus AS outerOrderStatus,
      order1688.productStatus,
      order1688.afterSaleStatusStr
    FROM
      pddAdUnit
    INNER JOIN itemOrder
      ON pddAdUnit.goodsId = itemOrder.productId
    LEFT JOIN order1688
      ON itemOrder.outerOrderId = order1688.orderId
    WHERE
      pddAdUnit.adId = ?
  `, id)
  if err != nil {
    log.Println("order-list-by-ad-id-query-error: ", err)
  }
  defer rows.Close()
  var orderList []interface{}
  for rows.Next() {
    var (
      orderId string
      orderStatus string
      orderStatusStr string
      userPaidAmount float64
      platformDiscount float64
      afterSaleStatus sql.NullInt32
      paymentTime string
      actualPayment sql.NullFloat64
      outerOrderStatus sql.NullInt32
      productStatus sql.NullString
      afterSaleStatusStr sql.NullString
    )
    if err := rows.Scan(
      &orderId,
      &orderStatus,
      &orderStatusStr,
      &userPaidAmount,
      &platformDiscount,
      &afterSaleStatus,
      &paymentTime,
      &actualPayment,
      &outerOrderStatus,
      &productStatus,
      &afterSaleStatusStr,
    ); err != nil {
      log.Println("order-list-by-ad-id-scan-error: ", err)
      log.Println("id: ", id)
    }
    order := map[string]interface{}{
      "orderId": orderId,
      "orderStatus": orderStatus,
      "orderStatusStr": orderStatusStr,
      "userPaidAmount": userPaidAmount,
      "platformDiscount": platformDiscount,
      "afterSaleStatus": afterSaleStatus.Int32,
      "paymentTime": paymentTime,
      "actualPayment": actualPayment.Float64,
      "outerOrderStatus": outerOrderStatus.Int32,
      "productStatus": productStatus.String,
      "afterSaleStatusStr": afterSaleStatusStr.String,
    }
    orderList = append(orderList, order)
  }
  json.NewEncoder(w).Encode(orderList)
}
