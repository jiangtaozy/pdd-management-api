/*
 * Maintained by jemo from 2024.3.29 to now
 * Created by jemo on 2024.3.29 14:18:09
 * 员工表
 */

package database

const employee = `
  CREATE TABLE IF NOT EXISTS employee (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL COMMENT '姓名',
    entryDate DATE NOT NULL COMMENT '入职日期',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
//insert into employee (name, entryDate) values ('符江涛', '2019-08-02');
//insert into employee (name, entryDate) values ('符淑娟', '2023-10-01');
//insert into employee (name, entryDate) values ('屈凤彩', '2023-10-20');
//insert into employee (name, entryDate) values ('李梦瑶', '2024-03-11');
