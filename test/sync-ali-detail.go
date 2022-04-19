/*
 * Maintained by jemo from 2022.4.15 to now
 * Created by jemo on 2022.4.15 15:53:40
 * SyncAliDetail Test
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
  //data, err := os.ReadFile("../html/ali-detail.html")
  data, err := os.ReadFile("../html/ali-detail-1.html")
  if err != nil {
    log.Println("sync-ali-detail-test-read-file-error: ", err)
  }
  handle.SyncAliDetail(string(data))
}
