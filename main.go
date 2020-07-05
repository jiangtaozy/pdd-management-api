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
  mux.Handle("/", http.FileServer(http.Dir("/home/jemo/workspace/pdd/pdd-management-web/build")))
  handler := cors.Default().Handler(mux)
  log.Fatal(http.ListenAndServe(port, handler))
}
