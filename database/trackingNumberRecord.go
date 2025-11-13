/*
 * Maintained by jemo from 2025.08.08 to now
 * Created by jemo on 2025.08.08 20:59:43
 * 扫描快递单记录
 */

package database

const createTrackingNumberRecord =`
  CREATE TABLE IF NOT EXISTS trackingNumberRecord (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    trackingNumber VARCHAR(50) COMMENT '快递单号',
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
