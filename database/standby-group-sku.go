/*
 * Maintained by jemo from 2024.2.28 to now
 * Created by jemo on 2024.2.28 16:27:41
 * 备选货号组合sku
 */
package database

const standbyGroupSku = `
  CREATE TABLE IF NOT EXISTS standbyGroupSku (
    temuSkuId VARCHAR(20) NOT NULL,
    aliSkuId VARCHAR(20) NOT NULL,
    quantity INTEGER,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
    PRIMARY KEY(temuSkuId, aliSkuId)
  );
 `
