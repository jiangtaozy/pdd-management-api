/*
 * Maintained by jemo from 2025.08.25 to now
 * Created by jemo on 2025.08.25 16:48:27
 * 拼多多未上架
 */

package database

const createPddNotShelf =`
  CREATE TABLE IF NOT EXISTS pddNotShelf (
    skuId VARCHAR(20) COMMENT 'sku ID',
    supplierId VARCHAR(20) COMMENT '供应商ID',
    skuExtCode VARCHAR(20) NOT NULL PRIMARY KEY COMMENT 'sku货号',
    displayImage VARCHAR(200) COMMENT '图像',
    productSkuSpecList VARCHAR(200) COMMENT 'sku spec',
    inventoryQuantity INTEGER COMMENT '库存数量',
    isOnShelf BOOLEAN DEFAULT 0 COMMENT '是否已上架',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
