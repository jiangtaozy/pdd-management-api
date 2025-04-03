/*
 * Maintained by jemo from 2025.02.06 to now
 * Created by jemo on 2025.02.06 15:54:42
 * 拼多多电子面单数据-用于计算仓库库存数据
 */

package database

const createPddExpressBill =`
  CREATE TABLE IF NOT EXISTS pddExpressBill (
    waybillCode VARCHAR(50) COMMENT '电子面单号',
    shipName VARCHAR(50) COMMENT '快递公司',
    branchName VARCHAR(100) COMMENT '网点名称',
    showStatus TINYINT COMMENT '单号状态, 2: 已使用, 3: 回收中',
    createTime DATETIME COMMENT '取号时间',
    receiptAddress VARCHAR(200) COMMENT '收货地址',
    accountId VARCHAR(50) COMMENT '单号申请者',
    userId VARCHAR(50) COMMENT '单号使用者',
    mallId VARCHAR(50) COMMENT '店铺ID',
    orderId VARCHAR(50) COMMENT '拼多多订单号',
    PRIMARY KEY(waybillCode, orderId)
  );
`
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773336487414791', '河北定州市公司明月店镇寄存点分部', 2, '250124-430838987133717');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('SF0267078411003', '河北定州市公司明月店镇寄存点分部', 2, '250124-430838987133717');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('SF0267078411003', '河北定州市公司明月店镇寄存点分部', 2, '250126-538192167302731');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('YT7522221549608', '河北定州市公司明月店镇寄存点分部', 2, '250119-215849377300309');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('JT9000327748223', '杭州余杭海创园网点', 2, '250109-250368541021961');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773336485832800', '杭州余杭海创园网点', 2, '250124-057357156660300');

//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339396650469', '杭州余杭海创园网点', 2, '250217-203098765663977');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339384206734', '杭州余杭海创园网点', 2, '250217-076766319001940');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339384366602', '杭州余杭海创园网点', 2, '250217-125955029381278');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339383589323', '杭州余杭海创园网点', 2, '250217-219991304610607');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339384366640', '杭州余杭海创园网点', 2, '250217-583280955751303');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339384899497', '杭州余杭海创园网点', 2, '250217-489905242502509');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339384520941', '杭州余杭海创园网点', 2, '250217-058426672742282');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339211275343', '杭州余杭海创园网点', 2, '250216-274999622012059');
//insert into pddExpressBill (waybillCode, branchName, showStatus, orderId) values ('773339212204810', '杭州余杭海创园网点', 2, '250215-421045235350350');
