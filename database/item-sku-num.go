/*
 * Maintained by jemo from 2022.4.16 to now
 * Created by jemo on 2022.4.16 16:29:21
 * 商品SKU编码列表
 */

package database

const createItemSkuNum =`
  CREATE TABLE IF NOT EXISTS itemSkuNum (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    shortSkuName VARCHAR(20) unique COMMENT '简化sku名称',
    shortSkuNum VARCHAR(20) unique COMMENT '简化sku编码',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
