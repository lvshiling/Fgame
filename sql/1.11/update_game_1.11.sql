set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 

-- ----------------------------
-- create by cjb 2019-02-11
-- Table structure for t_hongbao 红包数据
-- ----------------------------
CREATE TABLE `t_hongbao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `hongBaoType` int(11) NOT NULL COMMENT "红包类型",
  `sendId` bigint(20) DEFAULT 0 COMMENT "发红包玩家id",
  `awardList` text(2000) NOT NULL COMMENT "红包奖励列表",
  `snatchLog` text(2000) NOT NULL COMMENT "红包领取记录",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by cjb 2019-02-11
-- Table structure for t_player_hongbao 玩家红包数据
-- ----------------------------
CREATE TABLE `t_player_hongbao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `snatchCount` int(11) NOT NULL COMMENT "抢红包次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by cjb 2019-02-14
-- Table structure for t_player_chat 玩家聊天数据
-- ----------------------------
CREATE TABLE `t_player_chat` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chatCount` int(11) NOT NULL COMMENT "聊天次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by zrc 2019-02-19
-- Table structure for t_player_activity_pk 玩家活动pk数据
-- ----------------------------
CREATE TABLE `t_player_activity_pk` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType` int(11) NOT NULL COMMENT "活动类型",
  `killedNum` int(11) NOT NULL COMMENT "被杀数",
  `lastKilledTime` bigint(20) NOT NULL COMMENT "上次被杀时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
