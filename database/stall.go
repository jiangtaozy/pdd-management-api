/*
 * Maintained by jemo from 2020.5.28 to now
 * Created by jemo on 2020.5.28 17:34:30
 * Stall
 * 档口
 */

package database

const createStall =`
  CREATE TABLE IF NOT EXISTS stall (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    cityName VARCHAR(20) NOT NULL DEFAULT '',
    mallName VARCHAR(20) NOT NULL DEFAULT '',
    floor TINYINT NOT NULL DEFAULT 1,
    stallNumber VARCHAR(5) DEFAULT '',
    phone VARCHAR(11) NOT NULL DEFAULT '',
    telephone VARCHAR(13) NOT NULL DEFAULT '',
    wechat VARCHAR(20) NOT NULL DEFAULT '',
    qq VARCHAR(12) NOT NULL DEFAULT '',
    dataUrl VARCHAR(30) NOT NULL DEFAULT '',
    stallUrl VARCHAR(50) DEFAULT ''
  );
`
