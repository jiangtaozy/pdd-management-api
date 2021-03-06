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
  "encoding/json"
  "net/http"
)

func SaveNetworkData(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("save-network-data-err: ", err)
  }
  requestText := body["requestText"].(string)
  requestBody := make(map[string]interface{})
  err = json.Unmarshal([]byte(requestText), &requestBody)
  if err != nil {
    log.Println("save-network-data-unmarshal-request-text-error: ", err)
  }
  responseContent := body["responseContent"].(string)
  responseBody := make(map[string]interface{})
  err = json.Unmarshal([]byte(responseContent), &responseBody)
  if err != nil {
    log.Println("save-network-data-unmarshal-response-content-error: ", err)
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
  now := time.Now()
  diff := now.Sub(start)
  log.Println("diff: ", diff)
  io.WriteString(w, "ok")
}
