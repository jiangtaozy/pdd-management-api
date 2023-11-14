/*
 * Maintained by jemo from 2023.7.4 to now
 * Created by jemo on 2023.7.4 16:33:24
 */

package database

const createWarehouse =`
  CREATE TABLE IF NOT EXISTS warehouse (
    warehouseId INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '仓库编号',
    warehouseName VARCHAR(20) unique COMMENT '仓库名称',
    warehouseAddress VARCHAR(100) unique COMMENT '仓库地址',
    receiver VARCHAR(20) unique COMMENT '收货人',
    phone VARCHAR(20) unique COMMENT '手机号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
