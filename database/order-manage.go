/*
 * Maintained by jemo from 2023.3.28 to now
 * Created by jemo on 2023.3.28 17:30:41
 * 备货单管理
 */
package database

const orderManage = `
  CREATE TABLE IF NOT EXISTS orderManage (
    subPurchaseOrderSn VARCHAR(20) COMMENT '订单号',
    productSkcId VARCHAR(20) COMMENT 'skc id',
    productSn VARCHAR(20) COMMENT 'skc 货号',
    purchaseTime DATETIME COMMENT '下单时间',
    status INTEGER COMMENT '订单状态，1: 待发货, 2: 已送货，3: 已收货，7: 已入库, 8: 已作废, null: 待审核',
    productSkuId VARCHAR(20) COMMENT 'sku id',
    extCode VARCHAR(20) COMMENT 'sku 货号',
    className VARCHAR(20) COMMENT '属性',
    purchaseQuantity INTEGER COMMENT '下单数量',
    deliverQuantity INTEGER COMMENT '送货数量',
    waitReceiveNum INTEGER COMMENT '待收货数量',
    waitQcNum INTEGER COMMENT '待质检数量',
    waitOnShelfNum INTEGER,
    waitInStock INTEGER,
    realReceiveAuthenticQuantity INTEGER COMMENT '实际接收数量',
    defectiveQuantity INTEGER,
    availableSaleDays DECIMAL(10,2),
    availableSaleDaysFromInventory DECIMAL(10,2),
    onShelfWaitingQcQuantity INTEGER,
    onShelfQcQuantity INTEGER,
    qcWaitingWithdrawQuantity INTEGER,
    qcWithdrawQuantity INTEGER,
    qcResult INTEGER,
    skuLackNum INTEGER,
    accumulateDeliverNum INTEGER,
    accumulateInboundNum INTEGER,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(subPurchaseOrderSn, productSkuId)
  );
 `
