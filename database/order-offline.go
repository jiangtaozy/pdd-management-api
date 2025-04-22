/*
 * Maintained by jemo from 2023.4.25 to now
 * Created by jemo on 2023.4.25 14:39:34
 * offline order
 * 线下订单
 */

package database

const orderOffline =`
  CREATE TABLE IF NOT EXISTS orderOffline (
    orderId BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    productTitle VARCHAR(150) COMMENT '货品标题',
    skcCode VARCHAR(20) COMMENT 'skc货号',
    skuCode VARCHAR(20) COMMENT 'sku货号',
    skuName VARCHAR(300) COMMENT 'sku名称，如【香槟】钻樱桃抓夹',
    price DECIMAL(10,2) COMMENT '单价(元)',
    amount INTEGER UNSIGNED COMMENT '数量',
    totalPrice DECIMAL(10,2) COMMENT '货品总价(元)',
    shippingFare DECIMAL(10,2) COMMENT '运费(元)',
    discount DECIMAL(10,2) COMMENT '涨价或折扣(元)',
    actualPayment DECIMAL(10,2) COMMENT '实付款(元)',
    orderStatus TINYINT COMMENT '订单状态，0: 待付款，1: 待发货，2: 待收货，3: 已收货，4: 交易成功，5: 已退换货，6: 交易关闭',
    orderStatusStr VARCHAR(20) COMMENT '订单状态描述',
    courierCompany VARCHAR(200) COMMENT '物流公司',
    trackingNumber VARCHAR(200) COMMENT '运单号',
    isReceived BOOLEAN DEFAULT 0 COMMENT '是否已收货',
    receivedTime DATETIME COMMENT '收货时间',
    receivedWarehouseId INTEGER UNSIGNED COMMENT '收货仓库id',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`

//alter table orderOffline add skuName VARCHAR(300) COMMENT 'sku名称，如【香槟】钻樱桃抓夹';
