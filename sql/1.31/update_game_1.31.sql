 set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


 -- create by jzy 2019-07-15
 -- Table structure for t_player_wushuangweapon_slot 无双神器位置
 -- ----------------------------
CREATE TABLE `t_player_wushuangweapon_slot` ( 
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id", 
   `slotId` int(11) NOT NULL COMMENT "部位ID",
   `itemId` int(11) NOT NULL COMMENT "物品ID",
   `level` int(11) NOT NULL COMMENT "等级",
   `experience` bigint(20) NOT NULL COMMENT "经验",
   `bind` int(11) NOT NULL COMMENT "绑定类型",
   `isActive` int(11) NOT NULL COMMENT "是否激活",
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",  
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
    PRIMARY KEY (`id`),  
   KEY(`playerId`), 
      INDEX playerIdIndex (`playerId`) 
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4; 



-- create by xubin 2019-07-15
-- ----------------------------
-- Table structure for t_player_jieyi 玩家结义数据
-- ----------------------------
CREATE TABLE `t_player_jieyi` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `name` varchar(100) DEFAULT "" COMMENT "结义威名",
    `rank` int(11) NOT NULL COMMENT "结义排行",  
    `jieYiId` bigint(20) NOT NULL COMMENT "结义id",
    `jieYiDaoJu` int(11) NOT NULL COMMENT "结义道具类型",
    `tokenType` int(11) NOT NULL COMMENT "信物类型",
    `tokenLev` int(11) NOT NULL COMMENT "信物等级",
    `tokenPro` int(11) NOT NULL COMMENT "信物升级进度值",
    `tokenNum` int(11) NOT NULL COMMENT "信物升级次数",
    `shengWeiZhi` int(11) NOT NULL COMMENT "玩家个人声威值",
    `nameLev` int(11) NOT NULL COMMENT "威名等级",
    `namePro` int(11) NOT NULL COMMENT "威名升级进度值",
    `nameNum` int(11) NOT NULL COMMENT "威名升级次数",
    `lastQiuYuanTime` bigint(20) NOT NULL COMMENT "上一次求援时间",
    `lastDropTime` bigint(20) NOT NULL COMMENT "上一次掉落声威值时间",
    `lastInviteTime` bigint(20) NOT NULL COMMENT "上一次邀请时间",
    `lastPostTime` bigint(20) NOT NULL COMMENT "上一次发布留言时间",
     `lastLeaveTime` bigint(20) NOT NULL COMMENT "上一次离开结义时间",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`),
        INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-07-15
-- ----------------------------
-- Table structure for t_jieyi 结义数据
-- ----------------------------
CREATE TABLE `t_jieyi` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
    `serverId` bigint(11) NOT NULL COMMENT "服务器id",
    `originServerId` bigint(11) NOT NULL COMMENT "起始服务器id",
    `name` varchar(100) DEFAULT "" COMMENT "结义威名",
    `memberNum` int(11) NOT NULL COMMENT "结义成员数量(弃用)",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-07-15
-- ----------------------------
-- Table structure for t_jieyi_member 结义成员数据
-- ----------------------------
   CREATE TABLE `t_jieyi_member` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
     `serverId` bigint(11) NOT NULL COMMENT "服务器id",
     `jieYiId` bigint(20) NOT NULL COMMENT "结义id",
     `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `name` varchar(50) DEFAULT "" COMMENT "玩家名字",
     `level` int(11) NOT NULL COMMENT "玩家等级",
     `role` int(11) NOT NULL COMMENT "玩家角色",
     `sex` int(11) NOT NULL COMMENT "玩家性别",
     `zhuanSheng` int(11) NOT NULL COMMENT "玩家转生级别",
     `force` bigint(20) NOT NULL COMMENT "玩家战力",
     `tokenType` int(11) NOT NULL COMMENT "信物类型",
     `tokenLev` int(11) NOT NULL COMMENT "信物等级",
     `tokenPro` int(11) NOT NULL COMMENT "信物升级进度值",
     `tokenNum` int(11) NOT NULL COMMENT "信物升级次数",
     `jieYiDaoJu` int(11) NOT NULL COMMENT "结义道具类型",
     `shengWeiZhi` int(11) NOT NULL COMMENT "玩家个人声威值",
     `nameLev` int(11) NOT NULL COMMENT "威名等级",
     `jieYiTime` bigint(20) NOT NULL COMMENT "加入结义时间",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`),
     KEY(`playerId`),
       INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-07-15
-- ----------------------------
-- Table structure for t_jieyi_leave_word 结义留言数据
-- ----------------------------
CREATE TABLE `t_jieyi_leave_word` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `serverId` bigint(11) NOT NULL COMMENT "服务器id",
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `name` varchar(50) DEFAULT "" COMMENT "玩家名字",
    `level` int(11) NOT NULL COMMENT "玩家等级",
    `role` int(11) NOT NULL COMMENT "玩家角色",
    `sex` int(11) NOT NULL COMMENT "玩家性别",
    `force` bigint(20) NOT NULL COMMENT "玩家战力",
    `leaveWord` varchar(50) DEFAULT "" COMMENT "结义留言",
    `lastPostTime` bigint(20) NOT NULL COMMENT "上一次发布留言时间",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


     -- ----------------------------
  -- Table structure for t_jieyi_invite 结义邀请数据
  -- ----------------------------
   CREATE TABLE `t_jieyi_invite` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
     `serverId` bigint(11) NOT NULL COMMENT "服务器id",
     `state` int(11) NOT NULL COMMENT "邀请状态",
     `daoJu` int(11) NOT NULL COMMENT "道具类型",
     `inviteDaoJu` int(11) NOT NULL COMMENT "邀请人道具类型",
     `inviteToken` int(11) NOT NULL COMMENT "邀请人信物类型",
     `inviteTokenLev` int(11) NOT NULL COMMENT "邀请人信物等级",
     `nameLev` int(11) NOT NULL COMMENT "威名等级",
     `inviteId` bigint(20) NOT NULL COMMENT "邀请人",
     `inviteeId` bigint(20) NOT NULL COMMENT "被邀请人",
     `name` varchar(50) DEFAULT "" COMMENT "结义名字",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`)
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



 
-- create by jzy 2019-07-24
alter Table `t_player_cache` add column `wushuangListInfo` text(2000) NOT NULL COMMENT "无双神器信息";
UPDATE `t_player_cache` SET `wushuangListInfo`= "[]";

