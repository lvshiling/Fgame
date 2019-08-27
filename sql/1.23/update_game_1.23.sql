set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 
-- create by zrc 2019-06-04
-- ----------------------------
-- Table structure for t_player_week 玩家周卡数据
-- ----------------------------
CREATE TABLE `t_player_week` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `seniorExpireTime` bigint(20) DEFAULT 0  COMMENT "高级周卡过期时间",
  `seniorLastDayRewTime` bigint(20) DEFAULT 0  COMMENT "高级每日奖励领取时间",
  `seniorCycDay` int(11) NOT NULL COMMENT "高级循环领取天数",
  `juniorExpireTime` bigint(20) DEFAULT 0  COMMENT "初级周卡过期时间",
  `juniorLastDayRewTime` bigint(20) DEFAULT 0  COMMENT "初级每日奖励领取时间",
  `juniorCycDay` int(11) NOT NULL COMMENT "初级循环领取天数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


