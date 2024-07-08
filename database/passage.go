/*
 * Maintained by jemo from 2023.11.22 to now
 * Created by jemo on 2023.11.22 14:57:46
 * 通道列表
 */

package database

const createPassage =`
  CREATE TABLE IF NOT EXISTS passage (
    passageId INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '通道编号',
    passageName VARCHAR(20) COMMENT '通道名称',
    roomId INTEGER COMMENT '房间编号',
    warehouseId INTEGER COMMENT '仓库编号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`

//insert into passage (passageName, roomId, warehouseId, createdAt) values ('门口通道', 1, 2, now());
//insert into passage (passageName, roomId, warehouseId, createdAt) values ('靠窗通道', 1, 2, now());
//insert into passage (passageName, roomId, warehouseId, createdAt) values ('1通道', 2, 1, now());
//insert into passage (passageName, roomId, warehouseId, createdAt) values ('1通道', 3, 1, now());
//insert into passage (passageName, roomId, warehouseId, createdAt) values ('1通道', 4, 1, now());
