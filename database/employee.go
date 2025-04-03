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
    isOn BOOLEAN DEFAULT true COMMENT '是否在职',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
//insert into employee (name, entryDate) values ('符江涛', '2019-08-02');
//insert into employee (name, entryDate) values ('符淑娟', '2023-10-01');
//insert into employee (name, entryDate) values ('屈凤彩', '2023-10-20');
//insert into employee (name, entryDate) values ('李梦瑶', '2024-03-11');
//insert into employee (name, entryDate) values ('庄娅琪', '2024-09-02');
//insert into employee (name, entryDate) values ('钟朝林', '2024-09-27');
//insert into employee (name, entryDate) values ('刘锦颐', '2024-11-01');

//alter table employee modify isOn BOOLEAN DEFAULT true COMMENT '是否在职';
//insert into employee (name, entryDate) values ('杨柳青', '2025-03-03');
//insert into employee (name, entryDate) values ('刘梦梦', '2025-03-03');
//insert into employee (name, entryDate) values ('黄丹莹', '2025-03-04');
