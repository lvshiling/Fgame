set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


 -- create by xb 2019-07-4
-- ----------------------------
-- Table structure for t_player_charm_add_log  魅力增加日志
-- ----------------------------
 CREATE TABLE `t_player_charm_add_log` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `sendId` bigint(20) NOT NULL COMMENT "赠送者id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `charm` int(11) NOT NULL COMMENT "魅力值",
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
   PRIMARY KEY (`id`),
   KEY(`playerId`), 
     INDEX playerIdIndex (`playerId`)  
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

