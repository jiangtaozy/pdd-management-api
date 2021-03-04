/*
 * Maintained by jemo from 2021.2.27 to now
 * Created by jemo on 2021.2.27 17:51:13
 * Get Women Image
 * 获取女装网图片
 */

package handle

import (
  "io"
  "os"
  "log"
  "fmt"
  "image"
  _ "image/jpeg"
  _ "image/png"
  "net/http"
  "crypto/md5"
  "path/filepath"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func GetWomenImage(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("get-women-image.go-decode-body-err: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  id := body["id"].(float64)
  db := database.DB
  rows, err := db.Query(`
    SELECT
      id,
      img800
    FROM
      womenItemMainImage
    WHERE
      searchId = ?
  `, id)
  if err != nil {
    log.Println("get-women-image.go-query-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer rows.Close()
  update, err := db.Prepare(`
    UPDATE
      womenItemMainImage
    SET
      imgPath = ?,
      isLongImage = ?
    WHERE
      id = ?
  `)
  if err != nil {
    log.Println("get-women-image.go-update-prepare-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer update.Close()
  for rows.Next() {
    var (
      imgId int64
      img800 string
    )
    err = rows.Scan(
      &imgId,
      &img800,
    )
    if err != nil {
      log.Println("get-women-image.go-scan-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    hash, err := getMd5String(img800)
    if err != nil {
      log.Println("get-women-image.go-get-hash-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    path := filepath.Join(hash[0: 2], hash[2: 4], hash[4: 6])
    suffix := filepath.Ext(img800)
    fileName := hash + suffix
    allPath := filepath.Join("../image", path)
    allFilePath := filepath.Join(allPath, fileName)
    // 下载图片
    response, err := http.Get(img800)
    if err != nil {
      log.Println("get-women-image.go-http-get-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    defer response.Body.Close()
    err = os.MkdirAll(allPath, os.ModePerm)
    if err != nil {
      log.Println("get-women-image.go-mkdir-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    file, err := os.Create(allFilePath)
    if err != nil {
      log.Println("get-women-image.go-create-file-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    defer file.Close()
    _, err = io.Copy(file, response.Body)
    if err != nil {
      log.Println("get-women-image.go-copy-file-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    // 保存路径
    imgPath := filepath.Join(path, fileName)
    reader, err := os.Open(allFilePath)
    if err != nil {
      log.Println("get-women-image.go-open-image-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    im, _, err := image.DecodeConfig(reader)
    if err != nil {
      log.Println("get-women-image.go-image-decode-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    isLongImage := im.Width != im.Height
    _, err = update.Exec(
      imgPath,
      isLongImage,
      imgId,
    )
    if err != nil {
      log.Println("get-women-image.go-update-exec-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  }
  io.WriteString(w, "ok")
}

func getMd5String(str string) (string, error) {
  m := md5.New()
  _, err := io.WriteString(m, str)
  if err != nil {
    log.Println("get-md5-string-error: ", err)
    return "", err
  }
  arr := m.Sum(nil)
  return fmt.Sprintf("%x", arr), nil
}
