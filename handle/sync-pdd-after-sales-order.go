/*
 * Maintained by jemo from 2021.7.13 to now
 * Created by jemo on 2021.7.13 11:26:06
 * Sync Pdd After Sales Order
 * 同步拼多多售后订单
 */

package handle

import (
  "log"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncPddAfterSalesOrder(requestBody map[string]interface{}, responseBody map[string]interface{}) {
  result := responseBody["result"].(map[string]interface{})
  list := result["list"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddAfterSalesOrder (
      afterSalesReasonCode,
      afterSalesReasonDesc,
      afterSalesStatus,
      afterSalesTitle,
      afterSalesType,
      assignedProcessorId,
      assignedProcessorName,
      createdAt,
      expireActionDesc,
      expireRemainTime,
      goodsName,
      goodsNumber,
      goodsSpec,
      id,
      mallRemark,
      mallRemarkTag,
      mallRemarkTagName,
      orderAmount,
      orderSn,
      receiveAmount,
      refundAmount,
      rejectChatTipDesc,
      rejectChatTipExpireRemainTime,
      remarkStatus,
      returnCouponAmount,
      reverseShippingId,
      reverseTrackingNumber,
      sellerAfterSalesShippingStatus,
      sellerAfterSalesShippingStatusDesc,
      thumbUrl,
      ticketId,
      uid,
      version
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) 
  `)
  if err != nil {
    log.Println("sync-pdd-after-sales-order-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      pddAfterSalesOrder
    SET
      afterSalesReasonCode = ?,
      afterSalesReasonDesc = ?,
      afterSalesStatus = ?,
      afterSalesTitle = ?,
      afterSalesType = ?,
      assignedProcessorId = ?,
      assignedProcessorName = ?,
      createdAt = ?,
      expireActionDesc = ?,
      expireRemainTime = ?,
      goodsName = ?,
      goodsNumber = ?,
      goodsSpec = ?,
      mallRemark = ?,
      mallRemarkTag = ?,
      mallRemarkTagName = ?,
      orderAmount = ?,
      orderSn = ?,
      receiveAmount = ?,
      refundAmount = ?,
      rejectChatTipDesc = ?,
      rejectChatTipExpireRemainTime = ?,
      remarkStatus = ?,
      returnCouponAmount = ?,
      reverseShippingId = ?,
      reverseTrackingNumber = ?,
      sellerAfterSalesShippingStatus = ?,
      sellerAfterSalesShippingStatusDesc = ?,
      thumbUrl = ?,
      ticketId = ?,
      uid = ?,
      version = ?
    WHERE
      id = ?
  `)
  if err != nil {
    log.Println("sync-pdd-after-sales-order-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    data := list[i].(map[string]interface{})
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        pddAfterSalesOrder
      WHERE
        id = ?
    `, data["id"]).Scan(&count)
    if err != nil {
      log.Println("sync-pdd-after-sales-order-count-error: ", err)
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        data["afterSalesReasonCode"],
        data["afterSalesReasonDesc"],
        data["afterSalesStatus"],
        data["afterSalesTitle"],
        data["afterSalesType"],
        data["assignedProcessorId"],
        data["assignedProcessorName"],
        data["createdAt"],
        data["expireActionDesc"],
        data["expireRemainTime"],
        data["goodsName"],
        data["goodsNumber"],
        data["goodsSpec"],
        data["id"],
        data["mallRemark"],
        data["mallRemarkTag"],
        data["mallRemarkTagName"],
        data["orderAmount"],
        data["orderSn"],
        data["receiveAmount"],
        data["refundAmount"],
        data["rejectChatTipDesc"],
        data["rejectChatTipExpireRemainTime"],
        data["remarkStatus"],
        data["returnCouponAmount"],
        data["reverseShippingId"],
        data["reverseTrackingNumber"],
        data["sellerAfterSalesShippingStatus"],
        data["sellerAfterSalesShippingStatusDesc"],
        data["thumbUrl"],
        data["ticketId"],
        data["uid"],
        data["version"],
      )
      if err != nil {
        log.Println("sync-pdd-after-sales-order-insert-exec-error: ", err)
        log.Println("data: ", data)
      }
    } else {
      _, err = stmtUpdate.Exec(
        data["afterSalesReasonCode"],
        data["afterSalesReasonDesc"],
        data["afterSalesStatus"],
        data["afterSalesTitle"],
        data["afterSalesType"],
        data["assignedProcessorId"],
        data["assignedProcessorName"],
        data["createdAt"],
        data["expireActionDesc"],
        data["expireRemainTime"],
        data["goodsName"],
        data["goodsNumber"],
        data["goodsSpec"],
        data["mallRemark"],
        data["mallRemarkTag"],
        data["mallRemarkTagName"],
        data["orderAmount"],
        data["orderSn"],
        data["receiveAmount"],
        data["refundAmount"],
        data["rejectChatTipDesc"],
        data["rejectChatTipExpireRemainTime"],
        data["remarkStatus"],
        data["returnCouponAmount"],
        data["reverseShippingId"],
        data["reverseTrackingNumber"],
        data["sellerAfterSalesShippingStatus"],
        data["sellerAfterSalesShippingStatusDesc"],
        data["thumbUrl"],
        data["ticketId"],
        data["uid"],
        data["version"],
        data["id"],
      )
      if err != nil {
        log.Println("sync-pdd-after-sales-order-update-exec-error: ", err)
        log.Println("data: ", data)
      }
    }
  }
}
