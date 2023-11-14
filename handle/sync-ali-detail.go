/*
 * Maintained by jemo from 2021.12.13 to now
 * Created by jemo on 2021.12.13 21:44:00
 * 同步阿里巴巴详情页
 */

package handle

import (
  "fmt"
  "log"
  "regexp"
  "strings"
  "strconv"
  "database/sql"
  "encoding/json"
  "github.com/PuerkitoBio/goquery"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncAliDetail(html string) {
  doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
  if err != nil {
    log.Println("sync-ali-detail-goquery-err: ", err)
  }
  script := doc.Find("script").Text()
  reg, _ := regexp.Compile("window.__INIT_DATA=(.*)")
  findString := reg.FindStringSubmatch(script)
  if len(findString) == 2 {
    var data map[string]interface{}
    err := json.Unmarshal([]byte(findString[1]), &data)
    if err != nil {
      log.Println("sync-ali-detail-json-decode-err: ", err)
    }
    globalData := data["globalData"].(map[string]interface{})
    tempModel := globalData["tempModel"].(map[string]interface{})
    offerId := tempModel["offerId"].(float64)
    offerTitle := tempModel["offerTitle"].(string)
    orderParamModel := globalData["orderParamModel"].(map[string]interface{})
    orderParam := orderParamModel["orderParam"].(map[string]interface{})
    skuParam := orderParam["skuParam"].(map[string]interface{})
    skuRangePrices := skuParam["skuRangePrices"].([]interface{})
    maxPrice := skuRangePrices[0].(map[string]interface{})["price"].(string)
    companyName := ""
    if tempModel["companyName"] != nil {
      companyName = tempModel["companyName"].(string)
    }
    offerBaseInfo := globalData["offerBaseInfo"].(map[string]interface{})
    sellerWinportUrl := offerBaseInfo["sellerWinportUrl"].(string)
    SaveItem(offerId, offerTitle, maxPrice, companyName, sellerWinportUrl)
    skuModel := globalData["skuModel"].(map[string]interface{})
    skuInfoMap, ok := skuModel["skuInfoMap"].(map[string]interface{})
    if(ok) {
      SaveItemSku(offerId, skuInfoMap, maxPrice)
    }
    skuProps, ok := skuModel["skuProps"].([]interface{})
    if(ok) {
      SaveItemSkuProps(offerId, skuProps)
    }
  }
}

func SaveItemSku(offerId float64, skuInfoMap map[string]interface{}, maxPrice string) {
    db := database.DB
    originalId := fmt.Sprintf("%.f", offerId)
    var searchId int
    err := db.QueryRow("SELECT searchId FROM item WHERE originalId = ?", originalId).Scan(&searchId)
    if err == sql.ErrNoRows {
      log.Println("sync-ali-detail-select-search-id-error: ", err)
    } else {
      removeAllItemSku, err := db.Prepare("UPDATE itemSku SET isDeleted = true WHERE searchId = ?")
      if err != nil {
        log.Println("sync-ali-detail-remove-all-sku-prepare-error: ", err)
      }
      defer removeAllItemSku.Close()
      _, err = removeAllItemSku.Exec(searchId)
      if err != nil {
        log.Println("sync-ali-detail-remove-all-sku-exec-error: ", err)
      }
      for skuName, skuMapInterface := range skuInfoMap {
        skuMap := skuMapInterface.(map[string]interface{})
        specId := skuMap["specId"].(string)
        specAttrs := skuMap["specAttrs"].(string)
        saleCount := int(skuMap["saleCount"].(float64))
        canBookCount := int(skuMap["canBookCount"].(float64))
        skuId := int(skuMap["skuId"].(float64))
        isPromotionSku := skuMap["isPromotionSku"].(bool)
        var price int
        if skuMap["price"] != nil {
          priceFloat, err := strconv.ParseFloat(skuMap["price"].(string), 64)
          if err == nil {
            price = int(priceFloat * 100)
          }
        } else {
          maxPriceFloat, err := strconv.ParseFloat(maxPrice, 64)
          if err == nil {
            price = int(maxPriceFloat * 100)
          }
        }
        var discountPrice int
        if skuMap["discountPrice"] != nil {
          discountPriceFloat, err := strconv.ParseFloat(skuMap["discountPrice"].(string), 64)
          if err == nil {
            discountPrice = int(discountPriceFloat * 100)
          }
        }
        var itemSkuCount int
        err = db.QueryRow("SELECT COUNT(*) FROM itemSku WHERE searchId = ? and skuId = ?", searchId, skuId).Scan(&itemSkuCount)
        if err != nil {
          log.Println("sync-ali-detail-count-item-sku-error: ", err)
        }
        if itemSkuCount == 0 {
          insertItemSku, err := db.Prepare("INSERT INTO itemSku (searchId, skuName, specId, specAttrs, price, saleCount, discountPrice, canBookCount, skuId, isPromotionSku, isDeleted) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
          if err != nil {
            log.Println("sync-ali-detail-insert-item-sku-prepare-error: ", err)
          }
          defer insertItemSku.Close()
          _, err = insertItemSku.Exec(searchId, skuName, specId, specAttrs, price, saleCount, discountPrice, canBookCount, skuId, isPromotionSku, false)
          if err != nil {
            log.Println("sync-ali-detail-insert-item-sku-exec-error: ", err)
          }
        } else {
          updateItemSku, err := db.Prepare("UPDATE itemSku SET skuName = ?, specId = ?, specAttrs = ?, price = ?, saleCount = ?, discountPrice = ?, canBookCount = ?, isPromotionSku = ?, isDeleted = false WHERE searchId = ? and skuId = ?")
          if err != nil {
            log.Println("sync-ali-detail-update-item-sku-prepare-error: ", err)
          }
          defer updateItemSku.Close()
          _, err = updateItemSku.Exec(skuName, specId, specAttrs, price, saleCount, discountPrice, canBookCount, isPromotionSku, searchId, skuId)
          if err != nil {
            log.Println("sync-ali-detail-update-item-sku-exec-error: ", err)
          }
        }
      }
    }
}

func SaveItem(offerId float64, offerTitle string, maxPrice string, companyName string, sellerWinportUrl string) {
  originalId := fmt.Sprintf("%.f", offerId)
  detailUrl := "https://detail.1688.com/offer/" + originalId + ".html"
  db := database.DB
  // insert supplier
  var supplierCount int
  err := db.QueryRow("SELECT COUNT(*) FROM supplier WHERE name = ?", companyName).Scan(&supplierCount)
  if err != nil {
    log.Println("sync-ali-detail-count-supplier-error: ", err)
  }
  if supplierCount == 0 {
    insertSupplier, err := db.Prepare("INSERT INTO supplier (name, url, siteType) VALUES(?, ?, ?)")
    if err != nil {
      log.Println("sync-ali-detail-insert-supplier-prepare-error: ", err)
    }
    defer insertSupplier.Close()
    _, err = insertSupplier.Exec(companyName, sellerWinportUrl, 1)
    if err != nil {
      log.Println("sync-ali-detail-insert-supplier-exec-error: ", err)
    }
  }
  var supplierId int64
  err = db.QueryRow("SELECT id FROM supplier WHERE name = ?", companyName).Scan(&supplierId)
  if err != nil {
    log.Println("sync-ali-detail-query-supplier-id-error: ", err)
  }
  // insert item
  var itemCount int
  err = db.QueryRow("SELECT COUNT(*) FROM item WHERE originalId = ?", originalId).Scan(&itemCount)
  if err != nil {
    log.Println("sync-ali-detail-count-item-error: ", err)
  }
  if itemCount == 0 {
    insertItem, err := db.Prepare("INSERT INTO item (name, price, detailUrl, siteType, originalId, supplierId, forSell, keyName) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
      log.Println("sync-ali-detail-insert-item-prepare-error: ", err)
    }
    defer insertItem.Close()
    _, err = insertItem.Exec(offerTitle, maxPrice, detailUrl, 1, originalId, supplierId, true, offerTitle)
    if err != nil {
      log.Println("sync-ali-detail-insert-item-exec-error: ", err)
    }
  } else if itemCount == 1 {
    updateItem, err := db.Prepare("UPDATE item SET price = ?  WHERE originalId = ?")
    if err != nil {
      log.Println("sync-ali-detail-update-item-prepare-error: ", err)
    }
    defer updateItem.Close()
    _, err = updateItem.Exec(maxPrice, originalId)
    if err != nil {
      log.Println("sync-ali-detail-update-item-exec-error: ", err)
    }
  }
}

// 保存skuProp
func SaveItemSkuProps(offerId float64, skuProps []interface{}) {
  db := database.DB
  originalId := fmt.Sprintf("%.f", offerId)
  var searchId int
  err := db.QueryRow("SELECT searchId FROM item WHERE originalId = ?", originalId).Scan(&searchId)
  if err == sql.ErrNoRows {
    log.Println("sync-ali-detail-select-search-id-error: ", err)
    return
  }
  insert, err := db.Prepare("insert into itemSkuProp (searchId, propName, propValue) values (? , ?, ?) on duplicate key update propValue = ?")
  if err != nil {
    log.Println("sync-ali-detail-save-sku-prop-insert-prepare-error: ", err)
    return
  }
  defer insert.Close()
  for i := 0; i < len(skuProps); i++ {
    propMap := skuProps[i].(map[string]interface{})
    prop := propMap["prop"].(string)
    value := propMap["value"].([]interface{})
    valueJson, err := json.Marshal(value)
    if err != nil {
      log.Println("sync-ali-detail-marshal-error: ", err)
      return
    }
    _, err = insert.Exec(searchId, prop, valueJson, valueJson)
    if err != nil {
      log.Println("sync-ali-detail-save-sku-prop-insert-exec-error: ", err)
    }
  }
}
