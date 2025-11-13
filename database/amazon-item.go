/*
 * Maintained by jemo from 2025.06.13 to now
 * Created by jemo on 2025.06.13 17:27:28
 * 亚马逊商品列表
 */

package database

const createAmazonItem =`
  CREATE TABLE IF NOT EXISTS amazonItem (
    sku VARCHAR(10) COMMENT 'sku货号',
    asin VARCHAR(15) PRIMARY KEY COMMENT 'asin',
    fnSku VARCHAR(15) COMMENT '条码id',
    price DECIMAL(10,2) COMMENT '零售价',
    currencyUnit VARCHAR(15) COMMENT '零售价币种',
    available INTEGER COMMENT '有货（亚马逊物流）',
    unfulfillable INTEGER COMMENT '不可售',
    inbound INTEGER COMMENT '入库',
    reserved INTEGER COMMENT '预留',
    amount DECIMAL(10,2) COMMENT '最近30天，销售额，美元',
    unitsSold INTEGER COMMENT '最近30天，售出件数',
    state VARCHAR(32) COMMENT '状态: 缺货，在售，报价缺失，不可售，已禁止显示搜索结果',
    lastChanged DATETIME COMMENT '最后修改时间',
    created DATETIME COMMENT '创建时间',
    title VARCHAR(320) COMMENT '标题',
    imageUrl VARCHAR(100) COMMENT '图片',
    barcode VARCHAR(30) COMMENT '条码'
  );
`

//alter table amazonItem add barcode VARCHAR(30) COMMENT '条码';
//alter table amazonItem drop column isDelivery;
//alter table amazonItem drop column deliveryVolume;
