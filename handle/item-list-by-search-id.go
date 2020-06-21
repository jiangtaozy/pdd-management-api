/*
 * Maintained by jemo from 2020.5.14 to now
 * Created by jemo on 2020.5.14 11:47:51
 * Item List by Search Id
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func ItemListBySearchId(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  idArray := query["searchId"]
  searchId := idArray[0]
  db := database.DB
  rows, err := db.Query("SELECT item.id, item.name, item.price, item.imgUrl, item.detailUrl, item.siteType, item.originalId, item.supplierId, item.saleQuantity, item.quantitySumMonth, item.gmv30dRt, item.searchId, item.suitPrice, item.shippingPrice, item.forSell, item.imgUrlOf100x100, supplier.name, supplier.creditLevel, supplier.shopRepurchaseRate, supplier.province, supplier.city FROM item LEFT JOIN supplier ON item.supplierId = supplier.id WHERE item.searchId = ?", searchId)
  if err != nil {
    log.Println("item-list-query-error: ", err)
  }
  defer rows.Close()
  var itemList []interface{}
  for rows.Next() {
    var (
      id int64
      name string
      price float64
      imgUrl string
      detailUrl string
      siteType int32
      originalId string
      supplierId int64
      saleQuantity int64
      quantitySumMonth int64
      gmv30dRt float64
      searchId int64
      suitPrice float64
      shippingPrice float64
      forSell bool
      imgUrlOf100x100 string
      supplierName string
      supplierCreditLevel int32
      supplierShopRepurchaseRate float32
      supplierProvince string
      supplierCity string
    )
    if err := rows.Scan(&id, &name, &price, &imgUrl, &detailUrl, &siteType, &originalId, &supplierId, &saleQuantity, &quantitySumMonth, &gmv30dRt, &searchId, &suitPrice, &shippingPrice, &forSell, &imgUrlOf100x100, &supplierName, &supplierCreditLevel, &supplierShopRepurchaseRate, &supplierProvince, &supplierCity); err != nil {
      log.Println("item-list-scan-error: ", err)
    }
    item := map[string]interface{}{
      "id": id,
      "name": name,
      "price": price,
      "imgUrl": imgUrl,
      "detailUrl": detailUrl,
      "siteType": siteType,
      "originalId": originalId,
      "supplierId": supplierId,
      "saleQuantity": saleQuantity,
      "quantitySumMonth": quantitySumMonth,
      "gmv30dRt": gmv30dRt,
      "searchId": searchId,
      "suitPrice": suitPrice,
      "shippingPrice": shippingPrice,
      "suitShippingPrice": suitPrice + shippingPrice,
      "forSell": forSell,
      "imgUrlOf100x100": imgUrlOf100x100,
      "supplierName": supplierName,
      "supplierCreditLevel": supplierCreditLevel,
      "supplierShopRepurchaseRate": supplierShopRepurchaseRate,
      "supplierProvince": supplierProvince,
      "supplierCity": supplierCity,
    }
    itemList = append(itemList, item)
  }
  json.NewEncoder(w).Encode(itemList)
}
