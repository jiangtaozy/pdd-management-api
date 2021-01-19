/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 16:59:13
 * Pdd Competitor Item Sale
 * 拼多多竞争对手商品销量
 */

package database

const createPddCompetitorItemSale = `
  CREATE TABLE IF NOT EXISTS pddCompetitorItemSale (
    goodsId VARCHAR(12) NOT NULL COMMENT '商品id',
    date DATE NOT NULL COMMENT '日期',
    sale INTEGER UNSIGNED NOT NULL COMMENT '销量'
  );
`
