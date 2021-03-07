/*
 * Maintained by jemo from 2021.2.24 to now
 * Created by jemo on 2021.2.24 16:24:53
 * Sync Women On Shelf
 * 同步女装网上架数据
 */

package handle

import (
  "io"
  "log"
  "time"
  "math/rand"
  "net/http"
  "database/sql"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncWomenOnShelf(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      womenItem.searchId,
      item.detailUrl
    FROM
      womenItem
    LEFT JOIN
      item
    ON
      womenItem.searchId = item.searchId
    LEFT JOIN
      pddItem
    ON
      womenItem.searchId = pddItem.outGoodsSn
    WHERE
      pddItem.isOnsale IS TRUE
    GROUP BY
      womenItem.searchId
  `)
  if err != nil {
    log.Println("sync-women-on-shelf.go-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  for rows.Next() {
    var (
      searchId float64
      detailUrl sql.NullString
    )
    err = rows.Scan(
      &searchId,
      &detailUrl,
    )
    if err != nil {
      log.Println("sync-women-on-shelf.go-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    num := rand.Int31n(10)
    time.Sleep(time.Duration(num) * time.Second)
    CollyGetWomenDetailData(w, detailUrl.String, searchId)
  }
  io.WriteString(w, "ok")
}
