set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

-- ----------------------------
-- ylz create by 2018-11-14
-- Table structure for t_player_system_skill 玩家系统技能
-- ----------------------------
CREATE TABLE `t_player_system_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "系统技能类型",
  `subType` int(11) NOT NULL COMMENT "技能类型",
  `level` int(11) NOT NULL COMMENT "等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- xzk create by 2018-11-17
-- Table structure for t_player_activity_num_record  玩家活动抽奖次数
-- ----------------------------
CREATE TABLE `t_player_activity_num_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `times` int(11) NOT NULL COMMENT "次数",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime`   bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- ylz create by 2018-11-19
-- Table structure for t_player_foe  玩家仇人列表
-- ----------------------------
CREATE TABLE `t_player_foe` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `foeId` bigint(20) NOT NULL COMMENT "仇人id",
  `killTime`   bigint(20) DEFAULT 0 COMMENT "击杀时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk create by 2018-11-21
-- Table structure for t_open_activity_rewards_limit 活动奖励次数限制数据
-- ----------------------------
CREATE TABLE `t_open_activity_rewards_limit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",
  `timesMap` varchar(512) NOT NULL COMMENT "领奖次数map",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk create by 2018-11-21
-- Table structure for t_open_activity_discount_limit 折扣商店次数限制数据
-- ----------------------------
CREATE TABLE `t_open_activity_discount_limit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",
  `discountDay` int(11) NOT NULL COMMENT "折扣日",
  `timesMap` varchar(1024) NOT NULL COMMENT "购买次数map",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- zrc create by 2018-11-22
alter table `t_player_xianfu` add column `group` int(11) DEFAULT 0 COMMENT "当前波数";

-- xzk create by 2018-11-23
alter table `t_open_activity_discount_limit` add column `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间";
alter table `t_open_activity_discount_limit` add column `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间";
alter table `t_open_activity_rewards_limit` add column `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间";
alter table `t_open_activity_rewards_limit` add column `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间";


-- xzk create by 2018-11-26
alter table `t_player_alliance` add column `lastMemberCallTime` bigint(20) DEFAULT 0 COMMENT "上次仙盟召集时间";

-- xzk create by 2018-11-27
alter table `t_alliance` add column `lastTransportRefreshTime` bigint(20) NOT NULL COMMENT "上次押镖次数刷新时间";