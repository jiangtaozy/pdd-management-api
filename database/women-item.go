/*
 * Maintained by jemo from 2021.1.21 to now
 * Created by jemo on 2021.1.21 17:02:28
 * Women Item
 * 女装网商品
 */

package database

const createWomenItem =`
  CREATE TABLE IF NOT EXISTS womenItem (
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    isLightningDelivery BOOLEAN COMMENT '是否闪电发货',
    isCloudWarehouse BOOLEAN COMMENT '是否云仓',
    isWechatMerchant BOOLEAN COMMENT '是否微商',
    isBigImg BOOLEAN COMMENT '是否精修大图',
    isHotSale BOOLEAN COMMENT '是否热卖',
    title VARCHAR(100) COMMENT '标题',
    keywords VARCHAR(50) COMMENT '关键词',
    price INTEGER UNSIGNED COMMENT '价格(分)',
    presentPrice INTEGER UNSIGNED COMMENT '建议价格(分)',
    goodsNo VARCHAR(20) COMMENT '货号',
    createdTime DATE COMMENT '创建日期',
    updatedTime DATE COMMENT '更新日期',
    totalSale INTEGER UNSIGNED COMMENT '总销量',
    showCount INTEGER UNSIGNED COMMENT '展示量',
    weight INTEGER UNSIGNED COMMENT '重量(克)',
    deliveryRate FLOAT UNSIGNED COMMENT '48小时发货率(%)',
    deliveryTime INTEGER UNSIGNED COMMENT '平均发货时效(小时)',
    productId INTEGER UNSIGNED COMMENT '女装网商品id',
    isReturn BOOLEAN COMMENT '是否8天包退',
    isOriginalImage BOOLEAN COMMENT '是否原图保证',
    isPowerMerchant BOOLEAN COMMENT '是否实力商家',
    isFactory BOOLEAN COMMENT '是否工厂认证',
    isLimitPrice BOOLEAN COMMENT '是否厂商控价',
    video VARCHAR(150) COMMENT '视频'
  );
`
