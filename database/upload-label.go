/*
 * Maintained by jemo from 2024.8.29 to now
 * Created by jemo on 2024.8.29 10:20:04
 * 上传标签
 */

package database

const uploadLabel = `
  CREATE TABLE IF NOT EXISTS uploadLabel (
    id INTEGER UNSIGNED NOT NULL UNIQUE KEY AUTO_INCREMENT,
    originalName VARCHAR(200) PRIMARY KEY,
    suffix VARCHAR(10),
    fileName VARCHAR(30),
    createdAtTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAtTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
