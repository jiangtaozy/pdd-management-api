/*
 * Maintained by jemo from 2020.9.10 to now
 * Created by jemo on 2020.9.10 16:37:32
 * Dy Item List
 * 抖音商品列表
 */

package handle

import (
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/openapi-fxg"
)

func DyItemList(w http.ResponseWriter, r *http.Request) {
  // get access token
  shopId := "973906"
  accessToken := GetAccessToken(shopId)
  itemList := openapiFxg.ProductList(
    appId,
    appSecret,
    accessToken,
  )
  json.NewEncoder(w).Encode(itemList)
}

