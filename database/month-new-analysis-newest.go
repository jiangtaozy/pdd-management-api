/*
 * Maintained by jemo from 2024.12.13 to now
 * Created by jemo on 2024.12.13 11:09:34
 * 月上新分析-最新数据
 */

package database

const createMonthNewAnalysisNewest =`
  CREATE TABLE IF NOT EXISTS monthNewAnalysisNewest (
    newMonth DATE NOT NULL COMMENT '上新月份',
    newQuantity INTEGER COMMENT '上新数量',
    thirtyDaysTotalSales INTEGER COMMENT '30天总销量',
    perDayTotalSales DECIMAL(10,2) COMMENT '30天总销量平均每天总销量',
    perDayTotalSalesNewQuantityRatio DECIMAL(10,4) COMMENT '30天总销量平均每天总销量与上新数量比率',
    regularSalesQuantity INTEGER COMMENT '平销款数量，每天销量大于等于1，小于5',
    regularSalesRatio DECIMAL(10,4) COMMENT '平销款比率',
    hotSalesQuantity INTEGER COMMENT '热销款数量，每天销量大于等于5',
    hotSalesRatio DECIMAL(10,4) COMMENT '热销款比率',
    salesQuantity INTEGER COMMENT '销售款数量，每天销量大于等于1',
    salesRatio DECIMAL(10,4) COMMENT '销售款比率',
    poorSalesQuantity INTEGER COMMENT '滞销款数量，每天销量小于1',
    poorSalesRatio DECIMAL(10,4) COMMENT '滞销款比率',
    salesPerDayAmount INTEGER COMMENT '销售款每天销售金额，单位分，每天销量大于等于1',
    salesPerDayGrossProfit INTEGER COMMENT '销售款每天销售毛利润，单位分，每天销量大于等于1',
    salesPerDayGrossProfitNewQuantityRatio DECIMAL(10,4) COMMENT '销售款每天销售毛利润与上新数量比率，单位分，每天销量大于等于1',
    lastThirtyDaysSalesAmount INTEGER COMMENT '近30天销售额，单位分',
    lastThirtyDaysSalesAmountNewQuantityRatio INTEGER COMMENT '近30天销售额与上新数量比率，单位分',
    lastThirtyDaysGrossProfit INTEGER COMMENT '近30天毛利润，单位分',
    lastThirtyDaysGrossProfitNewQuantityRatio INTEGER COMMENT '近30天毛利润与上新数量比率，单位分',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
