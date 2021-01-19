/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 14:17:56
 * Pdd Competitor
 * 拼多多竞争对手
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddCompetitorSave(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("pdd-competitor-save-decode-body-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  db := database.DB
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      pddCompetitor
    WHERE
      id = ?
  `, body["id"]).Scan(&count)
  if err != nil {
    log.Println("pdd-competitor-save-count-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  if count == 0 {
    stmtInsert, err := db.Prepare(`
      INSERT INTO pddCompetitor (
        name,
        advantage,
        disadvantage
      ) VALUES(?, ?, ?)
    `)
    if err != nil {
      log.Println("pdd-competitor-save-insert-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    _, err = stmtInsert.Exec(
      body["name"],
      body["advantage"],
      body["disadvantage"],
    )
    if err != nil {
      log.Println("pdd-competitor-save-insert-exec-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  } else {
    stmtUpdate, err := db.Prepare(`
      UPDATE
        pddCompetitor
      SET
        name = ?,
        advantage = ?,
        disadvantage = ?
      WHERE id = ?
    `)
    if err != nil {
      log.Println("pdd-competitor-save-update-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    _, err = stmtUpdate.Exec(
      body["name"],
      body["advantage"],
      body["disadvantage"],
      body["id"],
    )
    if err != nil {
      log.Println("pdd-competitor-save-update-exec-error", err)
      http.Error(w, err.Error(), 500)
      return
    }
  }
  io.WriteString(w, "ok")
}
