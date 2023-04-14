/*
 * Maintained by jemo from 2023.3.29 to now
 * Created by jemo on 2023.3.29 11:49:44
 * 多多跨境已加入站点商品sku
 */
package database

const temuItemSku = `
  CREATE TABLE IF NOT EXISTS temuItemSku (
    supplierId VARCHAR(20),
    productId VARCHAR(20),
    skcId VARCHAR(20),
    skuId VARCHAR(20) NOT NULL,
    hasUploadBom BOOLEAN,
    productPropertyList VARCHAR(100),
    supplierPrice VARCHAR(10),
    priceReviewStatus INTEGER,
    supplierPriceCurrencyType VARCHAR(10),
    exchangeSupplierPrice VARCHAR(10),
    exchangeRate VARCHAR(10),
    sensitiveList VARCHAR(50),
    extCode VARCHAR(20),
    len INTEGER,
    width INTEGER,
    height INTEGER,
    weight INTEGER,
    skuPreviewImage VARCHAR(200),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(skuId)
  );
 `
