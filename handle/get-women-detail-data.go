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
  "strings"
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
  CollyGetWomenDetailData(w, detailUrl, id)
  io.WriteString(w, "ok")
}

func CollyGetWomenDetailData(w http.ResponseWriter, detailUrl string, id float64) {
  collector := colly.NewCollector()
  collector.OnResponse(func(r *colly.Response) {
    //log.Println("string(r.Body): ", string(r.Body))
  })
  collector.OnHTML("body", func(e *colly.HTMLElement) {
    if strings.Index(e.Text, "跳转中") >= 0 {
      hrefRegx := regexp.MustCompile("window.location.href =\"(.*)\"")
      href := hrefRegx.FindStringSubmatch(e.Text)
      collector.Visit("https://www.hznzcn.com" + href[1])
      return
    }
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
    var presentPrice float64
    if presentPriceStr != "" {
      presentPrice, err = strconv.ParseFloat(presentPriceStr, 32)
      if err != nil {
        log.Println("get-women-detail-data-parse-present-price-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
      presentPrice = math.Round(presentPrice * 100)
    }
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
    video, _ := e.DOM.Find("#J_playVideo").Attr("videourl")
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
        isLimitPrice,
        video
      ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
        isLimitPrice = ?,
        video = ?
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
        video,
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
        video,
        id,
      )
      if err != nil {
        log.Println("get-women-detail-data-update-women-item-exec-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    }
    // 保存womenItemSku
    spec, _ := e.DOM.Find("#SpecPrice").Attr("value")
    specList := strings.Split(spec, "|")
    stmtInsertSku, err := db.Prepare(`
      INSERT INTO womenItemSku (
        searchId,
        productId,
        skuDesc,
        skuKey,
        price,
        isOnShelf,
        stock,
        skuColor,
        skuSize
      ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `)
    if err != nil {
      log.Println("get-women-detail-data-insert-women-item-sku-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    defer stmtInsertSku.Close()
    stmtUpdateSku, err := db.Prepare(`
      UPDATE
        womenItemSku
      SET
        productId = ?,
        skuKey = ?,
        price = ?,
        isOnShelf = ?,
        stock = ?,
        skuColor = ?,
        skuSize = ?
      WHERE
        searchId = ?
      AND
        skuDesc = ?
    `)
    if err != nil {
      log.Println("get-women-detail-data-update-women-item-sku-prepare-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    defer stmtUpdateSku.Close()
    for i := 0; i < len(specList); i++ {
      sku := specList[i]
      keyList := strings.Split(sku, ";")
      skuDesc := keyList[0]
      skuDescList := strings.Split(skuDesc, ",")
      skuColor := skuDescList[0]
      skuSize := skuDescList[1]
      if skuSize == "XXL" {
        skuSize = "2XL"
      } else if skuSize == "XXXL" {
        skuSize = "3XL"
      } else if skuSize == "XXXXL" {
        skuSize = "4XL"
      } else if skuSize == "L(120-135斤)" {
        skuSize = "L"
      } else if skuSize == "M（110-125斤）" {
        skuSize = "M"
      } else if skuSize == "S（100-115斤）" {
        skuSize = "S"
      } else if skuSize == "XL（130-160斤）" {
        skuSize = "XL"
      } else if skuSize == "XS(100斤以内)" {
        skuSize = "XS"
      }
      skuKey := keyList[1]
      priceStr := keyList[2]
      isOnShelf := keyList[3]
      stock := keyList[4]
      price, err := strconv.ParseFloat(priceStr, 32);
      if err != nil {
        log.Println("get-women-detail-data-parse-sku-price-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
      price = math.Round(price *  100)
      var skuCount int
      err = db.QueryRow(`
        SELECT
          COUNT(*)
        FROM
          womenItemSku
        WHERE
          searchId = ?
        AND
          skuDesc = ?
      `, id, skuDesc).Scan(&skuCount)
      if err != nil {
        log.Println("get-women-detail-data-women-item-sku-count-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
      if skuCount == 0 {
        _, err = stmtInsertSku.Exec(
          id,
          productId,
          skuDesc,
          skuKey,
          price,
          isOnShelf,
          stock,
          skuColor,
          skuSize,
        )
        if err != nil {
          log.Println("get-women-detail-data-insert-women-item-sku-exec-error: ", err)
          http.Error(w, err.Error(), 500)
          return
        }
      } else {
        _, err = stmtUpdateSku.Exec(
          productId,
          skuKey,
          price,
          isOnShelf,
          stock,
          skuColor,
          skuSize,
          id,
          skuDesc,
        )
        if err != nil {
          log.Println("get-women-detail-data-update-women-item-sku-exec-error: ", err)
          http.Error(w, err.Error(), 500)
          return
        }
      }
    }
    // 保存主图
    mainImgs := e.DOM.Find("#imageMenu ul li a img")
    mainImgs.Each(func(i int, s *goquery.Selection) {
      img225, _ := s.Attr("src")
      img500, _ := s.Attr("name")
      img800, _ := s.Attr("id")
      err = SaveWomenItemMainImage(id, productId, img225, img500, img800)
      if err != nil {
        log.Println("get-women-detail-data-save-women-item-main-image-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    })
    // 保存颜色
    colors := e.DOM.Find("#em0 span")
    colors.Each(func(i int, s *goquery.Selection) {
      color, _ := s.Attr("text")
      thumbnail, _ := s.Attr("thumbnailaddress")
      hrthumbnail, _ := s.Attr("hrthumbnail")
      original, _ := s.Attr("originaladdress")
      err = SaveWomenItemColor(id, productId, color, thumbnail, hrthumbnail, original)
      if err != nil {
        log.Println("get-women-detail-data-save-women-item-color-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    })
    // 保存尺码
    sizes := e.DOM.Find("#em1 table tr")
    sizes.Each(func(i int, s *goquery.Selection) {
      size, _ := s.Attr("sizevalue")
      err = SaveWomenItemSize(id, productId, size)
      if err != nil {
        log.Println("get-women-detail-data-save-women-item-size-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    })
    // 保存属性
    attributes := e.DOM.Find("#tab0_detail #props ul li")
    attributes.Each(func(i int, s *goquery.Selection) {
      value, _ := s.Attr("title")
      value = strings.Trim(value, " ")
      key := s.Find("span").Text()
      key = strings.Trim(key, "：")
      err = SaveWomenItemAttribute(id, productId, key, value)
      if err != nil {
        log.Println("get-women-detail-data-save-women-item-attribute-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    })
    // 保存详情图
    detailImgs := e.DOM.Find("#detail_img").Find("img")
    detailImgs.Each(func(i int, s *goquery.Selection) {
      src, _ := s.Attr("data-original")
      err = SaveWomenItemDetailImage(id, productId, src)
      if err != nil {
        log.Println("get-women-detail-data-save-women-item-detail-image-error: ", err)
        http.Error(w, err.Error(), 500)
        return
      }
    })
    // 保存item中womenProductId
    err = SaveWomenProductId(id, productId)
    if err != nil {
      log.Println("get-women-detail-data-save-women-product-id-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  })
  collector.OnError(func(_ *colly.Response, err error) {
    log.Println("get-women-detail-data-collector-on-error:", err)
    http.Error(w, err.Error(), 500)
    return
  })
  collector.Visit(detailUrl)
  collector.Wait()
}

func SaveWomenItemMainImage(searchId float64, productId string, img225 string, img500 string, img800 string) error {
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO womenItemMainImage (
      searchId,
      productId,
      img225,
      img500,
      img800
    ) VALUES (?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("get-women-detail-data-save-main-image-insert-prepare-error: ", err)
    return err
  }
  defer stmtInsert.Close()
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      womenItemMainImage
    WHERE
      searchId = ?
    AND
      img800 = ?
  `, searchId, img800).Scan(&count)
  if err != nil {
    log.Println("get-women-detail-data-save-main-image-count-error: ", err)
    return err
  }
  if count == 0 {
    _, err = stmtInsert.Exec(
      searchId,
      productId,
      img225,
      img500,
      img800,
    )
    if err != nil {
      log.Println("get-women-detail-data-save-main-image-insert-exec-error: ", err)
      return err
    }
  }
  return nil
}

func SaveWomenItemColor(searchId float64, productId string, color string, thumbnail string, hrthumbnail string, original string) error {
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO womenItemColor (
      searchId,
      productId,
      color,
      thumbnail,
      hrthumbnail,
      original
    ) VALUES (?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("get-women-detail-data-save-color-insert-prepare-error: ", err)
    return err
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      womenItemColor
    SET
      productId = ?,
      thumbnail = ?,
      hrthumbnail = ?,
      original = ?
    WHERE
      searchId = ?
    AND
      color = ?
  `)
  if err != nil {
    log.Println("get-women-detail-data-save-color-update-prepare-error: ", err)
    return err
  }
  defer stmtUpdate.Close()
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      womenItemColor
    WHERE
      searchId = ?
    AND
      color = ?
  `, searchId, color).Scan(&count)
  if err != nil {
    log.Println("get-women-detail-data-save-color-count-error: ", err)
    return err
  }
  if count == 0 {
    _, err = stmtInsert.Exec(
      searchId,
      productId,
      color,
      thumbnail,
      hrthumbnail,
      original,
    )
    if err != nil {
      log.Println("get-women-detail-data-save-color-insert-exec-error: ", err)
      return err
    }
  } else {
    _, err = stmtUpdate.Exec(
      productId,
      thumbnail,
      hrthumbnail,
      original,
      searchId,
      color,
    )
    if err != nil {
      log.Println("get-women-detail-data-save-color-update-exec-error: ", err)
      return err
    }
  }
  return nil
}

func SaveWomenItemSize(searchId float64, productId string, size string) error {
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO womenItemSize (
      searchId,
      productId,
      size
    ) VALUES (?, ?, ?)
  `)
  if err != nil {
    log.Println("get-women-detail-data-save-size-insert-prepare-error: ", err)
    return err
  }
  defer stmtInsert.Close()
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      womenItemSize
    WHERE
      searchId = ?
    AND
      size = ?
  `, searchId, size).Scan(&count)
  if err != nil {
    log.Println("get-women-detail-data-save-size-count-error: ", err)
    return err
  }
  if count == 0 {
    _, err = stmtInsert.Exec(
      searchId,
      productId,
      size,
    )
    if err != nil {
      log.Println("get-women-detail-data-save-size-insert-exec-error: ", err)
      return err
    }
  }
  return nil
}

func SaveWomenItemAttribute(searchId float64, productId string, key string, value string) error {
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO womenItemAttribute (
      searchId,
      productId,
      attributeKey,
      attributeValue
    ) VALUES (?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("get-women-detail-data-save-attribute-insert-prepare-error: ", err)
    return err
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      womenItemAttribute
    SET
      productId = ?,
      attributeValue = ?
    WHERE
      searchId = ?
    AND
      attributeKey = ?
  `)
  if err != nil {
    log.Println("get-women-detail-data-save-attribute-update-prepare-error: ", err)
    return err
  }
  defer stmtUpdate.Close()
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      womenItemAttribute
    WHERE
      searchId = ?
    AND
      attributeKey = ?
  `, searchId, key).Scan(&count)
  if err != nil {
    log.Println("get-women-detail-data-save-attribute-count-error: ", err)
    return err
  }
  if count == 0 {
    _, err = stmtInsert.Exec(
      searchId,
      productId,
      key,
      value,
    )
    if err != nil {
      log.Println("get-women-detail-data-save-attribute-insert-exec-error: ", err)
      return err
    }
  } else {
    _, err = stmtUpdate.Exec(
      productId,
      value,
      searchId,
      key,
    )
    if err != nil {
      log.Println("get-women-detail-data-save-attribute-update-exec-error: ", err)
      return err
    }
  }
  return nil
}

func SaveWomenItemDetailImage(searchId float64, productId string, src string) error {
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO womenItemDetailImage (
      searchId,
      productId,
      src
    ) VALUES (?, ?, ?)
  `)
  if err != nil {
    log.Println("get-women-detail-data-save-detail-image-insert-prepare-error: ", err)
    return err
  }
  defer stmtInsert.Close()
  var count int
  err = db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      womenItemDetailImage
    WHERE
      searchId = ?
    AND
      src = ?
  `, searchId, src).Scan(&count)
  if err != nil {
    log.Println("get-women-detail-data-save-detail-image-count-error: ", err)
    return err
  }
  if count == 0 {
    _, err = stmtInsert.Exec(
      searchId,
      productId,
      src,
    )
    if err != nil {
      log.Println("get-women-detail-data-save-detail-image-insert-exec-error: ", err)
      return err
    }
  }
  return nil
}

func SaveWomenProductId(searchId float64, productId string) error {
  db := database.DB
  stmtUpdate, err := db.Prepare(`
    UPDATE
      item
    SET
      womenProductId = ?
    WHERE
      searchId = ?
  `)
  if err != nil {
    log.Println("get-women-detail-data-save-women-product-id-update-prepare-error: ", err)
    return err
  }
  defer stmtUpdate.Close()
  _, err = stmtUpdate.Exec(
    productId,
    searchId,
  )
  if err != nil {
    log.Println("get-women-detail-data-save-women-product-id-update-exec-error: ", err)
    return err
  }
  return nil
}
