/*
 * Maintained by jemo from 2021.1.26 to now
 * Created by jemo on 2021.1.26 17:02:32
 * Women Item SKU
 * 女装网商品SKU
 */

package database

const createWomenItemSku =`
  CREATE TABLE IF NOT EXISTS womenItemSku (
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    productId INTEGER UNSIGNED COMMENT '女装网商品id',
    skuDesc VARCHAR(30) NOT NULL COMMENT 'sku描述，如红色,M',
    skuKey VARCHAR(30) NOT NULL COMMENT 'sku关键字，如YYG82590-5',
    price INTEGER UNSIGNED COMMENT 'sku价格，单位：分',
    isOnShelf BOOLEAN COMMENT '是否上架',
    stock INTEGER UNSIGNED COMMENT 'sku库存'
  );
`
