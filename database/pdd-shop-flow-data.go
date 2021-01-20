/*
 * Maintained by jemo from 2021.1.20 to now
 * Created by jemo on 2021.1.20 14:00:23
 * Pdd Shop Flow Data
 * 拼多多店铺流量数据
 */

package database

const createPddShopFlowData =`
  CREATE TABLE IF NOT EXISTS pddShopFlowData (
    cfmOrdrCnt INTEGER,
    cfmOrdrCntIsPercent BOOLEAN,
    cfmOrdrCntPct INTEGER,
    cfmOrdrUsrCnt INTEGER,
    cfmOrdrUsrCntIsPercent BOOLEAN,
    cfmOrdrUsrCntPct INTEGER,
    goodsFavCnt INTEGER,
    goodsFavCntIsPercent BOOLEAN,
    goodsFavCntPct FLOAT,
    gpv INTEGER,
    gpvIsPercent BOOLEAN,
    gpvPct FLOAT,
    guv INTEGER,
    guvIsPercent BOOLEAN,
    guvPct FLOAT,
    payOrdrAmt FLOAT,
    payOrdrAmtIsPercent BOOLEAN,
    payOrdrAmtPct FLOAT,
    payOrdrCnt INTEGER,
    payOrdrCntIsPercent BOOLEAN,
    payOrdrCntPct INTEGER,
    payOrdrUsrCnt INTEGER,
    payOrdrUsrCntIsPercent BOOLEAN,
    payOrdrUsrCntPct INTEGER,
    payUvRto FLOAT,
    payUvRtoIsPercent BOOLEAN,
    payUvRtoPct FLOAT,
    vstGoodsCnt INTEGER,
    vstGoodsCntIsPercent BOOLEAN,
    vstGoodsCntPct FLOAT,
    date DATE NOT NULL,
    INDEX pddShopFlowDataIndex (date)
  );
`
