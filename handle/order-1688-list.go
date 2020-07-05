/*
 * Maintained by jemo from 2020.7.5 to now
 * Created by jemo on 2020.7.5 15:56:08
 * 1688 Order List
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func Order1688List(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query("SELECT orderId, sellerCompany, totalPrice, shippingFare, discount, actualPayment, orderStatus, orderCreatedTime, orderPaymentTime, receiver, shippingAddress, postcode, phone, productTitle, price, amount, courierCompany, trackingNumber from order1688")
  if err != nil {
    log.Println("order-1688-list-query-error: ", err)
  }
  defer rows.Close()
  var orderList []interface{}
  for rows.Next() {
    var (
      orderId string
      sellerCompany string
      totalPrice float64
      shippingFare float64
      discount float64
      actualPayment float64
      orderStatus string
      orderCreatedTime string
      orderPaymentTime string
      receiver string
      shippingAddress string
      postcode string
      phone string
      productTitle string
      price float64
      amount int64
      courierCompany sql.NullString
      trackingNumber sql.NullString
    )
    err := rows.Scan(&orderId, &sellerCompany, &totalPrice, &shippingFare, &discount, &actualPayment, &orderStatus, &orderCreatedTime, &orderPaymentTime, &receiver, &shippingAddress, &postcode, &phone, &productTitle, &price, &amount, &courierCompany, &trackingNumber)
    if err != nil {
      log.Println("order-1688-list-scan-error: ", err)
    }
    order := map[string]interface{}{
      "orderId": orderId,
      "sellerCompany": sellerCompany,
      "totalPrice": totalPrice,
      "shippingFare": shippingFare,
      "discount": discount,
      "actualPayment": actualPayment,
      "orderStatus": orderStatus,
      "orderCreatedTime": orderCreatedTime,
      "orderPaymentTime": orderPaymentTime,
      "receiver": receiver,
      "shippingAddress": shippingAddress,
      "postcode": postcode,
      "phone": phone,
      "productTitle": productTitle,
      "price": price,
      "amount": amount,
      "courierCompany": courierCompany.String,
      "trackingNumber": trackingNumber.String,
    }
    orderList = append(orderList, order)
  }
  json.NewEncoder(w).Encode(orderList)
}
