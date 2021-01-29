/*
 * Maintained by jemo from 2021.1.29 to now
 * Created by jemo on 2021.1.29 17:42:05
 * Women Item Detail Image
 * 女装网商品详情图
 */

package database

const createWomenItemDetailImage =`
  CREATE TABLE IF NOT EXISTS womenItemDetailImage (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    productId INTEGER UNSIGNED COMMENT '女装网商品id',
    src VARCHAR(150) NOT NULL COMMENT '详情图'
  );
`

