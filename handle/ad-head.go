/*
 * Maintained by jemo from 2020.7.14 to now
 * Created by jemo on 2020.7.14 16:41:18
 * Ad Head
 * 推广团长
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func AdHead(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("ad-head-decode-body-error: ", err)
  }
  db := database.DB
  var count int
  err = db.QueryRow("SELECT COUNT(*) FROM adHead WHERE id = ?", body["id"]).Scan(&count)
  if err != nil {
    log.Println("ad-head-count-error: ", err)
  }
  if count == 0 {
    stmtInsert, err := db.Prepare("INSERT INTO adHead (headId, headName, dodokCommission, headCommission, coupon, wechatNickname, wechatNumber, pddNickname) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
      log.Println("ad-head-insert-prepare-error: ", err)
    }
    _, err = stmtInsert.Exec(body["headId"], body["headName"], body["dodokCommission"], body["headCommission"], body["coupon"], body["wechatNickname"], body["wechatNumber"], body["pddNickname"])
    if err != nil {
      log.Println("ad-head-insert-exec-error: ", err)
    }
  } else {
    stmtUpdate, err := db.Prepare("UPDATE adHead SET headId = ?, headName = ?, dodokCommission = ?, headCommission = ?, coupon = ?, wechatNickname = ?, wechatNumber = ?, pddNickname = ? WHERE id = ?")
    if err != nil {
      log.Println("ad-head-update-prepare-error: ", err)
    }
    _, err = stmtUpdate.Exec(body["headId"], body["headName"], body["dodokCommission"], body["headCommission"], body["coupon"], body["wechatNickname"], body["wechatNumber"], body["pddNickname"], body["id"])
    if err != nil {
      log.Println("ad-head-update-exec-error", err)
    }
  }
  io.WriteString(w, "ok")
}
