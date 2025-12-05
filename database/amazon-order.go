/*
 * Maintained by jemo from 2025.11.18 to now
 * Created by jemo on 2025.11.18 17:44:11
 * 亚马逊订单
 */

package database

const createAmazonOrder =`
  CREATE TABLE IF NOT EXISTS amazonOrder (
    amazonOrderId VARCHAR(30) NOT NULL PRIMARY KEY COMMENT '订单号',
    homeMarketplaceId VARCHAR(20),
    orderDate DATETIME COMMENT '下单时间',
    orderFulfillmentStatus VARCHAR(20) COMMENT '订单状态',
    salesChannel VARCHAR(30)
  );
`
