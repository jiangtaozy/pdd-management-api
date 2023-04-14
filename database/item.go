/*
 * Maintained by jemo from 2022.4.18 to now
 * Created by jemo on 2022.4.18 10:47:24
 * 商品列表
 */
package database

const item = `
  CREATE TABLE IF NOT EXISTS item (
    name VARCHAR(200) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    imgUrl VARCHAR(200) NOT NULL DEFAULT '',
    detailUrl VARCHAR(60) NOT NULL DEFAULT '',
    siteType TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1: 1688, 2: hznzcn, 3: 线下微信',
    originalId VARCHAR(30) NOT NULL DEFAULT '',
    supplierId INTEGER UNSIGNED NOT NULL DEFAULT 0,
    saleQuantity INTEGER UNSIGNED NOT NULL DEFAULT 0,
    quantitySumMonth INTEGER UNSIGNED NOT NULL DEFAULT 0,
    gmv30dRt DECIMAL(12,2) UNSIGNED NOT NULL DEFAULT 0,
    searchId BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    suitPrice DECIMAL(10,2) NOT NULL DEFAULT 0,
    shippingPrice DECIMAL(10,2) NOT NULL DEFAULT 0,
    forSell BOOLEAN NOT NULL DEFAULT false,
    imgUrlOf290x290 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf120x120 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf270x270 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf100x100 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf150x150 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf220x220 VARCHAR(100) NOT NULL DEFAULT '',
    womenProductId INTEGER UNSIGNED,
    keyName VARCHAR(200),
    itemTypeKey INTEGER COMMENT '类型key',
    itemNum VARCHAR(20) COMMENT '编码',
    length DECIMAL(10,2) UNSIGNED COMMENT '长(cm)',
    width DECIMAL(10,2) UNSIGNED COMMENT '宽(cm)',
    height DECIMAL(10,2) UNSIGNED COMMENT '高(cm)',
    weight INTEGER UNSIGNED COMMENT '重量(g)',
    crossBorderUrl VARCHAR(100) COMMENT '跨境平台地址',
    notForSellReason VARCHAR(300) COMMENT '下架原因',
    INDEX forSellIndex(forSell)
  );
 `
