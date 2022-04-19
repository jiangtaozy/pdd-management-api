/*
 * Maintained by jemo from 2022.4.18 to now
 * Created by jemo on 2022.4.18 10:48:30
 * 供应商列表
 */

package database

const supplier = `
  CREATE TABLE IF NOT EXISTS supplier (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    memberId VARCHAR(30) NOT NULL DEFAULT '',
    creditLevel TINYINT UNSIGNED NOT NULL DEFAULT 0,
    shopRepurchaseRate FLOAT UNSIGNED NOT NULL DEFAULT 0,
    province VARCHAR(20) NOT NULL DEFAULT '',
    city VARCHAR(60) NOT NULL DEFAULT '',
    url VARCHAR(150) NOT NULL DEFAULT '',
    siteType TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1: 1688, 2: hznzcn',
    mallName VARCHAR(20),
    floor TINYINT,
    stallNumber VARCHAR(5),
    phone VARCHAR(11),
    telephone VARCHAR(13),
    wechat VARCHAR(20),
    qq VARCHAR(12),
    dataUrl VARCHAR(30)
  );
`
