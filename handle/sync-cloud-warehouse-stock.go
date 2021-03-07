/*
 * Maintained by jemo from 2021.2.17 to now
 * Created by jemo on 2021.2.17 13:43:54
 * Sync Cloud Warehouse Stock
 * 同步云仓库存
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

func SyncCloudWarehouseStock(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      womenItem.searchId,
      womenItem.productId
    FROM
      womenItem
    LEFT JOIN
      pddItem
    ON
      womenItem.searchId = pddItem.outGoodsSn
    WHERE
      womenItem.isCloudWarehouse IS TRUE
      AND
      pddItem.isOnsale IS TRUE
  `)
  if err != nil {
    log.Println("sync-cloud-warehouse-stock.go-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  for rows.Next() {
    var (
      searchId float64
      productId sql.NullString
    )
    err = rows.Scan(
      &searchId,
      &productId,
    )
    if err != nil {
      log.Println("sync-cloud-warehouse-stock.go-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    num := rand.Int31n(10)
    time.Sleep(time.Duration(num) * time.Second)
    err = FetchWomenCloudWarehouseStock(searchId, productId.String)
    if err != nil {
      log.Println("sync-cloud-warehouse-stock.go-fetch-err: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  }
  io.WriteString(w, "ok")
}
