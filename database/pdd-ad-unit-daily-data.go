/*
 * Maintained by jemo from 2020.5.27 to now
 * Created by jemo on 2020.5.27 16:21:52
 * pdd ad unit daily data
 */

package database

const createPddAdUnitDailyData =`
  CREATE TABLE IF NOT EXISTS pddAdUnitDailyData (
    adId INTEGER UNSIGNED NOT NULL,
    impression INTEGER UNSIGNED NOT NULL COMMENT '曝光量',
    click INTEGER UNSIGNED NOT NULL COMMENT '点击量',
    ctr FLOAT UNSIGNED NOT NULL COMMENT '点击率',
    transactionCost INTEGER UNSIGNED NOT NULL COMMENT '交易花费(分)',
    spend INTEGER UNSIGNED NOT NULL COMMENT '花费(厘)',
    roi FLOAT UNSIGNED NOT NULL COMMENT '投入产出比',
    orderNum INTEGER UNSIGNED NOT NULL COMMENT '订单量',
    cpc FLOAT UNSIGNED NOT NULL COMMENT 'Cost Per Click 平均点击花费(厘)',
    cvr FLOAT UNSIGNED NOT NULL COMMENT 'Click Conversion Rate 点击转化率',
    gmv INTEGER UNSIGNED NOT NULL COMMENT '交易额(分)',
    cpm FLOAT UNSIGNED NOT NULL COMMENT '千次曝光花费(厘)',
    mallFavNum INTEGER UNSIGNED NOT NULL COMMENT '店铺关注量',
    goodsFavNum INTEGER UNSIGNED NOT NULL COMMENT '商品收藏量',
    inquiryNum INTEGER UNSIGNED NOT NULL DEFAULT 0,
    uniqueView INTEGER UNSIGNED NOT NULL,
    rankAverage INTEGER UNSIGNED NOT NULL,
    rankMedian INTEGER UNSIGNED NOT NULL,
    avgPayAmount INTEGER UNSIGNED NOT NULL,
    appActivateNum INTEGER UNSIGNED NOT NULL,
    costPerAppActivate INTEGER UNSIGNED NOT NULL,
    appActivateRate INTEGER UNSIGNED NOT NULL,
    appRegisterNum FLOAT UNSIGNED NOT NULL,
    costPerAppRegister INTEGER UNSIGNED NOT NULL,
    appPayNum INTEGER UNSIGNED NOT NULL,
    costPerAppPay INTEGER UNSIGNED NOT NULL,
    date DATE NOT NULL,
    entityId INTEGER UNSIGNED,
    dimensionType TINYINT UNSIGNED
  );
`
