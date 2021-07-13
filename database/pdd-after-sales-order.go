/*
 * Maintained by jemo from 2021.7.13 to now
 * Created by jemo on 2021.7.13 10:15:17
 * 拼多多售后订单
 */

package database

const createPddAfterSalesOrder = `
  CREATE TABLE IF NOT EXISTS pddAfterSalesOrder (
    afterSalesReasonCode INTEGER UNSIGNED COMMENT '售后原因编码',
    afterSalesReasonDesc VARCHAR(50) COMMENT '售后原因描述',
    afterSalesStatus INTEGER UNSIGNED COMMENT '售后状态',
    afterSalesTitle VARCHAR(50) COMMENT '售后标题',
    afterSalesType INTEGER UNSIGNED COMMENT '售后类型',
    assignedProcessorId VARCHAR(50),
    assignedProcessorName VARCHAR(50),
    createdAt INTEGER UNSIGNED COMMENT '创建时间',
    expireActionDesc VARCHAR(100) COMMENT '超期描述',
    expireRemainTime INTEGER UNSIGNED COMMENT '超期剩余时间',
    goodsName VARCHAR(200) COMMENT '商品名称',
    goodsNumber INTEGER UNSIGNED COMMENT '商品数量',
    goodsSpec VARCHAR(50) COMMENT '商品规格',
    id BIGINT UNSIGNED COMMENT '售后id',
    mallRemark VARCHAR(50),
    mallRemarkTag VARCHAR(50),
    mallRemarkTagName VARCHAR(50),
    orderAmount INTEGER UNSIGNED COMMENT '订单金额(分)',
    orderSn VARCHAR(30) NOT NULL COMMENT '拼多多订单号',
    receiveAmount INTEGER UNSIGNED COMMENT '订单金额(分)',
    refundAmount INTEGER UNSIGNED COMMENT '订单金额(分)',
    rejectChatTipDesc VARCHAR(50),
    rejectChatTipExpireRemainTime INTEGER UNSIGNED,
    remarkStatus VARCHAR(50),
    returnCouponAmount INTEGER UNSIGNED COMMENT '退还优惠券金额(分)',
    reverseShippingId INTEGER UNSIGNED,
    reverseTrackingNumber VARCHAR(50),
    sellerAfterSalesShippingStatus INTEGER UNSIGNED COMMENT '卖家售后物流状态',
    sellerAfterSalesShippingStatusDesc VARCHAR(50) COMMENT '卖家售后物流描述',
    thumbUrl VARCHAR(200) COMMENT '图片地址',
    ticketId VARCHAR(50),
    uid BIGINT UNSIGNED COMMENT '用户id',
    version INTEGER UNSIGNED
  );
`
