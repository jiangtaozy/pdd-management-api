/*
 * Maintained by jemo from 2023.7.4 to now
 * Created by jemo on 2023.7.4 17:00:16
 */

package database

const createWarehouseSkuQuantityRecord =`
  CREATE TABLE IF NOT EXISTS warehouseSkuQuantityRecord (
    warehouseId INTEGER NOT NULL COMMENT '仓库编号',
    skuCode VARCHAR(20) COMMENT 'sku货号',
    skuId VARCHAR(20) NOT NULL COMMENT 'skuId',
    beforeQuantity INTEGER COMMENT '修改前数量',
    afterQuantity INTEGER COMMENT '修改后数量',
    deviationQuantity INTEGER COMMENT '数量差值',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
