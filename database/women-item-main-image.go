/*
 * Maintained by jemo from 2021.1.28 to now
 * Created by jemo on 2021.1.28 17:04:47
 * Women Item Main Image
 * 女装网商品主图
 */

package database

const createWomenItemMainImage =`
  CREATE TABLE IF NOT EXISTS womenItemMainImage (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    productId INTEGER UNSIGNED COMMENT '女装网商品id',
    src VARCHAR(150) NOT NULL COMMENT '主图'
  );
`
