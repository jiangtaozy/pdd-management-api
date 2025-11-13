/*
 * Maintained by jemo from 2025.06.23 to now
 * Created by jemo on 2025.06.23 15:18:46
 * 亚马逊发货单
 */

package database

const createAmazonDelivery =`
  CREATE TABLE IF NOT EXISTS amazonDelivery (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '发货单编号',
    sku VARCHAR(10) COMMENT 'sku货号',
    asin VARCHAR(15) COMMENT 'asin',
    fnSku VARCHAR(15) COMMENT '条码id',
    quantity INTEGER COMMENT '发货数量',
    isShip TINYINT UNSIGNED COMMENT '是否已发货',
    isStockout TINYINT UNSIGNED COMMENT '是否缺货，不发',
    createdAtTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAtTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`


//alter table amazonDelivery add isStockout TINYINT UNSIGNED COMMENT '是否缺货，不发';

//select * from amazonDelivery where sku = '28638-01'\G;
//update amazonDelivery set isStockout = 1 where sku = '28638-01';

//select * from amazonDelivery where sku = '26996-04'\G;
