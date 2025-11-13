/*
 * Maintained by jemo from 2025.08.22 to now
 * Created by jemo on 2025.08.22 16:07:13
 * 拼多多降价记录
 * 本地库存大于0，拼多多7天销量为0，temu 7天销量为0，上架30天
 */

package database

const createPddReduceRecord =`
  CREATE TABLE IF NOT EXISTS pddReduceRecord (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    pddId VARCHAR(20),
    activityGroupPrice INTEGER UNSIGNED COMMENT '活动价，分',
    groupPrice INTEGER UNSIGNED COMMENT '拼单价，分',
    outSkuSn VARCHAR(100) COMMENT '外部sku编码',
    skuId BIGINT UNSIGNED NOT NULL,
    spec VARCHAR(200) COMMENT 'sku描述，碎花 S',
    skuThumbUrl VARCHAR(200) COMMENT '图片地址',
    isReduced BOOLEAN DEFAULT 0 COMMENT '是否已降价',
    supplierId VARCHAR(20),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
