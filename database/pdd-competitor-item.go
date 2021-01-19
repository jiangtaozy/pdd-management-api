/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 15:11:59
 * Pdd Competitor Item
 * 拼多多竞争对手商品
 */

package database

const createPddCompetitorItem = `
  CREATE TABLE IF NOT EXISTS pddCompetitorItem (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT '商品名称',
    price INTEGER UNSIGNED NOT NULL COMMENT '单位：分',
    goodsId VARCHAR(12) NOT NULL COMMENT '商品id',
    competitorId INTEGER UNSIGNED NOT NULL COMMENT '竞争对手id',
    relatedItemId BIGINT UNSIGNED NOT NULL COMMENT '关联商品id'
  );
`
