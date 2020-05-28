/*
 * Maintained by jemo from 2020.5.26 to now
 * Created by jemo on 2020.5.26 17:41:57
 * pdd ad plan
 */

package database

const createPddAdPlan =`
  CREATE TABLE IF NOT EXISTS pddAdPlan (
    mallId INTEGER UNSIGNED NOT NULL,
    planId INTEGER UNSIGNED NOT NULL,
    planName VARCHAR(64) NOT NULL,
    stickTime DATETIME,
    isStick BOOLEAN NOT NULL DEFAULT false,
    scenesType TINYINT UNSIGNED NOT NULL COMMENT '0: 多多搜索，1: 聚焦展位，2: 多多场景'
  );
`
