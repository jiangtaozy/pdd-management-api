/*
 * Maintained by jemo from 2020.9.10 to now
 * Created by jemo on 2020.9.10 16:37:32
 * Dy Item List
 * 抖音商品列表
 */

package handle

import (
  "io"
  "log"
  "net/http"
)

func DyItemList(w http.ResponseWriter, r *http.Request) {
  // get access token
  shopId := "973906"
  accessToken := GetAccessToken(shopId)
  log.Println("accessToken: ", accessToken)

  io.WriteString(w, "ok")
}
