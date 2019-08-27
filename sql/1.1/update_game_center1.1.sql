set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game_center`;

-- ZRC 2018-11-08
alter table t_user add column `gm` int(11) DEFAULT 0 COMMENT "gm";


-- ZRC 2018-11-13
-- 合并记录
CREATE TABLE `t_merge_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `fromServerId` int(11) DEFAULT 0 COMMENT "源服务器id",
  `toServerId` int(11) DEFAULT 0 COMMENT "目的服务器id", 
  `finalServerId` int(11) DEFAULT 0 COMMENT "最终服务器id",
  `mergeTime` bigint(20) DEFAULT 0 COMMENT "合服时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`fromServerId`),
  INDEX `fromServerIdIndex`(`fromServerId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ZRC 2018-11-14
-- 兑换记录
CREATE TABLE `t_redeem_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT "id",
  `redeemId` int(11) DEFAULT 0 COMMENT "兑换id",
  `redeemCode` varchar(256) DEFAULT "" COMMENT "兑换码",
  `platformId` int(11) DEFAULT 0 COMMENT "中心平台id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `sdkType` int(11) DEFAULT 0 COMMENT "sdk类型",
  `platformUserId` varchar(256) DEFAULT "" COMMENT "sdk用户id",
  `userId` bigint(20) DEFAULT 0 COMMENT "玩家id",
  `playerId` bigint(20) DEFAULT 0 COMMENT "角色id",
  `playerLevel` int(11) DEFAULT 0 COMMENT "玩家等级",
  `playerVipLevel` int(11) DEFAULT 0 COMMENT "玩家vip等级",
  `playerName` varchar(256) DEFAULT "" COMMENT "玩家名字",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`redeemCode`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;