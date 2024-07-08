/*
 * Maintained by jemo from 2023.11.22 to now
 * Created by jemo on 2023.11.22 15:22:13
 * 货架列表
 */

package database

const createShelf =`
  CREATE TABLE IF NOT EXISTS shelf (
    shelfId INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '货架编号',
    shelfName VARCHAR(20) COMMENT '货架名称',
    passageId INTEGER COMMENT '通道编号',
    roomId INTEGER COMMENT '房间编号',
    warehouseId INTEGER COMMENT '仓库编号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`

//insert into shelf (shelfName, passageId, roomId, warehouseId, createdAt) values ('A', 1, 1, 2, now());
//insert into shelf (shelfName, passageId, roomId, warehouseId, createdAt) values ('B', 1, 1, 2, now());
//insert into shelf (shelfName, passageId, roomId, warehouseId, createdAt) values ('C', 2, 1, 2, now());
//insert into shelf (shelfName, passageId, roomId, warehouseId, createdAt) values ('A', 3, 2, 1, now());
//insert into shelf (shelfName, passageId, roomId, warehouseId, createdAt) values ('B', 4, 3, 1, now());
//insert into shelf (shelfName, passageId, roomId, warehouseId, createdAt) values ('C', 5, 4, 1, now());
