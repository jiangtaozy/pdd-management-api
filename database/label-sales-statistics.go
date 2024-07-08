/*
 * Maintained by jemo from 2023.6.3 to now
 * Created by jemo on 2023.6.3 14:33:41
 * 标签销售统计
 */
package database

const labelSaleManage = `
  CREATE TABLE IF NOT EXISTS labelSalesStatistics (
    day DATE NOT NULL,
    label INTEGER COMMENT '标签，0: 清仓款(近7天日均单量等于0), 1: 滞销款(小于1), 2: 平常款(大于1), 3: 旺款(大于5), 4: 爆款(大于10), 5: 爆旺款(大于100)',
    skuCount INTEGER COMMENT 'sku数量',
    inCartNumber7d INTEGER COMMENT '按近7日用户加购数量',
    dayInCart INTEGER COMMENT '按近7日用户加购数量计算每天加购量',
    lastSevenDaysSaleVolume INTEGER COMMENT '按近7天销量',
    daySale INTEGER COMMENT '每天销量-近7天销量/7',
    lastSevenDaysSalesAmount INTEGER COMMENT '近7天销售金额单位分',
    daySalesAmount INTEGER COMMENT '每天销售金额-近7天销售金额/7-单位分',
    lastSevenDaysGrossProfit INTEGER COMMENT '近7天毛利润-单位分',
    dayGrossProfit INTEGER COMMENT '每天毛利润-近7天毛利润/7-单位分',
    warehouseInventoryNum INTEGER COMMENT '仓内可用库存',
    totalLocalSurplusStock INTEGER COMMENT '本地剩余库存',
    totalInventory INTEGER COMMENT '总库存',
    totalPurchaseQuantity INTEGER COMMENT '当前历史总进货量',
    totalPurchaseAmount INTEGER COMMENT '当前历史总进货金额 = 当前历史总进货量* 进货价',
    totalSaleQuantity INTEGER COMMENT '当前历史总销量 = 当前历史总进货量 - 当前总库存',
    totalSaleAmount INTEGER COMMENT '当前历史总销售额 = 当前历史总销量 * 供货价',
    totalGrossProfit INTEGER COMMENT '当前历史总毛利润 = 当前历史总销量 * (供货价 - 进货价)',
    totalNetProfit INTEGER COMMENT '当前历史总净利润 = 当前历史总销售额 - 当前历史总进货金额',
    totalInventoryAmount INTEGER COMMENT '当前总库存金额 = 总库存 * 进货价',
    cumulativeTurnoverRate DECIMAL(6,4) COMMENT '累计资金周转率 = 累计总销售额 / 累计总进货金额',
    monthlyTurnoverRate DECIMAL(6,4) COMMENT '近三十天资金周转率 = 近一个月销售额 / ((累计进货量 - 累计销量 + 近一个月销量) * 进货价)',
    sevenDayTurnoverRate DECIMAL(6,4) COMMENT '近七天周转率 = 近七天销售额 / ((累计进货量 - 累计销量 + 近七天销量) * 进货价)',
    stockDays DECIMAL(10,2) COMMENT '库存天数 = 总库存 / 近7天销量 * 7',
    returnQuantity INTEGER COMMENT '退货库存',
    returnAmount INTEGER COMMENT '退货库存金额 = 退货库存 * 进货价',
    lastThirtyDaysSaleVolume INTEGER COMMENT '近30天销量',
    lastThirtyDaysSalesAmount INTEGER COMMENT '近30天销售额',
    lastThirtyDaysGrossProfit INTEGER COMMENT '近30天毛利润单位分',
    lastThirtyDaysPurchaseAmount INTEGER COMMENT '近30天进货金额单位分',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(day, label)
  );
 `
