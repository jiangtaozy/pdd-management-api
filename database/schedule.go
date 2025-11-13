/*
 * Maintained by jemo from 2025.09.23 to now
 * Created by jemo on 2025.09.23 09:32:01
 * 公司待办事项
 */
package database

const schedule = `
  CREATE TABLE IF NOT EXISTS schedule (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    supplierId VARCHAR(20) COMMENT 'temu店铺id',
    skcCode VARCHAR(20) COMMENT 'skc货号',
    skuCode VARCHAR(20) COMMENT 'sku货号',
    chargePersonId INT UNSIGNED COMMENT '负责人ID',
    content VARCHAR(300) COMMENT '内容',
    isCompleted BOOLEAN DEFAULT 0 COMMENT '是否完成',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
 `
