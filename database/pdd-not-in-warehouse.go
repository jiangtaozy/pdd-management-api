/*
 * Maintained by jemo from 2025.09.02 to now
 * Created by jemo on 2025.09.02 15:26:54
 * 拼多多未入库
 */

package database

const createPddNotInWarehouse =`
  CREATE TABLE IF NOT EXISTS pddNotInWarehouse (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    pddId VARCHAR(20),
    skuId VARCHAR(20),
    outGoodsSn VARCHAR(20) COMMENT 'skc货号',
    outSkuSn VARCHAR(20) COMMENT 'sku货号',
    skuQuantity INTEGER UNSIGNED COMMENT '拼多多库存',
    inventoryQuantity INTEGER COMMENT '系统库存',
    spec VARCHAR(200) COMMENT '属性',
    skuThumbUrl VARCHAR(200) COMMENT '图片地址',
    goodsName VARCHAR(100),
    isInWarehouse BOOLEAN DEFAULT 0 COMMENT '是否已入库',
    realQuantity INTEGER UNSIGNED COMMENT '实际库存',
    isRealQuantityZero BOOLEAN DEFAULT 0 COMMENT '是否实际库存为0',
    isCompleted BOOLEAN DEFAULT 0 COMMENT '是否已处理完成',
    createdAtTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAtTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
