/*
 * Maintained by jemo from 2021.1.29 to now
 * Created by jemo on 2021.1.29 17:16:22
 * Women Item Attribute
 * 女装网商品属性
 */

package database

const createWomenItemAttribute =`
  CREATE TABLE IF NOT EXISTS womenItemAttribute (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    productId INTEGER UNSIGNED COMMENT '女装网商品id',
    attributeKey VARCHAR(50) NOT NULL COMMENT '属性',
    attributeValue VARCHAR(50) NOT NULL COMMENT '属性值'
  );
`
