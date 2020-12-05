/*
 * Maintained by jemo from 2020.11.25 to now
 * Created by jemo on 2020.11.25 19:48:17
 * Douyin Item
 */

package database

const createDyItem = `
  CREATE TABLE IF NOT EXISTS dyItem (
    checkStatus TINYINT UNSIGNED NOT NULL COMMENT '商品审核状态：1未提审 2审核中 3审核通过 4审核驳回 5封禁',
    cosRatio TINYINT UNSIGNED,
    createTime DATETIME COMMENT '创建时间',
    description VARCHAR(4000) COMMENT '详情',
    discountPrice INTEGER UNSIGNED COMMENT '折扣价',
    extra VARCHAR(300),
    firstCid INTEGER UNSIGNED,
    img VARCHAR(200) COMMENT '主图',
    marketPrice INTEGER UNSIGNED COMMENT '市场价',
    mobile VARCHAR(11) COMMENT '手机号',
    name VARCHAR(60) COMMENT '标题',
    outProductId INTEGER UNSIGNED COMMENT '外部Id',
    payType TINYINT UNSIGNED COMMENT '支付类型',
    productId BIGINT UNSIGNED,
    productIdStr VARCHAR(20) COMMENT '商品Id',
    recommendRemark VARCHAR(100) COMMENT '推荐语',
    secondCid INTEGER UNSIGNED,
    settlementPrice INTEGER UNSIGNED COMMENT '结算价格',
    specId INTEGER UNSIGNED,
    status TINYINT UNSIGNED COMMENT '商品上下架状态：0上架 1下架',
    thirdCid INTEGER UNSIGNED,
    updateTime DATETIME COMMENT '更新时间',
    usp VARCHAR(100)
  );
`
