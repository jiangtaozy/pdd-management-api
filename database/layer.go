/*
 * Maintained by jemo from 2023.11.22 to now
 * Created by jemo on 2023.11.22 15:53:43
 * 层号列表
 */

package database

const createlayer =`
  CREATE TABLE IF NOT EXISTS layer (
    layerId INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '货架层编号',
    layerName VARCHAR(20) COMMENT '货架层名称',
    shelfId INTEGER COMMENT '货架编号',
    passageId INTEGER COMMENT '通道编号',
    roomId INTEGER COMMENT '房间编号',
    warehouseId INTEGER COMMENT '仓库编号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`

//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第1层', 1, 1, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第2层', 1, 1, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第3层', 1, 1, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第4层', 1, 1, 1, 2, now());
//
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第1层', 2, 1, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第2层', 2, 1, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第3层', 2, 1, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第4层', 2, 1, 1, 2, now());
//
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第1层', 3, 2, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第2层', 3, 2, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第3层', 3, 2, 1, 2, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('第4层', 3, 2, 1, 2, now());

//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('1', 7, 3, 2, 1, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('2', 7, 3, 2, 1, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('3', 7, 3, 2, 1, now());
//insert into layer (layerName, shelfId, passageId, roomId, warehouseId, createdAt) values ('4', 7, 3, 2, 1, now());

//update layer set layerName = '1' where layerId = 25;
//update layer set layerName = '2' where layerId = 26;
//update layer set layerName = '3' where layerId = 27;
//update layer set layerName = '4' where layerId = 28;
