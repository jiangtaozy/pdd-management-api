/*
 * Maintained by jemo from 2024.4.24 to now
 * Created by jemo on 2024.4.24 09:51:01
 * 每周上新统计
 */
package database

const weekNewStatistics = `
  CREATE TABLE IF NOT EXISTS weekNewStatistics (
    week VARCHAR(7) NOT NULL,
    day DATE NOT NULL,
    quantity INTEGER COMMENT '每周上新数量',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(day)
  );
 `
