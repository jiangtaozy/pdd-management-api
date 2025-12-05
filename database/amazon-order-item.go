/*
 * Maintained by jemo from 2025.11.18 to now
 * Created by jemo on 2025.11.18 17:58:56
 * 亚马逊订单商品
 */

package database

const createAmazonOrderItem =`
  CREATE TABLE IF NOT EXISTS amazonOrderItem (
    amazonOrderId VARCHAR(30) COMMENT '订单号',
    asin VARCHAR(20),
    billingCountry VARCHAR(20),
    customerOrderItemCode VARCHAR(100) NOT NULL PRIMARY KEY,
    extendedTitle VARCHAR(300),
    imageUrl VARCHAR(100),
    orderItemId VARCHAR(30),
    quantityCanceled INTEGER,
    quantityOrdered INTEGER,
    quantityShipped INTEGER,
    quantityUnShipped INTEGER,
    sellerSku VARCHAR(20),
    unitPriceAmount FLOAT COMMENT '价格',
    unitPriceCurrencyCode VARCHAR(10)
  );
`
