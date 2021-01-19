/*
 * Maintained by jemo from 2021.1.19 to now
 * Created by jemo on 2021.1.19 14:30:42
 * Pdd Competitor List
 */

package handle

import (
  "encoding/json"
  "database/sql"
  "log"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func PddCompetitorList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      id,
      name,
      advantage,
      disadvantage
    FROM
      pddCompetitor
  `)
  if err != nil {
    log.Println("pdd-competitor-list-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  var list []interface{}
  for rows.Next() {
    var (
      id int64
      name string
      advantage sql.NullString
      disadvantage sql.NullString
    )
    if err := rows.Scan(
      &id,
      &name,
      &advantage,
      &disadvantage,
    ); err != nil {
      log.Println("pdd-competitor-list-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    pddCompetitor := map[string]interface{}{
      "id": id,
      "name": name,
      "advantage": advantage.String,
      "disadvantage": disadvantage.String,
    }
    list = append(list, pddCompetitor)
  }
  json.NewEncoder(w).Encode(list)
}
