/*
 * Maintained by jemo from 2025.06.13 to now
 * Created by jemo on 2025.06.13 17:49:17
 * 亚马逊销售数据
 */

package database

const createAmazonSalesData =`
  CREATE TABLE IF NOT EXISTS amazonSalesData (
    day DATE NOT NULL,
    sku VARCHAR(10) COMMENT 'sku货号',
    asin VARCHAR(15) COMMENT 'asin',
    fnSku VARCHAR(15) COMMENT '条码id',
    price DECIMAL(10,2) COMMENT '零售价',
    amount DECIMAL(10,2) COMMENT '最近30天，销售额',
    currencyUnit VARCHAR(15) COMMENT '零售价币种',
    unitsSold INTEGER COMMENT '最近30天，售出件数',
    available INTEGER COMMENT '有货（亚马逊物流）',
    unfulfillable INTEGER COMMENT '',
    inbound INTEGER COMMENT '',
    reserved INTEGER COMMENT '',
    state VARCHAR(32) COMMENT '状态',
    PRIMARY KEY(day, asin)
  );
`
