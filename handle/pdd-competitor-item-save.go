/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 15:36:26
 * Pdd Competitor Item
 * 拼多多竞争对手商品
 */

package handle

import (
  "io"
  "encoding/json"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddCompetitorItemSave(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("pdd-competitor-item-save-decode-body-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  db := database.DB
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      pddCompetitorItem
    WHERE
      id = ?
  `, body["id"]).Scan(&count)
  if err != nil {
    log.Println("pdd-competitor-item-save-count-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  if count == 0 {
    stmtInsert, err := db.Prepare(`
      INSERT INTO pddCompetitorItem (
        name,
        price,
        goodsId,
        competitorId,
        relatedItemId
      ) VALUES(?, ?, ?, ?, ?)
    `)
    if err != nil {
      log.Println("pdd-competitor-item-save-insert-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    _, err = stmtInsert.Exec(
      body["name"],
      body["price"],
      body["goodsId"],
      body["competitorId"],
      body["relatedItemId"],
    )
    if err != nil {
      log.Println("pdd-competitor-item-save-insert-exec-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  } else {
    stmtUpdate, err := db.Prepare(`
      UPDATE
        pddCompetitorItem
      SET
        name = ?,
        price = ?,
        goodsId = ?,
        competitorId = ?,
        relatedItemId = ?
      WHERE id = ?
    `)
    if err != nil {
      log.Println("pdd-competitor-item-save-update-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    _, err = stmtUpdate.Exec(
      body["name"],
      body["price"],
      body["goodsId"],
      body["competitorId"],
      body["relatedItemId"],
      body["id"],
    )
    if err != nil {
      log.Println("pdd-competitor-item-save-update-exec-error", err)
      http.Error(w, err.Error(), 500)
      return
    }
  }
  io.WriteString(w, "ok")
}
