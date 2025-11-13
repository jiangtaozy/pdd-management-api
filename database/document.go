/*
 * Maintained by jemo from 2024.9.3 to now
 * Created by jemo on 2024.9.3 08:30:39
 * 文档
 */

package database

const document = `
  CREATE TABLE IF NOT EXISTS document (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    classification VARCHAR(50) COMMENT '文档分类',
    title VARCHAR(200) COMMENT '文档标题',
    content TEXT COMMENT '文档内容',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
//alter table document add classification VARCHAR(50) COMMENT '文档分类';
