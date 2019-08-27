set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


-- ----------------------------
-- ylz create by 2018-11-09
-- Table structure for t_marry_pre_wed  玩家婚礼预定档次(预定不成功返还)
-- ----------------------------
CREATE TABLE `t_marry_pre_wed` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `period` int(11) NOT NULL COMMENT "场次",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `playerName` varchar(100) NOT NULL COMMENT "玩家名字",
  `peerId` bigint(20) NOT NULL COMMENT "对方id",
  `grade` int(11) NOT NULL COMMENT "酒席档次",
  `hunCheGrade` int(11) NOT NULL COMMENT "婚车档次",
  `sugarGrade` int(11) NOT NULL COMMENT "喜糖档次",
  `status` int(11) NOT NULL COMMENT "状态 1进行中 2失败",
  `holdTime` bigint(20) NOT NULL COMMENT "举办时间",
  `preWedTime` bigint(20) DEFAULT 0 COMMENT "预定时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- xzk create by 2018-11-09
-- Table structure for t_player_tianshu 玩家天书数据
-- ----------------------------
CREATE TABLE `t_player_tianshu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "天书类型",
  `level` int(11) NOT NULL COMMENT "天书等级",
  `isReceive` int(11) NOT NULL COMMENT "是否领取",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
    KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- cjb create by 2018-11-12
alter table `t_player_cache` add column `massacreInfo` text(2000) NOT NULL COMMENT "戮仙刃";
UPDATE `t_player_cache` SET `massacreInfo`= "{}";

-- xzk create by 2018-11-12
-- Table structure for t_player_myboss 玩家个人BOSS数据
-- ----------------------------
CREATE TABLE `t_player_myboss` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `attendMap` varchar(512) NOT NULL COMMENT "参与次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
    KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk create by 2018-11-13
alter table `t_player_gold_equip_slot` add column `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据";

-- cjb create by 2018-11-13
-- Table structure for t_player_addition_sys_slot 玩家附加系统装备槽数据
-- ----------------------------
CREATE TABLE `t_player_addition_sys_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `level` int(11) NOT NULL COMMENT "等级",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- xzk create by 2018-11-13
alter table `t_player_item` add column `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据";

alter table `t_alliance_depot` add column `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据";

alter table `t_player_tower` add column `lastResetTime` bigint(20) NOT NULL COMMENT "上次打宝时间重置时间";


