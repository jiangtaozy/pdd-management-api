/*
 * Maintained by jemo from 2023.7.4 to now
 * Created by jemo on 2023.7.4 16:33:24
 */

package database

const createWarehouseSkuQuantity =`
  CREATE TABLE IF NOT EXISTS warehouseSkuQuantity (
    warehouseId INTEGER NOT NULL COMMENT '仓库编号',
    skuCode VARCHAR(20) COMMENT 'sku货号',
    skuId VARCHAR(20) NOT NULL COMMENT 'skuId',
    receivedQuantity INTEGER COMMENT '已收货数量',
    notReceivedQuantity INTEGER COMMENT '未收货数量',
    deliverQuantity INTEGER COMMENT '已发货数量',
    inventoryQuantity INTEGER COMMENT '库存数量',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(warehouseId, skuId)
  );
`
