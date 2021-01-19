/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 11:39:06
 * Pdd Competitor
 * 拼多多竞争对手
 */

package database

const createPddCompetitor = `
  CREATE TABLE IF NOT EXISTS pddCompetitor (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL COMMENT '店铺名称',
    advantage VARCHAR(100) COMMENT '优点',
    disadvantage VARCHAR(100) COMMENT '缺点'
  );
`
