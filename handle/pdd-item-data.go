/*
 * Maintained by jemo from 2020.8.27 to now
 * Created by jemo on 2020.8.27 9:57:59
 * Pdd Item Data
 */

package handle

import (
  "log"
  "time"
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddItemData(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  idList := query["id"]
  id := idList[0]
  db := database.DB
  // unit data
  rows, err := db.Query(`
    SELECT
      unit.adId,
      unit.adName,
      data.impression,
      data.click,
      data.spend,
      data.orderNum,
      data.gmv,
      data.mallFavNum,
      data.goodsFavNum,
      data.date
    FROM pddAdUnit AS unit
    INNER JOIN pddAdUnitDailyData AS data
      ON data.adId = unit.adId
    WHERE
      unit.goodsId = ?
    ORDER BY
      data.date ASC
  `, id)
  if err != nil {
    log.Println("pdd-item-data-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var unitList []interface{}
  for rows.Next() {
    var (
      adId int64
      adName string
      impression sql.NullInt64
      click sql.NullInt64
      spend sql.NullInt64
      orderNum sql.NullInt64
      gmv sql.NullInt64
      mallFavNum sql.NullInt64
      goodsFavNum sql.NullInt64
      date time.Time
    )
    err := rows.Scan(
      &adId,
      &adName,
      &impression,
      &click,
      &spend,
      &orderNum,
      &gmv,
      &mallFavNum,
      &goodsFavNum,
      &date,
    )
    if err != nil {
      log.Println("pdd-item-data-scan-data-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    unit := map[string]interface{}{
      "adId": adId,
      "adName": adName,
      "impression": impression.Int64,
      "click": click.Int64,
      "spend": spend.Int64,
      "orderNum": orderNum.Int64,
      "gmv": gmv.Int64,
      "mallFavNum": mallFavNum.Int64,
      "goodsFavNum": goodsFavNum.Int64,
      "date": date,
    }
    unitList = append(unitList, unit)
  }
  // order data
  orderRows, err := db.Query(`
    SELECT
      itemOrder.orderId,
      itemOrder.orderStatusStr,
      itemOrder.platformDiscount,
      itemOrder.userPaidAmount,
      itemOrder.paymentTime,
      order1688.actualPayment
    FROM itemOrder
    LEFT JOIN order1688
      ON itemOrder.outerOrderId = order1688.orderId
    WHERE itemOrder.productId = ?
      AND itemOrder.orderStatus = 1
      AND (itemOrder.afterSaleStatus IS NULL
      OR itemOrder.afterSaleStatus = 6
      OR itemOrder.afterSaleStatus = 12
      OR itemOrder.afterSaleStatus = 16)
  `, id)
  if err != nil {
    log.Println("pdd-item-data-query-order-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer orderRows.Close()
  var orderList []interface{}
  for orderRows.Next() {
    var (
      orderId string
      orderStatusStr string
      platformDiscount int64
      userPaidAmount int64
      paymentTime time.Time
      actualPayment sql.NullFloat64
    )
    err := orderRows.Scan(
      &orderId,
      &orderStatusStr,
      &platformDiscount,
      &userPaidAmount,
      &paymentTime,
      &actualPayment,
    )
    if err != nil {
      log.Println("pdd-item-data-scan-order-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    order := map[string]interface{}{
      "orderId": orderId,
      "orderStatusStr": orderStatusStr,
      "platformDiscount": platformDiscount,
      "userPaidAmount": userPaidAmount,
      "paymentTime": paymentTime,
      "actualPayment": actualPayment.Float64,
    }
    orderList = append(orderList, order)
  }
  // item data
  var dataList []interface{}
  for i := 0; i < len(unitList); i++ {
    unit := unitList[i].(map[string]interface{})
    date := unit["date"]
    var hasUnitInDataList = false
    for j := 0; j < len(dataList); j++ {
      data := dataList[j].(map[string]interface{})
      dataDate := data["date"]
      if date == dataDate {
        data["impression"] = data["impression"].(int64) + unit["impression"].(int64)
        data["click"] = data["click"].(int64) + unit["click"].(int64)
        data["spend"] = data["spend"].(int64) + unit["spend"].(int64)
        data["orderNum"] = data["orderNum"].(int64) + unit["orderNum"].(int64)
        data["gmv"] = data["gmv"].(int64) + unit["gmv"].(int64)
        data["mallFavNum"] = data["mallFavNum"].(int64) + unit["mallFavNum"].(int64)
        data["goodsFavNum"] = data["goodsFavNum"].(int64) + unit["goodsFavNum"].(int64)
        hasUnitInDataList = true
        break
      }
    }
    if !hasUnitInDataList {
      dataList = append(dataList, unit)
    }
  }
  for i := 0; i < len(orderList); i++ {
    order := orderList[i].(map[string]interface{})
    date := order["paymentTime"].(time.Time)
    for j := 0; j < len(dataList); j++ {
      data := dataList[j].(map[string]interface{})
      dateTime := data["date"].(time.Time)
      if err != nil {
        log.Println("pdd-item-data-parse-data-date-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
      if dateTime.Year() == date.Year() &&
        dateTime.Month() == date.Month() &&
        dateTime.Day() == date.Day() {
        if data["platformDiscount"] == nil {
          data["platformDiscount"] = order["platformDiscount"]
        } else {
          data["platformDiscount"] = data["platformDiscount"].(int64) + order["platformDiscount"].(int64)
        }
        if data["userPaidAmount"] == nil {
          data["userPaidAmount"] = order["userPaidAmount"]
        } else {
          data["userPaidAmount"] = data["userPaidAmount"].(int64) + order["userPaidAmount"].(int64)
        }
        if data["actualPayment"] == nil {
          data["actualPayment"] = order["actualPayment"]
          data["realOrderNum"] = 1
        } else {
          data["actualPayment"] = data["actualPayment"].(float64) + order["actualPayment"].(float64)
          data["realOrderNum"] = data["realOrderNum"].(int) + 1
        }
      }
    }
  }
  json.NewEncoder(w).Encode(dataList)
}
