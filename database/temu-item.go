/*
 * Maintained by jemo from 2023.3.29 to now
 * Created by jemo on 2023.3.29 10:46:53
 * 多多跨境已加入站点商品
 */
package database

const temuItem = `
  CREATE TABLE IF NOT EXISTS temuItem (
    supplierId VARCHAR(20),
    supplierName VARCHAR(20),
    productId VARCHAR(20),
    goodsId VARCHAR(20),
    isDress BOOLEAN,
    sampleNeeded BOOLEAN,
    mustSkipSample BOOLEAN,
    leafCategoryName VARCHAR(20),
    leafCategoryId INTEGER,
    fullCategoryName VARCHAR(300),
    carouselImageUrlList VARCHAR(1500),
    supplierPrice VARCHAR(20),
    supplierPriceCurrencyType VARCHAR(10),
    exchangeSupplierPrice VARCHAR(20),
    exchangeRate VARCHAR(20),
    productName VARCHAR(300),
    productPropertyList VARCHAR(1000),
    productCreatedAt DATETIME,
    productUpdatedAt DATETIME,
    productSource INTEGER,
    inSiteSimilarType INTEGER,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(productId)
  );
 `
