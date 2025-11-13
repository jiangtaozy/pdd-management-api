/*
 * Maintained by jemo from 2025.08.25 to now
 * Created by jemo on 2025.08.25 16:15:32
 * 亚马逊广告
 */

package database

const createAmazonAd =`
  CREATE TABLE IF NOT EXISTS amazonAd (
    id VARCHAR(30) NOT NULL PRIMARY KEY,
    state VARCHAR(20),
    name VARCHAR(200),
    calculatedStatus VARCHAR(30),
    sku VARCHAR(20),
    asin VARCHAR(20),
    asinPriceMillicents INTEGER UNSIGNED,
    asinPriceCurrencyCode VARCHAR(10),
    productImage VARCHAR(200),
    createdAt DATE DEFAULT CURRENT_DATE
  );
`
