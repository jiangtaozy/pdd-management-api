/*
 * Maintained by jemo from 2023.7.11 to now
 * Created by jemo on 2023.7.11 11:44:01
 * 发货单列表
 */
package database

const orderDelivery = `
  CREATE TABLE IF NOT EXISTS orderDelivery (
    supplierId VARCHAR(20) COMMENT '店铺id',
    subPurchaseOrderSn VARCHAR(20) COMMENT '备货单号',
    deliveryOrderSn VARCHAR(20) COMMENT '发货单号',
    deliveryMethod INTEGER COMMENT '发货方式，1: 自送, 2: 平台推荐服务商',
    subWarehouseId VARCHAR(20) COMMENT '收货仓库id',
    subWarehouseName VARCHAR(100) COMMENT '收货仓库',
    expressDeliverySn VARCHAR(20) COMMENT '快递单号',
    expressCompany VARCHAR(20) COMMENT '快递公司',
    purchaseTime DATETIME COMMENT '下单时间',
    deliverTime DATETIME COMMENT '发货时间',
    receiveTime DATETIME COMMENT '收货时间',
    status INTEGER COMMENT '发货单状态，0:待装箱发货，1: 待仓库收货, 2: 已收货, 5: 已取消, 6: 部分收货',
    deliveryAddressLabel VARCHAR(50) COMMENT '发货仓库',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(deliveryOrderSn)
  );
 `
