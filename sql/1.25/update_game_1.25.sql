set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


-- create by cjb 2019-06-17
alter table `t_player_lingtong_info` add column `starLevel` int(11) NOT NULL COMMENT "升星等级";
alter table `t_player_lingtong_info` add column `starNum` int(11) NOT NULL COMMENT "升星次数";
alter table `t_player_lingtong_info` add column `starPro` int(11) NOT NULL COMMENT "升星进度值";

-- create by xb 2019-06-17
alter table `t_player_title` add column `starLev` int(11) NOT NULL COMMENT "升星等级";
alter table `t_player_title` add column `starNum` int(11) NOT NULL COMMENT "升星次数";
alter table `t_player_title` add column `starBless` int(11) NOT NULL COMMENT "升星祝福值";

-- create by xzk 2019-06-18
-- ----------------------------
-- Table structure for t_player_qixue 玩家泣血枪数据
-- ----------------------------
CREATE TABLE `t_player_qixue` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `currLevel` int(11) NOT NULL COMMENT "当前阶数",
  `currStar` int(11) NOT NULL COMMENT "当前星数",
  `lastTime` bigint(20) NOT NULL COMMENT "上次被击杀掉落杀气时间",
  `shaLuNum` bigint(20) NOT NULL COMMENT "杀戮心数量",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数升级次数",
  `power` bigint(20) NOT NULL COMMENT "戮仙刃战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by cjb 2019-06-18
alter table `t_player_gold_equip_slot` add column `newStLevel` int(11) NOT NULL COMMENT "新强化等级";
alter table `t_player_goldequip_setting` add column `isCheckOldSt` int(11) NOT NULL COMMENT "是否检查过老强化等级";


-- create by xzk 2019-06-19
alter table `t_player_arena` add column `jiFenCount` int(11) NOT NULL COMMENT "累计积分";
alter table `t_player_arena` add column `jiFenDay` int(11) NOT NULL COMMENT "每日积分";
alter table `t_player_arena` add column `arenaTime` bigint(20) NOT NULL COMMENT "积分更新时间";
alter table `t_player_arena` add column `winCount` int(11) NOT NULL COMMENT "连胜次数";
alter table `t_player_arena` add column `failCount` int(11) NOT NULL COMMENT "连败次数";
alter table `t_player_arena` add column `rankRewTime` bigint(20) NOT NULL COMMENT "上次周榜奖励时间";
alter table `t_player_arena` add column `dayMaxWinCount` int(11) NOT NULL COMMENT "当天最高连胜";
alter table `t_player_arena` add column `dayWinCount` int(11) NOT NULL COMMENT "当天连胜";


-- create by xb 2019-06-24
CREATE TABLE `t_chuangshi_yugao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` bigint(20) NOT NULL COMMENT "服务器索引",
  `num` int(11) NOT NULL COMMENT "人数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by xb 2019-06-24
CREATE TABLE `t_player_chuangshi_yugao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `isJoin` int(11) NOT NULL COMMENT "是否参加",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


 