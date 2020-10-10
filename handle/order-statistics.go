/*
 * Maintained by jemo from 2020.7.7 to now
 * Created by jemo on 2020.7.7 10:52:06
 * Order statistics
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "database/sql"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func OrderStatistics(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      itemOrder.mallId,
      itemOrder.orderId,
      itemOrder.productTotalPrice,
      itemOrder.storeDiscount,
      itemOrder.platformDiscount,
      itemOrder.userPaidAmount,
      itemOrder.paymentTime,
      itemOrder.afterSaleStatus,
      itemOrder.afterSaleApplyTime,
      itemOrder.orderStatusStr,
      order1688.actualPayment
    FROM
      itemOrder AS itemOrder
    LEFT JOIN
      order1688 AS order1688
    ON
      itemOrder.outerOrderId = order1688.orderId
    WHERE
      itemOrder.orderStatus = 1
      AND order1688.actualPayment IS NOT NULL
  `)
  if err != nil {
    log.Println("order-statistics-query-error: ", err)
  }
  defer rows.Close()
  var orderList []interface{}
  for rows.Next() {
    var (
      mallId sql.NullString
      orderId sql.NullString
      productTotalPrice int64
      storeDiscount int64
      platformDiscount int64
      userPaidAmount int64
      paymentTime string
      afterSaleStatus sql.NullInt64
      afterSaleApplyTime sql.NullString
      orderStatusStr sql.NullString
      actualPayment float64
    )
    err := rows.Scan(
      &mallId,
      &orderId,
      &productTotalPrice,
      &storeDiscount,
      &platformDiscount,
      &userPaidAmount,
      &paymentTime,
      &afterSaleStatus,
      &afterSaleApplyTime,
      &orderStatusStr,
      &actualPayment,
    )
    if err != nil {
      log.Println("order-statistics-scan-error: ", err)
    }
    order := map[string]interface{}{
      "mallId": mallId.String,
      "orderId": orderId.String,
      "productTotalPrice": productTotalPrice,
      "storeDiscount": storeDiscount,
      "platformDiscount": platformDiscount,
      "userPaidAmount": userPaidAmount,
      "paymentTime": paymentTime,
      "afterSaleStatus": afterSaleStatus.Int64,
      "afterSaleApplyTime": afterSaleApplyTime.String,
      "orderStatusStr": orderStatusStr.String,
      "actualPayment": actualPayment,
    }
    orderList = append(orderList, order)
  }
  err = json.NewEncoder(w).Encode(orderList)
  if err != nil {
    log.Println("order-statistics-encode-err: ", err)
  }
}
