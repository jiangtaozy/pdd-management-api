/*
 * Maintained by jemo from 2024.11.26 to now
 * Created by jemo on 2024.11.26 09:21:55
 * 月上新
 */

package database

const createMonthNew =`
  CREATE TABLE IF NOT EXISTS monthNew (
    month DATE NOT NULL PRIMARY KEY,
    newQuantity INTEGER COMMENT '上新数量',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );
`
