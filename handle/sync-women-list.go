/*
 * Maintained by jemo from 2021.11.22 to now
 * Created by jemo on 2021.11.22 21:32:40
 * 同步女装网商品列表
 */

package handle

import (
  "log"
  "strings"
  "strconv"
  "github.com/PuerkitoBio/goquery"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncWomenList(html string) {
  doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
  if err != nil {
    log.Println("sync-women-list-goquery-err: ", err)
  }
  // 女装网新品/云仓/闪电发货
  doc.Find("#productList_Div ul li").Each(func(i int, s *goquery.Selection) {
    ImportWomenItem(s)
  })
  // 女装网店铺商品列表
  doc.Find("#BrandProductListDiv ul li").Each(func(i int, s *goquery.Selection) {
    ImportWomenItem(s)
  })
}

func ImportWomenItem(s *goquery.Selection) {
  productid, _ := s.Attr("productid")
  title, _ := s.Find(".insideBox .pic a img").Attr("title")
  // 每日新款
  keyName := s.Find(".insideBox .outWidth .dsrs a").Text()
  // 云仓
  if keyName == "" {
    keyName = s.Find(".insideBox .dsrs a").Text()
  }
  itemUrl, _ := s.Find(".insideBox .pic a").Attr("href")
  imgUrl, _ := s.Find(".insideBox .pic a img").Attr("data-original")
  priceStr := s.Find(".rowPri .price").Text()
  findSelector := ".insideBox .pic .ProductListItem_HoverInfo_" + productid
  brandname, _ := s.Find(findSelector).Attr("brandname")
  brandurl, _ := s.Find(findSelector).Attr("brandurl")
  db := database.DB
  // insert search item
  priceStr = strings.Replace(priceStr, "￥", "", -1)
  priceStr = strings.Replace(priceStr, "补贴价", "", -1)
  price, err := strconv.ParseFloat(strings.Trim(priceStr, " "), 64)
  if err != nil {
    log.Println("women-item-list-url-parse-float-error: ", err)
  }
  // insert supplier
  var supplierCount int
  err = db.QueryRow("SELECT COUNT(*) FROM supplier WHERE name = ?", brandname).Scan(&supplierCount)
  if err != nil {
    log.Println("women-item-list-url-count-supplier-error: ", err)
  }
  if supplierCount == 0 {
    insertSupplier, err := db.Prepare("INSERT INTO supplier (name, url, siteType) VALUES(?, ?, ?)")
    if err != nil {
      log.Println("women-item-list-url-insert-supplier-prepare-error: ", err)
    }
    defer insertSupplier.Close()
    _, err = insertSupplier.Exec(brandname, brandurl, 2)
    if err != nil {
      log.Println("women-item-list-url-insert-supplier-exec-error: ", err)
    }
  }
  var supplierId int64
  err = db.QueryRow("SELECT id FROM supplier WHERE name = ?", brandname).Scan(&supplierId)
  if err != nil {
    log.Println("women-item-list-url-query-supplier-id-error: ", err)
  }
  // insert item
  var itemCount int
  err = db.QueryRow("SELECT COUNT(*) FROM item WHERE womenProductId = ?", productid).Scan(&itemCount)
  if err != nil {
    log.Println("women-item-list-url-count-item-error: ", err)
  }
  if itemCount == 0 {
    insertItem, err := db.Prepare("INSERT INTO item (name, price, imgUrl, detailUrl, siteType, supplierId, suitPrice, forSell, womenProductId, keyName) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
      log.Println("women-item-list-url-insert-item-prepare-error: ", err)
    }
    defer insertItem.Close()
    _, err = insertItem.Exec(title, price, imgUrl, itemUrl, 2, supplierId, price, true, productid, keyName)
    if err != nil {
      log.Println("women-item-list-url-insert-item-exec-error: ", err)
    }
  } else if itemCount == 1 {
    updateItem, err := db.Prepare("UPDATE item SET price = ?, imgUrl = ?, detailUrl = ?, suitPrice = ?, keyName = ? WHERE womenProductId = ?")
    if err != nil {
      log.Println("women-item-list-url-update-item-prepare-error: ", err)
    }
    defer updateItem.Close()
    _, err = updateItem.Exec(price, imgUrl, itemUrl, price, keyName, productid)
    if err != nil {
      log.Println("women-item-list-url-update-item-exec-error: ", err)
    }
  } else {
    stmtDeleteItem, err := db.Prepare("DELETE FROM item WHERE womenProductId = ?")
    if err != nil {
      log.Println("women-item-list-url-delete-item-prepare-error: ", err)
    }
    _, err = stmtDeleteItem.Exec(productid)
    if err != nil {
      log.Println("women-item-list-url-delete-item-exec-error: ", err)
    }
  }
}
