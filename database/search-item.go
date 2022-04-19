/*
 * Maintained by jemo from 2022.4.18 to now
 * Created by jemo on 2022.4.18 10:51:31
 * 搜索商品列表
 */

package database

const searchItem = `
  CREATE TABLE IF NOT EXISTS searchItem (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL
  );
 `
