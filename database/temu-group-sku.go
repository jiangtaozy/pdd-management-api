/*
 * Maintained by jemo from 2023.3.29 to now
 * Created by jemo on 2023.3.29 15:05:45
 * 多多跨境组合sku
 */
package database

const temuGroupSku = `
  CREATE TABLE IF NOT EXISTS temuGroupSku (
    temuSkuId VARCHAR(20) NOT NULL,
    aliSkuId VARCHAR(20) NOT NULL,
    quantity INTEGER,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(temuSkuId, aliSkuId)
  );
 `
