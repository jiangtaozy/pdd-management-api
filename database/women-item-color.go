/*
 * Maintained by jemo from 2021.1.30 to now
 * Created by jemo on 2021.1.30 16:25:17
 * Women Item Color
 * 女装网商品颜色
 */

package database

const createWomenItemColor =`
  CREATE TABLE IF NOT EXISTS womenItemColor (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    productId INTEGER UNSIGNED COMMENT '女装网商品id',
    color VARCHAR(36) NOT NULL COMMENT '颜色',
    thumbnail VARCHAR(150) COMMENT '颜色图225*225',
    hrthumbnail VARCHAR(150) COMMENT '颜色图500*500',
    original VARCHAR(150) COMMENT '颜色图800*800'
  );
`
