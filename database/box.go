/*
 * Maintained by jemo from 2023.11.27 to now
 * Created by jemo on 2023.11.27 14:13:03
 * 箱号列表
 */

package database

const createBox =`
  CREATE TABLE IF NOT EXISTS box (
    boxId INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '箱号编号',
    boxName VARCHAR(20) COMMENT '箱号名称',
    rowId INTEGER COMMENT '货架排编号',
    layerId INTEGER COMMENT '货架层编号',
    shelfId INTEGER COMMENT '货架编号',
    passageId INTEGER COMMENT '通道编号',
    roomId INTEGER COMMENT '房间编号',
    warehouseId INTEGER COMMENT '仓库编号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
