set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

-- create by ylz 2019-02-18
alter table `t_player_skill` add column `tianFuInfo` varchar(256) DEFAULT "{}" COMMENT "天赋信息";

-- create by cjb 2019-02-18
alter table `t_player_addition_sys_level` add column `lingLevel` int(11) NOT NULL COMMENT "化灵等级";
alter table `t_player_addition_sys_level` add column `lingNum` int(11) NOT NULL COMMENT "化灵次数";
alter table `t_player_addition_sys_level` add column `lingPro` int(11) NOT NULL COMMENT "化灵进度";

-- create by xzk 2019-02-21
alter table `t_alliance_member` add column  `vip` int(11) NOT NULL COMMENT "vip等级";
alter table `t_player_cache` add column `isHuiYuan` int(11) NOT NULL COMMENT "是否永久至尊会员"; 

-- create by xzk 2019-02-25
-- Table structure for t_player_friend_add_rew  玩家赞赏 
-- ----------------------------
CREATE TABLE `t_player_friend_admire` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `friId` bigint(20) NOT NULL COMMENT "好友id",
  `admireTimes` int(11) NOT NULL COMMENT "赞赏次数",  
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
