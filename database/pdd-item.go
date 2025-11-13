/*
 * Maintained by jemo from 2020.5.24 to now
 * Created by jemo on 2020.5.24 10:46:39
 * pdd item
 */

package database

const createPddItem =`
  CREATE TABLE IF NOT EXISTS pddItem (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    mallId VARCHAR(20),
    quantity INTEGER UNSIGNED NOT NULL DEFAULT 0,
    score INTEGER UNSIGNED NOT NULL DEFAULT 0,
    resource VARCHAR(50) NOT NULL DEFAULT '',
    priority VARCHAR(50),
    skuGroupPriceMin INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '单位：分',
    skuGroupPriceMax INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '单位：分',
    skuPriceMin INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '单位：分',
    skuPriceMax INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '单位：分',
    originSkuGroupPriceMin INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '单位：分',
    originSkuGroupPriceMax INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '单位：分',
    pddId BIGINT UNSIGNED NOT NULL,
    goodsName VARCHAR(60) NOT NULL,
    goodsSn VARCHAR(30),
    goodsType TINYINT UNSIGNED NOT NULL DEFAULT 0,
    catId SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    catName VARCHAR(30) NOT NULL DEFAULT '',
    eventType TINYINT UNSIGNED NOT NULL DEFAULT 0,
    displayPriority VARCHAR(10),
    reserveQuantity SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    hdThumbUrl VARCHAR(200) NOT NULL DEFAULT '',
    imageUrl VARCHAR(200),
    thumbUrl VARCHAR(200) NOT NULL DEFAULT '',
    isOnsale BOOLEAN NOT NULL DEFAULT false,
    soldQuantity INTEGER UNSIGNED NOT NULL DEFAULT 0,
    offSaleCheck VARCHAR(50),
    onsaleCheck VARCHAR(50),
    editCheck VARCHAR(50),
    grayGoodsHighQualityRefund BOOLEAN,
    shareDesc VARCHAR(50),
    goodsDesc VARCHAR(50),
    activityTitle VARCHAR(50),
    isCardSecret BOOLEAN NOT NULL DEFAULT false,
    marketPrice INTEGER UNSIGNED NOT NULL DEFAULT 0 COMMENT '单位：分',
    outGoodsSn VARCHAR(20),
    soldQuantityForThirtyDays INTEGER UNSIGNED NOT NULL DEFAULT 0,
    favCnt INTEGER UNSIGNED NOT NULL DEFAULT 0,
    multiTreasure BOOLEAN NOT NULL DEFAULT false,
    multiTreasureStatus TINYINT UNSIGNED NOT NULL DEFAULT 0,
    ifNewGoods BOOLEAN NOT NULL DEFAULT false,
    tagCategoryList VARCHAR(50),
    titleScrFlag BOOLEAN,
    titleDesc VARCHAR(50),
    propsScrFlag BOOLEAN,
    propsDesc VARCHAR(50),
    galleryScrFlag BOOLEAN,
    galleryDesc VARCHAR(50),
    galleryVideoScrFlag BOOLEAN,
    longGraphScrFlag BOOLEAN,
    whiteGraphScrFlag BOOLEAN,
    materialScrFlag VARCHAR(50),
    goodsInfoScr VARCHAR(10),
    goodsInfoScoreUpdateInfoHint VARCHAR(50),
    createdAt TIMESTAMP NULL,
    preSaleTime INTEGER UNSIGNED NOT NULL DEFAULT 0,
    shipmentLimitSecond INTEGER UNSIGNED NOT NULL DEFAULT 0,
    isGroupPreSale BOOLEAN NOT NULL DEFAULT false,
    isPreSale BOOLEAN COMMENT '是否预售',
    guideTarget INTEGER UNSIGNED NOT NULL DEFAULT 0,
    overSell BOOLEAN NOT NULL DEFAULT false,
    marketLabels VARCHAR(300),
    labels VARCHAR(300),
    skuType TINYINT UNSIGNED NOT NULL DEFAULT 0,
    isMoreSku VARCHAR(50),
    skuList VARCHAR(50),
    rejectStatus VARCHAR(50),
    rejectReason VARCHAR(50),
    grossProfit DECIMAL(10,2) COMMENT '累计毛利润',
    promotionCost DECIMAL(10,2) COMMENT '累计推广费用',
    netProfit DECIMAL(10,2) COMMENT '累计净利润',
    INDEX pddIdIndex(pddId),
    UNIQUE KEY(pddId)
    INDEX outGoodsSnIndex(outGoodsSn)
  );
`

//alter table pddItem modify column grayGoodsHighQualityRefund BOOLEAN;
//alter table pddItem modify column marketLabels VARCHAR(300);
//alter table pddItem modify column labels VARCHAR(300);

//alter table pddItem add mallId VARCHAR(20);
