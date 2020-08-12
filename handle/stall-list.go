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
  rows, err := db.Query("SELECT id, name, city, mallName, floor, stallNumber, phone, telephone, wechat, qq, dataUrl, url, siteType FROM supplier")
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
