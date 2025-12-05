/*
 * Maintained by jemo from 2025.11.18 to now
 * Created by jemo on 2025.11.18 18:11:40
 * 亚马逊退款订单产品
 */

package database

const createAmazonRefundItem =`
  CREATE TABLE IF NOT EXISTS amazonRefundItem (
    orderItemCode VARCHAR(100) NOT NULL PRIMARY KEY COMMENT 'customerOrderItemCode',
    itemQuantity INTEGER,
    amazonOrderId VARCHAR(30),
    refundStatus VARCHAR(30)
  );
`
