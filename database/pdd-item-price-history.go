/*
 * Maintained by jemo from 2021.1.9 to now
 * Created by jemo on 2021.1.9 17:02:18
 * Pdd item price history
 * 拼多多商品价格历史
 */

package database

const createPddItemPriceHistory = `
  CREATE TABLE IF NOT EXISTS pddItemPriceHistory (
    pddId BIGINT UNSIGNED NOT NULL COMMENT '商品id',
    date DATE NOT NULL COMMENT '记录日期',
    skuGroupPriceMin INTEGER UNSIGNED NOT NULL COMMENT '单位：分',
    skuGroupPriceMax INTEGER UNSIGNED NOT NULL COMMENT '单位：分',
    INDEX pddItemPriceHistoryIndex (pddId, date)
  );
`
