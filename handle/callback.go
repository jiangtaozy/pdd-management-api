/*
 * Maintained by jemo from 2020.4.28 to now
 * Created by jemo on 2020.4.28 11:52:41
 * Callback Handle
 */

package handle

import (
  "log"
  "bytes"
  "encoding/json"
  "net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  if len(query["code"]) > 0 {
    code := query["code"][0]
    const oauthUrl = "https://open-api.pinduoduo.com/oauth/token"
    requestBody, err := json.Marshal(map[string]string{
      "client_id": "",
      "code": code,
      "grant_type": "authorization_code",
      "client_secret": "",
    })
    if err != nil {
      log.Fatalln(err)
    }
    resp, err := http.Post(oauthUrl, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
      log.Fatalln(err)
    }
    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    /*
    {
      error_response: {
        error_code: 10013,
        error_msg: code 已过期
      }
    }
    {
      "scope": [
        "pdd.logistics.companies.get",
        "pdd.logistics.online.send",
        "pdd.order.information.get",
        "pdd.order.number.list.get",
        "pdd.order.number.list.increment.get",
        "pdd.order.status.get",
        "pdd.refund.list.increment.get",
        "pdd.refund.status.check",
        "pdd.virtual.mobile.charge.notify"
      ],
      "access_token": "19eeaea897914e95907488d2f94b21ef26ef2687",
      "expires_in": 86400,
      "refresh_token": "75cd5b38e60d46c7a566521c792e0e7e9e341b4f",
      "owner_id": "213213"
      "owner_name": "pdd213213"
    }
    */
    errorResponse := result["error_response"].(map[string]interface{})
    if errorResponse != nil {
      errorCode := errorResponse["error_code"]
      errorMsg := errorResponse["error_msg"]
      log.Println("errorCode: ", errorCode)
      log.Println("errorMsg: ", errorMsg)
    } else {
      scope := result["scope"]
      accessToken := result["access_token"]
      expiresIn := result["expires_in"]
      refreshToken := result["refresh_token"]
      ownerId := result["owner_id"]
      ownerName := result["owner_name"]
      log.Println("scope: ", scope)
      log.Println("accessToken: ", accessToken)
      log.Println("expiresIn: ", expiresIn)
      log.Println("refreshToken: ", refreshToken)
      log.Println("ownerId: ", ownerId)
      log.Println("ownerName: ", ownerName)
    }
  }
  http.Redirect(w, r, "http://localhost:7000", 301)
}
