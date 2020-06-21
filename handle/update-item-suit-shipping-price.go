/*
 * Maintained by jemo from 2020.5.14 to now
 * Created by jemo on 2020.5.14 17:30:51
 * Update Item Suit Price And Shipping Price
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func UpdateItemSuitShippingPrice(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("update-item-suit-shipping-price-data-error: ", err)
  }
  var id = body["id"]
  var suitPrice = body["suitPrice"]
  var shippingPrice = body["shippingPrice"]
  var forSell = body["forSell"]
  db := database.DB
  stmtUpdate, err := db.Prepare("UPDATE item SET suitPrice = ?, shippingPrice = ?, forSell = ? WHERE id = ?")
  if err != nil {
    log.Println("update-item-suit-shipping-price-prepare-error: ", err)
  }
  _, err = stmtUpdate.Exec(suitPrice, shippingPrice, forSell, id)
  if err != nil {
    log.Println("update-item-suit-shipping-price-exec-error: ", err)
  }
  io.WriteString(w, "ok")
}
