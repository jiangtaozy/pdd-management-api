/*
 * Maintained by jemo from 2020.10.10 to now
 * Created by jemo on 2020.10.10 21:24:42
 * merge order id
 * 合并订单号
 */

package database

const createMergeOrderId =`
  CREATE TABLE IF NOT EXISTS mergeOrderId (
    mainOrderId VARCHAR(30) NOT NULL COMMENT '主订单号',
    mergedOrderId VARCHAR(30) NOT NULL COMMENT '被合并订单号',
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (mainOrderId, mergedOrderId)
  );
`
