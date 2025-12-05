/*
 * Maintained by jemo from 2025.11.20 to now
 * Created by jemo on 2025.11.20 16:26:09
 * 亚马逊广告组
 */

package database

const createAmazonAdGroup =`
  CREATE TABLE IF NOT EXISTS amazonAdGroup (
    adGroupName VARCHAR(200),
    campaignExternalId VARCHAR(30),
    adGroupId VARCHAR(30) NOT NULL PRIMARY KEY,
    statusName VARCHAR(50)
  );
`
