/*
 * Maintained by jemo from 2024.5.27 to now
 * Created by jemo on 2024.5.27 16:22:34
 * 机会商品
 */

package database

const chanceItem = `
  CREATE TABLE IF NOT EXISTS chanceItem (
    goodsId VARCHAR(200),
    creatorName VARCHAR(100),
    description VARCHAR(1000),
    type TINYINT UNSIGNED,
    createdAt DATETIME,
    bidingInvitationOrderId VARCHAR(200),
    imageUrlList TEXT,
    startTime DATETIME,
    catIdList VARCHAR(200),
    imageWatermarkCheckStatus TINYINT UNSIGNED,
    requirement VARCHAR(200),
    catNameList VARCHAR(300),
    name VARCHAR(300),
    endTime DATETIME,
    isSelect TINYINT UNSIGNED COMMENT '是否选款-0: 不选款, 1: 选款',
    isFind TINYINT UNSIGNED COMMENT '是否找到同款-0: 未找到, 1: 找到',
    code VARCHAR(20) COMMENT '货号',
    createdAtTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAtTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    unique key goodsIdBindingOrderUniqueKey (goodsId, bidingInvitationOrderId)
  );
`
//alter table chanceItem modify column code VARCHAR(20) COMMENT '货号';
