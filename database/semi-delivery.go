/*
 * Maintained by jemo from 2025.08.04 to now
 * Created by jemo on 2025.08.04 17:03:15
 * temu半托管发货单
 */

package database

const createSemiDelivery =`
  CREATE TABLE IF NOT EXISTS semiDelivery (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '发货单编号',
    skuId VARCHAR(20),
    sku VARCHAR(10) COMMENT 'sku货号',
    quantity INTEGER COMMENT '发货数量',
    isShip TINYINT UNSIGNED COMMENT '是否发货',
    isStockout TINYINT UNSIGNED COMMENT '是否缺货，不发货',
    createdAtTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAtTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
