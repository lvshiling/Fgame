set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


 -- create by  2019-06-28
alter table `t_alliance` add column `lastMergeTime` bigint(20) NOT NULL COMMENT "合帮时间";

-- ----------------------------
-- Table structure for t_player_mingge_buchang 玩家命格补偿数据
-- ----------------------------
CREATE TABLE `t_player_mingge_buchang` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `buchang` int(11) NOT NULL COMMENT "补偿",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


