set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';

DROP DATABASE IF EXISTS `game_center`;
CREATE DATABASE `game_center` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `game_center`;

-- ----------------------------
-- Table structure for t_user 用户表
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) NOT NULL COMMENT "平台",
  `platformUserId` varchar(100) NOT NULL COMMENT "平台用户id",
  `name` varchar(100) DEFAULT "" COMMENT "名字",
  `password` varchar(100) DEFAULT "" COMMENT "密码",
  `phoneNum` varchar(20) DEFAULT "" COMMENT "手机号码",
  `idCard` varchar(25) DEFAULT "" COMMENT "身份证号码",
  `realName` varchar(10) DEFAULT "" COMMENT "真实姓名",
  `realNameState` int(11) DEFAULT 0 COMMENT "认证状态",
  `gm` int(11) DEFAULT 0 COMMENT "gm状态0:普通用户1:gm用户",
  `forbid` int DEFAULT 0 COMMENT "禁号 0正常 1禁号",
  `forbidTime` bigint DEFAULT 0 COMMENT "封号时间",
  `forbidEndTime` bigint DEFAULT 0 COMMENT "封号结束时间",
  `forbidName` varchar(256) DEFAULT "" COMMENT "封号人", 
  `forbidText` varchar(256) DEFAULT "" COMMENT "禁号原因",
  `ip` varchar(256) DEFAULT "" COMMENT "ip",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
   INDEX platformUserIdIndex (`platform`, `platformUserId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_platform 
-- ----------------------------
DROP TABLE IF EXISTS `t_platform`;
CREATE TABLE `t_platform` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(50) DEFAULT '0' COMMENT '平台名字',
  `signKey` varchar(512) DEFAULT '' COMMENT '签名key',
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  `sdkType` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_server 游戏服务器
-- ----------------------------
DROP TABLE IF EXISTS `t_server`;
CREATE TABLE `t_server` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverType` int(11) DEFAULT 0 COMMENT "服务器类型0:单游戏服 1:临服游戏 2:战场游戏服 3:平台游戏服 4:跨平台游戏服", 	
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `parentServerId` int(11) DEFAULT 0 COMMENT "归属服务器,id为服务器表的serverId",
  `platform` int(11) DEFAULT 0 COMMENT "平台id", 
  `name` varchar(100) DEFAULT "" COMMENT "名字",
  `startTime` bigint(20) DEFAULT 0 COMMENT "开始时间",
  `serverIp` varchar(256) DEFAULT "" COMMENT "服务器ip",
  `serverPort` varchar(256) DEFAULT "" COMMENT "服务器端口",
  `serverRemoteIp`  varchar(256) DEFAULT "" COMMENT "远程服务器ip",
  `serverRemotePort`   varchar(256) DEFAULT "" COMMENT "远程服务器端口",
  `serverDBIp` varchar(256) DEFAULT "" COMMENT "数据库ip",
  `serverDBPort` varchar(256) DEFAULT "" COMMENT "数据库端口",
  `serverDBName` varchar(256) DEFAULT "" COMMENT "数据库名字",
  `serverDBUser` varchar(256) DEFAULT "" COMMENT "数据库用户名",
  `serverDBPassword` varchar(256) DEFAULT "" COMMENT "数据库密码",
  `serverTag` int(11) DEFAULT 0 COMMENT "服务器标签(0:无 1：新服 2：热服)",
  `serverStatus` int(11) DEFAULT 0 COMMENT "服务器状态(0:流畅 1:爆满 2:维护)",
  `preShow` int(11) DEFAULT 0 COMMENT "提前展示(0:不提前展示 1:提前展示)",
  `jiaoYiZhanQuServerId` int(11) DEFAULT 0 COMMENT '交易战区号';
  `pingTaiFuServerId` int(11) DEFAULT 0 COMMENT '全平台服Id';
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;




-- ----------------------------
-- Table structure for t_group 组
-- ----------------------------
DROP TABLE IF EXISTS `t_group`;
CREATE TABLE `t_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `groupId` int(11) DEFAULT 0 COMMENT "组",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_region 自定义战区
-- ----------------------------
DROP TABLE IF EXISTS `t_region`;
CREATE TABLE `t_region` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `regionId` int(11) DEFAULT 0 COMMENT "战区id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player 玩家
-- ----------------------------
DROP TABLE IF EXISTS `t_player`;
CREATE TABLE `t_player` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `userId` bigint(20)  NOT NULL  COMMENT "用户id",
  `serverId` int(11) NOT NULL  COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL  COMMENT "玩家id",
  `playerName` varchar(100) NOT NULL  COMMENT "玩家名字",
  `role` int(11) NOT NULL COMMENT "角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `level` int(11) NOT NULL COMMENT "等级",
  `zhuanShu` int(11) NOT NULL COMMENT "转数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`userId`),
  INDEX playerIdIndex (`userId`) USING BTREE,
   INDEX playerIndex (`playerId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_tulong_rank 跨服屠龙排行榜
-- ----------------------------
DROP TABLE IF EXISTS `t_tulong_rank`;
CREATE TABLE `t_tulong_rank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `areaId` int(11) DEFAULT 0 COMMENT "区id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `allianceId` bigint(20) DEFAULT 0 COMMENT "仙盟id",
  `allianceName` varchar(100) DEFAULT "" COMMENT "仙盟名字",
  `killNum` int(11) DEFAULT 0 COMMENT "击杀数量",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "最后操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;






-- ----------------------------
-- Table structure for t_treasurebox_log 跨服宝箱日志
-- ----------------------------
DROP TABLE IF EXISTS `t_treasurebox_log`;
CREATE TABLE `t_treasurebox_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `areaId` int(11) DEFAULT 0 COMMENT "区id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `playerName` varchar(100) DEFAULT "" COMMENT "玩家名字",
  `itemInfo` varchar(1024) DEFAULT "[]" COMMENT "获得物品",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "最后操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_chat_set 聊天限制配置
-- ----------------------------
DROP TABLE IF EXISTS `t_chat_set`;
CREATE TABLE `t_chat_set` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` int(11) DEFAULT NULL COMMENT '平台id',
  `serverId` int(11) DEFAULT NULL COMMENT '服务器序号',
  `minVip` int(11) DEFAULT NULL COMMENT '最低vip等级',
  `minPlayerlevel` int(11) DEFAULT NULL COMMENT '最低玩家等级',
  `startTime` varchar(8) DEFAULT NULL COMMENT '开始时间,10:20:30',
  `endTime` varchar(8) DEFAULT NULL COMMENT '结束时间',
  `updateTime` bigint(20) DEFAULT '0' COMMENT '更新时间',
  `createTime` bigint(20) DEFAULT '0' COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT '0' COMMENT '删除时间',
  `worldVip` int(11) DEFAULT NULL COMMENT '世界vip等级',
  `worldPlayerLevel` int(11) DEFAULT NULL COMMENT '世界玩家等级',
  `pChatVip` int(11) DEFAULT NULL COMMENT '私聊vip等级',
  `pChatPlayerLevel` int(11) DEFAULT NULL COMMENT '私聊玩家等级',
  `guildVip` int(11) DEFAULT NULL COMMENT '公会vip',
  `guildPlayerLevel` int(11) DEFAULT NULL COMMENT '公会玩家等级',
  `sdkType` int(11) DEFAULT 0 COMMENT 'SDK类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_order 订单
-- ----------------------------
DROP TABLE IF EXISTS `t_order`;
CREATE TABLE `t_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `orderId` varchar(256) DEFAULT "" COMMENT "订单id",
  `sdkOrderId` varchar(256) DEFAULT "" COMMENT "sdk订单id",
  `status` int(11) DEFAULT 0 COMMENT "订单状态(0:初始化 1:充值成功 2:发货成功 3:取消)", 
  `sdkType` int(11) DEFAULT 0 COMMENT 'sdkType',
  `devicePlatform` int(11) DEFAULT 0 COMMENT "设备平台",
  `platformUserId` varchar(256) DEFAULT "" COMMENT "平台用户id",
  `serverId` int(11) DEFAULT 0 COMMENT '服务器序号',
  `userId` bigint(20) DEFAULT 0 COMMENT '用户id',
  `playerId` bigint(20) DEFAULT 0 COMMENT '角色id',
  `playerLevel` int(11) DEFAULT 0 COMMENT "角色等级",
  `playerName` varchar(50) DEFAULT 0 COMMENT "角色名字",
  `gold` int(11) DEFAULT 0 COMMENT "元宝",
  `chargeId` int(11) DEFAULT 0 COMMENT '充值档次',
  `money` int(11) DEFAULT 0 COMMENT '钱',
  `receivePayTime` bigint(20) COMMENT "接收充值成功时间",
  `test` int(11) COMMENT "测试订单",
  `createTime` bigint(20) DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) DEFAULT 0 COMMENT '删除时间',
  `updateTime` bigint(20) DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX orderIdIndex (`orderId`) USING BTREE,
  INDEX sdkOrderIdIndex (`sdkOrderId`) USING BTREE,
  INDEX playerIdIndex (`playerId`) USING BTREE,
  INDEX createTimeIndex (`createTime`) USING BTREE,
  INDEX updateTimeIndex (`updateTime`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `t_redeem`;
CREATE TABLE `t_redeem`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `giftBagName` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '礼包名字',
  `giftBagDesc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '礼包文本',
  `giftBagContent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '礼包内容',
  `redeemNum` int(11) NULL DEFAULT NULL COMMENT '兑换码数量',
  `redeemPlayerUseNum` int(11) NULL DEFAULT NULL COMMENT '兑换码个人使用次数,0表示无限次',
  `redeemServerUseNum` int(11) NULL DEFAULT NULL COMMENT '兑换码全服使用次数,0表示无限',
  `sdkTypes` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '兑换码sdk类别，以英文逗号隔开，为空表示全渠道',
  `sendType` int(11) NULL DEFAULT NULL COMMENT '兑换发送方式,1:直接发放给玩家，2以邮件形式发放给玩家',
  `startTime` bigint(20) NULL DEFAULT NULL COMMENT '生效开始时间',
  `endTime` bigint(20) NULL DEFAULT NULL COMMENT '生效结束时间',
  `minPlayerLevel` int(11) NULL DEFAULT NULL COMMENT '生效最低等级，0表示不限制',
  `minVipLevel` int(11) NULL DEFAULT NULL COMMENT '生效最低VIP等级,0表示不限制',
  `createFlag` int(11) NULL DEFAULT NULL COMMENT '码生成标志',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '兑换码配置表' ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `t_redeem_code`;
CREATE TABLE `t_redeem_code`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '兑换码Id',
  `redeemCode` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '兑换码',
  `redeemId` int(11) NULL DEFAULT NULL COMMENT '兑换码设置id,来自表t_redeem的主键id',
  `useNum` int(11) NULL DEFAULT NULL COMMENT '实际兑换次数',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `IX_t_redeem_code_code`(`redeemCode`) USING BTREE,
  INDEX `IX_t_redeem_code_pid`(`redeemId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '礼包兑换码' ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `t_redeem_platform`;
CREATE TABLE `t_redeem_platform`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` int(11) NULL DEFAULT NULL,
  `redeemId` int(11) NULL DEFAULT NULL COMMENT '兑换码设置id,来自表t_redeem的主键id',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `IX_t_redeem_platform_redeemId`(`redeemId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '对应中心平台id' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for t_merge_record 合服记录
-- ----------------------------
DROP TABLE IF EXISTS `t_merge_record`;
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
-- ----------------------------
-- Table structure for t_redeem_record 兑换码记录
-- ----------------------------
DROP TABLE IF EXISTS `t_redeem_record`;
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
  KEY(`redeemCode`),
  INDEX `redeemCodeIndex`(`redeemCode`) USING BTREE
  INDEX `playerIdIndex`(`playerId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `t_notice_login`;
CREATE TABLE `t_notice_login`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` int(11) NULL DEFAULT NULL COMMENT '平台id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '公告内容',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登陆公告' ROW_FORMAT = Dynamic;



-- ----------------------------
-- Table structure for t_shenmo_rank 神魔战场排行榜
-- ----------------------------
DROP TABLE IF EXISTS `t_shenmo_rank`;
CREATE TABLE `t_shenmo_rank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `allianceId` bigint(20) DEFAULT 0 COMMENT "仙盟id",
  `allianceName` varchar(100) DEFAULT "" COMMENT "仙盟名字",
  `jiFenNum` int(11) DEFAULT 0 COMMENT "本周积分数量",
  `lastJiFenNum` int(11) DEFAULT 0 COMMENT "上周积分数量",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "最后操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_shenmo_rank_time 神魔战场排行榜
-- ----------------------------
DROP TABLE IF EXISTS `t_shenmo_rank_time`;
CREATE TABLE `t_shenmo_rank_time` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "上周时间戳",
  `thisTime` bigint(11) DEFAULT 0 COMMENT "本周时间戳",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- 开始 2019-01-11 中心封禁Ip表  by cjy
DROP TABLE IF EXISTS `t_ipforbid`;
CREATE TABLE `t_ipforbid`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip',
  `forbid` int(11) NULL DEFAULT 0 COMMENT '禁号 0正常 1禁号',
  `forbidTime` bigint(20) NULL DEFAULT 0 COMMENT '封号时间',
  `forbidEndTime` bigint(20) NULL DEFAULT 0 COMMENT '封号结束时间',
  `forbidName` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '封号人',
  `forbidText` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '禁号原因',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `ix_t_ipforbid_ip`(`ip`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Ip封禁表' ROW_FORMAT = Dynamic;



-- ----------------------------
-- Table structure for t_trade_item 交易行 zrc
-- ----------------------------
DROP TABLE IF EXISTS `t_trade_item`;
CREATE TABLE `t_trade_item`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `platform` int(11) NOT NULL COMMENT '平台',
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `tradeId` bigint(20) NOT NULL COMMENT "本地商品id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `playerName` varchar(100) NOT NULL COMMENT "玩家名字",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量",
  `level` int(11) NOT NULL COMMENT "等级",
  `gold` bigint(20) NOT NULL COMMENT "价格",
  `propertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `status` int(11) NOT NULL COMMENT "状态0:上架,1:售出,2:下架",
   `buyPlayerPlatform` int(11) COMMENT "购买者平台",
  `buyPlayerServerId` int(11) COMMENT "购买者服务器id",
  `buyPlayerId` bigint(20) COMMENT "购买者id",
  `buyPlayerName` varchar(500) COMMENT "购买者名字",
  `updateTime` bigint(20) NOT NULL COMMENT '更新时间',
  `createTime` bigint(20) NOT NULL COMMENT '创建时间',
  `deleteTime` bigint(20) NOT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_platform_marryprice 结婚价格 zrc
-- ----------------------------
DROP TABLE IF EXISTS `t_platform_marryprice`;
CREATE TABLE `t_platform_marryprice`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` bigint(11) NULL DEFAULT NULL COMMENT '中心平台Id',
  `kindType` int(11) NULL DEFAULT NULL COMMENT '启用类型1现实版，2廉价版',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) 
) ENGINE = InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ;


-- ----------------------------
-- Table structure for t_client_version 客户端版本 zrc
-- ----------------------------
DROP TABLE IF EXISTS `t_client_version`;
CREATE TABLE `t_client_version`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `iosVersion` varchar(50) DEFAULT NULL COMMENT 'ios版本',
  `androidVersion` varchar(50) DEFAULT NULL COMMENT '安卓版本',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`)
)  ENGINE = InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ;

-- 平台服务器ip  by zrc
DROP TABLE IF EXISTS `t_platform_server_config`;
CREATE TABLE `t_platform_server_config`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tradeServerIp` varchar(50) DEFAULT NULL COMMENT '交易服务器ip',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`)
)  ENGINE = InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ;

-- 兑换数量  by zrc
DROP TABLE IF EXISTS `t_redeem_use_num`;
CREATE TABLE `t_redeem_use_num`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '兑换码Id',
  `redeemId` int(11) NULL DEFAULT NULL COMMENT '兑换码设置id,来自表t_redeem的主键id',
  `useNum` int(11) NULL DEFAULT NULL COMMENT '实际兑换次数',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `IX_t_redeem_id`(`redeemId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '兑换数量' ROW_FORMAT = Dynamic;

-- 中心平台配置选项，by cjy 20190529
DROP TABLE IF EXISTS `t_platform_setting`;
CREATE TABLE `t_platform_setting`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` bigint(11) NULL DEFAULT NULL COMMENT '中心平台Id',
  `settingContent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '设置结果,json',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `t_platform_setting_platform`(`platformId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `t_jiaoyi_zhanqu`;
CREATE TABLE `t_jiaoyi_zhanqu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` int(11) NULL DEFAULT 0 COMMENT '平台id',
  `serverId` int(11) NULL DEFAULT NULL COMMENT '交易战区id',
  `jiaoYiName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '交易战区名',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `t_platform_chatset`;
CREATE TABLE `t_platform_chatset`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` int(11) NULL DEFAULT 0 COMMENT '平台id',
  `minVip` int(11) NULL DEFAULT NULL COMMENT '最低vip等级',
  `minPlayerlevel` int(11) NULL DEFAULT NULL COMMENT '最低玩家等级',
  `startTime` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '开始时间,10:20:30',
  `endTime` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '结束时间',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  `worldVip` int(11) NULL DEFAULT NULL COMMENT '世界vip等级',
  `worldPlayerLevel` int(11) NULL DEFAULT NULL COMMENT '世界玩家等级',
  `pChatVip` int(11) NULL DEFAULT NULL COMMENT '私聊vip等级',
  `pChatPlayerLevel` int(11) NULL DEFAULT NULL COMMENT '私聊玩家等级',
  `guildVip` int(11) NULL DEFAULT NULL COMMENT '公会vip',
  `guildPlayerLevel` int(11) NULL DEFAULT NULL COMMENT '公会玩家等级',
  `sdkType` int(11) NULL DEFAULT NULL,
  `teamVip` int(11) NULL DEFAULT 0 COMMENT '组队vip等级',
  `teamPlayerLevel` int(11) NULL DEFAULT 0 COMMENT '组队玩家等级',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;



-- ----------------------------
-- Table structure for t_arenapvp_bazhu 比武大会历届冠军
-- ----------------------------
DROP TABLE IF EXISTS `t_arenapvp_bazhu`;
CREATE TABLE `t_arenapvp_bazhu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) NOT NULL COMMENT "平台id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerPlatform` int(11) NOT NULL COMMENT "玩家平台id",
  `playerServerId` int(11) NOT NULL COMMENT "玩家服务器id",
  `raceNumber` int(11) NOT NULL COMMENT "x届",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `playerName` varchar(11) NOT NULL COMMENT "玩家名字",
  `sex` int(11) NOT NULL COMMENT "性别", 
  `role` int(11) NOT NULL COMMENT "角色",
  `fashionId` int(11) NOT NULL COMMENT "时装id", 
  `wingId` int(11) NOT NULL COMMENT "战翼id",
  `weaponId` int(11) NOT NULL COMMENT "兵魂id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_feedbackfee_exchange 兑换记录
-- ----------------------------
DROP TABLE IF EXISTS `t_feedbackfee_exchange`;
CREATE TABLE `t_feedbackfee_exchange` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) NOT NULL COMMENT "平台id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `exchangeId` bigint(20) NOT NULL COMMENT "兑换id",
  `expiredTime` bigint(20) NOT NULL COMMENT "过期时间",
  `money` int(32) NOT NULL COMMENT "钱(分)", 
  `code` varchar(50) NOT NULL COMMENT "码", 
  `status` int(11) NOT NULL COMMENT "状态(0:初始化,1:过期,2:成功,3:通知)",
  `wxId` varchar(100) NOT NULL COMMENT "微信领取id",
  `orderId` varchar(100) NOT NULL COMMENT "订单id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_arena_boss 定时boss
-- ----------------------------
DROP TABLE IF EXISTS `t_arena_boss`;
CREATE TABLE `t_arena_boss` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
      `platform` int(11) NOT NULL COMMENT "平台id",
    `serverId` int(11) NOT NULL COMMENT "服务器id",
    `bossId` int(11) NOT NULL COMMENT "boss",
    `mapId` int(11) NOT NULL COMMENT "地图id",
    `lastKillTime` bigint(20) NOT NULL COMMENT "上次击杀时间",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
  


/*
  
-- ----------------------------
-- Table structure for t_chuangshi_shenwang_signup 神王候选者报名
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_shenwang_signup`;
CREATE TABLE `t_chuangshi_shenwang_signup` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `campType` int(11) DEFAULT 0 COMMENT "阵营", 
  `playerServerId` int(11) DEFAULT 0 COMMENT "玩家服务器id",
  `playerId` bigint(20) DEFAULT 0 COMMENT "玩家id", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_chuangshi_shenwang_vote 神王候选者列表
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_shenwang_vote`;
CREATE TABLE `t_chuangshi_shenwang_vote` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `playerServerId` int(11) DEFAULT 0 COMMENT "玩家服务器id",
  `campType` int(11) DEFAULT 0 COMMENT "阵营",
  `playerId` bigint(20) DEFAULT 0 COMMENT "玩家id",
  `ticketNum` int(11) DEFAULT 0 COMMENT "投票次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_chuangshi_king_toupiao_record 投票记录
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_king_toupiao_record`;
CREATE TABLE `t_chuangshi_king_toupiao_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `campType` int(11) DEFAULT 0 COMMENT "阵营",
  `playerServerId` int(11) DEFAULT 0 COMMENT "游戏服务器id",
  `playerId` bigint(20) DEFAULT 0 COMMENT "玩家id",
  `playerName` varchar(50) DEFAULT "" COMMENT "玩家名字",
  `houXuanPlatform` int(11) DEFAULT 0 COMMENT "候选者平台",
  `houXuanGameServerId` int(11) DEFAULT 0 COMMENT "候选者游戏服务器id",
  `houXuanPlayerId` bigint(20) DEFAULT 0 COMMENT "候选者玩家id",
  `houXuanPlayerName` varchar(50) DEFAULT "" COMMENT "候选者玩家名字",
  `lastVoteTime` bigint(20) DEFAULT 0 COMMENT "上次投票时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_chuangshi_camp 创世阵营
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_camp`;
CREATE TABLE `t_chuangshi_camp` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `campType` int(11) DEFAULT 0 COMMENT "阵营",
  `kingId` bigint(20) DEFAULT 0 COMMENT "神王id",
  `force` bigint(20) DEFAULT 0 COMMENT "阵营总战力",
  `shenWangStatus` int(11) DEFAULT 0 COMMENT "神王竞选阶段",
  `jifen` bigint(20) DEFAULT 0 COMMENT "库存积分",
  `diamonds` bigint(20) DEFAULT 0 COMMENT "库存钻石",
  `payJifen` bigint(20) DEFAULT 0 COMMENT "工资积分", 
  `payDiamonds` bigint(20) DEFAULT 0 COMMENT "工资钻石",
  `lastShouYiTime` bigint(20) DEFAULT 0 COMMENT "上次工资时间",
  `targetMap` varchar(512) DEFAULT "{}" COMMENT "攻城目标",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_chuangshi_city 创世城池
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_city`;
CREATE TABLE `t_chuangshi_city` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `campType` int(11) DEFAULT 0 COMMENT "当前阵营",
  `originalCamp` int(11) DEFAULT 0 COMMENT "初始阵营",
  `typ` int(11) DEFAULT 0 COMMENT "城池类型",
  `index` int(11) DEFAULT 0 COMMENT "城池索引",
  `ownerId` bigint(20) DEFAULT 0 COMMENT "城主id",
  `jifen` bigint(20) DEFAULT 0 COMMENT "库存积分",
  `diamonds` bigint(20) DEFAULT 0 COMMENT "库存钻石",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_chuangshi_city_jianshe 创世城池建设
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_city_jianshe`;
CREATE TABLE `t_chuangshi_city_jianshe` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `cityId` bigint(20) DEFAULT 0 COMMENT "城池id",
  `jianSheType` int(11) DEFAULT 0 COMMENT "建设类型",
  `jianSheLevel` int(11) DEFAULT 0 COMMENT "建设等级",
  `jianSheExp` int(11) DEFAULT 0 COMMENT "建设经验",
  `skillLevelSet` int(11) DEFAULT 0 COMMENT "当前使用技能（天气台专用）",
  `skillMap` varchar(512) DEFAULT "{}" COMMENT "技能激活记录（天气台专用）",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_chuangshi_member 创世成员
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_member`;
CREATE TABLE `t_chuangshi_member` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `playerPlatform` int(11) DEFAULT 0 COMMENT "玩家平台id",
  `playerServerId` int(11) DEFAULT 0 COMMENT "玩家服务器id",
  `playerId` bigint(20) DEFAULT 0 COMMENT "玩家id",
  `playerName` varchar(50) DEFAULT "" COMMENT "玩家名字", 
  `playerLevel` int(11) DEFAULT 0 COMMENT "等级",
  `playerZhuanSheng` int(11) DEFAULT 0 COMMENT "转生等级",
  `playerJifen` bigint(20) DEFAULT 0 COMMENT "玩家创世积分",
  `online` int(11) DEFAULT 0 COMMENT "玩家离线状态0离线1在线",
  `allianceId` bigint(20) DEFAULT 0 COMMENT "仙盟id",
  `allianceName` varchar(50) DEFAULT "" COMMENT "仙盟名",
  `force` bigint(20) DEFAULT 0 COMMENT "战力",
  `scheduleJifen` bigint(20) DEFAULT 0 COMMENT "分配的积分",
  `scheduleDiamonds` bigint(20) DEFAULT 0 COMMENT "分配的钻石",
  `pos` int(11) DEFAULT 0 COMMENT "阵营职位",
  `campType` int(11) DEFAULT 0 COMMENT "阵营",
  `alPos` int(11) DEFAULT 0 COMMENT "仙盟职位",
  `sex` int(11) NOT NULL COMMENT "性别",  
  `role` int(11) NOT NULL COMMENT "角色",
  `fashionId` int(11) NOT NULL COMMENT "时装id", 
  `wingId` int(11) NOT NULL COMMENT "战翼id", 
  `weaponId` int(11) NOT NULL COMMENT "兵魂id",
  `guanZhiLevel` int(11) NOT NULL COMMENT "官职系统等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_chuangshi_camp_log 创世阵营战报
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_camp_log`;
CREATE TABLE `t_chuangshi_camp_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `platform` int(11) DEFAULT 0 COMMENT "平台id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `campType` int(11) DEFAULT 0 COMMENT "阵营",
  `type` int(11) DEFAULT 0 COMMENT "战报类型", 
  `content` varchar(512) DEFAULT 0 COMMENT "战报内容",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

*/