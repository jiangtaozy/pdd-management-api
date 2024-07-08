/*
 * Maintained by jemo from 2023.11.28 to now
 * Created by jemo on 2023.11.28 16:26:42
 * 包裹出入库记录
 */

package database

const createPackageSkuRecord =`
  CREATE TABLE IF NOT EXISTS packageSkuRecord (
    packageId INTEGER NOT NULL COMMENT '包裹编号',
    labelCode VARCHAR(20) NOT NULL COMMENT '商品条码号',
    quantity INTEGER NOT NULL COMMENT '数量',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
