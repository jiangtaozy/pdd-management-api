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
    itemOrder.productTotalPrice,
    itemOrder.storeDiscount,
    itemOrder.platformDiscount,
    itemOrder.userPaidAmount,
    itemOrder.paymentTime,
    order1688.actualPayment
    FROM itemOrder AS itemOrder
    LEFT JOIN order1688 AS order1688
    ON itemOrder.outerOrderId = order1688.orderId
    WHERE
    (itemOrder.afterSaleStatus <> 5 OR itemOrder.afterSaleStatus IS NULL)
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
      productTotalPrice int64
      storeDiscount int64
      platformDiscount int64
      userPaidAmount int64
      paymentTime string
      actualPayment float64
    )
    err := rows.Scan(&mallId, &productTotalPrice, &storeDiscount, &platformDiscount, &userPaidAmount, &paymentTime, &actualPayment)
    if err != nil {
      log.Println("order-statistics-scan-error: ", err)
    }
    order := map[string]interface{}{
      "mallId": mallId.String,
      "productTotalPrice": productTotalPrice,
      "storeDiscount": storeDiscount,
      "platformDiscount": platformDiscount,
      "userPaidAmount": userPaidAmount,
      "paymentTime": paymentTime,
      "actualPayment": actualPayment,
    }
    orderList = append(orderList, order)
  }
  err = json.NewEncoder(w).Encode(orderList)
  if err != nil {
    log.Println("order-statistics-encode-err: ", err)
  }
}

