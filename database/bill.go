/*
 * Maintained by jemo from 2020.10.12 to now
 * Created by jemo on 2020.10.12 17:01:54
 * Bill
 * 货款明细
 */

package database

const bill = `
  CREATE TABLE IF NOT EXISTS bill (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY,
    amount INTEGER NOT NULL COMMENT '收支金额(分)',
    classId TINYINT NOT NULL COMMENT '财务类型，1:交易收入, 2:优惠券结算, 3:退款, 5:技术服务费, 7:扣款, 8:其他, 9:多多进宝, 10: 转账, 11: 其他软件服务',
    createdAt DATETIME NOT NULL COMMENT '入帐时间',
    extraInfo VARCHAR(20),
    financeId TINYINT NOT NULL,
    flowId BIGINT,
    mallId INTEGER UNSIGNED NOT NULL,
    note VARCHAR(50) COMMENT '备注',
    orderSn VARCHAR(32) COMMENT '商户订单号',
    relatedId BIGINT,
    shouldNotShowOrderSnInMms BOOLEAN,
    type INTEGER
  );
`
