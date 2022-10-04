/*
 * Maintained by jemo from 2020.5.20 to now
 * Created by jemo on 2020.5.20 11:01:53
 * Order List
 */

package handle

import (
  "log"
  "time"
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func OrderList(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  db := database.DB
  rows, err := db.Query(`
    SELECT
      itemOrder.id,
      itemOrder.mallId,
      itemOrder.productName,
      itemOrder.orderId,
      itemOrder.outerOrderId,
      itemOrder.orderStatus,
      itemOrder.orderStatusStr,
      itemOrder.productTotalPrice,
      itemOrder.storeDiscount,
      itemOrder.platformDiscount,
      itemOrder.postage,
      itemOrder.userPaidAmount,
      itemOrder.numberOfProducts,
      itemOrder.receiver,
      itemOrder.phone,
      itemOrder.province,
      itemOrder.city,
      itemOrder.district,
      itemOrder.street,
      itemOrder.paymentTime,
      itemOrder.joinSuccessTime,
      itemOrder.orderConfirmationTime,
      itemOrder.commitmentDeliveryTime,
      itemOrder.deliveryTime,
      itemOrder.confirmDeliveryTime,
      itemOrder.productId,
      itemOrder.productSku,
      itemOrder.skuId,
      itemOrder.trackingNumber,
      itemOrder.courierCompany,
      itemOrder.merchantNotes,
      itemOrder.afterSaleStatus,
      itemOrder.buyerMessage,
      item.detailUrl,
      item.originalId,
      order1688.orderStatus AS outerOrderStatus,
      order1688.actualPayment,
      order1688.productStatus,
      order1688.afterSaleStatusStr,
      order1688.receiver AS shippingName,
      order1688.shippingAddress,
      order1688.phone AS shippingPhone
    FROM itemOrder
    LEFT JOIN pddItem
      ON itemOrder.productId = pddItem.pddId
    LEFT JOIN order1688
      ON itemOrder.outerOrderId = order1688.orderId
    LEFT JOIN item
      ON pddItem.outGoodsSn = item.searchId
    WHERE
      (item.forSell = TRUE
      OR
      item.forSell IS NULL)
      and itemOrder.paymentTime > SUBDATE(CURDATE(),INTERVAL 1 month)
    ORDER BY itemOrder.paymentTime DESC
  `)
  if err != nil {
    log.Println("order-list-query-error: ", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  defer rows.Close()
  var orderList []interface{}
  for rows.Next() {
    var (
      id int64
      mallId sql.NullInt64
      productName string
      orderId string
      outerOrderId sql.NullString
      orderStatus sql.NullString
      orderStatusStr string
      productTotalPrice float64
      storeDiscount float64
      platformDiscount float64
      postage float64
      userPaidAmount float64
      numberOfProducts int64
      receiver sql.NullString
      phone sql.NullString
      province sql.NullString
      city sql.NullString
      district sql.NullString
      street sql.NullString
      paymentTime sql.NullString
      joinSuccessTime sql.NullString
      orderConfirmationTime sql.NullString
      commitmentDeliveryTime sql.NullString
      deliveryTime sql.NullString
      confirmDeliveryTime sql.NullString
      productId string
      productSku sql.NullString
      skuId sql.NullString
      trackingNumber sql.NullString
      courierCompany sql.NullString
      merchantNotes sql.NullString
      afterSaleStatus sql.NullInt32
      buyerMessage sql.NullString
      detailUrl sql.NullString
      originalId sql.NullString
      outerOrderStatus sql.NullInt32
      actualPayment sql.NullFloat64
      productStatus sql.NullString
      afterSaleStatusStr sql.NullString
      shippingName sql.NullString
      shippingAddress sql.NullString
      shippingPhone sql.NullString
    )
    if err := rows.Scan(
      &id,
      &mallId,
      &productName,
      &orderId,
      &outerOrderId,
      &orderStatus,
      &orderStatusStr,
      &productTotalPrice,
      &storeDiscount,
      &platformDiscount,
      &postage,
      &userPaidAmount,
      &numberOfProducts,
      &receiver,
      &phone,
      &province,
      &city,
      &district,
      &street,
      &paymentTime,
      &joinSuccessTime,
      &orderConfirmationTime,
      &commitmentDeliveryTime,
      &deliveryTime,
      &confirmDeliveryTime,
      &productId,
      &productSku,
      &skuId,
      &trackingNumber,
      &courierCompany,
      &merchantNotes,
      &afterSaleStatus,
      &buyerMessage,
      &detailUrl,
      &originalId,
      &outerOrderStatus,
      &actualPayment,
      &productStatus,
      &afterSaleStatusStr,
      &shippingName,
      &shippingAddress,
      &shippingPhone,
    ); err != nil {
      log.Println("order-list-scan-error: ", err)
      w.WriteHeader(http.StatusInternalServerError)
      return
    }
    order := map[string]interface{}{
      "id": id,
      "mallId": mallId.Int64,
      "productName": productName,
      "orderId": orderId,
      "outerOrderId": outerOrderId.String,
      "orderStatus": orderStatus.String,
      "orderStatusStr": orderStatusStr,
      "productTotalPrice": productTotalPrice,
      "storeDiscount": storeDiscount,
      "platformDiscount": platformDiscount,
      "postage": postage,
      "userPaidAmount": userPaidAmount,
      "numberOfProducts": numberOfProducts,
      "receiver": receiver.String,
      "phone": phone.String,
      "province": province.String,
      "city": city.String,
      "district": district.String,
      "street": street.String,
      "paymentTime": paymentTime.String,
      "joinSuccessTime": joinSuccessTime.String,
      "orderConfirmationTime": orderConfirmationTime.String,
      "commitmentDeliveryTime": commitmentDeliveryTime.String,
      "deliveryTime": deliveryTime.String,
      "confirmDeliveryTime": confirmDeliveryTime.String,
      "productId": productId,
      "productSku": productSku.String,
      "skuId": skuId.String,
      "trackingNumber": trackingNumber.String,
      "courierCompany": courierCompany.String,
      "merchantNotes": merchantNotes.String,
      "afterSaleStatus": afterSaleStatus.Int32,
      "buyerMessage": buyerMessage.String,
      "detailUrl": detailUrl.String,
      "originalId": originalId.String,
      "outerOrderStatus": outerOrderStatus.Int32,
      "actualPayment": actualPayment.Float64,
      "productStatus": productStatus.String,
      "afterSaleStatusStr": afterSaleStatusStr.String,
      "shippingName": shippingName.String,
      "shippingAddress": shippingAddress.String,
      "shippingPhone": shippingPhone.String,
    }
    orderList = append(orderList, order)
  }
  now := time.Now()
  diff := now.Sub(start)
  log.Println("time: ", diff)
  json.NewEncoder(w).Encode(orderList)
}
