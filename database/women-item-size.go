/*
 * Maintained by jemo from 2021.1.30 to now
 * Created by jemo on 2021.1.30 16:57:45
 * Women Item Size
 * 女装网商品尺码
 */

package database

const createWomenItemSize =`
  CREATE TABLE IF NOT EXISTS womenItemSize (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    productId INTEGER UNSIGNED COMMENT '女装网商品id',
    size VARCHAR(20) NOT NULL COMMENT '尺码'
  );
`
