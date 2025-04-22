/*
 * Maintained by jemo from 2022.4.15 to now
 * Created by jemo on 2022.4.15 10:56:18
 * 阿里1688商品SKU
 */

package database

const createItemSku =`
  CREATE TABLE IF NOT EXISTS itemSku (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    skuName VARCHAR(300) NOT NULL COMMENT 'sku名称，如【香槟】钻樱桃抓夹',
    specId VARCHAR(50) COMMENT 'specId',
    specAttrs VARCHAR(50) COMMENT 'spec名称',
    price INTEGER UNSIGNED COMMENT 'sku价格，单位：分',
    saleCount INTEGER UNSIGNED COMMENT '已售数量',
    discountPrice INTEGER UNSIGNED COMMENT '折扣价，单位：分',
    canBookCount INTEGER UNSIGNED COMMENT 'sku库存',
    retailPrice INTEGER UNSIGNED COMMENT '零售价，单位：分',
    skuId VARCHAR(50) COMMENT 'skuId',
    isPromotionSku BOOLEAN COMMENT '是否推广',
    shortSkuName VARCHAR(20) COMMENT '简化sku名称',
    shortSkuNum VARCHAR(20) COMMENT '简化sku编码',
    isDeleted BOOLEAN NOT NULL DEFAULT 0 COMMENT '是否已删除',
    skuExtCode VARCHAR(10) COMMENT 'sku货号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
//alter table itemSku add skuExtCode VARCHAR(10) COMMENT 'sku货号';
