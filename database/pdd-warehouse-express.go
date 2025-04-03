/*
 * Maintained by jemo from 2025.02.06 to now
 * Created by jemo on 2025.02.06 16:49:48
 * 拼多多仓库对应快递网点
 */

package database

const createPddWarehouseExpress =`
  CREATE TABLE IF NOT EXISTS pddWarehouseExpress (
    warehouseId INTEGER UNSIGNED COMMENT '仓库编号',
    warehouseName VARCHAR(20) COMMENT '仓库名称',
    shipName VARCHAR(50) COMMENT '快递公司/亚马逊发货人',
    branchName VARCHAR(100) COMMENT '网点名称/亚马逊发货地址',
    PRIMARY KEY(warehouseId, branchName)
  );
`

//insert into pddWarehouseExpress (warehouseId, warehouseName, shipName, branchName) values (1, '闫家庄', '韵达快递', '河北定州市公司明月店镇寄存点分部');
//insert into pddWarehouseExpress (warehouseId, warehouseName, shipName, branchName) values (2, '海智中心', '极兔速递', '杭州余杭海创园网点');


//insert into pddWarehouseExpress (warehouseId, warehouseName, shipName, branchName) values (2, '海智中心', '符江涛', '仓前街道 海智中心3幢406室');
//insert into pddWarehouseExpress (warehouseId, warehouseName, shipName, branchName) values (2, '海智中心', '中通快递', '杭州未来科技城');
//insert into pddWarehouseExpress (warehouseId, warehouseName, shipName, branchName) values (2, '海智中心', '申通快递', '浙江杭州余杭区科技城公司');
