/*
 * Maintained by jemo from 2021.2.11 to now
 * Created by jemo on 2021.2.11 16:43:57
 * Pdd item sku
 */

package database


const createPddItemSku =`
  CREATE TABLE IF NOT EXISTS pddItemSku (
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
    spec VARCHAR(20) COMMENT 'sku描述，碎花 S'
  );
`
