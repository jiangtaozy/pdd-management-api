/*
 * Maintained by jemo from 2021.1.20 to now
 * Created by jemo on 2021.1.20 11:03:53
 * Pdd Goods Flow Data
 * 拼多多商品流量数据
 */

package database

const createPddGoodsFlowData =`
  CREATE TABLE IF NOT EXISTS pddGoodsFlowData (
    adStrategy INTEGER,
    cate3AvgGoodsVcr FLOAT,
    cate3IsPgvAbove BOOLEAN,
    cate3PctGoodsVcr FLOAT,
    cfmOrdrCnt INTEGER,
    cfmOrdrCntPpr INTEGER,
    cfmOrdrCntPprIsPercent BOOLEAN,
    cfmOrdrCntYtd INTEGER,
    cfmOrdrGoodsQty INTEGER,
    cfmOrdrGoodsQtyPpr INTEGER,
    cfmOrdrGoodsQtyPprIsPercent BOOLEAN,
    cfmOrdrGoodsQtyYtd INTEGER,
    cfmOrdrRtoPpr INTEGER,
    cfmOrdrRtoPprIsPercent BOOLEAN,
    goodsFavCnt INTEGER COMMENT '商品收藏用户数',
    goodsFavCntPpr FLOAT,
    goodsFavCntPprIsPercent BOOLEAN,
    goodsFavCntYtd INTEGER,
    goodsId BIGINT UNSIGNED NOT NULL,
    goodsLabel VARCHAR(50),
    goodsName VARCHAR(60),
    goodsPv INTEGER COMMENT '商品浏览量',
    goodsPvPpr FLOAT,
    goodsPvPprIsPercent BOOLEAN,
    goodsPvYtd INTEGER,
    goodsUv INTEGER COMMENT '商品访客数',
    goodsUvPpr FLOAT,
    goodsUvPprIsPercent BOOLEAN,
    goodsUvYtd INTEGER,
    goodsVcr FLOAT COMMENT '支付转化率',
    goodsVcrPpr FLOAT,
    goodsVcrPprIsPercent BOOLEAN,
    goodsVcrYtd INTEGER,
    hdThumbUrl VARCHAR(200),
    imprUsrCnt INTEGER,
    imprUsrCntPpr FLOAT,
    imprUsrCntPprIsPercent BOOLEAN,
    imprUsrCntYtd INTEGER,
    isCreated1m BOOLEAN,
    isNewstyle BOOLEAN,
    payOrdrAmt FLOAT COMMENT '支付金额',
    payOrdrAmtPpr FLOAT,
    payOrdrAmtPprIsPercent BOOLEAN,
    payOrdrAmtYtd FLOAT,
    payOrdrCnt INTEGER COMMENT '支付订单数',
    payOrdrCntPpr FLOAT,
    payOrdrCntPprIsPercent BOOLEAN,
    payOrdrCntYtd INTEGER,
    payOrdrGoodsQty INTEGER,
    payOrdrGoodsQtyPpr FLOAT,
    payOrdrGoodsQtyPprIsPercent BOOLEAN,
    payOrdrGoodsQtyYtd INTEGER,
    payOrdrUsrCnt INTEGER,
    payOrdrUsrCntPpr FLOAT,
    payOrdrUsrCntPprIsPercent BOOLEAN,
    payOrdrUsrCntYtd INTEGER,
    pctGoodsVcr INTEGER,
    pctGoodsVcrYtd INTEGER,
    statDate DATE NOT NULL,
    url VARCHAR(200),
    INDEX pddGoodsFlowDataIndex (goodsId, statDate)
  );
`
