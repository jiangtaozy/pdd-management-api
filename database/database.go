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

func ConnectDB() {
  var err error
  DB, err = sql.Open("mysql", dbUrl)
  if err != nil {
    log.Println("db-initdb-open-error: ", err)
  }
}

func InitDB() {
  ConnectDB()
  //execSQL(item)
  //execSQL(supplier)
  //execSQL(searchItem)
  //execSQL(itemOrder)
  //execSQL(createPddItem)
  //execSQL(createPddAdPlan)
  //execSQL(createPddAdUnit)
  //execSQL(createPddAdUnitDailyData)
  //execSQL(createStall)
  //execSQL(addColumnAlter)
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
    imgUrl VARCHAR(100) NOT NULL DEFAULT '',
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
    imgUrlOf220x220 VARCHAR(100) NOT NULL DEFAULT ''
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
    mallName VARCHAR(20) NOT NULL DEFAULT '',
    floor TINYINT NOT NULL DEFAULT 1,
    stallNumber VARCHAR(5) DEFAULT '',
    phone VARCHAR(11) NOT NULL DEFAULT '',
    telephone VARCHAR(13) NOT NULL DEFAULT '',
    wechat VARCHAR(20) NOT NULL DEFAULT '',
    qq VARCHAR(12) NOT NULL DEFAULT '',
    dataUrl VARCHAR(30) NOT NULL DEFAULT ''
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
    productName VARCHAR(60) NOT NULL COMMENT '拼多多商品',
    orderId VARCHAR(30) NOT NULL COMMENT '拼多多订单号',
    orderStatus VARCHAR(20) NOT NULL COMMENT '拼多多订单状态',
    productTotalPrice DECIMAL(10,2) NOT NULL COMMENT '拼多多商品总价(元)',
    storeDiscount DECIMAL(10,2) NOT NULL COMMENT '店铺优惠折扣(元)',
    platformDiscount DECIMAL(10,2) NOT NULL COMMENT '平台优惠折扣(元)',
    postage DECIMAL(10,2) NOT NULL COMMENT '邮费(元)',
    onsiteInstallationFee DECIMAL(10,2) NOT NULL COMMENT '上门安装费(元)',
    homeDeliveryFee DECIMAL(10,2) NOT NULL COMMENT '送货入户费(元)',
    homeDeliveryAndInstallationFee DECIMAL(10,2) NOT NULL COMMENT '送货入户并安装费(元)',
    userPaidAmount DECIMAL(10,2) NOT NULL COMMENT '用户实付金额(元)',
    merchantReceivedAmount DECIMAL(10,2) NOT NULL COMMENT '商家实收金额(元)',
    numberOfProducts INTEGER UNSIGNED NOT NULL COMMENT '商品数量(件)',
    idCardName VARCHAR(60) NOT NULL DEFAULT '' COMMENT '身份证姓名',
    identificationNumber VARCHAR(60) NOT NULL DEFAULT '' COMMENT '身份证号码',
    receiver VARCHAR(60) NOT NULL COMMENT '收货人',
    phone VARCHAR(11) NOT NULL COMMENT '手机',
    whetherUnderReview VARCHAR(30) NOT NULL DEFAULT '' COMMENT '是否审核中',
    province VARCHAR(20) NOT NULL DEFAULT '' COMMENT '省',
    city VARCHAR(60) NOT NULL DEFAULT '' COMMENT '市',
    district VARCHAR(60) NOT NULL DEFAULT '' COMMENT '区',
    street VARCHAR(60) NOT NULL DEFAULT '' COMMENT '街道',
    paymentTime DATETIME COMMENT '支付时间',
    joinSuccessTime DATETIME COMMENT '拼单成功时间',
    orderConfirmationTime DATETIME COMMENT '订单确认时间',
    commitmentDeliveryTime DATETIME COMMENT '承诺发货时间',
    deliveryTime DATETIME COMMENT '发货时间',
    confirmDeliveryTime DATETIME COMMENT '确认收货时间',
    productId VARCHAR(20) NOT NULL COMMENT '商品id',
    productSku VARCHAR(30) NOT NULL COMMENT '商品规格',
    userBuyPhone VARCHAR(11) NOT NULL DEFAULT '' COMMENT '用户购买手机号',
    skuId VARCHAR(30) NOT NULL DEFAULT '' COMMENT '样式ID',
    merchantCodeSkuDimension VARCHAR(30) NOT NULL DEFAULT '' COMMENT '商家编码-SKU维度',
    merchantCodeProductDimension VARCHAR(30) NOT NULL DEFAULT '' COMMENT '商家编码-商品维度',
    trackingNumber VARCHAR(30) NOT NULL DEFAULT '' COMMENT '快递单号',
    courierCompany VARCHAR(30) NOT NULL DEFAULT '' COMMENT '快递公司',
    haitaoCustomsOrderNumber VARCHAR(30) NOT NULL DEFAULT '' COMMENT '海淘清关订单号',
    paymentId VARCHAR(30) NOT NULL DEFAULT '' COMMENT '支付ID',
    paymentMethod VARCHAR(30) NOT NULL DEFAULT '' COMMENT '支付方式',
    whetherDrawOrZeroYuanTry VARCHAR(10) NOT NULL DEFAULT '' COMMENT '是否抽奖或0元试用',
    whetherShunfengAddPrice VARCHAR(10) NOT NULL DEFAULT '' COMMENT '是否顺丰加价',
    merchantNotes VARCHAR(60) NOT NULL DEFAULT '' COMMENT '商家备注',
    afterSaleStatus VARCHAR(30) NOT NULL DEFAULT '' COMMENT '售后状态',
    buyerMessage VARCHAR(100) NOT NULL DEFAULT '' COMMENT '买家留言',
    relatedGoodsCode VARCHAR(30) NOT NULL DEFAULT '' COMMENT '关联货品编码',
    goodsName VARCHAR(60) NOT NULL DEFAULT '' COMMENT '货品名称',
    goodsType VARCHAR(30) NOT NULL DEFAULT '' COMMENT '货品类型',
    goodsChild VARCHAR(30) NOT NULL DEFAULT '' COMMENT '子货品',
    warehouseName VARCHAR(30) NOT NULL DEFAULT '' COMMENT '仓库名称',
    warehouseAddress VARCHAR(60) NOT NULL DEFAULT '' COMMENT '仓库所在地址',
    whetherStoreMention VARCHAR(30) NOT NULL DEFAULT '' COMMENT '是否门店自提',
    storeName VARCHAR(30) NOT NULL DEFAULT '' COMMENT '门店名称',
    storeCustomCode VARCHAR(30) NOT NULL DEFAULT '' COMMENT '门店自定义编码',
    travelInformation VARCHAR(30) NOT NULL DEFAULT '' COMMENT '旅行类信息',
    consumerInformation VARCHAR(30) NOT NULL DEFAULT '' COMMENT '消费者资料'
  );
`

/*
const alter = `
  ALTER TABLE pddItem
  CHANGE COLUMN
  promotion_goods promotionGoods VARCHAR(50)
  ;
`
*/
const alter = `
  ALTER TABLE pddAdUnitDailyData
  MODIFY COLUMN
  inquiryNum INTEGER UNSIGNED NOT NULL DEFAULT 0
  ;
`
/*
const alter = `
  ALTER TABLE supplier
  MODIFY COLUMN url VARCHAR(150) NOT NULL DEFAULT ''
  ;
`
*/
const addColumnAlter = `
  ALTER TABLE supplier
  ADD mallName VARCHAR(20) NOT NULL DEFAULT '',
  ADD floor TINYINT NOT NULL DEFAULT 1,
  ADD stallNumber VARCHAR(5) DEFAULT '',
  ADD phone VARCHAR(11) NOT NULL DEFAULT '',
  ADD telephone VARCHAR(13) NOT NULL DEFAULT '',
  ADD wechat VARCHAR(20) NOT NULL DEFAULT '',
  ADD qq VARCHAR(12) NOT NULL DEFAULT '',
  ADD dataUrl VARCHAR(30) NOT NULL DEFAULT ''
  ;
`
