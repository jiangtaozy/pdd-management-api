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
    chargeContent VARCHAR(2000) COMMENT '负责工作内容',
    phone VARCHAR(15) COMMENT '手机号',
    password VARCHAR(150) COMMENT '密码-argon2',
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

//alter table employee add chargeContent VARCHAR(2000) COMMENT '负责工作内容';

//alter table employee add password VARCHAR(150) COMMENT '密码-argon2';
//alter table employee add phone VARCHAR(15) COMMENT '手机号';

//update employee set phone = '18794769375', password = '$argon2id$v=19$m=65536,t=3,p=4$3KptiF8VNjnSSTqSqxGyxw$DYNjOXFAQw5DGR0TYZHBMuCycb8iB+3/6vCETYdJnSY' where id = 1 and name = '符江涛';
