/*
 * Maintained by jemo from 2024.3.4 to now
 * Created by jemo on 2024.3.4 15:51:18
 * 每日上新统计
 */
package database

const dayNewStatistics = `
  CREATE TABLE IF NOT EXISTS dayNewStatistics (
    day DATE NOT NULL,
    quantity INTEGER COMMENT '每日上新数量',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(day)
  );
 `
