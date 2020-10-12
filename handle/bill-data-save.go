/*
 * Maintained by jemo from 2020.10.12 to now
 * Created by jemo on 2020.10.12 17:34:36
 * Bill data save
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

func BillDataSave(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  decoder := json.NewDecoder(r.Body)
  decoder.UseNumber()
  err := decoder.Decode(&body)
  if err != nil {
    log.Println("bill-data-save-json-decode-error: ", err)
  }
  result := body["result"].(map[string]interface{})
  data := result["data"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO bill (
      id,
      amount,
      classId,
      createdAt,
      extraInfo,
      financeId,
      flowId,
      mallId,
      note,
      orderSn,
      relatedId,
      shouldNotShowOrderSnInMms,
      type
    )
    VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("bill-data-save-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  for i := 0; i < len(data); i++ {
    bill := data[i].(map[string]interface{})
    for key, val := range bill {
      n, ok := val.(json.Number)
      if !ok {
        continue
      }
      if i, err := n.Int64(); err == nil {
        bill[key] = i
        continue
      }
      if f, err := n.Float64(); err == nil {
        bill[key] = f
        continue
      }
    }
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        bill
      WHERE
        id = ?
    `, bill["id"]).Scan(&count)
    if err != nil {
      log.Println("bill-data-save-count-error: ", err)
    }
    createdAt := time.Unix(bill["createdAt"].(int64), 0)
    if count == 0 {
      _, err = stmtInsert.Exec(
        bill["id"],
        bill["amount"],
        bill["classId"],
        createdAt,
        bill["extraInfo"],
        bill["financeId"],
        bill["flowId"],
        bill["mallId"],
        bill["note"],
        bill["orderSn"],
        bill["relatedId"],
        bill["shouldNotShowOrderSnInMms"],
        bill["type"],
      )
      if err != nil {
        log.Println("bill-data-save-insert-error: ", err)
        log.Println("bill: ", bill)
      }
    }
    io.WriteString(w, "ok")
  }
}
