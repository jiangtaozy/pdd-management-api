/*
 * Maintained by jemo from 2020.7.14 to now
 * Created by jemo on 2020.7.14 12:45:36
 * Ad Head
 */

package database

const createAdHead =`
  CREATE TABLE IF NOT EXISTS adHead (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    headId INTEGER UNSIGNED NOT NULL COMMENT '团长id',
    dodokCommission INTEGER UNSIGNED COMMENT '多多客佣金%',
    headCommission INTEGER UNSIGNED COMMENT '团长佣金%',
    coupon INTEGER UNSIGNED COMMENT '优惠券金额',
    wechatNickname VARCHAR(32) COMMENT '微信昵称',
    wechatNumber VARCHAR(20) COMMENT '微信号',
    pddNickname VARCHAR(32) COMMENT '拼多多昵称'
  );
`
