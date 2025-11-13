/*
 * Maintained by jemo from 2025.08.23 to now
 * Created by jemo on 2025.08.23 16:56:43
 * 亚马逊广告记录
 */

package database

const createAmazonAdRecord =`
  CREATE TABLE IF NOT EXISTS amazonAdRecord (
    id VARCHAR(30),
    state VARCHAR(20),
    name VARCHAR(200),
    calculatedStatus VARCHAR(30),
    sku VARCHAR(20),
    asin VARCHAR(20),
    asinPriceMillicents INTEGER UNSIGNED,
    asinPriceCurrencyCode VARCHAR(10),
    productImage VARCHAR(200),
    unitsSold INTEGER UNSIGNED,
    detailPageViews INTEGER UNSIGNED,
    impressions INTEGER UNSIGNED,
    clicks INTEGER UNSIGNED,
    spendMillicents INTEGER UNSIGNED,
    spendCurrencyCode VARCHAR(10),
    spendCoV INTEGER UNSIGNED,
    salesMillicents INTEGER UNSIGNED,
    salesCurrencyCode VARCHAR(10),
    multiTouchSalesMillicents INTEGER UNSIGNED,
    multiTouchSalesCurrencyCode VARCHAR(10),
    salesCoV INTEGER UNSIGNED,
    cpcMillicents INTEGER UNSIGNED,
    cpcCurrencyCode  VARCHAR(10),
    cpcCoV INTEGER UNSIGNED,
    ctr FLOAT,
    acos FLOAT,
    multiTouchAcos FLOAT,
    roas FLOAT,
    multiTouchRoas FLOAT,
    orders INTEGER UNSIGNED,
    multiTouchOrders INTEGER UNSIGNED,
    ntbCostPerOrder INTEGER UNSIGNED,
    ntbCostPerOrderFromClicks INTEGER UNSIGNED,
    ntbDetailPageVisits INTEGER UNSIGNED,
    ntbOrders INTEGER UNSIGNED,
    correctedAcos FLOAT,
    correctedRoas FLOAT,
    createdAt DATE DEFAULT CURRENT_DATE,
    unique key amazonAdUniqueKey (id, createdAt)
  );
`
