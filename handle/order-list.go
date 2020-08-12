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
  db := database.DB
  rows, err := db.Query("SELECT itemOrder.id, itemOrder.mallId, itemOrder.productName, itemOrder.orderId, itemOrder.orderStatus, itemOrder.orderStatusStr, itemOrder.productTotalPrice, itemOrder.storeDiscount, itemOrder.platformDiscount, itemOrder.postage, itemOrder.userPaidAmount, itemOrder.numberOfProducts, itemOrder.receiver, itemOrder.phone, itemOrder.province, itemOrder.city, itemOrder.district, itemOrder.street, itemOrder.paymentTime, itemOrder.joinSuccessTime, itemOrder.orderConfirmationTime, itemOrder.commitmentDeliveryTime, itemOrder.deliveryTime, itemOrder.confirmDeliveryTime, itemOrder.productId, itemOrder.productSku, itemOrder.skuId, itemOrder.trackingNumber, itemOrder.courierCompany, itemOrder.merchantNotes, itemOrder.afterSaleStatus, itemOrder.buyerMessage, item.detailUrl FROM itemOrder LEFT JOIN pddItem ON itemOrder.productId = pddItem.pddId LEFT JOIN item ON pddItem.outGoodsSn = item.searchId where item.forSell = TRUE OR item.forSell IS NULL ORDER BY itemOrder.paymentTime DESC")
  if err != nil {
    log.Println("order-list-query-error: ", err)
  }
  defer rows.Close()
  var orderList []interface{}
  for rows.Next() {
    var (
      id int64
      mallId int64
      productName string
      orderId string
      orderStatus string
      orderStatusStr string
      productTotalPrice float64
      storeDiscount float64
      platformDiscount float64
      postage float64
      userPaidAmount float64
      numberOfProducts int64
      receiver string
      phone sql.NullString
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
      afterSaleStatus sql.NullInt32
      buyerMessage string
      detailUrl sql.NullString
    )
    if err := rows.Scan(&id, &mallId, &productName, &orderId, &orderStatus, &orderStatusStr, &productTotalPrice, &storeDiscount, &platformDiscount, &postage, &userPaidAmount, &numberOfProducts, &receiver, &phone, &province, &city, &district, &street, &paymentTime, &joinSuccessTime, &orderConfirmationTime, &commitmentDeliveryTime, &deliveryTime, &confirmDeliveryTime, &productId, &productSku, &skuId, &trackingNumber, &courierCompany, &merchantNotes, &afterSaleStatus, &buyerMessage, &detailUrl); err != nil {
      log.Println("order-list-scan-error: ", err)
    }
    order := map[string]interface{}{
      "id": id,
      "mallId": mallId,
      "productName": productName,
      "orderId": orderId,
      "orderStatus": orderStatus,
      "orderStatusStr": orderStatusStr,
      "productTotalPrice": productTotalPrice,
      "storeDiscount": storeDiscount,
      "platformDiscount": platformDiscount,
      "postage": postage,
      "userPaidAmount": userPaidAmount,
      "numberOfProducts": numberOfProducts,
      "receiver": receiver,
      "phone": phone.String,
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
      "afterSaleStatus": afterSaleStatus.Int32,
      "buyerMessage": buyerMessage,
      "detailUrl": detailUrl.String,
    }
    orderList = append(orderList, order)
  }
  json.NewEncoder(w).Encode(orderList)
}
