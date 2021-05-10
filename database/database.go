/*
 * Maintained by jemo from 2020.5.8 to now
 * Created by jemo on 2020.5.8 16:13:10
 * database
 */

package database

import (
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
  ConnectDB()
  //execSQL(alter)
  //execSQL(order1688)
  //execSQL(item)
  //execSQL(supplier)
  //execSQL(searchItem)
  //execSQL(itemOrder)
  //execSQL(createPddItem)
  //execSQL(createPddAdPlan)
  //execSQL(createPddAdUnit)
  //execSQL(createPddAdUnitDailyData)
  //execSQL(createStall)
  //execSQL(createAdHead)
  //execSQL(createPddActivity)
  //execSQL(createAdUnitKeyword)
  //execSQL(createPddAdUnitHourlyData)
  //execSQL(createDyAccessToken)
  //execSQL(bill)
  //execSQL(createPddAdHourlyData)
  //execSQL(createDyItem)
  //execSQL(createDyOrder)
  //execSQL(createDyChildOrder)
  //execSQL(createPddItemPriceHistory)
  //execSQL(createPddCompetitor)
  //execSQL(createPddCompetitorItem)
  //execSQL(createPddCompetitorItemSale)
  //execSQL(createPddGoodsFlowData)
  //execSQL(createPddShopFlowData)
  //execSQL(createWomenItem)
  //execSQL(createWomenItemSku)
  //execSQL(createWomenItemMainImage)
  //execSQL(createWomenItemAttribute)
  //execSQL(createWomenItemDetailImage)
  //execSQL(createWomenItemColor)
  //execSQL(createWomenItemSize)
  //execSQL(createWomenItemCloudWarehouseSku)
  //execSQL(createPddItemSku)
}

func ConnectDB() {
  var err error
  DB, err = sql.Open("mysql", dbUrl)
  if err != nil {
    log.Println("db-initdb-open-error: ", err)
  }
}

func execSQL(sqlStmt string) {
  log.Print("db inited execSQL sqlStmt: ", sqlStmt)
  stmt, err := DB.Prepare(sqlStmt)
  if err != nil {
    log.Println("db-execSQL-prepare-error: ", err)
  }
  defer stmt.Close()
  _, err = stmt.Exec()
  if err != nil {
    log.Println("db-execSQL-exec-error: ", err)
  }
}

const item = `
  CREATE TABLE IF NOT EXISTS item (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    imgUrl VARCHAR(200) NOT NULL DEFAULT '',
    detailUrl VARCHAR(60) NOT NULL DEFAULT '',
    siteType TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1: 1688, 2: hznzcn',
    originalId VARCHAR(30) NOT NULL DEFAULT '',
    supplierId INTEGER UNSIGNED NOT NULL DEFAULT 0,
    saleQuantity INTEGER UNSIGNED NOT NULL DEFAULT 0,
    quantitySumMonth INTEGER UNSIGNED NOT NULL DEFAULT 0,
    gmv30dRt DECIMAL(12,2) UNSIGNED NOT NULL DEFAULT 0,
    searchId INTEGER UNSIGNED NOT NULL DEFAULT 0,
    suitPrice DECIMAL(10,2) NOT NULL DEFAULT 0,
    shippingPrice DECIMAL(10,2) NOT NULL DEFAULT 0,
    forSell BOOLEAN NOT NULL DEFAULT false,
    imgUrlOf290x290 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf120x120 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf270x270 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf100x100 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf150x150 VARCHAR(100) NOT NULL DEFAULT '',
    imgUrlOf220x220 VARCHAR(100) NOT NULL DEFAULT '',
    womenProductId INTEGER UNSIGNED,
    keyName VARCHAR(200)
  );
 `
const supplier = `
  CREATE TABLE IF NOT EXISTS supplier (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    memberId VARCHAR(30) NOT NULL DEFAULT '',
    creditLevel TINYINT UNSIGNED NOT NULL DEFAULT 0,
    shopRepurchaseRate FLOAT UNSIGNED NOT NULL DEFAULT 0,
    province VARCHAR(20) NOT NULL DEFAULT '',
    city VARCHAR(60) NOT NULL DEFAULT '',
    url VARCHAR(150) NOT NULL DEFAULT '',
    siteType TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '1: 1688, 2: hznzcn',
    mallName VARCHAR(20),
    floor TINYINT,
    stallNumber VARCHAR(5),
    phone VARCHAR(11),
    telephone VARCHAR(13),
    wechat VARCHAR(20),
    qq VARCHAR(12),
    dataUrl VARCHAR(30)
  );
`

const searchItem = `
  CREATE TABLE IF NOT EXISTS searchItem (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL
  );
 `

const itemOrder = `
  CREATE TABLE IF NOT EXISTS itemOrder (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    mallId INTEGER UNSIGNED COMMENT '店铺id',
    productName VARCHAR(60) NOT NULL COMMENT '拼多多商品',
    orderId VARCHAR(30) NOT NULL COMMENT '拼多多订单号',
    outerOrderId VARCHAR(30) COMMENT '外部订单号1688/女装网订单',
    orderStatus TINYINT NOT NULL COMMENT '拼多多订单状态，0: 待支付，1: 待发货/已发货，待签收/已签收/未发货，退款成功/已发货，退款成功，2: 已取消',
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
    receiver VARCHAR(60) NOT NULL COMMENT '收货人',
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
    productSku VARCHAR(30) NOT NULL COMMENT '商品规格',
    numberOfProducts INTEGER UNSIGNED NOT NULL COMMENT '商品数量(件)',
    skuId VARCHAR(30) COMMENT '样式ID',
    merchantCodeSkuDimension VARCHAR(30) COMMENT '商家编码-SKU维度',
    merchantCodeProductDimension VARCHAR(30) COMMENT '商家编码-商品维度',
    trackingNumber VARCHAR(30) COMMENT '快递单号',
    courierCompany VARCHAR(30) COMMENT '快递公司',
    merchantNotes VARCHAR(60) NOT NULL DEFAULT '' COMMENT '商家备注',
    afterSaleStatus TINYINT UNSIGNED COMMENT '售后状态: NULL无售后, 5退款成功, 6买家撤销，10商家同意退货退款，待买家发货，11用户已发货，待商家处理, 12售后取消，退款失败，16换货成功，18商家已发货，待消费者确认收货',
    afterSaleApplyTime DATETIME COMMENT '售后申请时间',
    buyerMessage VARCHAR(100) NOT NULL DEFAULT '' COMMENT '买家留言',
    goodsName VARCHAR(60) NOT NULL DEFAULT '' COMMENT '货品名称',
    goodsType VARCHAR(30) COMMENT '货品类型',
  );
`

/*
const alter = `
  ALTER TABLE pddAdPlan
  MODIFY COLUMN
  scenesType TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '0: 多多搜索，1: 聚焦展位，2: 多多场景'
  ;
`
*/

/*
const alter = `
  ALTER TABLE pddItem
  CHANGE COLUMN
  promotion_goods promotionGoods VARCHAR(50)
  ;
`
*/

/*
const alter = `
  ALTER TABLE dyAccessToken
  ADD
  createdAt TIMESTAMP NOT NULL COMMENT '创建时间'
  AFTER shopName
  ;
`
*/

/*
const alter = `
  ALTER TABLE itemOrder
  DROP COLUMN
  afterSaleStatus
  ;
`
*/

/*
const alter = `
  ALTER TABLE adUnitKeyword
  ADD INDEX adUnitKeywordIndex (mallId, adId, keywordId, date)
  ;
`
*/
