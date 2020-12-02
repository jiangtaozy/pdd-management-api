/*
 * Maintained by jemo from 2020.9.10 to now
 * Created by jemo on 2020.9.10 16:37:32
 * Dy Item List
 * 抖音商品列表
 */

package handle

import (
  "log"
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func DyItemList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      dyItem.checkStatus,
      dyItem.createTime,
      dyItem.discountPrice,
      dyItem.extra,
      dyItem.firstCid,
      dyItem.img,
      dyItem.marketPrice,
      dyItem.mobile,
      dyItem.name,
      dyItem.outProductId,
      dyItem.payType,
      dyItem.productId,
      dyItem.productIdStr,
      dyItem.recommendRemark,
      dyItem.secondCid,
      dyItem.settlementPrice,
      dyItem.specId,
      dyItem.status,
      dyItem.thirdCid,
      dyItem.updateTime,
      item.price,
      item.detailUrl,
      item.name AS womenName
    FROM
      dyItem
    LEFT JOIN
      item
    ON
      dyItem.outProductId = item.womenProductId
  `)
  if err != nil {
    log.Println("dy-item-list-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      checkStatus int64
      createTime sql.NullString
      discountPrice sql.NullInt64
      extra sql.NullString
      firstCid sql.NullInt64
      img sql.NullString
      marketPrice sql.NullInt64
      mobile sql.NullString
      name sql.NullString
      outProductId sql.NullInt64
      payType sql.NullInt64
      productId sql.NullInt64
      productIdStr sql.NullString
      recommendRemark sql.NullString
      secondCid sql.NullInt64
      settlementPrice sql.NullInt64
      specId sql.NullInt64
      status sql.NullInt64
      thirdCid sql.NullInt64
      updateTime sql.NullString
      price sql.NullFloat64
      detailUrl sql.NullString
      womenName sql.NullString
    )
    err = rows.Scan(
      &checkStatus,
      &createTime,
      &discountPrice,
      &extra,
      &firstCid,
      &img,
      &marketPrice,
      &mobile,
      &name,
      &outProductId,
      &payType,
      &productId,
      &productIdStr,
      &recommendRemark,
      &secondCid,
      &settlementPrice,
      &specId,
      &status,
      &thirdCid,
      &updateTime,
      &price,
      &detailUrl,
      &womenName,
    )
    if err != nil {
      log.Println("dy-item-list-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    item := map[string]interface{}{
      "checkStatus": checkStatus,
      "createTime": createTime.String,
      "discountPrice": discountPrice.Int64,
      "extra": extra.String,
      "firstCid": firstCid.Int64,
      "img": img.String,
      "marketPrice": marketPrice.Int64,
      "mobile": mobile.String,
      "name": name.String,
      "outProductId": outProductId.Int64,
      "payType": payType.Int64,
      "productId": productId.Int64,
      "productIdStr": productIdStr.String,
      "recommendRemark": recommendRemark.String,
      "secondCid": secondCid.Int64,
      "settlementPrice": settlementPrice.Int64,
      "specId": specId.Int64,
      "status": status.Int64,
      "thirdCid": thirdCid.Int64,
      "updateTime": updateTime.String,
      "price": price.Float64,
      "detailUrl": detailUrl.String,
      "womenName": womenName.String,
    }
    list = append(list, item)
  }
  json.NewEncoder(w).Encode(list)
}
