 set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


 -- create by xubin 2019-07-15
 alter table `t_player_goldequip_setting` add column  `fenJieZhuanShu` int(11) DEFAULT 0 COMMENT "分解转数";

 -- create by xubin 2019-07-29
-- ----------------------------
-- Table structure for t_player_new_first_charge_record 玩家新首充记录 
-- ----------------------------
CREATE TABLE `t_player_new_first_charge_record` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id", 
    `record`  varchar(1024)  NOT NULL COMMENT "领取记录",
    `startTime` bigint(20) NOT NULL COMMENT "开始时间",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`), 
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by xubin 2019-07-29
-- ----------------------------
-- Table structure for t_new_first_charge 新首充活动信息
-- ----------------------------
CREATE TABLE `t_new_first_charge` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `serverId` bigint(11) NOT NULL COMMENT "服务器id",
    `startTime` bigint(20) NOT NULL COMMENT "开始时间",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 
-- ----------------------------
-- Table structure for t_player_boss_relive boss复活
-- ----------------------------
CREATE TABLE `t_player_boss_relive` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `bossType` int(11) NOT NULL COMMENT "boss类型",
    `reliveTime` int(11) NOT NULL COMMENT "复活次数",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by jzy 2019-08-02
-- ----------------------------
-- Table structure for t_player_wushuang_settings 无双神器配置
-- ----------------------------
CREATE TABLE `t_player_wushuang_settings` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `itemId` int(11) NOT NULL COMMENT "物品id",
    `level` int(11) NOT NULL COMMENT "等级",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xzk 2019-08-05
alter table `t_player_arenapvp` add column `ticketFlag` int(11) NOT NULL COMMENT "是否购买门票";


-- ----------------------------
-- Table structure for t_new_first_charge_log 新首充活动记录信息
-- ----------------------------
CREATE TABLE `t_new_first_charge_log` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `serverId` bigint(11) NOT NULL COMMENT "服务器id",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;