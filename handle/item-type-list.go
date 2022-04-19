/*
 * Maintained by jemo from 2022.4.18 to now
 * Created by jemo on 2022.4.18 10:06:28
 * 商品类型列表
 */

package handle

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func ItemTypeList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query(`
    SELECT
      id,
      typeName,
      typeNum
    FROM
      itemType
  `)
  if err != nil {
    log.Println("item-type-list-query-error: ", err)
  }
  defer rows.Close()
  var itemTypeList []interface{}
  for rows.Next() {
    var (
      id int64
      typeName string
      typeNum string
    )
    if err := rows.Scan(&id, &typeName, &typeNum); err != nil {
      log.Println("item-type-list-scan-error: ", err)
    }
    itemType := map[string]interface{}{
      "id": id,
      "typeName": typeName,
      "typeNum": typeNum,
    }
    itemTypeList = append(itemTypeList, itemType)
  }
  json.NewEncoder(w).Encode(itemTypeList)
}
