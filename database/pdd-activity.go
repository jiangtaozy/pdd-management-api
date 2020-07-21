/*
 * Maintained by jemo from 2020.7.12 to now
 * Created by jemo on 2020.7.12 17:03:11
 * pdd activity
 */

package database

const createPddActivity =`
  CREATE TABLE IF NOT EXISTS pddActivity (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    activityId INTEGER UNSIGNED NOT NULL COMMENT '活动id',
    activityName VARCHAR(30) NOT NULL COMMENT '活动名称',
    activityType TINYINT UNSIGNED NOT NULL COMMENT '活动类型，3：限量',
    goodsId BIGINT UNSIGNED NOT NULL COMMENT '商品id',
    goodsName VARCHAR(60) NOT NULL COMMENT '商品名称',
    hdThumbUrl VARCHAR(200) NOT NULL DEFAULT '',
    minOnSaleGroupPrice INTEGER UNSIGNED NOT NULL COMMENT '销售最低价(分)',
    maxOnSaleGroupPrice INTEGER UNSIGNED NOT NULL COMMENT '销售最高价(分)',
    onlineQuantity INTEGER UNSIGNED NOT NULL COMMENT '数量',
    maxPreSalePrice INTEGER UNSIGNED NOT NULL COMMENT '活动最高价(分)',
    minPreSalePrice INTEGER UNSIGNED NOT NULL COMMENT '活动最低价(分)',
    discount INTEGER UNSIGNED NOT NULL COMMENT '折扣(千分之)',
    activityQuantity INTEGER UNSIGNED NOT NULL COMMENT '活动商品数量',
    activityStockQuantity INTEGER UNSIGNED NOT NULL COMMENT '活动优惠券数量',
    startTime TIMESTAMP NOT NULL COMMENT '开始时间(时间戳)',
    endTime TIMESTAMP NOT NULL COMMENT '结束时间(时间戳)',
    status TINYINT UNSIGNED NOT NULL COMMENT '状态，2：活动中',
    endOperationTime TIMESTAMP NOT NULL COMMENT '操作结束时间(时间戳)',
    newGoods BOOLEAN NOT NULL COMMENT '是否新品'
  );
`
