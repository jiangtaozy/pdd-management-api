/*
 * Maintained by jemo from 2022.11.09 to now
 * Created by jemo on 2022.11.09 11:43:48
 * 阿里1688商品SKU规则如颜色、尺码等
 */

package database

const createItemSkuProp =`
  CREATE TABLE IF NOT EXISTS itemSkuProp (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    searchId INTEGER UNSIGNED NOT NULL COMMENT '商品id',
    propName VARCHAR(30) NOT NULL COMMENT '规格名称，如颜色',
    propValue VARCHAR(8000) NOT NULL COMMENT '规格内容，如[imageUrl, name]',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    unique key propNameUniqueKey (searchId, propName)
  );
`
