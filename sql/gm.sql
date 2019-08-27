set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';

DROP DATABASE IF EXISTS `gamegm`;
CREATE DATABASE `gamegm` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `gamegm`;


DROP TABLE IF EXISTS `t_channel`;
CREATE TABLE `t_channel` (
  `channelId` int(11) NOT NULL AUTO_INCREMENT,
  `channelName` varchar(300) DEFAULT NULL,
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`channelId`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_charge_log
-- ----------------------------
DROP TABLE IF EXISTS `t_charge_log`;
CREATE TABLE `t_charge_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `playerId` bigint(20) DEFAULT NULL COMMENT '玩家角色id',
  `playerName` varchar(100) DEFAULT NULL,
  `channelId` int(11) DEFAULT NULL COMMENT '渠道Id',
  `platformId` int(11) DEFAULT NULL COMMENT '平台Id',
  `centerPlatformId` int(11) DEFAULT NULL COMMENT '中心平台Id',
  `serverId` int(11) DEFAULT NULL COMMENT '中心服务器主键Id',
  `serverName` varchar(100) DEFAULT NULL COMMENT '服务器名',
  `gold` int(11) DEFAULT NULL COMMENT '扶持元宝',
  `chargeTime` bigint(20) DEFAULT NULL COMMENT '扶持时间',
  `userName` varchar(100) DEFAULT '' COMMENT '扶持用户',
  `reason` varchar(500) DEFAULT NULL COMMENT '扶持原因',
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_gmuser
-- ----------------------------
DROP TABLE IF EXISTS `t_gmuser`;
CREATE TABLE `t_gmuser` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userName` varchar(50) DEFAULT NULL,
  `psd` varchar(200) DEFAULT NULL,
  `avator` varchar(300) DEFAULT NULL,
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  `privilege_level` int(11) NOT NULL DEFAULT '0' COMMENT '权健级别.  999 最高 0 最低',
  `channelId` int(11) DEFAULT NULL,
  `platformId` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_mail_apply
-- ----------------------------
DROP TABLE IF EXISTS `t_mail_apply`;
CREATE TABLE `t_mail_apply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mailType` int(11) DEFAULT NULL COMMENT '类型，1个人邮件，2全服邮件',
  `serverId` int(11) DEFAULT NULL COMMENT '服务器id，来自中心服中的服务器表的主键Id',
  `title` varchar(255) DEFAULT NULL COMMENT '邮件标题',
  `content` varchar(1000) DEFAULT NULL COMMENT '邮件内容',
  `playerlist` text COMMENT '玩家列表，以英文逗号隔开',
  `proplist` text COMMENT '道具列表',
  `bindFlag` int(11) DEFAULT 0 COMMENT "是否绑定，0否1是";
  `freezTime` int(11) DEFAULT NULL COMMENT '冻结时间，单位毫秒',
  `effectDays` int(11) DEFAULT NULL COMMENT '邮件有效天数',
  `roleStartTime` bigint(20) DEFAULT NULL COMMENT '角色创建开始时间',
  `roleEndTime` bigint(20) DEFAULT NULL COMMENT '角色创建结束时间',
  `minLevel` int(11) DEFAULT NULL COMMENT '最小等级限制',
  `maxLevel` int(255) DEFAULT NULL COMMENT '最大等级限制',
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  `mailUser` bigint(20) DEFAULT NULL COMMENT '邮件申请人',
  `mailTime` bigint(20) DEFAULT NULL COMMENT '邮件申请时间',
  `mailState` int(11) DEFAULT NULL COMMENT '邮件状态,1申请中,2审批通过，3审批不通过',
  `approveUser` bigint(20) DEFAULT NULL COMMENT '审批人',
  `approveTime` bigint(20) DEFAULT NULL COMMENT '审批时间',
  `approveReason` varchar(1000) DEFAULT NULL COMMENT '审批理由',
  `sendFlag` int(11) DEFAULT NULL COMMENT '发送状态,0未发送，1已发送',
  `sdkType` int(11) DEFAULT NULL,
  `centerPlatformId` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_notice
-- ----------------------------
DROP TABLE IF EXISTS `t_notice`;
CREATE TABLE `t_notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `channelId` int(11) DEFAULT NULL COMMENT '渠道id',
  `platformId` int(11) DEFAULT NULL COMMENT 'gm平台Id',
  `serverId` int(11) DEFAULT NULL COMMENT '中心服server表中的主键Id',
  `content` text COMMENT '内容',
  `beginTime` bigint(20) DEFAULT NULL COMMENT '开始时间',
  `endTime` bigint(20) DEFAULT NULL COMMENT '结束时间',
  `intervalTime` bigint(20) DEFAULT NULL,
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  `successFlag` int(11) DEFAULT NULL,
  `errorMsg` text,
  `serverName` varchar(100) DEFAULT NULL,
  `centerPlatformId` int(11) DEFAULT NULL,
  `userName` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_t_notice_time` (`createTime`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_platform
-- ----------------------------
DROP TABLE IF EXISTS `t_platform`;
CREATE TABLE `t_platform` (
  `platformId` int(11) NOT NULL AUTO_INCREMENT,
  `platformName` varchar(300) DEFAULT NULL,
  `channelId` int(11) DEFAULT NULL,
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  `centerPlatformId` bigint(20) DEFAULT NULL,
  `sdkType` int(11) DEFAULT NULL,
  `singKey` varchar(100) DEFAULT "",
  PRIMARY KEY (`platformId`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_server_support_pool
-- ----------------------------
DROP TABLE IF EXISTS `t_server_support_pool`;
CREATE TABLE `t_server_support_pool` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `serverId` int(11) DEFAULT NULL COMMENT '服务器Id，来自中心服的服务器主键id',
  `beginGold` int(11) DEFAULT NULL COMMENT '元宝数量',
  `curGold` int(11) DEFAULT NULL,
  `delGold` int(11) DEFAULT NULL COMMENT '总消耗元宝',
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  `sdkType` int(11) DEFAULT NULL,
  `centerPlatformId` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_user_sensitive
-- ----------------------------
DROP TABLE IF EXISTS `t_user_sensitive`;
CREATE TABLE `t_user_sensitive` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL,
  `content` text,
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

INSERT INTO t_gmuser(userName,psd,avator,privilege_level)values('super_admin','123456','',999);


DROP TABLE IF EXISTS `t_player_stats`;
CREATE TABLE `t_player_stats`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `beginTime` bigint(20) NULL DEFAULT NULL COMMENT '日期的开始时间',
  `statType` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '统计类目',
  `statCount` int(100) NULL DEFAULT NULL COMMENT '统计数量',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `ix_t_player_stats`(`beginTime`, `statType`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 154 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `t_server_online`;
CREATE TABLE `t_server_online`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` int(11) NULL DEFAULT NULL COMMENT '中心平台Id',
  `serverId` int(11) NULL DEFAULT NULL COMMENT '中心服务器序号Id',
  `playerId` bigint(20) NULL DEFAULT NULL COMMENT '玩家Id',
  `onLineIndex` int(11) NULL DEFAULT NULL COMMENT '第几天登陆的服务器，0表示注册的时候',
  `onLineTime` bigint(20) NULL DEFAULT NULL COMMENT '在线时间',
  `onLineDate` bigint(20) NULL DEFAULT NULL COMMENT '在线日期',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `ix_t_server_online_player`(`playerId`, `onLineIndex`) USING BTREE,
  INDEX `ix_t_server_online_time`(`onLineDate`, `playerId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 499 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `t_marry_set_log`;
CREATE TABLE `t_marry_set_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` bigint(20) NULL DEFAULT NULL COMMENT '中心平台',
  `serverId` int(11) NULL DEFAULT NULL COMMENT '中心服务器Id',
  `successFlag` int(11) NULL DEFAULT NULL COMMENT '成功失败,1成功，0失败',
  `failMsg` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '异常原因',
  `kindType` int(255) NULL DEFAULT NULL COMMENT '发送类型，1现实版，2廉价版',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- add by cjy，
DROP TABLE IF EXISTS `t_platform_supportpool_set`;
CREATE TABLE `t_platform_supportpool_set`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `centerPlatformId` bigint(20) NULL DEFAULT NULL COMMENT '中心平台ID',
  `supportGold` int(11) NULL DEFAULT NULL COMMENT '扶持初始化金额',
  `supportRate` int(11) NULL DEFAULT NULL COMMENT '扶持比例。',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
