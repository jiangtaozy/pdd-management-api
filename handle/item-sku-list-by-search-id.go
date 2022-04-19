/*
 * Maintained by jemo from 2022.4.16 to now
 * Created by jemo on 2022.4.16 14:50:21
 * Item Sku List by Search Id
 */

package handle

import (
  "log"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func ItemSkuListBySearchId(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  searchId := query["searchId"][0]
  db := database.DB
  rows, err := db.Query("select itemSku.id, itemSku.skuName, itemSku.price, itemSku.canBookCount, itemSku.shortSkuName, itemSkuNum.shortSkuNum from itemSku left join itemSkuNum on itemSku.shortSkuName = itemSkuNum.shortSkuName where searchId = ?", searchId)
  if err != nil {
    log.Println("item-sku-list-by-search-id-query-error: ", err)
  }
  defer rows.Close()
  var itemSkuList []interface{}
  for rows.Next() {
    var (
      id int64
      skuName string
      price int64
      canBookCount int64
      shortSkuName sql.NullString
      shortSkuNum sql.NullString
    )
    err = rows.Scan(&id, &skuName, &price, &canBookCount, &shortSkuName, &shortSkuNum)
    if err != nil {
      log.Println("item-sku-list-by-search-id-scan-error: ", err)
    }
    itemSku := map[string]interface{}{
      "id": id,
      "skuName": skuName,
      "price": price,
      "canBookCount": canBookCount,
      "shortSkuName": shortSkuName.String,
      "shortSkuNum": shortSkuNum.String,
    }
    itemSkuList = append(itemSkuList, itemSku)
  }
  json.NewEncoder(w).Encode(itemSkuList)
}
