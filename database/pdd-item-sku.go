/*
 * Maintained by jemo from 2021.2.11 to now
 * Created by jemo on 2021.2.11 16:43:57
 * Pdd item sku
 */

package database


const createPddItemSku =`
  CREATE TABLE IF NOT EXISTS pddItemSku (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    pddId BIGINT UNSIGNED NOT NULL,
    outGoodsSn VARCHAR(20),
    activityGroupPrice INTEGER UNSIGNED COMMENT '活动价，分',
    groupPrice INTEGER UNSIGNED COMMENT '拼单价，分',
    isOnsale BOOLEAN COMMENT '是否在售',
    normalPrice INTEGER UNSIGNED COMMENT '单买价，分',
    outSkuSn VARCHAR(20) COMMENT '外部sku编码',
    skuId BIGINT UNSIGNED NOT NULL,
    skuQuantity INTEGER UNSIGNED COMMENT '库存',
    skuSoldQuantity INTEGER UNSIGNED COMMENT '已售',
    spec VARCHAR(200) COMMENT 'sku描述，碎花 S',
    specColor VARCHAR(100) COMMENT 'sku颜色，碎花',
    specSize VARCHAR(100) COMMENT 'sku尺码，S',
    groupSku VARCHAR(1000) COMMENT '组合sku',
    isDeleted BOOLEAN DEFAULT 0 COMMENT '是否已删除',
    skuThumbUrl VARCHAR(200) COMMENT '图片地址'
  );
`

//alter table pddItemSku modify column spec VARCHAR(200) COMMENT 'sku描述，碎花 S';
//alter table pddItemSku modify column specColor VARCHAR(100) COMMENT 'sku颜色，碎花';
//alter table pddItemSku modify column specSize VARCHAR(100) COMMENT 'sku尺码，S';
