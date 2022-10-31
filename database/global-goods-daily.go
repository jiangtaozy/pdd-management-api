/*
 * Maintained by jemo from 2022.10.31 to now
 * Created by jemo on 2022.10.31 09:57:16
 * pdd ad total station goods daily
 * 拼多多全站推广商品每日数据
 */

package database

const createGlobalGoodsDaily =`
  CREATE TABLE IF NOT EXISTS globalGoodsDaily (
    date DATE NOT NULL COMMENT '日期',
    goodsId BIGINT UNSIGNED NOT NULL COMMENT '商品ID',
    goodsName VARCHAR(60) NOT NULL COMMENT '商品名称',
    spend FLOAT UNSIGNED NOT NULL COMMENT '花费(元)',
    gmv FLOAT UNSIGNED NOT NULL COMMENT '交易额(元)',
    roi FLOAT UNSIGNED NOT NULL COMMENT '实际投产比',
    dealNum INTEGER UNSIGNED NOT NULL COMMENT '成交笔数',
    costPerDeal FLOAT UNSIGNED NOT NULL COMMENT '每笔成交花费(元)',
    amountPerDeal FLOAT UNSIGNED NOT NULL COMMENT '每笔成交金额(元)',
    directAmount FLOAT UNSIGNED NOT NULL COMMENT '直接交易额(元)',
    indirectAmount FLOAT UNSIGNED NOT NULL COMMENT '间接交易额(元)',
    directDealNum INTEGER UNSIGNED NOT NULL COMMENT '直接成交笔数',
    indirectDealNum INTEGER UNSIGNED NOT NULL COMMENT '间接成交笔数',
    directAmountPerDeal FLOAT UNSIGNED NOT NULL COMMENT '每笔直接成交金额(元)',
    indirectAmountPerDeal FLOAT UNSIGNED NOT NULL COMMENT '每笔间接成交金额(元)',
    globalCostRate FLOAT UNSIGNED NOT NULL COMMENT '全站推广费比',
    impression INTEGER UNSIGNED NOT NULL COMMENT '曝光量',
    click INTEGER UNSIGNED NOT NULL COMMENT '点击量',
    clickRate FLOAT UNSIGNED NOT NULL COMMENT '点击率',
    clickConversionRate FLOAT UNSIGNED NOT NULL COMMENT '点击转化率',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(date, goodsId)
  );
`
