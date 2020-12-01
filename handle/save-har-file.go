/*
 * Maintained by jemo from 2020.5.18 to now
 * Created by jemo on 2020.5.18 14:44:07
 * Upload Pinduoduo Order File
 */

package handle

import (
  "io"
  "log"
  "time"
  "encoding/json"
  "net/http"
)

func SaveHarFile(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20) // 32M
  file, _, err := r.FormFile("file")
  if err != nil {
    log.Println("upload-pdd-order-file-form-file-err: ", err)
  }
  defer file.Close()
  var har map[string]interface{}
  err = json.NewDecoder(file).Decode(&har)
  if err != nil {
    log.Println("save-har-file-decode-error: ", err)
  }
  logMap := har["log"].(map[string]interface{})
  entries := logMap["entries"].([]interface{})
  start := time.Now()
  for i := 0; i < len(entries); i++ {
    entry := entries[i].(map[string]interface{})
    request := entry["request"].(map[string]interface{})
    requestBody := make(map[string]interface{})
    if request["postData"] != nil {
      postData := request["postData"].(map[string]interface{})
      postText := postData["text"].(string)
      requestMimeType := postData["mimeType"].(string)
      if requestMimeType == "application/json;charset=UTF-8" {
        err := json.Unmarshal([]byte(postText), &requestBody)
        if err != nil {
          log.Println("save-har-file-unmarshal-post-text-error: ", err)
        }
      }
    }
    response := entry["response"].(map[string]interface{})
    responseContent := response["content"].(map[string]interface{})
    responseBody := make(map[string]interface{})
    if responseContent["text"] != nil {
      responseText := responseContent["text"].(string)
      responseMimeType := responseContent["mimeType"]
      if responseMimeType == "application/json" {
        err := json.Unmarshal([]byte(responseText), &responseBody)
        if err != nil {
          log.Println("save-har-file-unmarshal-response-text-error: ", err)
        }
      }
    }
    url := request["url"]
    if url == "https://yingxiao.pinduoduo.com/venus/api/subway/keyword/listKeywordPage" {
      SaveListKeywordPage(responseBody)
      log.Println("url: ", url)
      now := time.Now()
      diff := now.Sub(start)
      log.Println("diff: ", diff)
      start = now
    }
    if url == "https://yingxiao.pinduoduo.com/mms-gateway/venus/api/unit/listPage" {
      SaveUnitListPage(requestBody, responseBody)
      log.Println("url: ", url)
    }
  }
  log.Println("\n\n***************************************************ok")
  io.WriteString(w, "ok")
}
