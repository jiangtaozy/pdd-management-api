/*
 * Maintained by jemo from 2021.1.20 to now
 * Created by jemo on 2021.1.20 17:56:54
 * Get Women Detail Data
 * 获取女装网详细数据
 */

package handle

import (
  "io"
  "log"
  "math"
  "regexp"
  "strconv"
  "encoding/json"
  "net/http"
  "github.com/gocolly/colly/v2"
  "github.com/PuerkitoBio/goquery"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func GetWomenDetailData(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  detailUrl := body["detailUrl"].(string)
  id := body["id"].(float64)
  if err != nil {
    log.Println("get-women-detail-data-decode-body-err: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  log.Println("detailUrl: ", detailUrl)
  log.Println("id: ", id)
  collector := colly.NewCollector()
  collector.OnHTML("body", func(e *colly.HTMLElement) {
    isLightningDelivery := e.DOM.Find(".sdfhLabelIco").Length()
    isCloudWarehouse := e.DOM.Find(".yuncang").Length()
    isWechatMerchant := e.DOM.Find(".weishang").Length()
    isBigImg := e.DOM.Find(".jingxiu").Length()
    isHotSale := e.DOM.Find(".remai").Length()
    title := e.DOM.Find(".detail-midtitle h1").Text()
    keywords := e.DOM.Find(".hasKeywords").Text()
    priceStr := e.DOM.Find("#productShopPrice").Text()
    price, err := strconv.ParseFloat(priceStr, 32)
    if err != nil {
      log.Println("get-women-detail-data-parse-price-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    price = math.Round(price * 100)
    presentPriceStr := e.DOM.Find(".presentPrice span").Text()
    presentPrice, err := strconv.ParseFloat(presentPriceStr, 32)
    if err != nil {
      log.Println("get-women-detail-data-parse-present-price-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    presentPrice = math.Round(presentPrice * 100)
    goodsNo := e.DOM.Find("#productCoodsNo").Text()
    createdTime := e.DOM.Find(`.midbrand-1[style="text-align: center;"] label`).Text()
    updatedTime := e.DOM.Find(".midbrand-2 label").Text()
    totalSaleStr := e.DOM.Find(".shopData-li2 span").Text()
    intRegx := regexp.MustCompile("[0-9]+")
    totalSale := intRegx.FindString(totalSaleStr)
    showCount := e.DOM.Find(".shopData-li4 span").First().Text()
    weightStr := e.DOM.Find(".shopData-li3 span").Text()
    weight := intRegx.FindString(weightStr)
    deliveryRateStr := e.DOM.Find("#Product_OOS_Button_DeliveryRate i").Text()
    floatRegx := regexp.MustCompile("[0-9]+([.][0-9]+)?")
    deliveryRateFloat := floatRegx.FindString(deliveryRateStr)
    var deliveryRate float64
    if deliveryRateFloat != "" {
      deliveryRate, err = strconv.ParseFloat(deliveryRateFloat, 32)
      if err != nil {
        log.Println("get-women-detail-data-parse-delivery-rate-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    }
    deliveryTimeStr := e.DOM.Find("#Product_OOS_Button_DTC7Da i").Text()
    deliveryTimeInt := intRegx.FindString(deliveryTimeStr)
    var deliveryTime int
    if deliveryTimeInt != "" {
      deliveryTime, err = strconv.Atoi(deliveryTimeInt)
      if err != nil {
        log.Println("get-women-detail-data-parse-delivery-time-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    }
    productId, _ := e.DOM.Find("#Product_ID").Attr("value")
    isReturn := e.DOM.Find(".adlRg .tuihuo").Length()
    isOriginalImage := e.DOM.Find(".adlRg .yuantu").Length()
    isPowerMerchant := e.DOM.Find(".adlRg .shili").Length()
    isFactory := e.DOM.Find(".adlRg .gongchang").Length()
    isLimitPrice := e.DOM.Find(".adlRg .jia").Length()
    // 保存 womenItem
    db := database.DB
    stmtInsertWomenItem, err := db.Prepare(`
      INSERT INTO womenItem (
        searchId,
        isLightningDelivery,
        isCloudWarehouse,
        isWechatMerchant,
        isBigImg,
        isHotSale,
        title,
        keywords,
        price,
        presentPrice,
        goodsNo,
        createdTime,
        updatedTime,
        totalSale,
        showCount,
        weight,
        deliveryRate,
        deliveryTime,
        productId,
        isReturn,
        isOriginalImage,
        isPowerMerchant,
        isFactory,
        isLimitPrice
      ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `)
    if err != nil {
      log.Println("get-women-detail-data-insert-women-item-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    defer stmtInsertWomenItem.Close()
    stmtUpdateWomenItem, err := db.Prepare(`
      UPDATE
        womenItem
      SET
        isLightningDelivery = ?,
        isCloudWarehouse = ?,
        isWechatMerchant = ?,
        isBigImg = ?,
        isHotSale = ?,
        title = ?,
        keywords = ?,
        price = ?,
        presentPrice = ?,
        goodsNo = ?,
        createdTime = ?,
        updatedTime = ?,
        totalSale = ?,
        showCount = ?,
        weight = ?,
        deliveryRate = ?,
        deliveryTime = ?,
        productId = ?,
        isReturn = ?,
        isOriginalImage = ?,
        isPowerMerchant = ?,
        isFactory = ?,
        isLimitPrice = ?
      WHERE
        searchId = ?
    `)
    if err != nil {
      log.Println("get-women-detail-data-update-women-item-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    defer stmtUpdateWomenItem.Close()
    var womenItemCount int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        womenItem
      WHERE
        searchId = ?
    `, id).Scan(&womenItemCount)
    if err != nil {
      log.Println("get-women-detail-data-women-item-count-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    if womenItemCount == 0 {
      _, err = stmtInsertWomenItem.Exec(
        id,
        isLightningDelivery,
        isCloudWarehouse,
        isWechatMerchant,
        isBigImg,
        isHotSale,
        title,
        keywords,
        price,
        presentPrice,
        goodsNo,
        createdTime,
        updatedTime,
        totalSale,
        showCount,
        weight,
        deliveryRate,
        deliveryTime,
        productId,
        isReturn,
        isOriginalImage,
        isPowerMerchant,
        isFactory,
        isLimitPrice,
      )
      if err != nil {
        log.Println("get-women-detail-data-insert-women-item-exec-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    } else {
      _, err = stmtUpdateWomenItem.Exec(
        isLightningDelivery,
        isCloudWarehouse,
        isWechatMerchant,
        isBigImg,
        isHotSale,
        title,
        keywords,
        price,
        presentPrice,
        goodsNo,
        createdTime,
        updatedTime,
        totalSale,
        showCount,
        weight,
        deliveryRate,
        deliveryTime,
        productId,
        isReturn,
        isOriginalImage,
        isPowerMerchant,
        isFactory,
        isLimitPrice,
        id,
      )
      if err != nil {
        log.Println("get-women-detail-data-update-women-item-exec-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    }
    // todo 规格
    spec, _ := e.DOM.Find("#SpecPrice").Attr("value")
    log.Println("spec: ", spec)
    // todo 主图
    mainImgs := e.DOM.Find("#imageMenu ul li a img")
    mainImgs.Each(func(i int, s *goquery.Selection) {
      img225, _ := s.Attr("src")
      log.Println("img225: ", img225)
      img500, _ := s.Attr("name")
      log.Println("img500: ", img500)
      img800, _ := s.Attr("id")
      log.Println("img800: ", img800)
    })
    // todo 属性
    attributes := e.DOM.Find("#tab0_detail #props ul li")
    attributes.Each(func(i int, s *goquery.Selection) {
      value, _ := s.Attr("title")
      log.Println("value: ", value)
      key := s.Find("span").Text()
      log.Println("key: ", key)
    })
    // todo 详情图
    detailImgs := e.DOM.Find("#detail_img").Find("img")
    detailImgs.Each(func(i int, s *goquery.Selection) {
      src, _ := s.Attr("data-original")
      log.Println("src: ", src)
    })

    io.WriteString(w, "ok")
  })
  collector.OnError(func(_ *colly.Response, err error) {
    log.Println("get-women-detail-data-collector-on-error:", err)
    http.Error(w, err.Error(), 500)
    return
  })
  collector.Visit(detailUrl)
}
