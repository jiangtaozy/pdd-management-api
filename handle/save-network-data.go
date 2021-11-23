/*
 * Maintained by jemo from 2020.11.19 to now
 * Created by jemo on 2020.11.19 10:36:19
 * Save Network Data
 */

package handle

import (
  "io"
  "log"
  "time"
  "strings"
  "regexp"
  "encoding/json"
  "net/http"
)

func SaveNetworkData(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("save-network-data-err: ", err)
  }
  requestBody := make(map[string]interface{})
  requestText := body["requestText"]
  if requestText != nil {
    err = json.Unmarshal([]byte(requestText.(string)), &requestBody)
    if err != nil {
      log.Println("save-network-data-unmarshal-request-text-error: ", err)
    }
  }
  responseContent := body["responseContent"].(string)
  responseBody := make(map[string]interface{})
  responseContentMimeType := body["responseContentMimeType"].(string)
  if responseContentMimeType == "application/json" {
    err = json.Unmarshal([]byte(responseContent), &responseBody)
    if err != nil {
      log.Println("save-network-data-unmarshal-response-content-error: ", err)
    }
  }
  url := body["requestUrl"].(string)
  start := time.Now()
  // 关键词
  if url == "https://yingxiao.pinduoduo.com/mms-gateway/venus/api/subway/keyword/listKeywordPage" {
    SaveListKeywordPage(responseBody)
    log.Println("url: ", url)
  }
  // 单元列表
  if url == "https://yingxiao.pinduoduo.com/mms-gateway/venus/api/unit/listPage" {
    SaveUnitListPage(requestBody, responseBody)
    log.Println("url: ", url)
  }
  // 小时数据
  if url == "https://yingxiao.pinduoduo.com/mms-gateway/apollo/api/report/queryHourlyReport" {
    SaveQueryHourlyReport(requestBody, responseBody)
    log.Println("url: ", url)
  }
  // 商品列表
  if url == "https://mms.pinduoduo.com/vodka/v2/mms/query/display/mall/goodsList" {
    SyncPddItem(requestBody, responseBody)
    log.Println("url: ", url)
  }
  // 订单列表
  if url == "https://mms.pinduoduo.com/mangkhut/mms/recentOrderList" {
    SyncPddOrder(requestBody, responseBody)
    log.Println("url: ", url)
  }
  // 流量数据
  if url == "https://mms.pinduoduo.com/sydney/api/goodsDataShow/queryGoodsDetailVOList" {
    SyncPddGoodsFlowData(requestBody, responseBody)
    log.Println("url: ", url)
  }
  // 流量详情
  if url == "https://mms.pinduoduo.com/sydney/api/goodsDataShow/queryGoodsPageOverView" {
    SyncPddShopFlowData(requestBody, responseBody)
    log.Println("url: ", url)
  }
  // 售后列表
  if url == "https://mms.pinduoduo.com/mercury/mms/afterSales/queryList" {
    SyncPddAfterSalesOrder(requestBody, responseBody)
    log.Println("url: ", url)
  }
  // 女装网订单列表
  if strings.Contains(url, "https://www.hznzcn.com/order/query_my_order_list") {
    SyncWomenOrder(requestBody, responseBody)
    log.Println("url: ", url)
  }
  // 女装网杭州女装新款列表
  isWomenList, _ := regexp.MatchString(`https://www.hznzcn.com/hz/gallery-[\d-]+-grid.html`, url)
  if isWomenList {
    SyncWomenList(responseContent)
  }
  now := time.Now()
  diff := now.Sub(start)
  log.Println("diff: ", diff)
  io.WriteString(w, "ok")
}
