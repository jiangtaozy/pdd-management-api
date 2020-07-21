/*
 * Maintained by jemo from 2020.7.21 to now
 * Created by jemo on 2020.7.21 16:37:16
 * Save pdd activity data
 */

package handle

import (
  "io"
  "time"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SavePddActivityData(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("save-pdd-activity-data-decode-err: ", err)
  }
  dataString := body["pddActivityData"].(string)
  dataMap := make(map[string]interface{})
  err = json.Unmarshal([]byte(dataString), &dataMap)
  if err != nil {
    log.Println("save-pdd-activity-data-json-unmarshal-error: ", err)
  }
  result := dataMap["result"].(map[string]interface{})
  list := result["marketing_activity_list"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare("INSERT INTO pddActivity (activityId, activityName, activityType, goodsId, goodsName, hdThumbUrl, maxOnSaleGroupPrice, minOnSaleGroupPrice, onlineQuantity, maxPreSalePrice, minPreSalePrice, discount, activityQuantity, activityStockQuantity, startTime, endTime, status, endOperationTime, newGoods) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Println("save-pdd-activity-data-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare("UPDATE pddActivity SET activityName = ?, activityType = ?, goodsId = ?, goodsName = ?, hdThumbUrl = ?, maxOnSaleGroupPrice = ?, minOnSaleGroupPrice = ?, onlineQuantity = ?, maxPreSalePrice = ?, minPreSalePrice = ?, discount = ?, activityQuantity = ?, activityStockQuantity = ?, startTime = ?, endTime = ?, status = ?, endOperationTime = ?, newGoods = ? WHERE activityId = ?")
  if err != nil {
    log.Println("save-pdd-activity-data-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    activity := list[i].(map[string]interface{})
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM pddActivity WHERE activityId = ?", activity["activity_id"]).Scan(&count)
    if err != nil {
      log.Println("save-pdd-activity-data-count-error: ", err)
    }
    activityId := activity["activity_id"]
    activityName := activity["activity_name"]
    activityType := activity["activity_type"]
    goodsId := activity["goods_id"]
    goodsName := activity["goods_name"]
    hdThumbUrl := activity["hd_thumb_url"]
    minOnSaleGroupPrice := activity["min_on_sale_group_price"]
    maxOnSaleGroupPrice := activity["max_on_sale_group_price"]
    onlineQuantity := activity["online_quantity"]
    priceInfo := activity["price_info"].([]interface{})[0].(map[string]interface{})
    maxPreSalePrice := priceInfo["max_pre_sale_price"]
    minPreSalePrice := priceInfo["min_pre_sale_price"]
    discount := priceInfo["discount"]
    activityQuantity := priceInfo["activity_quantity"]
    activityStockQuantity := priceInfo["activity_stock_quantity"]
    startTime := time.Unix(int64(activity["start_time"].(float64)), 0)
    endTime := time.Unix(int64(activity["end_time"].(float64)), 0)
    status := activity["status"]
    var endOperationTime time.Time
    if activity["end_operation_time"].(float64) != 0 {
      endOperationTime = time.Unix(int64(activity["end_operation_time"].(float64)), 0)
    }
    newGoods := activity["new_goods"]
    if count == 0 {
      _, err = stmtInsert.Exec(activityId, activityName, activityType, goodsId, goodsName, hdThumbUrl, maxOnSaleGroupPrice, minOnSaleGroupPrice, onlineQuantity, maxPreSalePrice, minPreSalePrice, discount, activityQuantity, activityStockQuantity, startTime, endTime, status, endOperationTime, newGoods)
      if err != nil {
        log.Println("save-pdd-activity-data-insert-exec-error: ", err)
      }
    } else {
      _, err = stmtUpdate.Exec(activityName, activityType, goodsId, goodsName, hdThumbUrl, maxOnSaleGroupPrice, minOnSaleGroupPrice, onlineQuantity, maxPreSalePrice, minPreSalePrice, discount, activityQuantity, activityStockQuantity, startTime, endTime, status, endOperationTime, newGoods, activityId)
      if err != nil {
        log.Println("save-pdd-activity-data-update-exec-error: ", err)
      }
    }
  }
  io.WriteString(w, "ok")
}
