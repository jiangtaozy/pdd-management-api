/*
 * Maintained by jemo from 2021.1.9 to now
 * Created by jemo on 2021.1.9 18:30:13
 * Pdd Item Price History Sync
 */

package handle

import (
  "log"
  "time"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddItemPriceHistorySync(pddId float64, skuGroupPriceMin float64, skuGroupPriceMax float64) {
  now := time.Now()
  today := now.Format("2006-01-02")
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddItemPriceHistory (
      pddId,
      date,
      skuGroupPriceMin,
      skuGroupPriceMax
    ) VALUES (
      ?, ?, ?, ?
    )
  `)
  if err != nil {
    log.Println("pdd-item-price-history-sync-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      pddItemPriceHistory
    SET
      skuGroupPriceMin = ?,
      skuGroupPriceMax = ?
    WHERE
      pddId = ?
      AND
      date = ?
  `)
  if err != nil {
    log.Println("pdd-item-price-history-sync-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      pddItemPriceHistory
    WHERE
      pddId = ?
      AND
      date = ?
  `, pddId, today).Scan(&count)
  if err != nil {
    log.Println("pdd-item-price-history-sync-query-count-error: ", err)
  }
  if count == 0 {
    _, err = stmtInsert.Exec(
      pddId,
      today,
      skuGroupPriceMin,
      skuGroupPriceMax,
    )
    if err != nil {
      log.Println("pdd-item-price-history-sync-insert-exec-error: ", err)
    }
  } else {
    _, err = stmtUpdate.Exec(
      skuGroupPriceMin,
      skuGroupPriceMax,
      pddId,
      today,
    )
    if err != nil {
      log.Println("pdd-item-price-history-sync-update-exec-error: ", err)
    }
  }
}
