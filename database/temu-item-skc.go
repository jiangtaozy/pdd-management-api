/*
 * Maintained by jemo from 2023.3.29 to now
 * Created by jemo on 2023.3.29 11:35:02
 * 多多跨境已加入站点商品skc
 */
package database

const temuItemSkc = `
  CREATE TABLE IF NOT EXISTS temuItemSkc (
    supplierId VARCHAR(20),
    productId VARCHAR(20),
    selectId VARCHAR(20),
    skcId VARCHAR(20) NOT NULL,
    selectStatus INTEGER COMMENT '1:未发布，7: 价格申报中，9: 价格已作废, 10: 待创建首单, 11: 已创建首单，12: 已发布到站点, 13: 已下架, 14: 待修改, 15: 修改完成, 17: 已终止',
    goodsSkcStatus INTEGER,
    color VARCHAR(20),
    colorName VARCHAR(20),
    supplierPrice VARCHAR(20),
    supplierPriceCurrencyType VARCHAR(10),
    exchangeSupplierPrice VARCHAR(10),
    exchangeRate VARCHAR(10),
    previewImgUrlList VARCHAR(500),
    sampleType INTEGER,
    sampleQcType INTEGER,
    qcInfo VARCHAR(200),
    createdTime DATETIME,
    selectedTime DATETIME,
    samplePostTime DATETIME,
    qcCompletedTime DATETIME,
    priceVerificationTime DATETIME,
    firstPurchaseTime DATETIME,
    addedToSiteTime DATETIME,
    unPublishedTime DATETIME,
    samplePostingFinishedTime DATETIME,
    terminatedTime DATETIME,
    buyerUid VARCHAR(20),
    buyerName VARCHAR(20),
    labelList VARCHAR(100),
    buyerSelectedSkuList VARCHAR(100),
    defaultSelectedSkuList VARCHAR(100),
    productSellMode INTEGER,
    reSampleTimes INTEGER,
    remarkList VARCHAR(100),
    selectCanceledTimes INTEGER,
    qcSampleAdjustment VARCHAR(200),
    sampleQcOrderStatus INTEGER,
    primarySecondarySkc INTEGER,
    primaryMultiColor INTEGER,
    productHotTag INTEGER,
    productSearchTag INTEGER,
    extCode VARCHAR(20),
    latestPriceComparingStatus VARCHAR(20),
    terminationReasonList VARCHAR(200),
    isNotSentFirstOrder BOOLEAN DEFAULT 0 COMMENT '是否不发首单-新增字段-核价太低等原因',
    hasNotice BOOLEAN DEFAULT 0 COMMENT '是否已通知发首单-新增字段',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(skcId)
  );
 `
