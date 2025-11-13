/*
 * Maintained by jemo from 2020.12.3 to now
 * Created by jemo on 2020.12.3 16:32:25
 * Douyin Order
 * 抖音小店订单-父订单
 */

package database

const createDyOrder = `
  CREATE TABLE IF NOT EXISTS dyOrder (
    orderId VARCHAR(20) NOT NULL PRIMARY KEY COMMENT '订单ID',
    shopId INTEGER UNSIGNED COMMENT '店铺ID',
    openId VARCHAR(10) COMMENT '买家抖音ID',
    orderStatus TINYINT UNSIGNED COMMENT '订单状态',
    orderStatusStr VARCHAR(30) COMMENT '订单状态',
    orderType TINYINT UNSIGNED COMMENT '订单类型 (0实物，2普通虚拟，4poi核销，5三方核销，6服务市场)',
    orderTag VARCHAR(30),
    childNum TINYINT UNSIGNED COMMENT '子订单数量',
    postAddrProvinceId VARCHAR(20) COMMENT '省ID',
    postAddrProvinceName VARCHAR(50) COMMENT '省',
    postAddrCityId VARCHAR(20) COMMENT '市ID',
    postAddrCityName VARCHAR(50) COMMENT '市',
    postAddrTownId VARCHAR(20) COMMENT '区ID',
    postAddrTownName VARCHAR(50) COMMENT '区',
    postAddrStreetId VARCHAR(20) COMMENT '乡ID',
    postAddrStreetName VARCHAR(50) COMMENT '乡区',
    postAddrDetail VARCHAR(100) COMMENT '详细地址',
    postCode VARCHAR(6) COMMENT '邮编',
    postReceiver VARCHAR(50) COMMENT '收件人姓名',
    postTel VARCHAR(11) COMMENT '收件人电话',
    buyerWords VARCHAR(50) COMMENT '买家备注',
    sellerWords VARCHAR(50) COMMENT '卖家备注',
    logisticsId INTEGER UNSIGNED COMMENT '物流公司ID',
    logisticsCode VARCHAR(30) COMMENT '物流单号',
    logisticsTime DATETIME COMMENT '发货时间。未发货时为"0"，已发货返回秒级时间戳',
    receiptTime DATETIME COMMENT '收货时间。未收货时为"0"，已发货返回秒级时间戳',
    createTime DATETIME COMMENT '订单创建时间，秒级时间戳',
    updateTime DATETIME COMMENT '订单更新时间',
    expShipTime DATETIME COMMENT '订单最晚发货时间',
    cancelReason VARCHAR(50) COMMENT '订单取消原因',
    payType TINYINT UNSIGNED COMMENT '支付类型 (0：货到付款，1：微信，2：支付宝）',
    payTime DATETIME COMMENT '支付时间',
    postAmount INTEGER UNSIGNED COMMENT '邮费金额 (单位: 分)',
    couponAmount INTEGER UNSIGNED COMMENT '平台优惠券金额 (单位: 分)',
    shopCouponAmount INTEGER UNSIGNED COMMENT '商家优惠券金额 (单位: 分)',
    couponMetaId VARCHAR(50) COMMENT '优惠券id',
    orderTotalAmount INTEGER UNSIGNED COMMENT '订单实付金额（不包含运费）分',
    isComment BOOLEAN COMMENT '是否评价 :1已评价，0未评价',
    urgeCnt TINYINT UNSIGNED COMMENT '催单次数',
    bType TINYINT UNSIGNED COMMENT '订单APP渠道，0:站外 1:火山 2:抖音 3:头条 4:西瓜 5:微信 6:闪购 7:头条lite版本 8:懂车帝 9:皮皮虾 11:抖音极速版 12:TikTok',
    subBType TINYINT UNSIGNED COMMENT '订单来源类型 0:未知 1:app 2:小程序 3:h5',
    cBiz TINYINT UNSIGNED COMMENT '订单业务类型，如下所示 1:鲁班广告 2:联盟 4:商城 8:自主经营 10:线索通支付表单 12:抖音门店 14:抖+ 15:穿山甲',
    isInsurance BOOLEAN COMMENT '是否有退货运费险',
    cType TINYINT UNSIGNED COMMENT '已废弃，无业务意义',
    cosRatio VARCHAR(5) COMMENT '已废弃，无业务意义',
    userName VARCHAR(20) COMMENT '买家名称',
    finalStatus TINYINT UNSIGNED COMMENT '订单状态',
    shippedNum TINYINT UNSIGNED,
    remark VARCHAR(100) COMMENT '备注'
  );
`
