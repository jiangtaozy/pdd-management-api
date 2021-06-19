/*
 * Maintained by jemo from 2020.9.7 to now
 * Created by jemo on 2020.9.7 11:37:52
 */

package handle

import (
  "io"
  "log"
  "strings"
  "net/http"
  "golang.org/x/net/html"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func HangAfterSaleOrderUpload(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20) // 32M
  file, _, err := r.FormFile("file")
  if err != nil {
    log.Println("hang-after-sale-order-upload-form-file-error: ", err)
  }
  defer file.Close()
  doc, err := html.Parse(file)
  if err != nil {
    log.Println("hang-after-sale-order-upload-parse-error: ", err)
  }
  var list []interface{}
  var f func(*html.Node)
  f = func(node *html.Node) {
    if node.Type == html.ElementNode && node.Data == "a" {
      for _, attr := range node.Attr {
        if attr.Key == "class" && attr.Val == "nz-orderClo" {
          text := node.FirstChild
          order := map[string]interface{}{
            "orderId": strings.TrimSpace(text.Data),
          }
          list = append(list, order)
        }
      }
    }
    if node.Type == html.ElementNode && node.Data == "td" {
      for _, attr := range node.Attr {
        if attr.Key == "class" && attr.Val == "lsRowTd6" {
          text := node.FirstChild
          var data = strings.TrimSpace(text.Data)
          // 售后审核不通过
          if text.NextSibling != nil {
            data += text.NextSibling.FirstChild.Data
          }
          list[len(list) - 1].(map[string]interface{})["afterSaleStatusStr"] = data
        }
      }
    }
    for c := node.FirstChild; c != nil; c = c.NextSibling {
      f(c)
    }
  }
  f(doc)
  db := database.DB
  stmtUpdate, err := db.Prepare(`
    UPDATE order1688
    SET
      afterSaleStatusStr = ?
    WHERE
      orderId = ?
  `)
  if err != nil {
    log.Println("hang-after-sale-order-upload-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(list); i++ {
    order := list[i].(map[string]interface{})
    orderId := order["orderId"]
    afterSaleStatusStr := order["afterSaleStatusStr"]
    _, err = stmtUpdate.Exec(afterSaleStatusStr, orderId)
    if err != nil {
      log.Println("hang-after-sale-order-upload-update-exec-error: ", err)
    }
  }
  io.WriteString(w, "ok")
}
