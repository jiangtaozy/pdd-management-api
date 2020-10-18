/*
 * Maintained by jemo from 2020.5.28 to now
 * Created by jemo on 2020.5.28 22:25:50
 * Stall List
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
  "database/sql"
)

func StallList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      supplier.id,
      supplier.name,
      supplier.city,
      supplier.mallName,
      supplier.floor,
      supplier.stallNumber,
      supplier.phone,
      supplier.telephone,
      supplier.wechat,
      supplier.qq,
      supplier.dataUrl,
      supplier.url,
      supplier.siteType
    FROM
      supplier
  `)
  /*
  rows, err := db.Query(`
    SELECT
      supplier.id,
      supplier.name,
      supplier.city,
      supplier.mallName,
      supplier.floor,
      supplier.stallNumber,
      supplier.phone,
      supplier.telephone,
      supplier.wechat,
      supplier.qq,
      supplier.dataUrl,
      supplier.url,
      supplier.siteType,
      item.name,
      pddItem.goodsName,
      adData.impression,
      adData.click,
      adData.spend,
      adData.orderNum,
      adData.gmv,
      adData.mallFavNum,
      adData.goodsFavNum,
      itemOrder.orderStatus,
      itemOrder.afterSaleStatus,
      itemOrder.orderStatusStr,
      itemOrder.userPaidAmount,
      itemOrder.platformDiscount,
      itemOrder.orderId,
      order1688.actualPayment
    FROM
      supplier
    LEFT JOIN item
      ON supplier.id = item.supplierId
    LEFT JOIN pddItem
      ON item.searchId = pddItem.outGoodsSn
    LEFT JOIN pddAdUnit
      ON pddItem.pddId = pddAdUnit.goodsId
    LEFT JOIN
      (SELECT adId, SUM(impression) impression, SUM(click) click, SUM(spend) spend, SUM(orderNum) orderNum, SUM(gmv) gmv, SUM(mallFavNum) mallFavNum, SUM(goodsFavNum) goodsFavNum
      FROM pddAdUnitDailyData
      GROUP BY adId) AS adData
      ON pddAdUnit.adId = adData.adId
    LEFT JOIN itemOrder AS itemOrder
      ON pddItem.pddId = itemOrder.productId
    LEFT JOIN order1688 AS order1688
      ON itemOrder.outerOrderId = order1688.orderId
  `)
  */
  if err != nil {
    log.Println("stall-list-query-error: ", err)
  }
  defer rows.Close()
  var stallList []interface{}
  for rows.Next() {
    var (
      id int64
      name string
      city string
      mallName sql.NullString
      floor sql.NullInt64
      stallNumber sql.NullString
      phone sql.NullString
      telephone sql.NullString
      wechat sql.NullString
      qq sql.NullString
      dataUrl sql.NullString
      url string
      siteType int64
    )
    if err := rows.Scan(&id, &name, &city, &mallName, &floor, &stallNumber, &phone, &telephone, &wechat, &qq, &dataUrl, &url, &siteType); err != nil {
      log.Println("stall-list-scan-error: ", err)
    }
    stall := map[string]interface{}{
      "id": id,
      "name": name,
      "city": city,
      "mallName": mallName.String,
      "floor": floor.Int64,
      "stallNumber": stallNumber.String,
      "phone": phone.String,
      "telephone": telephone.String,
      "wechat": wechat.String,
      "qq": qq.String,
      "dataUrl": dataUrl.String,
      "url": url,
      "siteType": siteType,
    }
    stallList = append(stallList, stall)
  }
  json.NewEncoder(w).Encode(stallList)
}
