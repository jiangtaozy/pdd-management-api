/*
 * Maintained by jemo from 2021.12.15 to now
 * Created by jemo on 2021.12.15 10:42:06
 * 同步阿里巴巴订单列表
 */

package handle

import (
  "log"
  "time"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncAliOrderList(html string) {
  doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
  if err != nil {
    log.Println("sync-ali-order-list-goquery-err: ", err)
  }
  doc.Find(".order-item").Each(func(i int, s *goquery.Selection) {
    ImportOrder(s)
  })
}

func ImportOrder(s *goquery.Selection) {
  db := database.DB
  orderId, _ := s.Find(".comment").Attr("data-orderid")
  sellerCompany, _ := s.Find(".bannerCorp").Attr("data-copytitle")
  totalPrice := strings.TrimSpace(s.Find(".s6 .total").Text())
  shippingFare := strings.ReplaceAll(strings.TrimSpace(s.Find(".s6 .fare").Text()), "含运费", "")
  actualPayment := totalPrice
  orderStatusStr := strings.TrimSpace(s.Find(".s7 div").First().Text())
  orderStatusMap := map[string]interface{}{
    "等待买家付款": 0,
    "等待卖家发货": 1,
    "等待买家确认收货": 2,
    "已收货": 3,
    "交易成功": 4,
    "交易关闭": 6,
    "已发货": 2,
    "退款中": 5,
  }
  orderStatus := orderStatusMap[orderStatusStr]
  dateStr := s.Find(".date").Text()
  const layout = "2006-01-02 15:04:05"
  location, _ := time.LoadLocation("Asia/Shanghai")
  orderCreatedTime, err := time.ParseInLocation(layout, dateStr, location)
  if err != nil {
    log.Println("sync-ali-order-list-parse-time-error: ", err)
  }
  productTitle := s.Find(".detail table tbody tr").First().Find(".productName").Text()
  price := s.Find(".detail table tbody tr").First().Find(".s3 div[title] span").Text()
  amount := strings.TrimSpace(s.Find(".detail table tbody tr").First().Find(".s4 div[title]").Text())
  orderType := 0 // 订单类型，0: 1688, 1: 女装网
  productSku := ""
  s.Find(".detail table tbody tr").Find(".trade-spec").Each(func(i int, spec *goquery.Selection) {
    if i != 0 {
      productSku += ","
    }
    productSku += strings.TrimSpace(spec.Text())
  })
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      order1688
    WHERE
      orderId = ?
  `, orderId).Scan(&count)
  if err != nil {
    log.Println("sync-ali-order-list-count-error: ", err)
  }
  if count == 0 {
    stmtInsert, err := db.Prepare(`
      insert into order1688 (
        orderId,
        sellerCompany,
        totalPrice,
        shippingFare,
        actualPayment,
        orderStatus,
        orderStatusStr,
        orderCreatedTime,
        productTitle,
        price,
        amount,
        orderType,
        productSku
      ) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `)
    if err != nil {
      log.Println("sync-ali-order-list-insert-prepare-error: ", err)
    }
    defer stmtInsert.Close()
    _, err = stmtInsert.Exec(
      orderId,
      sellerCompany,
      totalPrice,
      shippingFare,
      actualPayment,
      orderStatus,
      orderStatusStr,
      orderCreatedTime,
      productTitle,
      price,
      amount,
      orderType,
      productSku,
    )
    if err != nil {
      log.Println("sync-ali-order-list-insert-exec-error: ", err)
    }
  } else {
    stmtUpdate, err := db.Prepare(`
      update
        order1688
      set
        orderStatus = ?,
        orderStatusStr = ?
      where
        orderId = ?
    `)
    if err != nil {
      log.Println("sync-ali-order-list-update-prepare-error: ", err)
    }
    defer stmtUpdate.Close()
    _, err = stmtUpdate.Exec(
      orderStatus,
      orderStatusStr,
      orderId,
    )
    if err != nil {
      log.Println("sync-ali-order-list-update-exec-error: ", err)
    }
  }
}
