/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 17:10:18
 * Pdd Competitor Item Sale
 * 拼多多竞争对手商品销量
 */

package handle

import (
  "io"
  "log"
  "time"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddCompetitorItemSaleSave(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("pdd-competitor-item-sale-save-decode-body-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  date, err := time.Parse(time.RFC3339, body["date"].(string))
  if err != nil {
    log.Println("pdd-competitor-item-sale-save-parse-date-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  db := database.DB
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      pddCompetitorItemSale
    WHERE
      goodsId = ?
    AND
      date = ?
  `, body["goodsId"], date).Scan(&count)
  if err != nil {
    log.Println("pdd-competitor-item-sale-save-count-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  if count == 0 {
    stmtInsert, err := db.Prepare(`
      INSERT INTO pddCompetitorItemSale (
        goodsId,
        date,
        sale
      ) VALUES(?, ?, ?)
    `)
    if err != nil {
      log.Println("pdd-competitor-item-sale-save-insert-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    _, err = stmtInsert.Exec(
      body["goodsId"],
      date,
      body["sale"],
    )
    if err != nil {
      log.Println("pdd-competitor-item-sale-save-insert-exec-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  } else {
    stmtUpdate, err := db.Prepare(`
      UPDATE
        pddCompetitorItemSale
      SET
        sale = ?
      WHERE
        goodsId = ?
      AND
        date = ?
    `)
    if err != nil {
      log.Println("pdd-competitor-item-sale-save-update-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    _, err = stmtUpdate.Exec(
      body["sale"],
      body["goodsId"],
      date,
    )
    if err != nil {
      log.Println("pdd-competitor-item-sale-save-update-exec-error", err)
      http.Error(w, err.Error(), 500)
      return
    }
  }
  io.WriteString(w, "ok")
}
