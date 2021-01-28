/*
 * Maintained by jemo from 2021.1.28 to now
 * Created by jemo on 2021.1.28 16:56:37
 * Test
 * 测试接口
 */

package handle

import (
  "io"
  "log"
  "net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
  log.Println("ok")

  /*
  productId := 1
  id := 10000
  src := "https://img.hznzcn.com/images/crawler/20191122/16d770bcca4d4f578c778078f2091f0e.jpg"
  err := SaveWomenItemMainImage(id, productId, src)
  if err != nil {
    log.Println("test-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  */

  io.WriteString(w, "ok")
}
