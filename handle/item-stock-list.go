/*
 * Maintained by jemo from 2021.2.13 to now
 * Created by jemo on 2021.2.13 17:19:17
 * Item Stock List
 * 商品库存列表
 */

package handle

import (
  "log"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func ItemStockList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      pddSku.pddId,
      pddSku.outGoodsSn,
      pddSku.activityGroupPrice,
      pddSku.groupPrice,
      pddSku.isOnsale,
      pddSku.normalPrice,
      pddSku.outSkuSn,
      pddSku.skuQuantity,
      pddSku.skuSoldQuantity,
      pddSku.spec,
      pddSku.specColor,
      pddSku.specSize,
      womenSku.searchId,
      womenSku.productId,
      womenSku.skuDesc,
      womenSku.ycAvailNum,
      womenSku.ycStockTips,
      womenSku.skuColor,
      womenSku.skuSize,
      pddItem.isPreSale,
      womenItem.isCloudWarehouse
    FROM pddItemSku AS pddSku
    LEFT JOIN womenItemCloudWarehouseSku AS womenSku
      ON pddSku.outGoodsSn = womenSku.searchId
      AND pddSku.specColor = womenSku.skuColor
      AND pddSku.specSize = womenSku.skuSize
    LEFT JOIN pddItem
      ON pddSku.pddId = pddItem.pddId
    LEFT JOIN womenItem
      ON pddSku.outGoodsSn = womenItem.searchId
  `)
  if err != nil {
    log.Println("item-stock-list-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      pddId sql.NullInt64
      outGoodsSn sql.NullString
      activityGroupPrice sql.NullInt64
      groupPrice sql.NullInt64
      isOnSale sql.NullBool
      normalPrice sql.NullInt64
      outSkuSn sql.NullString
      skuQuantity sql.NullInt64
      skuSoldQuantity sql.NullInt64
      spec sql.NullString
      specColor sql.NullString
      specSize sql.NullString
      searchId sql.NullInt64
      productId sql.NullInt64
      skuDesc sql.NullString
      ycAvailNum sql.NullInt64
      ycStockTips sql.NullString
      skuColor sql.NullString
      skuSize sql.NullString
      isPreSale sql.NullBool
      isCloudWarehouse sql.NullBool
    )
    if err := rows.Scan(
      &pddId,
      &outGoodsSn,
      &activityGroupPrice,
      &groupPrice,
      &isOnSale,
      &normalPrice,
      &outSkuSn,
      &skuQuantity,
      &skuSoldQuantity,
      &spec,
      &specColor,
      &specSize,
      &searchId,
      &productId,
      &skuDesc,
      &ycAvailNum,
      &ycStockTips,
      &skuColor,
      &skuSize,
      &isPreSale,
      &isCloudWarehouse,
    ); err != nil {
      log.Println("item-stock-list-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    sku := map[string]interface{}{
      "pddId": pddId.Int64,
      "outGoodsSn": outGoodsSn.String,
      "activityGroupPrice": activityGroupPrice.Int64,
      "groupPrice": groupPrice.Int64,
      "isOnSale": isOnSale.Bool,
      "normalPrice": normalPrice.Int64,
      "outSkuSn": outSkuSn.String,
      "skuQuantity": skuQuantity.Int64,
      "skuSoldQuantity": skuSoldQuantity.Int64,
      "spec": spec.String,
      "specColor": specColor.String,
      "specSize": specSize.String,
      "searchId": searchId.Int64,
      "productId": productId.Int64,
      "skuDesc": skuDesc.String,
      "ycAvailNum": ycAvailNum.Int64,
      "ycStockTips": ycStockTips.String,
      "skuColor": skuColor.String,
      "skuSize": skuSize.String,
      "isPreSale": isPreSale.Bool,
      "isCloudWarehouse": isCloudWarehouse.Bool,
    }
    list = append(list, sku)
  }
  json.NewEncoder(w).Encode(list)
}
