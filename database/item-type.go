/*
 * Maintained by jemo from 2022.4.18 to now
 * Created by jemo on 2022.4.18 10:06:28
 * 商品类型列表
 */

package database

const createItemType =`
  CREATE TABLE IF NOT EXISTS itemType (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    typeName VARCHAR(20) unique COMMENT '商品类型名称',
    typeNum VARCHAR(20) unique COMMENT '商品类型编码',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
