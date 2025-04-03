/*
 * Maintained by jemo from 2025.01.13 to now
 * Created by jemo on 2025.01.13 09:59:42
 * shop
 */

package database

const createShop =`
  CREATE TABLE IF NOT EXISTS shop (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    shopId VARCHAR(20) COMMENT '店铺id',
    shopName VARCHAR(50) COMMENT '店铺名称'
  );
`
//insert into shop (shopId, shopName) values ('128462412746'   , 'k sauce'              );
//insert into shop (shopId, shopName) values ('3985500114551'  , 'mujx'                       );
//insert into shop (shopId, shopName) values ('5075079321130'  , 'abracelet'                  );
//insert into shop (shopId, shopName) values ('5075838821406'  , 'bhat'                       );
//insert into shop (shopId, shopName) values ('5076143310501'  , 'csocks'                     );
//insert into shop (shopId, shopName) values ('5667407834765'  , 'dbag'                       );
//insert into shop (shopId, shopName) values ('5736772513034'  , 'ebuckle'                    );
//insert into shop (shopId, shopName) values ('5837085643448'  , 'fglasses'                   );
//insert into shop (shopId, shopName) values ('6055227764828'  , 'gboys'                      );
//insert into shop (shopId, shopName) values ('6057084551138'  , 'hfan'                       );
//insert into shop (shopId, shopName) values ('6183426731163'  , 'igame'                      );
//insert into shop (shopId, shopName) values ('6199741427526'  , 'jsticker'                   );
//insert into shop (shopId, shopName) values ('634418215239967', 'k sauce local'    );
//insert into shop (shopId, shopName) values ('404105700701'   , 'Awesome Selection - k sauce');
