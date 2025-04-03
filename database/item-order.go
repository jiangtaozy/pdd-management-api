/*
 * Maintained by jemo from 2022.4.18 to now
 * Created by jemo on 2022.4.18 10:49:54
 * 商品订单列表，拼多多订单
 */

package database

const itemOrder = `
  CREATE TABLE IF NOT EXISTS itemOrder (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    mallId INTEGER UNSIGNED COMMENT '店铺id',
    productName VARCHAR(60) NOT NULL COMMENT '拼多多商品',
    orderId VARCHAR(30) NOT NULL UNIQUE COMMENT '拼多多订单号',
    outerOrderId VARCHAR(30) COMMENT '外部订单号1688/女装网订单',
    orderStatus TINYINT COMMENT '拼多多订单状态，0: 待支付，1: 待发货/已发货，待签收/已签收/未发货，退款成功/已发货，退款成功，2: 已取消',
    orderStatusStr VARCHAR(20) NOT NULL COMMENT '拼多多订单状态',
    productTotalPrice INTEGER UNSIGNED NOT NULL COMMENT '拼多多商品总价(分)',
    storeDiscount INTEGER UNSIGNED NOT NULL COMMENT '店铺优惠折扣(分)',
    platformDiscount INTEGER UNSIGNED NOT NULL COMMENT '平台优惠折扣(分)',
    postage INTEGER UNSIGNED NOT NULL COMMENT '邮费(分)',
    serviceAmount INTEGER UNSIGNED COMMENT '服务费(分)',
    onsiteInstallationFee INTEGER UNSIGNED COMMENT '上门安装费(分)',
    homeDeliveryFee INTEGER UNSIGNED COMMENT '送货入户费(分)',
    homeDeliveryAndInstallationFee INTEGER UNSIGNED COMMENT '送货入户并安装费(分)',
    userPaidAmount INTEGER UNSIGNED NOT NULL COMMENT '用户实付金额(分)',
    merchantReceiveAmount INTEGER UNSIGNED COMMENT '商家实收金额(分)',
    receiver VARCHAR(60) COMMENT '收货人',
    phone VARCHAR(11) COMMENT '手机',
    province VARCHAR(20) COMMENT '省',
    city VARCHAR(60) COMMENT '市',
    district VARCHAR(60) COMMENT '区',
    street VARCHAR(60) COMMENT '街道',
    paymentTime DATETIME COMMENT '支付时间',
    joinSuccessTime DATETIME COMMENT '拼单成功时间',
    orderConfirmationTime DATETIME COMMENT '订单确认时间',
    commitmentDeliveryTime DATETIME COMMENT '承诺发货时间',
    deliveryTime DATETIME COMMENT '发货时间',
    confirmDeliveryTime DATETIME COMMENT '确认收货时间',
    productId VARCHAR(20) NOT NULL COMMENT '商品id',
    productSku VARCHAR(30) COMMENT '商品规格',
    numberOfProducts INTEGER UNSIGNED NOT NULL COMMENT '商品数量(件)',
    skuId VARCHAR(30) COMMENT '样式ID',
    merchantCodeSkuDimension VARCHAR(30) COMMENT '商家编码-SKU维度',
    merchantCodeProductDimension VARCHAR(30) COMMENT '商家编码-商品维度',
    trackingNumber VARCHAR(30) COMMENT '快递单号',
    courierCompany VARCHAR(30) COMMENT '快递公司',
    merchantNotes VARCHAR(60) COMMENT '商家备注',
    afterSaleStatus TINYINT UNSIGNED COMMENT '售后状态编码: NULL无售后, 5退款成功, 6买家撤销，10商家同意退货退款，待买家发货，11用户已发货，待商家处理, 12售后取消，退款失败，16换货成功，18商家已发货，待消费者确认收货',
    afterSaleStatusStr VARCHAR(30) COMMENT '售后状态: 无售后或售后取消',
    afterSaleApplyTime DATETIME COMMENT '售后申请时间',
    buyerMessage VARCHAR(100) DEFAULT '' COMMENT '买家留言',
    goodsName VARCHAR(60) DEFAULT '' COMMENT '货品名称',
    goodsType VARCHAR(30) COMMENT '货品类型',
    privacyCode VARCHAR(10) COMMENT '隐私号的四位分机号',
    noteSku VARCHAR(1000) COMMENT '备注对应sku',
    cost DECIMAL(10,2) COMMENT '成本(元)',
    mergeOrderNum INTEGER UNSIGNED COMMENT '组合订单数',
    profit DECIMAL(10,2) COMMENT '利润(元)',
    actualIncome DECIMAL(10,2) COMMENT '实际收入(元)',
    actualCost DECIMAL(10,2) COMMENT '实际成本(元)',
    outerOrderStatus TINYINT COMMENT '订单状态，0: 待付款，1: 待发货，2: 待收货，3: 已收货，4: 交易成功，5: 已退换货，6: 交易关闭',
    printStatus BOOLEAN COMMENT '是否打印快递单',
    thumbUrl VARCHAR(200) COMMENT '预览图',
    unique orderIdUnique(orderId)
  );
`
/*
 已签收
 已发货，待签收
 已发货，待收货
 已收货

 已发货，退款成功
 未发货，退款成功
 已签收，退款成功
 已取消，退款成功
 已取消
 待支付
 待发货
*/

//select merchantCodeSkuDimension, sum(numberOfProducts) from itemOrder where paymentTime > '2025-03-05' group by merchantCodeSkuDimension order by sum(numberOfProducts) desc;

//alter table itemOrder add printStatus BOOLEAN COMMENT '是否打印快递单';
//alter table itemOrder add thumbUrl VARCHAR(200) COMMENT '预览图';

