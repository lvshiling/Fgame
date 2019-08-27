set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 

-- ----------------------------
--  create by zrc 2019-04-08
-- Table structure for t_player_trade_log 交易日志
-- ----------------------------
CREATE TABLE `t_player_trade_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `logType` int(11) NOT NULL COMMENT "日志类型0:出售1:购买",
  `tradeId` bigint(20) NOT NULL COMMENT "交易id",
  `sellServerId` int(11) NOT NULL COMMENT "出售服务器id",
  `sellPlayerId` bigint(20) NOT NULL COMMENT "出售玩家id",
  `sellPlayerName` varchar(100) NOT NULL COMMENT "出售玩家名字",
  `buyServerId` int(11) NOT NULL COMMENT "购买服务器id",
  `buyPlayerId` bigint(20) NOT NULL COMMENT "购买玩家id",
  `buyPlayerName` varchar(100) NOT NULL COMMENT "购买玩家名字",
  `getGold` int(11) NOT NULL COMMENT "获得的元宝",
  `gold` int(11) NOT NULL COMMENT "价格",
  `fee` int(11) NOT NULL COMMENT "手续费",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量",
  `propertyData` varchar(512) DEFAULT "{}"  COMMENT "属性",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by zrc 2019-04-08
alter table `t_trade_order` modify column `sellPlayerId` bigint(20);
alter table `t_trade_order` modify column `sellPlayerName` varchar(100);
alter table `t_trade_item` add column `system` int(11) COMMENT "系统回购";

-- create by xzk 2019-04-09
-- ----------------------------
-- Table structure for t_player_shenyu  神域之战
-- ----------------------------
CREATE TABLE `t_player_shenyu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `keyNum` int(11) NOT NULL COMMENT "钥匙数",
  `exp` bigint(20) NOT NULL COMMENT "获得经验",
  `itemInfo` varchar(1024) NOT NULL DEFAULT "{}" COMMENT "获得物品",
  `endTime` bigint(20) NOT NULL COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table `t_trade_order` add column `level` int(11) COMMENT "等级";
alter table `t_trade_item` add column `level` int(11) COMMENT "等级";
alter table `t_player_trade_log` add column `level` int(11) COMMENT "等级";
-- ----------------------------
-- Table structure for t_player_xian_tao  仙桃大会
-- ----------------------------
CREATE TABLE `t_player_xian_tao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `juniorPeachCount` int(11) NOT NULL COMMENT "百年仙桃数量",
  `highPeachCount` int(11) NOT NULL COMMENT "千年仙桃数量",
  `robCount` int(11) NOT NULL COMMENT "劫取次数",
  `beRobCount` int(11) NOT NULL COMMENT "被劫取次数",
  `endTime` bigint(20) NOT NULL COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by xzk 2019-04-16
alter table `t_player_found_back` add column `resLevel` int(11) NOT NULL COMMENT "资源等级";
