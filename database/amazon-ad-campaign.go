/*
 * Maintained by jemo from 2025.11.20 to now
 * Created by jemo on 2025.11.20 16:16:14
 * 亚马逊广告活动
 */

package database

const createAmazonAdCampaign =`
  CREATE TABLE IF NOT EXISTS amazonAdCampaign (
    campaignRetailer VARCHAR(20),
    campaignEffectiveBudget FLOAT,
    campaignId VARCHAR(30) NOT NULL PRIMARY KEY,
    currencyCode VARCHAR(20),
    campaignExternalId VARCHAR(30),
    bidType VARCHAR(20),
    startDate DATE,
    campaignType VARCHAR(20),
    campaignName VARCHAR(200)
  );
`
