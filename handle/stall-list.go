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
      mallName string
      floor int64
      stallNumber string
      phone string
      telephone string
      wechat string
      qq string
      dataUrl string
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
      "mallName": mallName,
      "floor": floor,
      "stallNumber": stallNumber,
      "phone": phone,
      "telephone": telephone,
      "wechat": wechat,
      "qq": qq,
      "dataUrl": dataUrl,
      "url": url,
      "siteType": siteType,
    }
    stallList = append(stallList, stall)
  }
  json.NewEncoder(w).Encode(stallList)
}
