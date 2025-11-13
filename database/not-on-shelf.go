/*
 * Maintained by jemo from 2025.09.03 to now
 * Created by jemo on 2025.09.03 17:46:26
 * 已选款，未上架
 */

package database

const createNotOnShelf =`
  CREATE TABLE IF NOT EXISTS notOnShelf (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    supplierId VARCHAR(20),
    skcCode VARCHAR(20) NOT NULL unique COMMENT 'skc货号',
    name VARCHAR(200),
    isOnShelf BOOLEAN DEFAULT 0 COMMENT '是否已上架',
    isNotToShelf BOOLEAN DEFAULT 0 COMMENT '是否不上架',
    isReceivedSample BOOLEAN DEFAULT 0 COMMENT '是否已收到样品',
    selectedAt DATETIME COMMENT '选款时间',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
