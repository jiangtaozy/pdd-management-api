/*
 * Maintained by jemo from 2024.11.26 to now
 * Created by jemo on 2024.11.26 10:15:43
 * 月上新分析
 */

package database

// 24年1月份上新，100款
// 24年10月份统计销量
const createMonthNewAnalysis =`
  CREATE TABLE IF NOT EXISTS monthNewAnalysis (
    newMonth DATE NOT NULL COMMENT '上新月份',
    saleMonth DATE NOT NULL COMMENT '销售月份',
    newQuantity INTEGER COMMENT '上新数量',
    thirtyDaysTotalSales INTEGER COMMENT '30天总销量',
    perDayTotalSales DECIMAL(10,2) COMMENT '30天总销量平均每天总销量',
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
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(newMonth, saleMonth)
  );
`
