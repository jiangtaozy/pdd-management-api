/*
 * Maintained by jemo from 2020.12.4 to now
 * Created by jemo on 2020.12.4 17:36:34
 * Douyin Order List
 */

package handle

import (
  "log"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func DyOrderList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      dyChildOrder.orderId,
      dyChildOrder.shopId,
      dyChildOrder.openId,
      dyChildOrder.orderStatus,
      dyChildOrder.orderType,
      dyChildOrder.orderTag,
      dyChildOrder.postAddrProvinceId,
      dyChildOrder.postAddrProvinceName,
      dyChildOrder.postAddrCityId,
      dyChildOrder.postAddrCityName,
      dyChildOrder.postAddrTownId,
      dyChildOrder.postAddrTownName,
      dyChildOrder.postAddrDetail,
      dyChildOrder.postCode,
      dyChildOrder.postReceiver,
      dyChildOrder.postTel,
      dyChildOrder.buyerWords,
      dyChildOrder.sellerWords,
      dyChildOrder.logisticsId,
      dyChildOrder.logisticsCode,
      dyChildOrder.logisticsTime,
      dyChildOrder.receiptTime,
      dyChildOrder.createTime,
      dyChildOrder.updateTime,
      dyChildOrder.expShipTime,
      dyChildOrder.cancelReason,
      dyChildOrder.payType,
      dyChildOrder.payTime,
      dyChildOrder.postAmount,
      dyChildOrder.couponAmount,
      dyChildOrder.shopCouponAmount,
      dyChildOrder.couponMetaId,
      dyChildOrder.totalAmount,
      dyChildOrder.isComment,
      dyChildOrder.urgeCnt,
      dyChildOrder.bType,
      dyChildOrder.subBType,
      dyChildOrder.isInsurance,
      dyChildOrder.cType,
      dyChildOrder.cosRatio,
      dyChildOrder.userName,
      dyChildOrder.finalStatus,
      dyChildOrder.shippedNum,
      dyChildOrder.code,
      dyChildOrder.comboAmount,
      dyChildOrder.comboId,
      dyChildOrder.comboNum,
      dyChildOrder.itemIds,
      dyChildOrder.outProductId,
      dyChildOrder.outSkuId,
      dyChildOrder.pid,
      dyChildOrder.platformFullAmount,
      dyChildOrder.productId,
      dyChildOrder.productName,
      dyChildOrder.productPic,
      dyChildOrder.specDesc,
      item.detailUrl
    FROM
      dyChildOrder
    LEFT JOIN
      item
    ON
      dyChildOrder.outProductId = item.womenProductId
  `)
  if err != nil {
    log.Println("dy-order-list-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      orderId string
      shopId sql.NullInt64
      openId sql.NullString
      orderStatus sql.NullInt64
      orderType sql.NullInt64
      orderTag sql.NullString
      postAddrProvinceId sql.NullString
      postAddrProvinceName sql.NullString
      postAddrCityId sql.NullString
      postAddrCityName sql.NullString
      postAddrTownId sql.NullString
      postAddrTownName sql.NullString
      postAddrDetail sql.NullString
      postCode sql.NullString
      postReceiver sql.NullString
      postTel sql.NullString
      buyerWords sql.NullString
      sellerWords sql.NullString
      logisticsId sql.NullInt64
      logisticsCode sql.NullString
      logisticsTime sql.NullString
      receiptTime sql.NullString
      createTime sql.NullString
      updateTime sql.NullString
      expShipTime sql.NullString
      cancelReason sql.NullString
      payType sql.NullInt64
      payTime sql.NullString
      postAmount sql.NullInt64
      couponAmount sql.NullInt64
      shopCouponAmount sql.NullInt64
      couponMetaId sql.NullString
      totalAmount sql.NullInt64
      isComment sql.NullBool
      urgeCnt sql.NullInt64
      bType  sql.NullInt64
      subBType sql.NullInt64
      isInsurance sql.NullBool
      cType sql.NullInt64
      cosRatio sql.NullString
      userName sql.NullString
      finalStatus sql.NullInt64
      shippedNum sql.NullInt64
      code sql.NullString
      comboAmount sql.NullInt64
      comboId sql.NullInt64
      comboNum sql.NullInt64
      itemIds sql.NullString
      outProductId sql.NullInt64
      outSkuId sql.NullInt64
      pid sql.NullString
      platformFullAmount sql.NullInt64
      productId sql.NullString
      productName sql.NullString
      productPic sql.NullString
      specDesc sql.NullString
      detailUrl sql.NullString
    )
    err = rows.Scan(
      &orderId,
      &shopId,
      &openId,
      &orderStatus,
      &orderType,
      &orderTag,
      &postAddrProvinceId,
      &postAddrProvinceName,
      &postAddrCityId,
      &postAddrCityName,
      &postAddrTownId,
      &postAddrTownName,
      &postAddrDetail,
      &postCode,
      &postReceiver,
      &postTel,
      &buyerWords,
      &sellerWords,
      &logisticsId,
      &logisticsCode,
      &logisticsTime,
      &receiptTime,
      &createTime,
      &updateTime,
      &expShipTime,
      &cancelReason,
      &payType,
      &payTime,
      &postAmount,
      &couponAmount,
      &shopCouponAmount,
      &couponMetaId,
      &totalAmount,
      &isComment,
      &urgeCnt,
      &bType,
      &subBType,
      &isInsurance,
      &cType,
      &cosRatio,
      &userName,
      &finalStatus,
      &shippedNum,
      &code,
      &comboAmount,
      &comboId,
      &comboNum,
      &itemIds,
      &outProductId,
      &outSkuId,
      &pid,
      &platformFullAmount,
      &productId,
      &productName,
      &productPic,
      &specDesc,
      &detailUrl,
    )
    if err != nil {
      log.Println("dy-order-list-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    order := map[string]interface{}{
      "orderId": orderId,
      "shopId": shopId.Int64,
      "openId": openId.String,
      "orderStatus": orderStatus.Int64,
      "orderType": orderType.Int64,
      "orderTag": orderTag.String,
      "postAddrProvinceId": postAddrProvinceId.String,
      "postAddrProvinceName": postAddrProvinceName.String,
      "postAddrCityId": postAddrCityId.String,
      "postAddrCityName": postAddrCityName.String,
      "postAddrTownId": postAddrTownId.String,
      "postAddrTownName": postAddrTownName.String,
      "postAddrDetail": postAddrDetail.String,
      "postCode": postCode.String,
      "postReceiver": postReceiver.String,
      "postTel": postTel.String,
      "buyerWords": buyerWords.String,
      "sellerWords": sellerWords.String,
      "logisticsId": logisticsId.Int64,
      "logisticsCode": logisticsCode.String,
      "logisticsTime": logisticsTime.String,
      "receiptTime": receiptTime.String,
      "createTime": createTime.String,
      "updateTime": updateTime.String,
      "expShipTime": expShipTime.String,
      "cancelReason": cancelReason.String,
      "payType": payType.Int64,
      "payTime": payTime.String,
      "postAmount": postAmount.Int64,
      "couponAmount": couponAmount.Int64,
      "shopCouponAmount": shopCouponAmount.Int64,
      "couponMetaId": couponMetaId.String,
      "totalAmount": totalAmount.Int64,
      "isComment": isComment.Bool,
      "urgeCnt": urgeCnt.Int64,
      "bType": bType.Int64,
      "subBType": subBType.Int64,
      "isInsurance": isInsurance.Bool,
      "cType": cType.Int64,
      "cosRatio": cosRatio.String,
      "userName": userName.String,
      "finalStatus": finalStatus.Int64,
      "shippedNum": shippedNum.Int64,
      "code": code.String,
      "comboAmount": comboAmount.Int64,
      "comboId": comboId.Int64,
      "comboNum": comboNum.Int64,
      "itemIds": itemIds.String,
      "outProductId": outProductId.Int64,
      "outSkuId": outSkuId.Int64,
      "pid": pid.String,
      "platformFullAmount": platformFullAmount.Int64,
      "productId": productId.String,
      "productName": productName.String,
      "productPic": productPic.String,
      "specDesc": specDesc.String,
      "detailUrl": detailUrl.String,
    }
    list = append(list, order)
  }
  json.NewEncoder(w).Encode(list)
}
