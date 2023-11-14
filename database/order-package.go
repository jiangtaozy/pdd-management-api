/*
 * Maintained by jemo from 2023.8.14 to now
 * Created by jemo on 2023.8.14 20:24:21
 * 发货单包裹列表
 */
package database

const orderPackage = `
  CREATE TABLE IF NOT EXISTS orderPackage (
    productSkuId VARCHAR(20) COMMENT 'sku id',
    skuNum INTEGER COMMENT '包裹数',
    packageSn VARCHAR(20) COMMENT '包裹号',
    subPurchaseOrderSn VARCHAR(20) COMMENT '备货单号',
    deliveryOrderSn VARCHAR(20) COMMENT '发货单号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(packageSn)
  );
 `
