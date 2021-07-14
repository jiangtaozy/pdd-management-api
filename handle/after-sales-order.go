/*
 * Maintained by jemo from 2021.7.13 to now
 * Created by jemo on 2021.7.13 15:50:59
 * After Sales Order
 * 售后订单
 */

package handle

import (
  "log"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AfterSalesOrder(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      pddAfterSalesOrder.afterSalesReasonDesc,
      pddAfterSalesOrder.afterSalesStatus,
      pddAfterSalesOrder.afterSalesTitle,
      pddAfterSalesOrder.afterSalesType,
      pddAfterSalesOrder.orderSn,
      pddAfterSalesOrder.sellerAfterSalesShippingStatus,
      pddAfterSalesOrder.sellerAfterSalesShippingStatusDesc,
      pddAfterSalesOrder.createdAt,
      itemOrder.outerOrderId,
      order1688.afterSaleStatusStr,
      order1688.orderStatusStr
    FROM
      pddAfterSalesOrder
    LEFT JOIN itemOrder
      ON pddAfterSalesOrder.orderSn = itemOrder.orderId
    LEFT JOIN order1688
      ON itemOrder.outerOrderId = order1688.orderId
    ORDER BY
      createdAt DESC
  `)
  if err != nil {
    log.Println("after-sales-order-query-error: ", err)
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      afterSalesReasonDesc sql.NullString
      afterSalesStatus sql.NullInt64
      afterSalesTitle sql.NullString
      afterSalesType sql.NullInt64
      orderSn string
      sellerAfterSalesShippingStatus sql.NullInt64
      sellerAfterSalesShippingStatusDesc sql.NullString
      createdAt sql.NullInt64
      outerOrderId sql.NullString
      womenAfterSaleStatusStr sql.NullString
      orderStatusStr sql.NullString
    )
    err := rows.Scan(
      &afterSalesReasonDesc,
      &afterSalesStatus,
      &afterSalesTitle,
      &afterSalesType,
      &orderSn,
      &sellerAfterSalesShippingStatus,
      &sellerAfterSalesShippingStatusDesc,
      &createdAt,
      &outerOrderId,
      &womenAfterSaleStatusStr,
      &orderStatusStr,
    )
    if err != nil {
      log.Println("after-sales-order-scan-error: ", err)
    }
    order := map[string]interface{}{
      "afterSalesReasonDesc": afterSalesReasonDesc.String,
      "afterSalesStatus": afterSalesStatus.Int64,
      "afterSalesTitle": afterSalesTitle.String,
      "afterSalesType": afterSalesType.Int64,
      "orderSn": orderSn,
      "sellerAfterSalesShippingStatus": sellerAfterSalesShippingStatus.Int64,
      "sellerAfterSalesShippingStatusDesc": sellerAfterSalesShippingStatusDesc.String,
      "createdAt": createdAt.Int64,
      "outerOrderId": outerOrderId.String,
      "womenAfterSaleStatusStr": womenAfterSaleStatusStr.String,
      "orderStatusStr": orderStatusStr.String,
    }
    list = append(list, order)
  }
  json.NewEncoder(w).Encode(list);
}
