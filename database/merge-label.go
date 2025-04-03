/*
 * Maintained by jemo from 2024.8.26 to now
 * Created by jemo on 2024.8.26 19:29:14
 * 标签，实拍图，合并标签
 */

package database

const mergeLabel = `
  CREATE TABLE IF NOT EXISTS mergeLabel (
    spuId VARCHAR(20),
    spuName VARCHAR(200),
    materialImgUrl VARCHAR(300),
    isSameSku TINYINT UNSIGNED,
    uploadStatus TINYINT UNSIGNED,
    checkTypeList VARCHAR(300),
    ruleNameList VARCHAR(3000),
    goodsId VARCHAR(30),
    labelFile VARCHAR(30),
    createdAtTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAtTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(goodsId)
  );
`
//alter table mergeLabel modify column ruleNameList VARCHAR(3000);
