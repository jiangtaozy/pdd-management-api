/*
 * Maintained by jemo from 2023.6.27 to now
 * Created by jemo on 2023.6.27 10:50:04
 * 退货包裹
 */
package database

const returnPackage = `
  CREATE TABLE IF NOT EXISTS returnPackage (
    returnSupplierPackageNo VARCHAR(40) COMMENT '退货包裹号',
    purchaseSubOrderSn VARCHAR(20) COMMENT '采购子单号',
    productSpuId VARCHAR(20) COMMENT 'spu id',
    productSkcId VARCHAR(20) COMMENT 'skc id',
    productSkuId VARCHAR(20) COMMENT 'sku ID',
    skcExtCode VARCHAR(20) COMMENT 'skc货号',
    skuExtCode VARCHAR(20) COMMENT 'sku货号',
    returnSupplierQuantity INTEGER COMMENT '退货数量',
    mainSaleSpec VARCHAR(20) COMMENT '主属性集-为空',
    secondarySaleSpec VARCHAR(40) COMMENT '次属性集',
    orderType INTEGER COMMENT '退货订单类型',
    orderTypeDesc VARCHAR(50) COMMENT '退货订单类型描述',
    reasonDesc VARCHAR(100) COMMENT '退货原因',
    remark VARCHAR(100) COMMENT '备注',
    outboundTime DATETIME COMMENT '出库时间',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(returnSupplierPackageNo, productSkcId, productSkuId)
  );
 `
