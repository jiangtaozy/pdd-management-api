/*
 * Maintained by jemo from 2024.2.28 to now
 * Created by jemo on 2024.2.28 11:29:06
 * 备选商品
 */
package database

const standbyItem = `
  CREATE TABLE IF NOT EXISTS standbyItem (
    searchId VARCHAR(20) NOT NULL COMMENT 'SKC货号',
    standbyId VARCHAR(20) NOT NULL COMMENT '备选SKC货号',
    UNIQUE INDEX searchIndex (searchId, standbyId)
  );
 `
/*
insert into standbyItem (searchId, standbyId) values('26722', '26905');
insert into standbyItem (searchId, standbyId) values('26722', '26906');
*/
