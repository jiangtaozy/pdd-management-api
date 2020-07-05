/*
 * Maintained by jemo from 2020.7.5 to now
 * Created by jemo on 2020.7.5 11:12:02
 * 1688 order
 */

package database

const order1688 =`
  CREATE TABLE IF NOT EXISTS order1688 (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    orderId VARCHAR(20) NOT NULL COMMENT '1688 订单号',
    sellerCompany VARCHAR(60) NOT NULL COMMENT '卖家公司名',
    totalPrice DECIMAL(10,2) NOT NULL COMMENT '货品总价(元)',
    shippingFare DECIMAL(10,2) NOT NULL COMMENT '运费(元)',
    discount DECIMAL(10,2) NOT NULL COMMENT '涨价或折扣(元)',
    actualPayment DECIMAL(10,2) NOT NULL COMMENT '实付款(元)',
    orderStatus VARCHAR(30) NOT NULL COMMENT '订单状态',
    orderCreatedTime DATETIME NOT NULL COMMENT '订单创建时间',
    orderPaymentTime DATETIME NOT NULL COMMENT '订单付款时间',
    receiver VARCHAR(30) NOT NULL COMMENT '收货人姓名',
    shippingAddress VARCHAR(100) NOT NULL COMMENT '收货地址',
    postcode VARCHAR(6) NOT NULL COMMENT '邮编',
    phone VARCHAR(11) NOT NULL COMMENT '联系手机',
    productTitle VARCHAR(150) NOT NULL COMMENT '货品标题',
    price DECIMAL(10,2) NOT NULL COMMENT '单价(元)',
    amount INTEGER UNSIGNED NOT NULL COMMENT '数量',
    courierCompany VARCHAR(30) COMMENT '物流公司',
    trackingNumber VARCHAR(30) COMMENT '运单号'
  );
`
