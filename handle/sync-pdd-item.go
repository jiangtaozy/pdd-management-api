/*
 * Maintained by jemo from 2021.1.9 to now
 * Created by jemo on 2021.1.9 17:30:23
 * Sync Pdd Item
 */

package handle

import (
  "log"
  "time"
  "strings"
  "encoding/json"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func SyncPddItem(requestBody map[string]interface{}, responseBody map[string]interface{}) {
  var result = responseBody["result"].(map[string]interface{})
  var goodsList = result["goods_list"].([]interface{})
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddItem (
      quantity,
      score,
      resource,
      priority,
      skuGroupPriceMin,
      skuGroupPriceMax,
      skuPriceMin,
      skuPriceMax,
      originSkuGroupPriceMin,
      originSkuGroupPriceMax,
      pddId,
      goodsName,
      goodsSn,
      goodsType,
      catId,
      catName,
      eventType,
      displayPriority,
      reserveQuantity,
      hdThumbUrl,
      imageUrl,
      thumbUrl,
      isOnsale,
      soldQuantity,
      offSaleCheck,
      onsaleCheck,
      editCheck,
      grayGoodsHighQualityRefund,
      shareDesc,
      goodsDesc,
      activityTitle,
      isCardSecret,
      marketPrice,
      outGoodsSn,
      soldQuantityForThirtyDays,
      favCnt,
      multiTreasure,
      multiTreasureStatus,
      ifNewGoods,
      tagCategoryList,
      titleScrFlag,
      titleDesc,
      propsScrFlag,
      propsDesc,
      galleryScrFlag,
      galleryDesc,
      galleryVideoScrFlag,
      longGraphScrFlag,
      whiteGraphScrFlag,
      materialScrFlag,
      goodsInfoScr,
      goodsInfoScoreUpdateInfoHint,
      createdAt,
      preSaleTime,
      shipmentLimitSecond,
      isGroupPreSale,
      isPreSale,
      guideTarget,
      overSell,
      marketLabels,
      labels,
      skuType,
      isMoreSku,
      skuList,
      rejectStatus,
      rejectReason
    ) VALUES(
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    )
  `)
  if err != nil {
    log.Println("upload-pdd-item-data-insert-prepare-error: ", err)
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      pddItem
    SET
      quantity = ?,
      score = ?,
      resource = ?,
      priority = ?,
      skuGroupPriceMin = ?,
      skuGroupPriceMax = ?,
      skuPriceMin = ?,
      skuPriceMax = ?,
      originSkuGroupPriceMin = ?,
      originSkuGroupPriceMax = ?,
      goodsName = ?,
      goodsSn = ?,
      goodsType = ?,
      catId = ?,
      catName = ?,
      eventType = ?,
      displayPriority = ?,
      reserveQuantity = ?,
      hdThumbUrl = ?,
      imageUrl = ?,
      thumbUrl = ?,
      isOnsale = ?,
      soldQuantity = ?,
      offSaleCheck = ?,
      onsaleCheck = ?,
      editCheck = ?,
      grayGoodsHighQualityRefund = ?,
      shareDesc = ?,
      goodsDesc = ?,
      activityTitle = ?,
      isCardSecret = ?,
      marketPrice = ?,
      outGoodsSn = ?,
      soldQuantityForThirtyDays = ?,
      favCnt = ?,
      multiTreasure = ?,
      multiTreasureStatus = ?,
      ifNewGoods = ?,
      tagCategoryList = ?,
      titleScrFlag = ?,
      titleDesc = ?,
      propsScrFlag = ?,
      propsDesc = ?,
      galleryScrFlag = ?,
      galleryDesc = ?,
      galleryVideoScrFlag = ?,
      longGraphScrFlag = ?,
      whiteGraphScrFlag = ?,
      materialScrFlag = ?,
      goodsInfoScr = ?,
      goodsInfoScoreUpdateInfoHint = ?,
      createdAt = ?,
      preSaleTime = ?,
      shipmentLimitSecond = ?,
      isGroupPreSale = ?,
      isPreSale = ?,
      guideTarget = ?,
      overSell = ?,
      marketLabels = ?,
      labels = ?,
      skuType = ?,
      isMoreSku = ?,
      skuList = ?,
      rejectStatus = ?,
      rejectReason = ?
    WHERE
      pddId = ?
  `)
  if err != nil {
    log.Println("upload-pdd-item-data-update-prepare-error: ", err)
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(goodsList); i++ {
    var goods = goodsList[i].(map[string]interface{})
    skuGroupPrice := goods["sku_group_price"].([]interface{})
    skuPrice := goods["sku_price"].([]interface{})
    var goodsInfoScoreDto map[string]interface{}
    originSkuGroupPrice := goods["origin_sku_group_price"].([]interface{})
    if goods["goods_info_score_dto"] != nil {
      goodsInfoScoreDto = goods["goods_info_score_dto"].(map[string]interface{})
    }
    tagCategoryList, err := json.Marshal(goods["tag_category_list"])
    if err != nil {
      log.Println("upload-pdd-item-data-marshal-error: ", err)
    }
    labels, err := json.Marshal(goods["labels"])
    if err != nil {
      log.Println("upload-pdd-item-data-marshal-labels-error: ", err)
    }
    createdAt := time.Unix(int64(goods["created_at"].(float64)), 0)
    id := goods["id"]
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM pddItem WHERE pddId = ?", id).Scan(&count)
    if err != nil {
      log.Println("upload-pdd-item-data-count-error: ", err)
    }
    var marketLabels string
    if goods["market_labels"] != nil {
      labels := goods["market_labels"].([]interface{})
      for i := 0; i < len(labels); i++ {
        label := labels[i].(map[string]interface{})
        marketLabels = marketLabels + label["tag_text"].(string) + ","
      }
    }
    // Data too long for column 'skuList'，所以存空
    skuList := ""
    if count == 0 {
      _, err = stmtInsert.Exec(
        goods["quantity"],
        goods["score"],
        goods["resource"],
        goods["priority"],
        skuGroupPrice[0],
        skuGroupPrice[1],
        skuPrice[0],
        skuPrice[1],
        originSkuGroupPrice[0],
        originSkuGroupPrice[1],
        goods["id"],
        goods["goods_name"],
        goods["goods_sn"],
        goods["goods_type"],
        goods["cat_id"],
        goods["cat_name"],
        goods["event_type"],
        goods["display_priority"],
        goods["reserve_quantity"],
        goods["hd_thumb_url"],
        goods["image_url"],
        goods["thumb_url"],
        goods["is_onsale"],
        goods["sold_quantity"],
        goods["off_sale_check"],
        goods["onsale_check"],
        goods["edit_check"],
        goods["gray_goods_high_quality_refund"],
        goods["share_desc"],
        goods["goods_desc"],
        goods["activity_title"],
        goods["is_card_secret"],
        goods["market_price"],
        goods["out_goods_sn"],
        goods["sold_quantity_for_thirty_days"],
        goods["fav_cnt"],
        goods["multi_treasure"],
        goods["multi_treasure_status"],
        goods["if_new_goods"],
        tagCategoryList,
        goodsInfoScoreDto["title_scr_flag"],
        goodsInfoScoreDto["title_desc"],
        goodsInfoScoreDto["props_scr_flag"],
        goodsInfoScoreDto["props_desc"],
        goodsInfoScoreDto["gallery_scr_flag"],
        goodsInfoScoreDto["gallery_desc"],
        goodsInfoScoreDto["gallery_video_scr_flag"],
        goodsInfoScoreDto["long_graph_scr_flag"],
        goodsInfoScoreDto["white_graph_scr_flag"],
        goodsInfoScoreDto["material_scr_flag"],
        goodsInfoScoreDto["goods_info_scr"],
        goodsInfoScoreDto["goods_info_score_update_info_hint"],
        createdAt,
        goods["pre_sale_time"],
        goods["shipment_limit_second"],
        goods["is_group_pre_sale"],
        goods["is_pre_sale"],
        goods["guide_target"],
        goods["over_sell"],
        marketLabels,
        labels,
        goods["sku_type"],
        goods["is_more_sku"],
        skuList,
        goods["reject_status"],
        goods["reject_reason"],
      )
      if err != nil {
        log.Println("upload-pdd-item-data-insert-exec-error: ", err)
        log.Println("err.Error(): ", err.Error())
      }
    } else {
      _, err = stmtUpdate.Exec(
        goods["quantity"],
        goods["score"],
        goods["resource"],
        goods["priority"],
        skuGroupPrice[0],
        skuGroupPrice[1],
        skuPrice[0],
        skuPrice[1],
        originSkuGroupPrice[0],
        originSkuGroupPrice[1],
        goods["goods_name"],
        goods["goods_sn"],
        goods["goods_type"],
        goods["cat_id"],
        goods["cat_name"],
        goods["event_type"],
        goods["display_priority"],
        goods["reserve_quantity"],
        goods["hd_thumb_url"],
        goods["image_url"],
        goods["thumb_url"],
        goods["is_onsale"],
        goods["sold_quantity"],
        goods["off_sale_check"],
        goods["onsale_check"],
        goods["edit_check"],
        goods["gray_goods_high_quality_refund"],
        goods["share_desc"],
        goods["goods_desc"],
        goods["activity_title"],
        goods["is_card_secret"],
        goods["market_price"],
        goods["out_goods_sn"],
        goods["sold_quantity_for_thirty_days"],
        goods["fav_cnt"],
        goods["multi_treasure"],
        goods["multi_treasure_status"],
        goods["if_new_goods"],
        tagCategoryList,
        goodsInfoScoreDto["title_scr_flag"],
        goodsInfoScoreDto["title_desc"],
        goodsInfoScoreDto["props_scr_flag"],
        goodsInfoScoreDto["props_desc"],
        goodsInfoScoreDto["gallery_scr_flag"],
        goodsInfoScoreDto["gallery_desc"],
        goodsInfoScoreDto["gallery_video_scr_flag"],
        goodsInfoScoreDto["long_graph_scr_flag"],
        goodsInfoScoreDto["white_graph_scr_flag"],
        goodsInfoScoreDto["material_scr_flag"],
        goodsInfoScoreDto["goods_info_scr"],
        goodsInfoScoreDto["goods_info_score_update_info_hint"],
        createdAt,
        goods["pre_sale_time"],
        goods["shipment_limit_second"],
        goods["is_group_pre_sale"],
        goods["is_pre_sale"],
        goods["guide_target"],
        goods["over_sell"],
        marketLabels,
        labels,
        goods["sku_type"],
        goods["is_more_sku"],
        skuList,
        goods["reject_status"],
        goods["reject_reason"],
        goods["id"],
      )
      if err != nil {
        log.Println("upload-pdd-item-data-update-exec-error: ", err)
      }
    }
    // sync pdd item price history
    PddItemPriceHistorySync(goods["id"].(float64), skuGroupPrice[0].(float64), skuGroupPrice[1].(float64));
    // 保存skuList
    SaveSkuList(goods["id"].(float64), goods["out_goods_sn"].(string), goods["sku_list"].([]interface{}))
  }
}

// 保存skuList
func SaveSkuList(pddId float64, outGoodsSn string, skuList []interface{}) {
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO pddItemSku (
      pddId,
      outGoodsSn,
      activityGroupPrice,
      groupPrice,
      isOnsale,
      normalPrice,
      outSkuSn,
      skuId,
      skuQuantity,
      skuSoldQuantity,
      spec,
      specColor,
      specSize
    )
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    log.Println("sync-pdd-item-save-sku-list-insert-prepare-error: ", err)
    return
  }
  defer stmtInsert.Close()
  stmtUpdate, err := db.Prepare(`
    UPDATE
      pddItemSku
    SET
      outGoodsSn = ?,
      activityGroupPrice = ?,
      groupPrice = ?,
      isOnsale = ?,
      normalPrice = ?,
      outSkuSn = ?,
      skuQuantity = ?,
      skuSoldQuantity = ?,
      spec = ?,
      specColor = ?,
      specSize = ?
    WHERE
      pddId = ?
    AND
      skuId = ?
  `)
  if err != nil {
    log.Println("sync-pdd-item-save-sku-list-update-prepare-error: ", err)
    return
  }
  defer stmtUpdate.Close()
  for i := 0; i < len(skuList); i++ {
    sku := skuList[i].(map[string]interface{})
    activityGroupPrice := sku["activityGroupPrice"]
    groupPrice := sku["groupPrice"]
    isOnsale := sku["isOnsale"]
    normalPrice := sku["normalPrice"]
    outSkuSn := sku["outSkuSn"]
    skuId := sku["skuId"]
    skuQuantity := sku["skuQuantity"]
    skuSoldQuantity := sku["skuSoldQuantity"]
    spec := sku["spec"].(string)
    specList := strings.Split(spec, " ")
    specColor := specList[0]
    specSize := ""
    if len(specList) > 1 {
      specSize = specList[1]
    }
    var count int
    err = db.QueryRow(`
      SELECT
        COUNT(*)
      FROM
        pddItemSku
      WHERE
        pddId = ?
      AND
        skuId = ?
    `, pddId, skuId).Scan(&count)
    if err != nil {
      log.Println("sync-pdd-item-save-sku-list-count-error: ", err)
      return
    }
    if count == 0 {
      _, err = stmtInsert.Exec(
        pddId,
        outGoodsSn,
        activityGroupPrice,
        groupPrice,
        isOnsale,
        normalPrice,
        outSkuSn,
        skuId,
        skuQuantity,
        skuSoldQuantity,
        spec,
        specColor,
        specSize,
      )
      if err != nil {
        log.Println("sync-pdd-item-save-sku-list-insert-exec-error: ", err)
        return
      }
    } else {
      _, err = stmtUpdate.Exec(
        outGoodsSn,
        activityGroupPrice,
        groupPrice,
        isOnsale,
        normalPrice,
        outSkuSn,
        skuQuantity,
        skuSoldQuantity,
        spec,
        specColor,
        specSize,
        pddId,
        skuId,
      )
      if err != nil {
        log.Println("sync-pdd-item-save-sku-list-update-exec-error: ", err)
        return
      }
    }
  }
}
