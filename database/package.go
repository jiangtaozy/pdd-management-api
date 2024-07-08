/*
 * Maintained by jemo from 2023.11.27 to now
 * Created by jemo on 2023.11.27 15:22:28
 * 包裹列表
 */

package database

const createPackage =`
  CREATE TABLE IF NOT EXISTS package (
    packageId INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '包裹编号',
    packageName VARCHAR(20) COMMENT '包裹名称',
    boxId INTEGER COMMENT '箱号编号',
    rowId INTEGER COMMENT '货架排编号',
    layerId INTEGER COMMENT '货架层编号',
    shelfId INTEGER COMMENT '货架编号',
    passageId INTEGER COMMENT '通道编号',
    roomId INTEGER COMMENT '房间编号',
    warehouseId INTEGER COMMENT '仓库编号',
    labelCode VARCHAR(20) COMMENT '商品条码号',
    quantity INTEGER COMMENT '数量',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
