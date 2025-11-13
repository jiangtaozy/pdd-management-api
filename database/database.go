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
  //execSQL(createPddAfterSalesOrder)
  //execSQL(createItemSku)
  //execSQL(createItemSkuNum)
  //execSQL(createItemType)
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

const alter = `
  alter table item
  ADD
    itemNum VARCHAR(20)
  AFTER itemTypeKey
  ;
`
/*
const alter = `
  ALTER TABLE item
  ADD INDEX
  forSellIndex(forSell)
  ;
`
*/

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
