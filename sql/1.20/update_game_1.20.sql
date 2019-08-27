set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

alter table `t_player_addition_sys_slot` add column `shenZhuLev` int(11) NOT NULL COMMENT "神铸等级";
alter table `t_player_addition_sys_slot` add column `shenZhuNum` int(11) NOT NULL COMMENT "神铸次数";
alter table `t_player_addition_sys_slot` add column `shenZhuPro` int(11) NOT NULL COMMENT "神铸进度";

-- ----------------------------
-- Table structure for t_player_activity_collect  个人采集
-- ----------------------------
CREATE TABLE `t_player_activity_collect` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType` int(11) NOT NULL COMMENT "活动类型", 
  `countMap` varchar(512) NOT NULL COMMENT "采集次数map", 
  `endTime` bigint(20) NOT NULL COMMENT "场景结束时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",  
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_trade_recycle  个人系统回购
-- ----------------------------
DROP TABLE IF EXISTS `t_player_trade_recycle`;
CREATE TABLE `t_player_trade_recycle` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `recycleGold` bigint(20) NOT NULL COMMENT "回收的元宝",
  `recycleTime` bigint(20) NOT NULL COMMENT "回收时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `t_trade_recycle`;
CREATE TABLE `t_trade_recycle` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `recycleGold` bigint(20) NOT NULL COMMENT "回收的元宝",
  `recycleTime` bigint(20) NOT NULL COMMENT "回收时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xb create by 2019-05-07
-- Table structure for t_player_addition_sys_awake  附加系统觉醒
-- ----------------------------
CREATE TABLE `t_player_addition_sys_awake` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型", 
  `isAwake` int(11) NOT NULL COMMENT "是否觉醒", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",  
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- cjb create by 2019-05-07
-- Table structure for t_player_addition_sys_tongling 玩家附加系统通灵数据
-- ----------------------------
CREATE TABLE `t_player_addition_sys_tongling` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型",
  `tongLingLev` int(11) NOT NULL COMMENT "通灵等级",
  `tongLingNum` int(11) NOT NULL COMMENT "通灵次数",
  `tongLingPro` int(11) NOT NULL COMMENT "通灵进度",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xzk 2019-05-08
alter table `t_player_alliance` add column `warPoint` int(11) DEFAULT 0 COMMENT "城战积分";


-- xzk create by 2019-05-09
-- Table structure for t_player_alliance_yuxi 玩家仙盟玉玺之战
-- ----------------------------
CREATE TABLE `t_player_alliance_yuxi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `isReceive` int(11) NOT NULL COMMENT "是否领取每日奖励",  
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`), 
   KEY(`playerId`),  
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by xzk 2019-05-08
alter table `t_player_alliance` add column `lastYuXiMemberCallTime` bigint(20) DEFAULT 0 COMMENT "上次玉玺之战仙盟召集时间"; 


-- ----------------------------
-- Table structure for t_activity_end_record  日常活动开启记录
-- ----------------------------
CREATE TABLE `t_activity_end_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `activityType` int(11) NOT NULL COMMENT "活动类型", 
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xzk 2019-05-15
alter table `t_alliance_hegemon` add column `defenceAllianceId` bigint(20) NOT NULL COMMENT "守方仙盟id"; 