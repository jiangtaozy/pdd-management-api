/*
 * Maintained by jemo from 2020.5.28 to now
 * Created by jemo on 2020.5.28 09:09:11
 * Ad Unit
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdUnit(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  adIdArray := query["adId"]
  adId := adIdArray[0]
  db := database.ConnectDB()
  var (
    adName string
    goodsId int64
    goodsName string
    thumbUrl sql.NullString
  )
  err := db.QueryRow("SELECT pddAdUnit.adId, pddAdUnit.adName, pddAdUnit.goodsId, pddAdUnit.goodsName, pddItem.thumbUrl FROM pddAdUnit LEFT JOIN pddItem ON pddAdUnit.goodsId = pddItem.pddId WHERE pddAdUnit.adId = ?", adId).Scan(&adId, &adName, &goodsId, &goodsName, &thumbUrl)
  if err != nil {
    log.Println("ad-unit-query-error: ", err)
  }
  json.NewEncoder(w).Encode(map[string]interface{}{
    "adId": adId,
    "adName": adName,
    "goodsId": goodsId,
    "goodsName": goodsName,
    "thumbUrl": thumbUrl.String,
  })
}
