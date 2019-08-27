set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 
 

-- create by xzk 2019-04-16
alter table `t_player` add column `systemCompensate` int(11) NOT NULL COMMENT "是否领取进阶系统补偿奖励";

-- create by zrc 2019-04-19
alter table `t_trade_order` add column `playerName` varchar(100) NOT NULL COMMENT "买家名字";
-- create by zrc 2019-04-19
alter table `t_trade_item` add column `originServerId` int(11) NOT NULL COMMENT "初始服务器id";
-- create by zrc 2019-04-19
alter table `t_trade_order` add column `buyServerId` int(11) NOT NULL COMMENT "购买服务器id";

-- create by xzk 2019-04-19
alter table `t_player_qiyu` add column `isHadNotice` int(11) NOT NULL COMMENT "是否前置提示过 0否1是";
alter table `t_player_cache` add column `pregnantInfo` text(2000) NOT NULL COMMENT "怀孕信息";
UPDATE `t_player_cache` SET `pregnantInfo`= "{}"; 


-- ----------------------------
-- Table structure for t_player_house  玩家房子
-- ----------------------------
CREATE TABLE `t_player_house` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `houseIndex` int(11) NOT NULL COMMENT "房子序号",
  `houseType` int(11) NOT NULL COMMENT "房子类型", 
  `level` int(11) NOT NULL COMMENT "当前等级",
  `maxLevel` int(11) NOT NULL COMMENT "历史最高等级",
  `dayTimes` int(11) NOT NULL COMMENT "每日维修次数",
  `isBroken` int(11) NOT NULL COMMENT "是否损坏",
  `lastBrokenTime` bigint(20) NOT NULL COMMENT "上次损坏时间",
  `isRent`  int(11) NOT NULL COMMENT "是否领取租金",
  `rentUpdateTime` bigint(20) NOT NULL COMMENT "租金更新时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4; 
