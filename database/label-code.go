/*
 * Maintained by jemo from 2023.1.15 to now
 * Created by jemo on 2023.1.15 13:31:31
 * 商品条码
 */
package database

const labelCode = `
  CREATE TABLE IF NOT EXISTS labelCode (
    supplierId VARCHAR(20) COMMENT '供应商ID',
    productSkcId VARCHAR(20) COMMENT 'skc ID',
    productSkuId VARCHAR(20) COMMENT 'sku ID',
    labelCode VARCHAR(20) NOT NULL UNIQUE COMMENT '商品条码',
    skcExtCode VARCHAR(20) COMMENT 'skc货号',
    skuExtCode VARCHAR(20) COMMENT 'sku货号',
    productId VARCHAR(20) COMMENT '商品ID',
    productName VARCHAR(200) COMMENT '商品名称',
    catId VARCHAR(10) COMMENT '类目ID',
    catName VARCHAR(100) COMMENT '类目名称',
    catType VARCHAR(20) COMMENT '类目类型',
    displayImage VARCHAR(200) COMMENT '图像',
    productSkcSpecList VARCHAR(200) COMMENT 'skc spec',
    skcEn VARCHAR(500) COMMENT 'skc spec en',
    productSkuSpecList VARCHAR(200) COMMENT 'sku spec',
    skuEn VARCHAR(500) COMMENT 'sku spec en',
    isFragile BOOLEAN COMMENT '是否易碎'
  );
 `
