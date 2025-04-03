/*
 * Maintained by jemo from 2025.02.11 to now
 * Created by jemo on 2025.02.11 14:51:08
 * 亚马逊货件详情
 */

package database

const amazonShipmentInfo = `
  CREATE TABLE IF NOT EXISTS amazonShipmentInfo (
    shipmentId VARCHAR(30) COMMENT '货件id',
    shipmentName VARCHAR(200),
    shipmentStatus VARCHAR(20),
    encryptedMerchantId VARCHAR(20),
    purchaseOrderId VARCHAR(20),
    destinationFC VARCHAR(20),
    shipFromAddressId VARCHAR(200),
    createdDate DATETIME,
    lastUpdatedDate DATETIME,
    needByDate DATETIME,
    expectedShipDate DATETIME,
    workflowType VARCHAR(20),
    workflowId VARCHAR(100),
    plannedShipmentCategory VARCHAR(20),
    onlyCasesAllowed  BOOLEAN,
    declaredCartonInformationSource VARCHAR(20),
    totalShippedQty INTEGER,
    totalLineItems INTEGER,
    totalReceivedQty INTEGER,
    totalCases INTEGER,
    totalLineItemsRequiringExpDate INTEGER,
    boxLevelPlacement BOOLEAN,
    merchantExemptedFromCLIPlannedServiceFee BOOLEAN,
    partneredImport BOOLEAN,
    fromAddressOwnerName VARCHAR(30),
    fromAddressLine1 VARCHAR(100),
    fromAddressLine2 VARCHAR(100),
    fromCity VARCHAR(20),
    fromStateOrRegion VARCHAR(20),
    fromDistrictOrCounty VARCHAR(20),
    fromPostalCode VARCHAR(20),
    fromCountryCode VARCHAR(10),
    fromPrimaryPhoneNumber VARCHAR(20),
    feesAmount FLOAT,
    feesCurrency VARCHAR(10),
    toName VARCHAR(100),
    toAddressLine1 VARCHAR(100),
    toAddressLine2 VARCHAR(100),
    toCity VARCHAR(30),
    toState VARCHAR(30),
    toDistrict VARCHAR(30),
    toCountryCode VARCHAR(10),
    toPostalCode VARCHAR(10),
    PRIMARY KEY(shipmentId)
  );
`
