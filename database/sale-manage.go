/*
 * Maintained by jemo from 2023.3.28 to now
 * Created by jemo on 2023.3.28 14:15:33
 * 销售管理
 */
package database

const saleManage = `
  CREATE TABLE IF NOT EXISTS saleManage (
    day DATE NOT NULL,
    supplierId VARCHAR(20) COMMENT '店铺id',
    productSkcId VARCHAR(20) COMMENT 'skc id',
    skcExtCode VARCHAR(20) COMMENT 'skc货号',
    buyerName VARCHAR(20) COMMENT '买手名',
    inCardNumber INTEGER COMMENT '用户累计加购数量',
    availableSaleDaysFromInventory DECIMAL(10,2) COMMENT '库存可售天数',
    predictSaleAdviceQuantity INTEGER,
    totalSaleVolume INTEGER,
    notVmiWaitDeliveryNum INTEGER COMMENT '非VMI待发货订单数',
    notVmiDeliveryDelayNum INTEGER COMMENT '非VMI发货延迟订单数',
    notVmiArrivalDelayNum INTEGER COMMENT '非VMI到货延迟订单数',
    notVmiTransportationNum INTEGER COMMENT '非VMI在途单数',
    skuExtCode VARCHAR(20) COMMENT 'SKU货号',
    className VARCHAR(20) COMMENT 'SKU属性',
    predictSaleWarehouseAvailableDays DECIMAL(10,2),
    predictTodaySaleVolume INTEGER,
    inCartNumber7d INTEGER COMMENT '近7日用户加购数量',
    sevenDaysSaleReference DECIMAL(10,2),
    lackQuantity INTEGER COMMENT '缺货数量',
    predictLastFiftyDaysSaleVolume INTEGER,
    sevenDaysReferenceSaleType INTEGER,
    productPriceAdjustStatus INTEGER,
    lastSevenDaysSaleVolume INTEGER COMMENT '近7天销量',
    vmiWaitDeliveryNum INTEGER COMMENT 'VMI待发货单数',
    vmiDeliveryDelayNum INTEGER COMMENT 'VMI发货延迟单数',
    vmiArrivalDelayNum INTEGER COMMENT 'VMI到货延迟单数',
    vmiTransportationNum INTEGER COMMENT 'VMI在途单数',
    suggestPurchaseNumUp INTEGER,
    nomsgSubsCntCntSth INTEGER COMMENT '已订阅待提醒到货',
    predictLastSevenDaysSaleVolume INTEGER,
    productSkuId VARCHAR(20) NOT NULL COMMENT 'sku id',
    lastThirtyDaysSaleVolume INTEGER COMMENT '近30天销量',
    goodsSkuId VARCHAR(20),
    isSubscribeArrivalRemind BOOLEAN,
    supplierPrice INTEGER COMMENT '申报价格',
    predictSaleAvailableDays DECIMAL(10,2),
    availableSaleDays DECIMAL(10,2) COMMENT '可售天数',
    todaySaleVolume INTEGER COMMENT '今天销量',
    predictSaleInventoryAvailableDays DECIMAL(10,2),
    adviceQuantity INTEGER COMMENT '建议备货量',
    waitOnShelfNum INTEGER,
    salesInventoryNum INTEGER,
    deliveryDelayNum INTEGER,
    warehouseInventoryNum INTEGER COMMENT '仓内可用库存',
    arrivalDelayNum INTEGER,
    waitApproveInventoryNum INTEGER COMMENT '待审核备货库存',
    waitQcNum INTEGER COMMENT '待质检',
    unavailableWarehouseInventoryNum INTEGER COMMENT '仓内暂不可用库存',
    waitInStock INTEGER,
    waitReceiveNum INTEGER COMMENT '已发货库存',
    waitDeliveryInventoryNum INTEGER COMMENT '已下单待发货库存',
    suggestPurchaseNumDown INTEGER,
    isAdjusted BOOLEAN,
    warehouseAvailableSaleDays DECIMAL(10,2) COMMENT '仓内库存可售天数',
    notReceivedAmount INTEGER COMMENT '未到货库存'
    totalLocalSurplusStock INTEGER COMMENT '本地剩余库存'
    totalInventory INTEGER COMMENT '总库存'
    lastThirtyDaysPredictThirtyDaysSale INTEGER COMMENT '按30天销量预估30天销量',
    lastSevenDaysPredictThirtyDaysSale INTEGER COMMENT '按7天销量预估30天销量',
    lastSevenDaysPredictThirtySevenDaysSale INTEGER COMMENT '按7天销量预估37天销量',
    isLastThirtyDaysNeedStock BOOLEAN COMMENT '按30天是否需要备货',
    lastThirtyDaysStockQuantity INTEGER COMMENT '按30天备货量',
    isLastSevenDaysNeedStock BOOLEAN COMMENT '按7天是否需要备货',
    lastSevenDaysStockQuantity INTEGER COMMENT '按7天备货量',
    isNeedApplyStock BOOLEAN COMMENT '是否需要申请备货单',
    needApplyStock INTEGER COMMENT '需要申请备货单数量',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(day, productSkuId)
  );
 `
