/*
 * Maintained by jemo from 2020.5.31 to now
 * Created by jemo on 2020.5.31 10:16:14
 * Women Item List Url
 * 保存女装网商品数据
 */

package handle

import (
  "encoding/json"
  "io"
  "log"
  "net/http"
  "strings"
  "strconv"
  "github.com/jiangtaozy/pdd-management-api/database"
  "github.com/gocolly/colly/v2"
)

func WomenItemListUrl(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("women-item-list-url-decode-err: ", err)
  }
  womenItemListUrl := body["womenItemListUrl"].(string)
  collector := colly.NewCollector()
  collector.OnHTML("#BrandProductListDiv ul li", func(e *colly.HTMLElement) {
    productid, _ := e.DOM.Attr("productid")
    title, _ := e.DOM.Find(".insideBox .pic a img").Attr("title")
    itemUrl, _ := e.DOM.Find(".insideBox .pic a").Attr("href")
    imgUrl, _ := e.DOM.Find(".insideBox .pic a img").Attr("data-original")
    priceStr := e.DOM.Find(".rowPri .price").Text()
    findSelector := ".insideBox .pic .ProductListItem_HoverInfo_" + productid
    brandname, _ := e.DOM.Find(findSelector).Attr("brandname")
    brandurl, _ := e.DOM.Find(findSelector).Attr("brandurl")
    db := database.DB
    // insert search item
    priceStr = strings.Replace(priceStr, "￥", "", -1)
    price, err := strconv.ParseFloat(strings.Trim(priceStr, " "), 64)
    if err != nil {
      log.Println("women-item-list-url-parse-float-error: ", err)
    }
    insertSearchItem, err := db.Prepare("INSERT INTO searchItem (name) VALUES(?)")
    if err != nil {
      log.Println("women-item-list-url-insert-search-item-exec-error: ", err)
    }
    defer insertSearchItem.Close()
    var searchItemCount int
    err = db.QueryRow("SELECT COUNT(*) FROM searchItem WHERE name = ?", title).Scan(&searchItemCount)
    if err != nil {
      log.Println("women-item-list-url-count-search-item-error: ", err)
    }
    if searchItemCount == 0 {
      _, err = insertSearchItem.Exec(title)
      if err != nil {
        log.Println("women-item-list-url-insert-search-item-exec-error: ", err)
      }
    }
    var searchId int64
    err = db.QueryRow("SELECT id FROM searchItem WHERE name = ?", title).Scan(&searchId)
    if err != nil {
      log.Println("women-item-list-url-query-search-item-id-error: ", err)
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
    err = db.QueryRow("SELECT COUNT(*) FROM item WHERE searchId = ? AND supplierId = ?", searchId, supplierId).Scan(&itemCount)
    if err != nil {
      log.Println("women-item-list-url-count-item-error: ", err)
    }
    if itemCount == 0 {
      insertItem, err := db.Prepare("INSERT INTO item (name, price, imgUrl, detailUrl, siteType, supplierId, searchId, suitPrice, forSell, womenProductId) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
      if err != nil {
        log.Println("women-item-list-url-insert-item-prepare-error: ", err)
      }
      defer insertItem.Close()
      _, err = insertItem.Exec(title, price, imgUrl, itemUrl, 2, supplierId, searchId, price, true, productid)
      if err != nil {
        log.Println("women-item-list-url-insert-item-exec-error: ", err)
      }
    } else {
      updateItem, err := db.Prepare("UPDATE item SET price = ?, imgUrl = ?, detailUrl = ?, suitPrice = ?, womenProductId = ? WHERE searchId = ? AND supplierId = ?")
      if err != nil {
        log.Println("women-item-list-url-update-item-prepare-error: ", err)
      }
      defer updateItem.Close()
      _, err = updateItem.Exec(price, imgUrl, itemUrl, price, productid, searchId, supplierId)
      if err != nil {
        log.Println("women-item-list-url-update-item-exec-error: ", err)
      }
    }
  })
  collector.OnError(func(_ *colly.Response, err error) {
    log.Println("women-item-list-url-collector-on-error:", err)
  })
  collector.Visit(womenItemListUrl)
  io.WriteString(w, "ok")
}
