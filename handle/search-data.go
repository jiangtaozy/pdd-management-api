/*
 * Maintained by jemo from 2020.5.8 to now
 * Created by jemo on 2020.5.8 9:43:56
 * Search Data
 */

package handle

import (
  "io"
  "strings"
  "encoding/json"
  "log"
  "net/http"
  "strconv"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SearchData(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("search-data-error: ", err)
  }
  var searchId = body["id"]
  var searchData = body["searchData"].(string)
  searchDataMap := make(map[string]interface{})
  err = json.Unmarshal([]byte(searchData), &searchDataMap)
  if err != nil {
    log.Println("search-data-json-unmarshal-error: ", err)
  }
  var data = searchDataMap["data"].(map[string]interface{})
  var offerList = data["offerList"].([]interface{})
  db := database.ConnectDB()
  stmtIns, err := db.Prepare("INSERT INTO supplier (name, memberId, creditLevel, shopRepurchaseRate, province, city, url, siteType) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Println("search-data-insert-prepare-error: ", err)
  }
  defer stmtIns.Close()
  stmtQuery, err := db.Prepare("SELECT id FROM supplier WHERE siteType = 1 AND memberId = ?")
  if err != nil {
    log.Println("search-data-select-prepare-error: ", err)
  }
  defer stmtQuery.Close()
  stmtInsertItem, err := db.Prepare("INSERT INTO item (name, price, imgUrl, detailUrl, siteType, originalId, supplierId, saleQuantity, quantitySumMonth, gmv30dRt, searchId, imgUrlOf290x290, imgUrlOf120x120, imgUrlOf270x270, imgUrlOf100x100, imgUrlOf150x150, imgUrlOf220x220) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
  for i := 0; i < len(offerList); i++ {
    // insert supplier
    var offer = offerList[i].(map[string]interface{})
    var company = offer["company"].(map[string]interface{})
    shopRepurchaseRateStr, ok := company["shopRepurchaseRate"].(string)
    var shopRepurchaseRate float64 = 0
    if ok {
      shopRepurchaseRateStr = strings.Replace(shopRepurchaseRateStr, "%", "", -1)
      shopRepurchaseRate, err = strconv.ParseFloat(shopRepurchaseRateStr, 64)
      if err != nil {
        log.Println("search-data-parse-float-error: ", err)
      }
      shopRepurchaseRate = shopRepurchaseRate * 0.01
    }
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM supplier WHERE siteType = 1 AND memberId = ?", company["memberId"]).Scan(&count)
    if err != nil {
      log.Println("search-data-count-error: ", err)
    }
    if count == 0 {
      _, err = stmtIns.Exec(company["name"], company["memberId"], company["creditLevel"], shopRepurchaseRate, company["province"], company["city"], company["url"], 1)
      if err != nil {
        log.Println("search-data-insert-exec-error: ", err)
      }
    }
    // insert item
    var supplierId int
    err = stmtQuery.QueryRow(company["memberId"]).Scan(&supplierId)
    if err != nil {
      log.Println("search-data-query-member-id-error: ", err)
    }
    originalIdFloat := offer["id"].(float64)
    originalId := strconv.FormatFloat(originalIdFloat, 'f', 0, 64)
    var itemCount int
    err = db.QueryRow("SELECT COUNT(*) FROM item WHERE siteType = 1 AND supplierId = ? AND originalId = ? AND searchId = ?", supplierId, originalId, searchId).Scan(&itemCount)
    if err != nil {
      log.Println("search-data-query-item-count-error: ", err)
    }
    if itemCount == 0 {
      information := offer["information"].(map[string]interface{})
      tradePrice := offer["tradePrice"].(map[string]interface{})
      offerPrice := tradePrice["offerPrice"].(map[string]interface{})
      var price float64 = 0
      if offerPrice["value"] != nil {
        value := offerPrice["value"].(map[string]interface{})
        price = value["integer"].(float64) + value["decimals"].(float64) * 0.01
      }
      image := offer["image"].(map[string]interface{})
      tradeQuantity := offer["tradeQuantity"].(map[string]interface{})
      _, err = stmtInsertItem.Exec(information["subject"], price, image["imgUrl"], information["detailUrl"], 1, originalId, supplierId, tradeQuantity["saleQuantity"], tradeQuantity["quantitySumMonth"], tradeQuantity["gmv30dRt"], searchId, image["imgUrlOf290x290"], image["imgUrlOf120x120"], image["imgUrlOf270x270"], image["imgUrlOf100x100"], image["imgUrlOf150x150"], image["imgUrlOf220x220"])
      if err != nil {
        log.Println("search-data-insert-item-error: ", err)
      }
    }
  }
  io.WriteString(w, "ok")
}
