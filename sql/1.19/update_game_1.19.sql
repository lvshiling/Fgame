set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

alter table `t_player_trade_log` add column `feeRate` int(11) NOT NULL COMMENT "手续费比例";

-- ----------------------------
-- Table structure for t_trade_recycle  系统回购
-- ----------------------------
CREATE TABLE `t_trade_recycle` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `recycleGold` int(11) NOT NULL COMMENT "回收的元宝",
  `recycleTime` bigint(20) NOT NULL COMMENT "回收时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_trade_recycle  个人系统回购
-- ----------------------------
CREATE TABLE `t_player_trade_recycle` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` int(11) NOT NULL COMMENT "玩家id",
  `recycleGold` int(11) NOT NULL COMMENT "回收的元宝",
  `recycleTime` bigint(20) NOT NULL COMMENT "回收时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



alter table `t_player_email` modify column `attachementInfo` text(5000) ;