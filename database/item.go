/*
 * Maintained by jemo from 2022.4.18 to now
 * Created by jemo on 2022.4.18 10:47:24
 * 商品列表
 */
package database

const item = `
  CREATE TABLE IF NOT EXISTS item (
    name VARCHAR(200) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    imgUrl VARCHAR(200) NOT NULL DEFAULT '',
    detailUrl VARCHAR(60) NOT NULL DEFAULT '',
    siteType TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1: 1688, 2: hznzcn, 3: 线下微信',
    originalId VARCHAR(30) NOT NULL DEFAULT '',
    supplierId INTEGER UNSIGNED NOT NULL DEFAULT 0,
    saleQuantity INTEGER UNSIGNED NOT NULL DEFAULT 0,
    quantitySumMonth INTEGER UNSIGNED NOT NULL DEFAULT 0,
    gmv30dRt DECIMAL(12,2) UNSIGNED NOT NULL DEFAULT 0,
    searchId BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    suitPrice DECIMAL(10,2) NOT NULL DEFAULT 0,
    shippingPrice DECIMAL(10,2) NOT NULL DEFAULT 0,
    forSell BOOLEAN NOT NULL DEFAULT false,
    imgUrlOf290x290 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf120x120 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf270x270 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf100x100 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf150x150 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf220x220 VARCHAR(100) NOT NULL DEFAULT '',
    womenProductId INTEGER UNSIGNED,
    keyName VARCHAR(200),
    itemTypeKey INTEGER COMMENT '类型key',
    itemNum VARCHAR(20) COMMENT '编码',
    length DECIMAL(10,2) UNSIGNED COMMENT '长(cm)',
    width DECIMAL(10,2) UNSIGNED COMMENT '宽(cm)',
    height DECIMAL(10,2) UNSIGNED COMMENT '高(cm)',
    weight INTEGER UNSIGNED COMMENT '重量(g)',
    crossBorderUrl VARCHAR(100) COMMENT '跨境平台地址',
    notForSellReason VARCHAR(300) COMMENT '下架原因',
    declaredPrice DECIMAL(10,2) COMMENT '申报价格',
    remark VARCHAR(1000) COMMENT '备注',
    INDEX forSellIndex(forSell),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
 `

//insert into item (name, price) values ('卡通随机多种图案补丁贴', 1);
//insert into item (name, price) values ('红色小米圆珠大颗黄色切面珠细锁骨项链', 9.2);
//insert into item (name, price) values ('金如意铃铛吊坠茶色切面小米珠锁骨链', 14.2);
//insert into item (name, price) values ('绿松石小米珠天然大颗碎石金属隔片锁骨链', 10);
//insert into item (name, price) values ('十字架吊坠绿松石小米珠天然大颗碎石金属隔片锁骨链', 13.5);
//insert into item (name, price) values ('马卡龙白圆柱石混色串珠弹力锁骨链', 14);
//insert into item (name, price) values ('红色切面小米珠金色圆隔珠不规则隔片细锁骨链', 3.2);
//insert into item (name, price) values ('黑色切面小米珠碎金银锁骨链', 6);

//alter table item modify column remark VARCHAR(1000) COMMENT '备注';
