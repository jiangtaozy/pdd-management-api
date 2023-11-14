/*
 * Maintained by jemo from 2023.11.9 to now
 * Created by jemo on 2023.11.9 16:51:59
 * 房间列表
 */

package database

const createRoom =`
  CREATE TABLE IF NOT EXISTS room (
    roomId INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '房间编号',
    roomName VARCHAR(20) unique COMMENT '房间名称',
    roomAddress VARCHAR(100) unique COMMENT '房间地址',
    warehouseId INTEGER COMMENT '仓库编号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`

//insert into room (roomName, roomAddress, warehouseId, createdAt)values ('905', '海智中心905室', 2, now());
