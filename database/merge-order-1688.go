/*
 * Maintained by jemo from 2023.4.5 to now
 * Created by jemo on 2023.4.5 11:51:25
 * merge order 1688
 */

package database

const mergeOrder1688 = `
  CREATE TABLE IF NOT EXISTS mergeOrder1688 (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    orderId VARCHAR(20) NOT NULL COMMENT '1688订单号',
    productTitle VARCHAR(150) COMMENT '货品标题',
    price DECIMAL(10,2) COMMENT '单价(元)',
    amount INTEGER UNSIGNED COMMENT '数量',
    skuId VARCHAR(30) COMMENT 'sku id',
    offerId VARCHAR(30) COMMENT 'offer id',
    unique orderIdSkuIdUnique(orderId, skuId),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
