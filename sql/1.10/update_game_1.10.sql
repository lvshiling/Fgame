set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 

 -- ----------------------------
-- create by ylz 2019-01-22
-- Table structure for t_player_kaifumubiao 玩家开服目标信息
-- ----------------------------
CREATE TABLE `t_player_kaifumubiao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `kaiFuDay` int(11) NOT NULL COMMENT "开服解锁时间",
  `finishNum` int(11) NOT NULL COMMENT "完成任务数",
  `isReward` int(11) NOT NULL COMMENT "是否领取过组奖励",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-22
-- Table structure for t_player_quest_crossday 玩家任务跨天信息
-- ----------------------------
CREATE TABLE `t_player_quest_crossday` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `crossDayTime` bigint(20) DEFAULT 0 COMMENT "跨天时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
 
-- create by xzk 2019-01-22
-- ----------------------------
-- Table structure for t_player_fei_sheng  玩家飞升表
-- ----------------------------
CREATE TABLE `t_player_fei_sheng` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `feiLevel`  int(11) NOT NULL COMMENT "飞升等级",
  `addRate`  int(11) NOT NULL COMMENT "增加的成功率",
  `gongDeNum`  bigint(20) NOT NULL COMMENT "功德值",
  `leftPotential`  int(11) NOT NULL COMMENT "剩余潜能点",
  `tiZhi`  int(11) NOT NULL COMMENT "体质点",
  `liDao`  int(11) NOT NULL COMMENT "力道点",
  `jinGu`  int(11) NOT NULL COMMENT "筋骨点",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by ylz 2019-01-23
alter table `t_player_arena` add column `totalRewardTime` int(11) DEFAULT 0 COMMENT "累计获胜次数";


-- create by ylz 2019-01-23
alter table `t_player_onearena` add column `robTime` bigint(20) DEFAULT 0 COMMENT "抢夺时间";


-- create by ylz 2019-01-23
alter table `t_player_alliance` add column `totalWinTime` int(11) DEFAULT 0 COMMENT "城战胜利次数";


-- ----------------------------
-- create by ylz 2019-01-24
-- Table structure for t_player_shenmo 玩家神魔数据
-- ----------------------------
CREATE TABLE `t_player_shenmo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `gongXunNum` int(11) NOT NULL COMMENT "玩家功勋数",
  `killNum` int(11) NOT NULL COMMENT "本次击杀",
  `endTime` bigint(20) DEFAULT 0 COMMENT "结束时间",
  `rewTime` bigint(20) DEFAULT 0 COMMENT "领取奖励的排行榜时间戳",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;