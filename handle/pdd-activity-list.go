/*
 * Maintained by jemo from 2020.7.21 to now
 * Created by jemo on 2020.7.21 18:19:57
 * Pdd Activity List
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddActivityList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT id, activityId, activityName, activityType, goodsId, goodsName, hdThumbUrl, maxOnSaleGroupPrice, minOnSaleGroupPrice, onlineQuantity, maxPreSalePrice, minPreSalePrice, discount, activityQuantity, activityStockQuantity, startTime, endTime, status, endOperationTime, newGoods FROM pddActivity`)
  if err != nil {
    log.Println("pdd-activity-list-query-error: ", err)
  }
  defer rows.Close()
  var activityList []interface{}
  for rows.Next() {
    var (
      id int64
      activityId int64
      activityName string
      activityType int64
      goodsId int64
      goodsName string
      hdThumbUrl string
      maxOnSaleGroupPrice int64
      minOnSaleGroupPrice int64
      onlineQuantity int64
      maxPreSalePrice int64
      minPreSalePrice int64
      discount int64
      activityQuantity int64
      activityStockQuantity int64
      startTime string
      endTime string
      status int64
      endOperationTime string
      newGoods bool
    )
    if err := rows.Scan(&id, &activityId, &activityName, &activityType, &goodsId, &goodsName, &hdThumbUrl, &maxOnSaleGroupPrice, &minOnSaleGroupPrice, &onlineQuantity, &maxPreSalePrice, &minPreSalePrice, &discount, &activityQuantity, &activityStockQuantity, &startTime, &endTime, &status, &endOperationTime, &newGoods); err != nil {
      log.Println("pdd-activity-list-scan-error: ", err)
    }
    activity := map[string]interface{}{
      "id": id,
      "activityId": activityId,
      "activityName": activityName,
      "activityType": activityType,
      "goodsId": goodsId,
      "goodsName": goodsName,
      "hdThumbUrl": hdThumbUrl,
      "maxOnSaleGroupPrice": maxOnSaleGroupPrice,
      "minOnSaleGroupPrice": minOnSaleGroupPrice,
      "onlineQuantity": onlineQuantity,
      "maxPreSalePrice": maxPreSalePrice,
      "minPreSalePrice": minPreSalePrice,
      "discount": discount,
      "activityQuantity": activityQuantity,
      "activityStockQuantity": activityStockQuantity,
      "startTime": startTime,
      "endTime": endTime,
      "status": status,
      "endOperationTime": endOperationTime,
      "newGoods": newGoods,
    }
    activityList = append(activityList, activity)
  }
  json.NewEncoder(w).Encode(activityList)
}
