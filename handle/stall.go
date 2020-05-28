/*
 * Maintained by jemo from 2020.5.28 to now
 * Created by jemo on 2020.5.28 21:44:13
 * Stall
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func Stall(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("stall-decode-body-error: ", err)
  }
  db := database.ConnectDB()
  var count int
  err = db.QueryRow("SELECT COUNT(*) FROM stall WHERE id = ?", body["id"]).Scan(&count)
  if err != nil {
    log.Println("stall-count-error: ", err)
  }
  if count == 0 {
    stmtInsert, err := db.Prepare("INSERT INTO stall (name, cityName, mallName, floor, stallNumber, phone, telephone, wechat, qq, dataUrl, stallUrl) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
      log.Println("stall-insert-prepare-error: ", err)
    }
    _, err = stmtInsert.Exec(body["name"], body["cityName"], body["mallName"], body["floor"], body["stallNumber"], body["phone"], body["telephone"], body["wechat"], body["qq"], body["dataUrl"], body["stallUrl"])
    if err != nil {
      log.Println("stall-insert-exec-error: ", err)
    }
  } else {
    stmtUpdate, err := db.Prepare("UPDATE stall SET name = ?, cityName = ?, mallName = ?, floor = ?, stallNumber = ?, phone = ?, telephone = ?, wechat = ?, qq = ?, dataUrl = ?, stallUrl = ? WHERE id = ?")
    if err != nil {
      log.Println("stall-update-prepare-error: ", err)
    }
    _, err = stmtUpdate.Exec(body["name"], body["cityName"], body["mallName"], body["floor"], body["stallNumber"], body["phone"], body["telephone"], body["wechat"], body["qq"], body["dataUrl"], body["stallUrl"], body["id"])
    if err != nil {
      log.Println("stall-update-exec-error", err)
    }
  }

  io.WriteString(w, "ok")
}
