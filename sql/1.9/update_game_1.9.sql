set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_info 玩家灵童信息
-- ----------------------------
CREATE TABLE `t_player_lingtong_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingTongId` int(11) NOT NULL COMMENT "灵童id",
  `lingTongName` varchar(50) NOT NULL COMMENT "灵童名字",
  `upgradeLevel` int(11) NOT NULL COMMENT "升级等级",
  `upgradeNum` int(11) NOT NULL COMMENT "升级次数",
  `upgradePro` int(11) NOT NULL COMMENT "升级进度值",
  `peiYangLevel` int(11) NOT NULL COMMENT "培养等级",
  `peiYangNum`  int(11) NOT NULL COMMENT "培养次数", 
  `peiYangPro`  bigint(20) NOT NULL COMMENT "培养进度值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong 玩家灵童
-- ----------------------------
CREATE TABLE `t_player_lingtong` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingTongId` int(11) NOT NULL COMMENT "出战灵童id",
  `level` int(11) NOT NULL COMMENT "灵童等级", 
  `power`  bigint(20) NOT NULL COMMENT "战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_fashion 玩家灵童时装信息
-- ----------------------------
CREATE TABLE `t_player_lingtong_fashion` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingTongId` int(11) NOT NULL COMMENT "灵童id",
  `fashionId` int(11) NOT NULL COMMENT "灵童时装id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_fashion_info 玩家灵童时装信息
-- ----------------------------
CREATE TABLE `t_player_lingtong_fashion_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `fashionId` int(11) NOT NULL COMMENT "灵童时装id",
  `upgradeLevel` int(11) NOT NULL COMMENT "升级等级",
  `upgradeNum` int(11) NOT NULL COMMENT "升级次数",
  `upgradePro` int(11) NOT NULL COMMENT "升级进度值",
  `isExpire` int(11) NOT NULL COMMENT "是否失效",     
	`activateTime` bigint(20) DEFAULT 0 COMMENT "激活时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_develop 玩家灵童养成信息
-- ----------------------------
CREATE TABLE `t_player_lingtong_develop` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `classType` int(11) NOT NULL COMMENT "养成系统类型",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `seqId` int(11) NOT NULL COMMENT "当前外观id",
  `unrealLevel` int(11) NOT NULL COMMENT "食幻化丹等级",
  `unrealNum` int(11) NOT NULL COMMENT "食幻化丹次数",
  `unrealPro` int(11) NOT NULL COMMENT "食幻化丹进度值",
  `culLevel`  int(11) NOT NULL COMMENT "食培养丹等级",
  `culNum` int(11) NOT NULL COMMENT "食培养丹次数",
  `culPro` int(11) NOT NULL COMMENT "食培养丹进度值",
  `tongLingLevel` int(11) NOT NULL COMMENT "食通灵丹等级",
  `tongLingNum` int(11) NOT NULL COMMENT "食通灵丹次数",
  `tongLingPro` int(11) NOT NULL COMMENT "食通灵丹进度值",
  `unrealInfo` varchar(256) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `hidden` int(11) default 0 COMMENT "是否隐藏外观",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_other 玩家灵童养成非进阶信息
-- ----------------------------
CREATE TABLE `t_player_lingtong_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `classType` int(11) NOT NULL COMMENT "养成系统类型",
  `type` int(11) NOT NULL COMMENT "皮肤类型",
  `seqId` int(11) NOT NULL COMMENT "皮肤id",
  `level` int(11) NOT NULL COMMENT "升星等级",
  `upNum` int(11) NOT NULL COMMENT "升星次数",
  `upPro` int(11) NOT NULL COMMENT "升星培养值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_power 玩家灵童养成战力信息
-- ----------------------------
CREATE TABLE `t_player_lingtong_power` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `classType` int(11) NOT NULL COMMENT "养成系统类型",
  `power`  bigint(20) NOT NULL COMMENT "战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by xzk 2019-01-10 
alter table `t_player_lingyu` add column `chargeVal` bigint(20) NOT NULL COMMENT "模块开启累计充值";

-- ----------------------------
-- create by ylz 2019-01-09
-- Table structure for t_player_lingtong_fashion_trial  玩家灵童时装试用卡阶数
-- ----------------------------
CREATE TABLE `t_player_lingtong_fashion_trial` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `trialFashionId`  int(11) NOT NULL COMMENT "时装id",
  `isExpire` int(11) NOT NULL COMMENT "是否失效 0否 1失效", 
  `activateTime` bigint(20) DEFAULT 0 COMMENT "激活时间",
  `durationTime` bigint(20) DEFAULT 0 COMMENT "持续时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by xzk 2019-09-11
-- Table structure for t_open_activity_boss_kill  运营活动BOSS首杀
-- ----------------------------
CREATE TABLE `t_open_activity_boss_kill` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `bossIdList` varchar(512) NOT NULL COMMENT "bossId列表",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by ylz 2019-01-11
alter table `t_player_cache` add column `allLingTongDevInfo` text(5000) NOT NULL COMMENT "灵童养成类信息";
UPDATE `t_player_cache` SET `allLingTongDevInfo`= "{}";

-- create by ylz 2019-01-11
alter table `t_player_cache` add column `lingTongInfo` text(5000) NOT NULL COMMENT "灵童信息";
UPDATE `t_player_cache` SET `lingTongInfo`= "{}";

-- create by ylz 2019-01-15
alter table `t_player_cache` add column `allSystemSkillInfo` text(5000) NOT NULL COMMENT "系统技能";
UPDATE `t_player_cache` SET `allSystemSkillInfo`= "{}";

-- create by cjb 2019-01-15
alter table `t_player_cache` add column `allAdditionSysInfo` text(5000) NOT NULL COMMENT "附加系统类信息";
UPDATE `t_player_cache` SET `allAdditionSysInfo`= "{}";

-- create by xzk 2019-01-16
alter table `t_player_cycle_charge_record` add column `preDayChargeNum` bigint(20) NOT NULL COMMENT "前一天充值数";

-- ----------------------------
-- create by xzk 2019-01-16
-- Table structure for t_player_cycle_cost_record  玩家每日消费记录
-- ----------------------------
CREATE TABLE `t_player_cycle_cost_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `costNum`  bigint(20) NOT NULL COMMENT "消费元宝数",
  `preDayCostNum`  bigint(20) NOT NULL COMMENT "上一天消费元宝数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by xzk 2019-01-16
-- Table structure for t_open_activity_alliance_cheer  城战助威记录
-- ----------------------------
CREATE TABLE `t_open_activity_alliance_cheer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `allianceId`  bigint(20) NOT NULL COMMENT "仙盟id",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)  
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;  

-- create by zrc 2019-01-16
alter table `t_player` add column `ip` varchar(50) NOT NULL COMMENT "ip";

-- ----------------------------
-- create by cjb 2019-01-18
-- Table structure for t_open_activity_crazybox_log  运营活动-疯狂宝箱日志
-- ----------------------------
CREATE TABLE `t_open_activity_crazybox_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",
  `playerName` varchar(20) NOT NULL COMMENT "玩家姓名",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;