/*
 * Maintained by jemo from 2020.9.22 to now
 * Created by jemo on 2020.9.22 17:30:32
 * Douyin Access Token
 */

package handle

import (
  "log"
  "time"
  "github.com/jiangtaozy/openapi-fxg"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func GetAccessToken(shopId string) string {
  db := database.DB
  var count int
  err := db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      dyAccessToken
    WHERE
      shopId = ?
  `, shopId).Scan(&count)
  if err != nil {
    log.Println("dy-item-list-count-access-token-error: ", err)
  }
  var token string
  if count > 0 {
    token = getAccessTokenFromDatabase(shopId)
  } else {
    token = getAccessTokenFromServer()
  }
  return token
}

func getAccessTokenFromDatabase(shopId string) string {
  db := database.DB
  var token string
  var expiresIn int
  var refreshToken string
  var createdAt time.Time
  err := db.QueryRow(`
    SELECT
      accessToken,
      expiresIn,
      refreshToken,
      createdAt
    FROM
      dyAccessToken
    WHERE
      shopId = ?
  `, shopId).Scan(
    &token,
    &expiresIn,
    &refreshToken,
    &createdAt,
  )
  if err != nil {
    log.Println("dy-item-list-query-access-token-error: ", err)
  }
  expiresTime := createdAt.Add(time.Second * time.Duration(expiresIn))
  now := time.Now()
  isNowBeforeExpiresTime := now.Before(expiresTime)
  if !isNowBeforeExpiresTime {
    // accessToken 过期
    refreshExpiresTime := createdAt.Add(time.Hour * 24 * 14)
    isNowBeforeRefreshExpiresTime := now.Before(refreshExpiresTime)
    if isNowBeforeRefreshExpiresTime {
      // refreshAccessToken 未过期
      token = refreshAccessTokenFromServer(refreshToken)
    } else {
      // refreshAccessToken 过期
      token = getAccessTokenFromServer()
    }
  }
  return token
}

func getAccessTokenFromServer() string {
  accessToken := openapiFxg.AccessToken(
    appId,
    appSecret,
  )
  token := saveAccessToken(accessToken)
  return token
}

func refreshAccessTokenFromServer(refreshToken string) string {
  accessToken := openapiFxg.RefreshAccessToken(
    appId,
    appSecret,
    refreshToken,
  )
  token := saveAccessToken(accessToken)
  return token
}

func saveAccessToken(accessToken map[string]interface{}) string {

  //  map[
  //    data:map[
  //      access_token:58169963-cf23-474d-af05-2c2362abde8a
  //      expires_in:604800
  //      refresh_token:1ba24fbe-4692-44a6-a397-31ff5a6f9536
  //      scope:SCOPE
  //      shop_id:973906
  //      shop_name:牧季衣坊
  //    ]
  //    err_no:0
  //    message:success
  //  ]

  data := accessToken["data"].(map[string]interface{})
  errNo := int(accessToken["err_no"].(float64))
  if errNo != 0 {
    log.Println("dy-item-list-access-token-err-no-error: ", accessToken)
  }
  token := data["access_token"].(string)
  expiresIn := int(data["expires_in"].(float64))
  refreshToken := data["refresh_token"].(string)
  scope := data["scope"]
  shopId := data["shop_id"].(float64)
  shopName := data["shop_name"]
  createdAt := time.Now()
  var count int
  db := database.DB
  err := db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      dyAccessToken
    WHERE
      shopId = ?
  `, shopId).Scan(&count)
  if err != nil {
    log.Println("dy-item-list-count-access-token-error: ", err)
  }
  if count == 0 {
    stmtInsert, err := db.Prepare(`
      INSERT INTO dyAccessToken (
        accessToken,
        expiresIn,
        refreshToken,
        scope,
        shopId,
        shopName,
        createdAt
      ) VALUES (?, ?, ?, ?, ?, ?, ?)
    `)
    if err != nil {
      log.Println("dy-item-list-prepare-insert-error: ", err)
    }
    defer stmtInsert.Close()
    _, err = stmtInsert.Exec(
      token,
      expiresIn,
      refreshToken,
      scope,
      shopId,
      shopName,
      createdAt,
    )
    if err != nil {
      log.Println("dy-item-list-insert-exec-error: ", err)
    }
  } else {
    stmtUpdate, err := db.Prepare(`
      UPDATE
        dyAccessToken
      SET
        accessToken = ?,
        expiresIn = ?,
        refreshToken = ?,
        scope = ?,
        shopName = ?,
        createdAt = ?
      WHERE
        shopId = ?
    `)
    if err != nil {
      log.Println("dy-item-list-prepare-update-error: ", err)
    }
    defer stmtUpdate.Close()
    _, err = stmtUpdate.Exec(
      token,
      expiresIn,
      refreshToken,
      scope,
      shopName,
      createdAt,
      shopId,
    )
    if err != nil {
      log.Println("dy-item-list-update-exec-error: ", err)
    }
  }
  return token
}
