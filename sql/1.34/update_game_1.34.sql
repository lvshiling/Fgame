 set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 
-- create by xzk 2019-08-13
-- ----------------------------
-- Table structure for t_arena_rank 3v3排行榜
-- ----------------------------
CREATE TABLE `t_arena_rank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `playerId` bigint(20) DEFAULT 0 COMMENT "玩家id",
  `playerName` varchar(100) DEFAULT "" COMMENT "玩家名字",
  `curWinCount` int(11) DEFAULT 0 COMMENT "本周连胜",
  `winCount` int(11) DEFAULT 0 COMMENT "本周最高连胜", 
  `lastWinCount` int(11) DEFAULT 0 COMMENT "上周连胜",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "最后操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_arena_rank_time 3v3排行榜
-- ----------------------------
CREATE TABLE `t_arena_rank_time` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "上周时间戳",
  `thisTime` bigint(11) DEFAULT 0 COMMENT "本周时间戳",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- create by jzy 2019-08-20
-- ----------------------------
-- Table structure for t_player_shangguzhiling 玩家上古之灵数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shangguzhiling`;
CREATE TABLE `t_player_shangguzhiling` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingShouType` int(11) NOT NULL COMMENT "灵兽类型",
  `level` int(11) NOT NULL COMMENT "等级",
  `experience` bigint(20) NOT NULL COMMENT "经验",
  `lingwen` varchar(512) DEFAULT "{}" COMMENT "灵纹",
  `uprankLevel` int(11) NOT NULL COMMENT "阶级",
  `uprankBless` bigint(20) NOT NULL COMMENT "祝福值",
  `uprankTimes` int(11) NOT NULL COMMENT "本阶尝试进阶次数",
  `linglian` varchar(512) DEFAULT "{}" COMMENT "灵炼",
  `linglianTimes` int(11) NOT NULL COMMENT "灵炼次数",
  `receiveTime` bigint(20) DEFAULT 0 COMMENT "上一次领取时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-08-21
CREATE TABLE `t_player_ring` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) DEFAULT 0 COMMENT "玩家id",
  `typ` bigint(11) DEFAULT 0 COMMENT "特戒类型",
  `bindType` bigint(11) DEFAULT 0 COMMENT "绑定类型",
  `itemId` bigint(11) DEFAULT 0 COMMENT "特戒物品id",
  `propertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`), 
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-08-21
CREATE TABLE `t_player_ring_baoku` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "宝库类型",
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

 -- create by xubin 2019-08-21
alter table `t_player_cache` add column `ringInfo` text(2000) NOT NULL COMMENT "特戒";
UPDATE `t_player_cache` SET `ringInfo`= "[]"; 