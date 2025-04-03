/*
 * Maintained by jemo from 2025.02.14 to now
 * Created by jemo on 2025.02.14 15:28:38
 * 欠货记录
 */

package database

const createShortDelivery =`
  CREATE TABLE IF NOT EXISTS shortDelivery (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    orderId VARCHAR(20) NOT NULL COMMENT '1688订单号',
    skcCode VARCHAR(10) COMMENT 'sku货号',
    skuCode VARCHAR(10) COMMENT 'skc货号',
    shortAmount INTEGER UNSIGNED COMMENT '欠货数量',
    shortQuantity INTEGER UNSIGNED COMMENT '欠货件数，2件套则为欠货数量除以2',
    courierCompany VARCHAR(200) COMMENT '补发物流公司',
    trackingNumber VARCHAR(200) COMMENT '补发运单号',
    remark VARCHAR(200) COMMENT '备注',
    isReissue BOOLEAN COMMENT '是否已补发',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
