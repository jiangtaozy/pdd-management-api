/*
 * Maintained by jemo from 2020.4.27 to now
 * Created by jemo on 2020.4.27 17:18:35
 * Main
 */

package main

import (
  "log"
  "net/http"
  "github.com/rs/cors"
  "github.com/jiangtaozy/pdd-management-api/handle"
  "github.com/jiangtaozy/pdd-management-api/database"
)

var port = ":7000"

func main() {
  database.InitDB()
  log.Println("listen at ", port)
  mux := http.NewServeMux()
  mux.HandleFunc("/callback", handle.Callback)
  mux.HandleFunc("/searchData", handle.SearchData)
  mux.HandleFunc("/searchTitle", handle.SearchTitle)
  mux.HandleFunc("/updateSearchTitle", handle.UpdateSearchTitle)
  mux.HandleFunc("/deleteSearchTitle", handle.DeleteSearchTitle)
  mux.HandleFunc("/searchTitleById", handle.SearchTitleById)
  mux.HandleFunc("/searchTitleList", handle.SearchTitleList)
  mux.HandleFunc("/itemListBySearchId", handle.ItemListBySearchId)
  mux.HandleFunc("/updateItemSuitShippingPrice", handle.UpdateItemSuitShippingPrice)
  mux.HandleFunc("/uploadPddOrderFile", handle.UploadPddOrderFile)
  mux.HandleFunc("/orderList", handle.OrderList)
  mux.HandleFunc("/uploadPddItemData", handle.UploadPddItemData)
  mux.HandleFunc("/pddGoods", handle.PddItemList)
  mux.HandleFunc("/uploadAdPlanData", handle.UploadAdPlanData)
  mux.HandleFunc("/adPlanList", handle.AdPlanList)
  mux.HandleFunc("/adUnitList", handle.AdUnitList)
  mux.HandleFunc("/adUnitListAll", handle.AdUnitListAll)
  mux.HandleFunc("/adUnitData", handle.AdUnitData)
  mux.HandleFunc("/adUnit", handle.AdUnit)
  mux.HandleFunc("/adUnitDataList", handle.AdUnitDataList)
  mux.HandleFunc("/stall", handle.Stall)
  mux.HandleFunc("/stallList", handle.StallList)
  mux.HandleFunc("/womenItemListUrl", handle.WomenItemListUrl)
  mux.HandleFunc("/countingItemByCreatedTime", handle.CountingItemByCreatedTime)
  mux.HandleFunc("/adDayData", handle.AdDayData)
  mux.HandleFunc("/saveAdUnitDataOfDate", handle.SaveAdUnitDataOfDate)
  mux.HandleFunc("/savePddOrderData", handle.SavePddOrderData)
  mux.HandleFunc("/upload1688OrderFile", handle.Upload1688OrderFile)
  mux.HandleFunc("/order1688List", handle.Order1688List)
  mux.HandleFunc("/uploadHznzcnOrderFile", handle.UploadHznzcnOrderFile)
  mux.HandleFunc("/orderStatistics", handle.OrderStatistics)
  mux.HandleFunc("/adHeadList", handle.AdHeadList)
  mux.HandleFunc("/adHead", handle.AdHead)
  mux.HandleFunc("/savePddActivityData", handle.SavePddActivityData)
  mux.HandleFunc("/pddActivityList", handle.PddActivityList)
  mux.HandleFunc("/itemOrderUpdate", handle.ItemOrderUpdate)
  mux.HandleFunc("/adUnitKeywordCreate", handle.AdUnitKeywordCreate)
  mux.HandleFunc("/pddItemData", handle.PddItemData)
  mux.HandleFunc("/adUnitHourlyDataSave", handle.AdUnitHourlyDataSave)
  mux.HandleFunc("/adUnitHourlyDataList", handle.AdUnitHourlyDataList)
  mux.HandleFunc("/hangAfterSaleOrderUpload", handle.HangAfterSaleOrderUpload)
  mux.HandleFunc("/dyItemList", handle.DyItemList)
  mux.HandleFunc("/afterSaleOrderDataSave", handle.AfterSaleOrderDataSave)
  mux.HandleFunc("/billDataSave", handle.BillDataSave)
  mux.HandleFunc("/billList", handle.BillList)
  mux.HandleFunc("/keywordList", handle.KeywordList)
  mux.HandleFunc("/saveHarFile", handle.SaveHarFile)
  mux.HandleFunc("/orderListByAdId", handle.OrderListByAdId)
  mux.HandleFunc("/saveNetworkData", handle.SaveNetworkData)
  mux.HandleFunc("/planHourlyDataList", handle.PlanHourlyDataList)
  mux.HandleFunc("/syncDyItemData", handle.SyncDyItemData)
  mux.HandleFunc("/mallTotalAdData", handle.MallTotalAdData)
  mux.HandleFunc("/syncDyOrderData", handle.SyncDyOrderData)
  mux.HandleFunc("/dyOrderList", handle.DyOrderList)
  mux.HandleFunc("/pddItemLastThreeDayPromoteList", handle.PddItemLastThreeDayPromoteList)
  mux.HandleFunc("/pddCompetitorSave", handle.PddCompetitorSave)
  mux.HandleFunc("/pddCompetitorList", handle.PddCompetitorList)
  mux.HandleFunc("/pddCompetitorItemSave", handle.PddCompetitorItemSave)
  mux.HandleFunc("/pddCompetitorItemList", handle.PddCompetitorItemList)
  mux.HandleFunc("/pddCompetitorItemSaleSave", handle.PddCompetitorItemSaleSave)
  mux.HandleFunc("/pddCompetitorItemSaleList", handle.PddCompetitorItemSaleList)
  mux.HandleFunc("/getWomenDetailData", handle.GetWomenDetailData)
  mux.HandleFunc("/pddItemPriceHistoryData", handle.PddItemPriceHistoryData)
  mux.HandleFunc("/getWomenCloudWarehouseStock", handle.GetWomenCloudWarehouseStock)
  mux.HandleFunc("/itemStockList", handle.ItemStockList)
  mux.HandleFunc("/syncCloudWarehouseStock", handle.SyncCloudWarehouseStock)
  mux.HandleFunc("/syncWomenOnShelf", handle.SyncWomenOnShelf)
  mux.HandleFunc("/test", handle.Test)
  mux.Handle("/", http.FileServer(http.Dir("/home/jemo/workspace/pdd/pdd-management-web/build")))
  handler := cors.Default().Handler(mux)
  log.Fatal(http.ListenAndServe(port, handler))
}
