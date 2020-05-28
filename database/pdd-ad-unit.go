/*
 * Maintained by jemo from 2020.5.27 to now
 * Created by jemo on 2020.5.27 9:48:36
 * pdd ad unit
 */

package database

const createPddAdUnit =`
  CREATE TABLE IF NOT EXISTS pddAdUnit (
    mallId INTEGER UNSIGNED NOT NULL,
    planId INTEGER UNSIGNED NOT NULL,
    adId INTEGER UNSIGNED NOT NULL,
    adName VARCHAR(80) NOT NULL,
    goodsId BIGINT UNSIGNED NOT NULL,
    goodsName VARCHAR(60) NOT NULL,
    scenesType TINYINT UNSIGNED NOT NULL COMMENT '0: 多多搜索，1: 聚焦展位，2: 多多场景'
  );
`
