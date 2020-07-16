/*
 * Maintained by jemo from 2020.7.14 to now
 * Created by jemo on 2020.7.14 15:29:00
 * Ad Head List
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdHeadList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query("SELECT id, headId, headName, dodokCommission, headCommission, coupon, wechatNickname, wechatNumber, pddNickname FROM adHead")
  if err != nil {
    log.Println("ad-head-list-query-error: ", err)
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      id int64
      headId sql.NullInt64
      headName sql.NullString
      dodokCommission sql.NullInt64
      headCommission sql.NullInt64
      coupon sql.NullInt64
      wechatNickname sql.NullString
      wechatNumber sql.NullString
      pddNickname sql.NullString
    )
    if err := rows.Scan(&id, &headId, &headName, &dodokCommission, &headCommission, &coupon, &wechatNickname, &wechatNumber, &pddNickname); err != nil {
      log.Println("ad-head-list-scan-error: ", err)
    }
    adHead := map[string]interface{}{
      "id": id,
      "headId": headId.Int64,
      "headName": headName.String,
      "dodokCommission": dodokCommission.Int64,
      "headCommission": headCommission.Int64,
      "coupon": coupon.Int64,
      "wechatNickname": wechatNickname.String,
      "wechatNumber": wechatNumber.String,
      "pddNickname": pddNickname.String,
    }
    list = append(list, adHead)
  }
  json.NewEncoder(w).Encode(list)
}
