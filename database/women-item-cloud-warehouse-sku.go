/*
 * Maintained by jemo from 2021.2.11 to now
 * Created by jemo on 2021.2.11 15:44:29
 * Women Item Cloud Warehouse Sku
 * 女装网云仓SKU
 */

package database

const createWomenItemCloudWarehouseSku =`
  CREATE TABLE IF NOT EXISTS womenItemCloudWarehouseSku (
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    productId INTEGER UNSIGNED COMMENT '女装网商品id',
    skuDesc VARCHAR(30) NOT NULL COMMENT 'sku，如红色,M',
    ycAvailNum INTEGER UNSIGNED COMMENT '云仓库存标志',
    ycStockTips VARCHAR(30) NOT NULL COMMENT '云仓库存描述，如供应商补货中',
    skuColor VARCHAR(20) NOT NULL COMMENT 'sku颜色，如红色',
    skuSize VARCHAR(20) NOT NULL COMMENT 'sku尺码，如M'
  );
`
