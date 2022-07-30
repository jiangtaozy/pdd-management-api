/*
 * Maintained by jemo from 2020.7.5 to now
 * Created by jemo on 2020.7.5 11:12:02
 * 1688 order
 */

package database

const order1688 =`
  CREATE TABLE IF NOT EXISTS order1688 (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    orderId VARCHAR(20) NOT NULL COMMENT '1688/女装网订单号',
    sellerCompany VARCHAR(60) COMMENT '卖家公司名',
    totalPrice DECIMAL(10,2) COMMENT '货品总价(元)',
    shippingFare DECIMAL(10,2) NOT NULL COMMENT '运费(元)',
    discount DECIMAL(10,2) COMMENT '涨价或折扣(元)',
    actualPayment DECIMAL(10,2) NOT NULL COMMENT '实付款(元)',
    orderStatus TINYINT NOT NULL COMMENT '订单状态，0: 待付款，1: 待发货，2: 待收货，3: 已收货，4: 交易成功，5: 已退换货，6: 交易关闭',
    orderStatusStr VARCHAR(20) COMMENT '订单状态描述',
    afterSaleStatusStr VARCHAR(100) COMMENT '退货状态',
    orderCreatedTime DATETIME NOT NULL COMMENT '订单创建时间',
    orderPaymentTime DATETIME COMMENT '订单付款时间',
    receiver VARCHAR(30) COMMENT '收货人姓名',
    shippingAddress VARCHAR(100) COMMENT '收货地址',
    postcode VARCHAR(6) COMMENT '邮编',
    phone VARCHAR(11) COMMENT '联系手机',
    productTitle VARCHAR(150) COMMENT '货品标题',
    price DECIMAL(10,2) COMMENT '单价(元)',
    amount INTEGER UNSIGNED COMMENT '数量',
    courierCompany VARCHAR(30) COMMENT '物流公司',
    trackingNumber VARCHAR(30) COMMENT '运单号',
    orderType TINYINT DEFAULT 0 COMMENT '订单类型，0: 1688, 1: 女装网',
    productSku VARCHAR(200) COMMENT '商品规格',
    orderTotalPrice DECIMAL(10,2) COMMENT '订单总金额',
    agentDeliveryFee DECIMAL(10,2) COMMENT '代发费',
    paymentMethod VARCHAR(10) COMMENT '付款方式',
    outerOrderId VARCHAR(20) COMMENT '外部订单号',
    deliveryTime DATETIME COMMENT '发货时间',
    productStatus VARCHAR(10) COMMENT '货物状态',
    distributionAmount INTEGER COMMENT '已配数量',
    deliveryAmount INTEGER COMMENT '已发数量',
    buyerCompany VARCHAR(50) COMMENT '买家公司名',
    buyerMember VARCHAR(50) COMMENT '买家会员名',
    sellerMember VARCHAR(50) COMMENT '卖家会员名',
    shipper VARCHAR(20) COMMENT '发货方',
    goodsType INTEGER COMMENT '货品种类',
    initiator VARCHAR(20) COMMENT '发起人登录名',
    unique orderIdUnique(orderId)
  );
`
