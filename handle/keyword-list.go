/*
 * Maintained by jemo from 2020.10.16 to now
 * Created by jemo on 2020.10.16 11:20:40
 * Keyword List
 * 某一推广单元关键词列表
 */

package handle

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func KeywordList(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  id := query["adId"][0]
  db := database.DB
  rows, err := db.Query(`
    SELECT
      impression,
      click,
      spend,
      orderNum,
      gmv,
      cpm,
      mallFavNum,
      goodsFavNum,
      bid,
      bidPremium,
      bidPremiumValue,
      qualityScore,
      keywordAdIdxOri,
      adId,
      keywordId,
      keyword,
      date
    FROM
      adUnitKeyword
    WHERE
      adId = ?
  `, id)
  if err != nil {
    log.Println("keyword-list-query-error: ", err)
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      impression int64
      click int64
      spend int64
      orderNum int64
      gmv int64
      cpm float64
      mallFavNum int64
      goodsFavNum int64
      bid int64
      bidPremium int64
      bidPremiumValue int64
      qualityScore int64
      keywordAdIdxOri float64
      adId int64
      keywordId int64
      keyword string
      date string
    )
    err = rows.Scan(
      &impression,
      &click,
      &spend,
      &orderNum,
      &gmv,
      &cpm,
      &mallFavNum,
      &goodsFavNum,
      &bid,
      &bidPremium,
      &bidPremiumValue,
      &qualityScore,
      &keywordAdIdxOri,
      &adId,
      &keywordId,
      &keyword,
      &date,
    )
    if err != nil {
      log.Println("keyword-list-scan-error: ", err)
    }
    obj := map[string]interface{}{
      "impression": impression,
      "click": click,
      "spend": spend,
      "orderNum": orderNum,
      "gmv": gmv,
      "cpm": cpm,
      "mallFavNum": mallFavNum,
      "goodsFavNum": goodsFavNum,
      "bid": bid,
      "bidPremium": bidPremium,
      "bidPremiumValue": bidPremiumValue,
      "qualityScore": qualityScore,
      "keywordAdIdxOri": keywordAdIdxOri,
      "adId": adId,
      "keywordId": keywordId,
      "keyword": keyword,
      "date": date,
    }
    list = append(list, obj)
  }
  json.NewEncoder(w).Encode(list)
}
