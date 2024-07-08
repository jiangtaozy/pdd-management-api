/*
 * Maintained by jemo from 2023.11.22 to now
 * Created by jemo on 2023.11.22 16:15:41
 * 排号列表
 */

package database

const createRow =`
  CREATE TABLE IF NOT EXISTS row (
    rowId INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '货架排编号',
    rowName VARCHAR(20) COMMENT '货架排名称',
    layerId INTEGER COMMENT '货架层编号',
    shelfId INTEGER COMMENT '货架编号',
    passageId INTEGER COMMENT '通道编号',
    roomId INTEGER COMMENT '房间编号',
    warehouseId INTEGER COMMENT '仓库编号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`

//A货架
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 1, 1, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 1, 1, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 1, 1, 1, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 2, 1, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 2, 1, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 2, 1, 1, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 3, 1, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 3, 1, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 3, 1, 1, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 4, 1, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 4, 1, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 4, 1, 1, 1, 2, now());
//
//B货架
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 5, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 5, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 5, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 5, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('5', 5, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('6', 5, 2, 1, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 6, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 6, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 6, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 6, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('5', 6, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('6', 6, 2, 1, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 7, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 7, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 7, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 7, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('5', 7, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('6', 7, 2, 1, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 8, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 8, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 8, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 8, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('5', 8, 2, 1, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('6', 8, 2, 1, 1, 2, now());
//
//C货架
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 9, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 9, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 9, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 9, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('5', 9, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('6', 9, 3, 2, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 10, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 10, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 10, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 10, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('5', 10, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('6', 10, 3, 2, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 11, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 11, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 11, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 11, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('5', 11, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('6', 11, 3, 2, 1, 2, now());
//
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 12, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 12, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 12, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 12, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('5', 12, 3, 2, 1, 2, now());
//insert into row (rowName, layerId, shelfId, passageId, roomId, warehouseId, createdAt) values ('6', 12, 3, 2, 1, 2, now());
