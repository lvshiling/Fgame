set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

-- ----------------------------
-- Table structure for t_player_supreme_title_wear 玩家穿戴至尊称号
-- ----------------------------
CREATE TABLE `t_player_supreme_title_wear` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `titleWear` int(11) NOT NULL COMMENT "穿戴称号id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_supreme_title 玩家至尊称号数据
-- ----------------------------
CREATE TABLE `t_player_supreme_title` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `titleId` int(11) NOT NULL COMMENT "称号id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by xzk 2019-02-26
-- Table structure for t_open_activity_xun_huan 玩家循环活动数据
-- ----------------------------
CREATE TABLE `t_open_activity_xun_huan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `arrGroup` int(11) NOT NULL COMMENT "随机活动组",
  `activityDay` int(20) NOT NULL COMMENT "当前活动天",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",  
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by ylz 2019-02-26
alter table `t_player_daily` add column  `dailyTag` int(11) DEFAULT 1 COMMENT "日环类型 1日环 2仙盟日常";


-- Table structure for t_player_equipbaoku 玩家装备宝库
-- ----------------------------
CREATE TABLE `t_player_equipbaoku` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `luckyPoints` int(11) NOT NULL COMMENT "幸运值",
  `attendPoints` int(11) NOT NULL COMMENT "积分",
  `totalAttendTimes` int(11) NOT NULL COMMENT "总参与次数",
  `lastSystemRefreshTime` bigint(20) DEFAULT 0 COMMENT "上次自动刷新时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_alliance_boss  仙盟boss
-- ----------------------------
CREATE TABLE `t_alliance_boss` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",
  `summonTime` bigint(20) NOT NULL COMMENT "召唤时间",
  `bossLevel` int(11) NOT NULL COMMENT "boss等级",
  `bossExp` int(11) NOT NULL COMMENT "boss经验",
  `isSummon` int(11) NOT NULL COMMENT "当日是否已经召唤过",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by cjb 2019-02-27
alter table `t_player_inventory` add column  `miBaoDepotNum` int(11) DEFAULT 0 COMMENT "秘宝仓库格子数";

-- ----------------------------
-- Table structure for t_player_equipbaoku_shop 玩家当日宝库商店购买道具(限购使用)
-- ----------------------------
CREATE TABLE `t_player_equipbaoku_shop` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `shopId` int(11) NOT NULL COMMENT "道具shopId",
  `dayCount` int(11) NOT NULL COMMENT "购买次数",
  `lastTime` bigint(20) NOT NULL COMMENT "最后一次购买时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by xzk 2019-03-01
alter table  `t_open_activity_xun_huan` add column `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间";  
alter table  `t_open_activity_xun_huan` add column  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间";


-- ----------------------------
-- create by zrc 2019-02-19
-- Table structure for t_player_activity_rank 玩家活动排行数据
-- ----------------------------
CREATE TABLE `t_player_activity_rank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType` int(11) NOT NULL COMMENT "活动类型",
  `rankMap` varchar(500) NOT NULL COMMENT "排行数据",
  `endTime` bigint(20) NOT NULL COMMENT "结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


