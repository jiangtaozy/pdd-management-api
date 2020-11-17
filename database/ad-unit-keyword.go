/*
 * Maintained by jemo from 2020.8.14 to now
 * Created by jemo on 2020.8.14 14:02:24
 * ad unit keyword
 */

package database

const createAdUnitKeyword =`
  CREATE TABLE IF NOT EXISTS adUnitKeyword (
    adId INTEGER UNSIGNED NOT NULL,
    mallId INTEGER UNSIGNED NOT NULL,
    impression INTEGER UNSIGNED NOT NULL COMMENT '曝光量',
    click INTEGER UNSIGNED NOT NULL COMMENT '点击量',
    ctr FLOAT UNSIGNED NOT NULL COMMENT '点击率',
    transactionCost INTEGER UNSIGNED NOT NULL COMMENT '交易花费(分)',
    spend INTEGER UNSIGNED NOT NULL COMMENT '花费(厘)',
    roi FLOAT UNSIGNED NOT NULL COMMENT '投入产出比',
    orderNum INTEGER UNSIGNED NOT NULL COMMENT '订单量',
    cpc FLOAT UNSIGNED NOT NULL COMMENT 'Cost Per Click 平均点击花费(厘)',
    cvr FLOAT UNSIGNED NOT NULL COMMENT 'Click Conversion Rate 点击转化率',
    gmv INTEGER UNSIGNED NOT NULL COMMENT '交易额(厘)',
    cpm FLOAT UNSIGNED NOT NULL COMMENT '千次曝光花费(厘)',
    mallFavNum INTEGER UNSIGNED NOT NULL COMMENT '店铺关注量',
    goodsFavNum INTEGER UNSIGNED NOT NULL COMMENT '商品收藏量',
    inquiryNum INTEGER UNSIGNED NOT NULL DEFAULT 0,
    uniqueView INTEGER UNSIGNED NOT NULL DEFAULT 0,
    rankAverage INTEGER UNSIGNED NOT NULL DEFAULT 0,
    rankMedian INTEGER UNSIGNED NOT NULL DEFAULT 0,
    avgPayAmount INTEGER UNSIGNED NOT NULL DEFAULT 0,
    appActivateNum INTEGER UNSIGNED NOT NULL DEFAULT 0,
    costPerAppActivate INTEGER UNSIGNED NOT NULL DEFAULT 0,
    appActivateRate INTEGER UNSIGNED NOT NULL DEFAULT 0,
    appRegisterNum INTEGER UNSIGNED NOT NULL DEFAULT 0,
    costPerAppRegister INTEGER UNSIGNED NOT NULL DEFAULT 0,
    appPayNum INTEGER UNSIGNED NOT NULL DEFAULT 0,
    costPerAppPay INTEGER UNSIGNED NOT NULL DEFAULT 0,
    date DATE NOT NULL,
    entityId BIGINT UNSIGNED,
    dimensionType TINYINT UNSIGNED,
    bid INTEGER UNSIGNED NOT NULL COMMENT '出价/厘',
    bidPremium INTEGER UNSIGNED NOT NULL COMMENT '精确匹配溢价/万分之一',
    bidPremiumValue INTEGER UNSIGNED NOT NULL COMMENT '精确匹配出价/厘',
    keyword VARCHAR(30) NOT NULL COMMENT '关键词',
    keywordAdIdx VARCHAR(20) NOT NULL COMMENT '90天平均排名',
    qualityScore INTEGER UNSIGNED NOT NULL COMMENT '质量分',
    keywordAdIdxOri FLOAT UNSIGNED,
    keywordId BIGINT UNSIGNED,
    keywordType INTEGER UNSIGNED,
    planStrategy INTEGER UNSIGNED,
    dataOperateStatus INTEGER UNSIGNED,
    INDEX adUnitKeywordIndex (mallId, adId, keywordId, date)
  );
`
