/*
 * Maintained by jemo from 2020.10.14 to now
 * Created by jemo on 2020.10.14 15:40:09
 * Bill list
 * 货款明细
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func BillList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
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
    FROM
      bill
    ORDER BY
      createdAt DESC
  `)
  if err != nil {
    log.Println("bill-list-query-error: ", err)
  }
  defer rows.Close()
  var billList []interface{}
  for rows.Next() {
    var (
      id int64
      amount int64
      classId int64
      createdAt string
      extraInfo sql.NullString
      financeId int64
      flowId int64
      mallId int64
      note sql.NullString
      orderSn sql.NullString
      relatedId sql.NullInt64
      shouldNotShowOrderSnInMms sql.NullBool
      billType sql.NullInt64
    )
    if err := rows.Scan(
      &id,
      &amount,
      &classId,
      &createdAt,
      &extraInfo,
      &financeId,
      &flowId,
      &mallId,
      &note,
      &orderSn,
      &relatedId,
      &shouldNotShowOrderSnInMms,
      &billType,
    ); err != nil {
      log.Println("bill-list-scan-error: ", err)
    }
    bill := map[string]interface{}{
      "id": id,
      "amount": amount,
      "classId": classId,
      "createdAt": createdAt,
      "extraInfo": extraInfo.String,
      "financeId": financeId,
      "flowId": flowId,
      "mallId": mallId,
      "note": note.String,
      "orderSn": orderSn.String,
      "relatedId": relatedId.Int64,
      "shouldNotShowOrderSnInMms": shouldNotShowOrderSnInMms.Bool,
      "type": billType.Int64,
    }
    billList = append(billList, bill)
  }
  json.NewEncoder(w).Encode(billList)
}
