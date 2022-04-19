/*
 * Maintained by jemo from 2021.12.15 to now
 * Created by jemo on 2021.12.15 11:47:37
 * SyncAliOrderList Test
 */

package main

import (
  "os"
  "log"
  "github.com/jiangtaozy/pdd-management-api/handle"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func main() {
  database.InitDB()
  data, err := os.ReadFile("../html/ali-order-list.html")
  if err != nil {
    log.Println("sync-ali-order-list-test-read-file-error: ", err)
  }
  handle.SyncAliOrderList(string(data))
}
