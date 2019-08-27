 set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

 -- create by xubin 2019-08-06
 alter table `t_player_lingtong` add column `basePower`  bigint(20) NOT NULL COMMENT "基础战力";
 alter table `t_player_shenqi` add column `power`  bigint(20) NOT NULL COMMENT "战力";

 -- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_goldequip 玩家元神金装数据
-- ----------------------------
 CREATE TABLE `t_player_goldequip` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_mingge 玩家命格数据
-- ----------------------------
CREATE TABLE `t_player_mingge` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_shenghen 玩家圣痕数据
-- ----------------------------
CREATE TABLE `t_player_shenghen` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_zhenfa_power 玩家阵法数据
-- ----------------------------
CREATE TABLE `t_player_zhenfa_power` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_tulong_equip 玩家屠龙装数据
-- ----------------------------
CREATE TABLE `t_player_tulong_equip` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_baby_power 玩家宝宝数据
-- ----------------------------
CREATE TABLE `t_player_baby_power` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



 -- create by jzy 2019-08-08
-- ----------------------------
-- Table structure for t_player_addition_sys_lingzhu 附加系统五行灵珠
-- ----------------------------
CREATE TABLE `t_player_addition_sys_lingzhu` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `sysType` bigint(11) NOT NULL COMMENT "附加系统类型",
    `lingZhuId` bigint(11) NOT NULL COMMENT "灵珠类型",
    `level` bigint(11) NOT NULL COMMENT "等级",
    `times` bigint(11) NOT NULL COMMENT "次数",
    `bless` bigint(20) NOT NULL COMMENT "祝福值",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


 -- create by xubin 2019-08-13
-- ----------------------------
-- Table structure for t_player_xianzun_card 玩家仙尊特权卡数据
-- ----------------------------
CREATE TABLE `t_player_xianzun_card` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `typ` bigint(11) NOT NULL COMMENT "仙尊特权卡类型",
    `IsActivite` bigint(11) NOT NULL COMMENT "是否激活",
    `isReceive` bigint(11) NOT NULL COMMENT "是否领取",
    `activiteTime` bigint(20) NOT NULL COMMENT "激活时间",
    `receiveTime` bigint(20) NOT NULL COMMENT "领取时间",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`), 
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xzk 2019/8/13
alter table `t_player_zhenxi_boss` add column `enterTimes` int(11) NOT NULL COMMENT "进入次数"; 


 -- create by jzy 2019-08-14
-- ----------------------------
-- Table structure for t_player_wushuang_buchang 玩家无双神器补偿邮件数据
-- ----------------------------
CREATE TABLE `t_player_wushuang_buchang` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `isSendEmail` bigint(11) NOT NULL COMMENT "是否发送邮件",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`), 
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by xubin 2019-08-14
alter table `t_player_cache` add column `xianZunCardInfo` text(2000) NOT NULL COMMENT "仙尊特权卡类型";
UPDATE `t_player_cache` SET `xianZunCardInfo`= "[]"; 


-- ----------------------------
-- Table structure for t_player_privilege_charge  玩家扶持充值记录
-- ----------------------------
CREATE TABLE `t_player_privilege_charge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chargeType`  int(11) NOT NULL COMMENT "平台类型",
  `chargeId` int(11) NOT NULL COMMENT "充值模板id",
  `chargeNum` bigint(20) NOT NULL COMMENT "元宝数量",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`), 
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;