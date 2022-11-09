/*
 * Maintained by jemo from 2022.11.03 to now
 * Created by jemo on 2022.11.03 11:26:40
 * daily profit
 * 每日利润
 */

package database

const createDailyProfit =`
  CREATE TABLE IF NOT EXISTS dailyProfit (
    date DATE NOT NULL PRIMARY KEY COMMENT '日期',
    income DECIMAL(10,2) COMMENT '收入(元)',
    cost DECIMAL(10,2) COMMENT '成本(元)',
    grossProfit DECIMAL(10,2) COMMENT '毛利润(元)',
    promotionCost DECIMAL(10,2) COMMENT '推广费用(元)',
    netProfit DECIMAL(10,2) COMMENT '净利润(元)'
  );
`
