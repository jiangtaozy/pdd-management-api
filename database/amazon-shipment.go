/*
 * Maintained by jemo from 2025.02.11 to now
 * Created by jemo on 2025.02.11 11:30:47
 * 亚马逊货件
 */

package database

const amazonShipment = `
  CREATE TABLE IF NOT EXISTS amazonShipment (
    shipmentId VARCHAR(30) COMMENT '货件id',
    eligible BOOLEAN,
    ineligibilityReason VARCHAR(100),
    estimatedTimestamp DATETIME,
    formerEstimatedTimestamp DATETIME,
    daysToWait INTEGER,
    shipmentStatus VARCHAR(30),
    hasDiscrepancy VARCHAR(30),
    merchantLaunched BOOLEAN,
    wmsrAllowed BOOLEAN,
    merchantSKU VARCHAR(20),
    title VARCHAR(400),
    fnsku VARCHAR(20),
    asin VARCHAR(20),
    upc VARCHAR(20),
    barcodeType VARCHAR(20),
    barcodeValue VARCHAR(20),
    isSortable BOOLEAN,
    eligibleForResearch BOOLEAN,
    expectedQuantity INTEGER,
    totalLocatedQuantity INTEGER,
    totalReceivedQuantity INTEGER,
    totalAdjustmentQuantity INTEGER,
    totalReimbursedQuantity INTEGER,
    totalClawbackQuantity INTEGER,
    totalAcknowledgedQuantity INTEGER,
    totalDeniedQuantity INTEGER,
    totalVirtuallyAdjustedQuantity INTEGER,
    otherReceivedQuantity INTEGER,
    otherReceiveStartDate DATETIME,
    otherReceiveEndDate DATETIME,
    iarAutoReconciliations VARCHAR(20),
    totalDiscrepancyQuantity INTEGER,
    caseId VARCHAR(20),
    itemCondition VARCHAR(20),
    reasonCode VARCHAR(20),
    PRIMARY KEY(shipmentId, fnsku)
  );
`
