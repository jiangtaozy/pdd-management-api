/*
 * Maintained by jemo from 2023.4.5 to now
 * Created by jemo on 2023.4.5 11:51:25
 * merge order 1688
 */

package database

/*
| mergeOrder1688 | CREATE TABLE `mergeOrder1688` (
  `id`           bigint(20)    unsigned             NOT     NULL    AUTO_INCREMENT,
  `orderId`      varchar(20)   NOT                  NULL    COMMENT '1688订单号',
  `productTitle` varchar(150)  DEFAULT              NULL    COMMENT '货品标题',
  `price`        decimal(10,2) DEFAULT              NULL    COMMENT '单价(元)',
  `amount`       int(10)       unsigned             DEFAULT NULL    COMMENT '数量',
  `skuId`        varchar(30)   DEFAULT              NULL    COMMENT 'SKU    ID',
  `offerId`      varchar(30)   DEFAULT              NULL    COMMENT 'Offer  ID',
  PRIMARY        KEY           (`id`),
  UNIQUE         KEY           `orderIdSkuIdUnique` (`orderId`,`skuId`)
) ENGINE=InnoDB AUTO_INCREMENT=2102 DEFAULT CHARSET=utf8mb4             |
*/
