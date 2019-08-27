set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `gamegm`;

-- CJY 2018-01-01
alter table t_mail_apply add COLUMN sdkType int(11) DEFAULT 0 COMMENT "sdk类型";
alter table t_mail_apply add COLUMN centerPlatformId int(11) DEFAULT 0 COMMENT "中心平台id";

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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

alter table t_server_support_pool add COLUMN sdkType int(11) DEFAULT 2 COMMENT "Sdk类型";
alter table t_server_support_pool add COLUMN centerPlatformId int(11) DEFAULT 1 COMMENT "中心平台id";

-- 2018-11-02
alter table t_notice add COLUMN successFlag int(11) DEFAULT 0 COMMENT "成功标志";
alter table t_notice add COLUMN errorMsg TEXT DEFAULT "" COMMENT "错误消息";
alter table t_notice add COLUMN serverName varchar(100) DEFAULT "" COMMENT "服务器名字";
alter table t_notice add COLUMN centerPlatformId int(11) DEFAULT 1 COMMENT "中心平台id";

CREATE INDEX ix_t_notice_time on t_notice(createTime);

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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

alter table t_notice add COLUMN userName VARCHAR(100) DEFAULT "" COMMENT "用户名字";


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

alter table t_mail_apply add COLUMN bindFlag int(11) DEFAULT 0 COMMENT "是否绑定，0否1是";


alter table t_server_support_pool add COLUMN orderGoldPer int(11) DEFAULT 0 COMMENT "订单元宝的百分比作为添加在扶持元宝上，单位%";
alter table t_server_support_pool add COLUMN orderGold int(11) DEFAULT 0 COMMENT "已经添加了订单的元宝";
alter table t_server_support_pool add COLUMN beginOrderTime bigint(20) DEFAULT 0 COMMENT "开始的统计的订单时间";
alter table t_server_support_pool add COLUMN curOrderTime bigint(20) DEFAULT 0 COMMENT "当前统计到的时间";

CREATE TABLE `t_server_daily_stats`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `serverId` int(11) NULL DEFAULT NULL COMMENT '服务器序号',
  `serverType` int(11) NULL DEFAULT NULL COMMENT '服务器类型',
  `platformId` int(11) NULL DEFAULT NULL COMMENT '中心平台Id',
  `curDate` bigint(20) NULL DEFAULT NULL COMMENT '当日时间戳',
  `maxOnlineNum` int(11) NULL DEFAULT NULL COMMENT '最高在线人数',
  `loginNum` int(11) NULL DEFAULT NULL COMMENT '登录人数',
  `orderPlayerNum` int(11) NULL DEFAULT NULL COMMENT '订单人数',
  `orderNum` int(11) NULL DEFAULT NULL COMMENT '订单数',
  `orderMoney` int(255) NULL DEFAULT NULL COMMENT '订单金额',
  `orderGold` int(255) NULL DEFAULT NULL COMMENT '订单元宝',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `ix_t_server_daily_stats_time`(`curDate`, `serverId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

alter table t_platform add COLUMN signKey varchar(100) DEFAULT '' COMMENT "签名key";


alter table t_mail_apply add COLUMN remark varchar(1000) DEFAULT '' COMMENT "备注";


-- add by cjy，
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
