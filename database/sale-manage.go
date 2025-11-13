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
    className VARCHAR(200) COMMENT 'SKU属性',
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
    supplierPrice INTEGER COMMENT '申报价格（单位分）',
    predictSaleAvailableDays DECIMAL(10,2) COMMENT '可售天数',
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
    isLastThirtyDaysNeedStock BOOLEAN COMMENT '按30天是否需要订货',
    lastThirtyDaysStockQuantity INTEGER COMMENT '按30天备货量',
    isLastSevenDaysNeedStock BOOLEAN COMMENT '按7天是否需要订货，是否订货，是否进货',
    lastSevenDaysStockQuantity INTEGER COMMENT '按7天订货量',
    isNeedApplyStock BOOLEAN COMMENT '是否需要申请备货单',
    needApplyStock INTEGER COMMENT '需要申请备货单数量',
    label INTEGER COMMENT '标签，-1: null未上架, 0: 清仓款(近7天日均单量等于0), 1: 滞销款(小于1), 2: 平常款(大于1), 3: 旺款(大于5), 4: 爆款(大于10), 5: 爆旺款(大于100)',
    growth BOOLEAN comment '0: 下降款(7日日均销量 < 30日日均销量 / 2), 1: 增长款(7日日均销量 > 30日日均销量 * 2), 2: 平常款(30日日均销量 / 2 < 7日日均销量 < 30日日均销量 * 2)',
    lastSevenDaysSalesAmount INTEGER COMMENT '近7天销售金额单位分',
    lastSevenDaysGrossProfit INTEGER COMMENT '近7天毛利润单位分',
    purchasePrice INTEGER COMMENT '进货价(单位分)',
    lastSevenDaysPurchaseAmount INTEGER COMMENT '近7天进货金额单位分',
    totalPurchaseQuantity INTEGER COMMENT '当前历史总进货量',
    totalPurchaseAmount INTEGER COMMENT '当前历史总进货金额 = 当前历史总进货量* 进货价',
    totalSaleQuantity INTEGER COMMENT '当前历史总销量 = 当前历史总进货量 - 当前总库存',
    totalSaleAmount INTEGER COMMENT '当前历史总销售额 = 当前历史总销量 * 供货价',
    totalGrossProfit INTEGER COMMENT '当前历史总毛利润 = 当前历史总销量 * (供货价 - 进货价)',
    totalNetProfit INTEGER COMMENT '当前历史总净利润 = 当前历史总销售额 - 当前历史总进货金额',
    totalInventoryAmount INTEGER COMMENT '当前总库存金额 = 总库存 * 进货价',
    cumulativeTurnoverRate DECIMAL(6,4) COMMENT '累计资金周转率 = 累计销售额 / 累计进货金额',
    monthlyTurnoverRate DECIMAL(6,4) COMMENT '近30天资金周转率 = 近一个月销售额 / ((累计进货量 - 累计销量 + 近一个月销量) * 进货价)',
    lastThirtyDaysSalesAmount INTEGER COMMENT '近30天销售额单位分',
    lastThirtyDaysGrossProfit INTEGER COMMENT '近30天毛利润单位分',
    lastThirtyDaysPurchaseAmount INTEGER COMMENT '近30天进货金额单位分',
    monthlyRemaining INTEGER COMMENT '近一个月剩余进货金额加销售进货金额 = ((累计进货量 - 累计销量 + 近一个月销量) * 进货价)',
    sevenDayTurnoverRate DECIMAL(6,4) COMMENT '七天周转率 = 近七天销售额 / ((累计进货量 - 累计销量 + 近七天销量) * 进货价)',
    sevenDayRemaining INTEGER COMMENT '七天剩余进货金额加销售进货金额 = ((累计进货量 - 累计销量 + 近七天销量) * 进货价)',
    stockDays DECIMAL(10,2) COMMENT '库存天数 = 总库存 / 近7天销量 * 7',
    returnQuantity INTEGER COMMENT '退货库存',
    returnAmount INTEGER COMMENT '退货库存金额 = 退货库存 * 进货价',
    safeInventoryDays INTEGER COMMENT '库存安全天数',
    isAdviceStock BOOLEAN COMMENT '是否建议备货',
    supplyStatus INTEGER COMMENT '备货状态，1: 暂时无法备货, 0: 正常备货/无货审核中',
    onSalesDurationOffline INTEGER COMMENT '加入站点时长(天)',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    isVerifyPrice BOOLEAN COMMENT '开款价格状态',
    purchaseConfig VARCHAR(20) COMMENT '备货逻辑',
    pddLastSevenDaysSaleVolume INTEGER COMMENT '拼多多近7天销量',
    PRIMARY KEY(day, productSkuId)
  );
 `
 /* 添加索引
alter table saleManage add index index_day(day);
alter table saleManage add index index_skcExtCode(skcExtCode);
alter table saleManage add index index_skuExtCode(skuExtCode);
alter table saleManage add index index_isLastSevenDaysNeedStock(isLastSevenDaysNeedStock);
alter table saleManage add purchaseConfig VARCHAR(20) COMMENT '备货逻辑';
//alter table saleManage add pddLastSevenDaysSaleVolume INTEGER COMMENT '拼多多近7天销量';
*/

//alter table saleManage modify column className VARCHAR(200) COMMENT 'SKU属性';
