/*
 * Maintained by jemo from 2024.3.29 to now
 * Created by jemo on 2024.3.29 14:56:18
 * 工作记录每天记录
 */

package database

const workRecord = `
  CREATE TABLE IF NOT EXISTS workRecord (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    employeeId INT NOT NULL COMMENT '员工id',
    recordDate DATE NOT NULL COMMENT '日期',
    recordContent TEXT COMMENT '记录内容',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
//insert into workRecord (employeeId, recordDate, recordContent) values ('1', '2024-03-29', '内容');
//insert into workRecord (employeeId, recordDate, recordContent) values ('1', '2024-03-28', '内容');
