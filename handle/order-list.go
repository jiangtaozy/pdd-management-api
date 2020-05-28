/*
 * Maintained by jemo from 2020.5.20 to now
 * Created by jemo on 2020.5.20 11:01:53
 * Order List
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func OrderList(w http.ResponseWriter, r *http.Request) {
  db := database.ConnectDB()
  rows, err := db.Query("SELECT itemOrder.id, itemOrder.productName, itemOrder.orderId, itemOrder.orderStatus, itemOrder.productTotalPrice, itemOrder.storeDiscount, itemOrder.platformDiscount, itemOrder.postage, itemOrder.userPaidAmount, itemOrder.merchantReceivedAmount, itemOrder.numberOfProducts, itemOrder.receiver, itemOrder.phone, itemOrder.whetherUnderReview, itemOrder.province, itemOrder.city, itemOrder.district, itemOrder.street, itemOrder.paymentTime, itemOrder.joinSuccessTime, itemOrder.orderConfirmationTime, itemOrder.commitmentDeliveryTime, itemOrder.deliveryTime, itemOrder.confirmDeliveryTime, itemOrder.productId, itemOrder.productSku, itemOrder.skuId, itemOrder.trackingNumber, itemOrder.courierCompany, itemOrder.merchantNotes, itemOrder.afterSaleStatus, itemOrder.buyerMessage, item.detailUrl FROM itemOrder LEFT JOIN pddItem ON itemOrder.productId = pddItem.pddId LEFT JOIN item ON pddItem.outGoodsSn = item.searchId where item.forSell = TRUE OR item.forSell IS NULL")
  if err != nil {
    log.Println("order-list-query-error: ", err)
  }
  defer rows.Close()
  var orderList []interface{}
  for rows.Next() {
    var (
      id int64
      productName string
      orderId string
      orderStatus string
      productTotalPrice float64
      storeDiscount float64
      platformDiscount float64
      postage float64
      userPaidAmount float64
      merchantReceivedAmount float64
      numberOfProducts int64
      receiver string
      phone string
      whetherUnderReview string
      province string
      city string
      district string
      street string
      paymentTime string
      joinSuccessTime string
      orderConfirmationTime string
      commitmentDeliveryTime string
      deliveryTime sql.NullString
      confirmDeliveryTime sql.NullString
      productId string
      productSku string
      skuId string
      trackingNumber string
      courierCompany string
      merchantNotes string
      afterSaleStatus string
      buyerMessage string
      detailUrl sql.NullString
    )
    if err := rows.Scan(&id, &productName, &orderId, &orderStatus, &productTotalPrice, &storeDiscount, &platformDiscount, &postage, &userPaidAmount, &merchantReceivedAmount, &numberOfProducts, &receiver, &phone, &whetherUnderReview, &province, &city, &district, &street, &paymentTime, &joinSuccessTime, &orderConfirmationTime, &commitmentDeliveryTime, &deliveryTime, &confirmDeliveryTime, &productId, &productSku, &skuId, &trackingNumber, &courierCompany, &merchantNotes, &afterSaleStatus, &buyerMessage, &detailUrl); err != nil {
      log.Println("order-list-scan-error: ", err)
    }
    order := map[string]interface{}{
      "id": id,
      "productName": productName,
      "orderId": orderId,
      "orderStatus": orderStatus,
      "productTotalPrice": productTotalPrice,
      "storeDiscount": storeDiscount,
      "platformDiscount": platformDiscount,
      "postage": postage,
      "userPaidAmount": userPaidAmount,
      "merchantReceivedAmount": merchantReceivedAmount,
      "numberOfProducts": numberOfProducts,
      "receiver": receiver,
      "phone": phone,
      "whetherUnderReview": whetherUnderReview,
      "province": province,
      "city": city,
      "district": district,
      "street": street,
      "paymentTime": paymentTime,
      "joinSuccessTime": joinSuccessTime,
      "orderConfirmationTime": orderConfirmationTime,
      "commitmentDeliveryTime": commitmentDeliveryTime,
      "deliveryTime": deliveryTime.String,
      "confirmDeliveryTime": confirmDeliveryTime.String,
      "productId": productId,
      "productSku": productSku,
      "skuId": skuId,
      "trackingNumber": trackingNumber,
      "courierCompany": courierCompany,
      "merchantNotes": merchantNotes,
      "afterSaleStatus": afterSaleStatus,
      "buyerMessage": buyerMessage,
      "detailUrl": detailUrl.String,
    }
    orderList = append(orderList, order)
  }
  json.NewEncoder(w).Encode(orderList)
}
