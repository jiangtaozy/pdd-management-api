/*
 * Maintained by jemo from 2020.8.29 to now
 * Created by jemo on 2020.8.29 07:06:57
 * pdd ad unit hourly data
 */

package database

const createPddAdUnitHourlyData = `
  CREATE TABLE IF NOT EXISTS pddAdUnitHourlyData (
    adId INTEGER UNSIGNED NOT NULL COMMENT '单元id',
    date DATE NOT NULL COMMENT '日期',
    hour TINYINT UNSIGNED NOT NULL COMMENT '小时',
    impression INTEGER UNSIGNED NOT NULL COMMENT '曝光量',
    click INTEGER UNSIGNED NOT NULL COMMENT '点击量',
    spend INTEGER UNSIGNED NOT NULL COMMENT '花费(厘)',
    orderNum INTEGER UNSIGNED NOT NULL COMMENT '订单量',
    gmv INTEGER UNSIGNED NOT NULL COMMENT '交易额(厘)',
    mallFavNum INTEGER UNSIGNED NOT NULL COMMENT '店铺关注量',
    goodsFavNum INTEGER UNSIGNED NOT NULL COMMENT '商品收藏量',
    inquiryNum INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '询单数',
    uniqueView INTEGER UNSIGNED NOT NULL DEFAULT 0,
    rankAverage INTEGER UNSIGNED NOT NULL DEFAULT 0,
    rankMedian INTEGER UNSIGNED NOT NULL DEFAULT 0
  );
`
