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
    classId TINYINT NOT NULL COMMENT '财务类型',
    createdAt DATETIME NOT NULL COMMENT '入帐时间',
    extraInfo VARCHAR(20),
    financeId TINYINT NOT NULL,
    flowId BIGINT NOT NULL,
    mallId INTEGER UNSIGNED NOT NULL,
    note VARCHAR(50) COMMENT '备注',
    orderSn VARCHAR(32) COMMENT '商户订单号',
    relatedId BIGINT,
    shouldNotShowOrderSnInMms BOOLEAN,
    type INTEGER
  );
`
