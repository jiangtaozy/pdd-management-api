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
    selectStatus INTEGER COMMENT '11: 已下首单，13: 已下架',
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
    qcInfo VARCHAR(20),
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
    qcSampleAdjustment VARCHAR(20),
    sampleQcOrderStatus INTEGER,
    primarySecondarySkc INTEGER,
    primaryMultiColor INTEGER,
    productHotTag INTEGER,
    productSearchTag INTEGER,
    extCode VARCHAR(20),
    latestPriceComparingStatus VARCHAR(20),
    terminationReasonList VARCHAR(200),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(skcId)
  );
 `
