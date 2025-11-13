/*
 * Maintained by jemo from 2025.09.05 to now
 * Created by jemo on 2025.09.05 10:24:12
 * 已上架，未发首单
 */

package database

const createNotSendFirstOrder =`
  CREATE TABLE IF NOT EXISTS notSendFirstOrder (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    supplierId VARCHAR(20),
    skcExtCode VARCHAR(20) COMMENT 'skc货号',
    skuExtCode VARCHAR(20) COMMENT 'sku货号',
    productName VARCHAR(200) COMMENT '商品名称',
    displayImage VARCHAR(200) COMMENT '图像',
    productSkuSpecList VARCHAR(200) COMMENT 'sku spec',
    isSendFirstOrder BOOLEAN DEFAULT 0 COMMENT '是否已发首单',
    isNotSend BOOLEAN DEFAULT 0 COMMENT '是否不发首单',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
