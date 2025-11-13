/*
 * Maintained by jemo from 2025.08.04 to now
 * Created by jemo on 2025.08.04 14:59:42
 * temu半托管商品sku
 */
package database

const semiItemSku = `
  CREATE TABLE IF NOT EXISTS semiItemSku (
    supplierId VARCHAR(20),
    productId VARCHAR(20),
    goodsId VARCHAR(20),
    skcId VARCHAR(20),
    goodsSkcId VARCHAR(20),
    skuPreviewImage VARCHAR(200),
    skuId VARCHAR(20) NOT NULL,
    goodsSkuId VARCHAR(20),
    selectStatus INTEGER COMMENT '10：未发布到站点，12：已发布到站点',
    extCode VARCHAR(20),
    len INTEGER,
    width INTEGER,
    height INTEGER,
    weight INTEGER COMMENT '毫克',
    productPropertyList VARCHAR(100),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(skuId)
  );
 `
