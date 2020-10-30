/*
 * Maintained by jemo from 2020.10.10 to now
 * Created by jemo on 2020.10.10 17:14:13
 * After sale order data save
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

func AfterSaleOrderDataSave(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("after-sale-order-data-save-decode-err: ", err)
  }
  afterSaleOrderDataString := body["afterSaleOrderData"].(string)
  afterSaleOrderDataMap := make(map[string]interface{})
  err = json.Unmarshal([]byte(afterSaleOrderDataString), &afterSaleOrderDataMap)
  if err != nil {
    log.Println("after-sale-order-data-save-json-unmarshal-error: ", err)
  }
  result := afterSaleOrderDataMap["result"].(map[string]interface{})
  afterSalesList := result["list"].([]interface{})
  db := database.DB
  stmtUpdate, err := db.Prepare(`
    UPDATE
      itemOrder
    SET
      afterSaleStatus = ?,
      afterSaleApplyTime = ?
    WHERE
      orderId = ?
  `)
  if err != nil {
    log.Println("after-sale-order-data-save-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(afterSalesList); i++ {
    order := afterSalesList[i].(map[string]interface{})
    afterSaleApplyTime := time.Unix(int64(order["createdAt"].(float64)), 0)
    _, err = stmtUpdate.Exec(
      order["afterSalesStatus"],
      afterSaleApplyTime,
      order["orderSn"],
    )
    if err != nil {
      log.Println("after-sale-order-data-save-update-exec-error: ", err)
    }
  }
  io.WriteString(w, "ok")
}