alter Table `t_player_gold_equip_slot` add column `castingSpiritInfo` varchar(500) NOT NULL COMMENT "神铸铸灵信息";
UPDATE `t_player_gold_equip_slot` SET `castingSpiritInfo`= "{}";
alter Table `t_player_gold_equip_slot` add column `forgeSoulInfo` varchar(500) NOT NULL COMMENT "神铸锻魂信息";
UPDATE `t_player_gold_equip_slot` SET `forgeSoulInfo`= "{}";

-- ----------------------------
-- Table structure for t_dingshi_boss 定时boss
-- ----------------------------
CREATE TABLE `t_dingshi_boss` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `serverId` int(11) NOT NULL COMMENT "服务器id",
    `bossId` int(11) NOT NULL COMMENT "boss",
    `mapId` int(11) NOT NULL COMMENT "地图id",
    `lastKillTime` bigint(20) NOT NULL COMMENT "上次击杀时间",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_zhenxi_boss 珍稀boss数据
-- ----------------------------
CREATE TABLE `t_player_zhenxi_boss` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `reliveTime` int(11) NOT NULL COMMENT "复活次数",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- -- create by xzk 2019-07-23
alter Table `t_register_setting` add column `auto` int(11) NOT NULL COMMENT "自动关闭过";

ALTER TABLE `t_player_cache` ADD INDEX `nameIndex` (`serverId`, `name`);


 