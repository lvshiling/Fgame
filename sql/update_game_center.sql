set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game_center`;

alter table t_order add column `devicePlatform` int(11) DEFAULT 0 COMMENT "设备平台";
alter table t_order add column  `platformUserId` varchar(256) DEFAULT "" COMMENT "平台用户id";
alter table t_order add column  `playerLevel` int(11) DEFAULT 0 COMMENT "角色等级";
alter table t_order add column  `playerName` varchar(50) DEFAULT 0 COMMENT "角色名字";
alter table t_order add column  `gold` int(11) DEFAULT 0 COMMENT "元宝";

alter table t_chat_set add COLUMN sdkType int(11) DEFAULT 0 COMMENT "聊天设置";

alter table t_tulong_rank add COLUMN platform int(11) DEFAULT 1 COMMENT "平台";
  

create table t_redeem
(
   id                   int not null auto_increment,
   giftBagName          varchar(100) comment '礼包名字',
   giftBagDesc          varchar(500) comment '礼包文本',
   giftBagContent       text comment '礼包内容',
   redeemNum            int comment '兑换码数量',
   redeemPlayerUseNum   int comment '兑换码个人使用次数,0表示无限次',
   redeemServerUseNum   int comment '兑换码全服使用次数,0表示无限',
   sdkTypes             varchar(100) comment '兑换码sdk类别，以英文逗号隔开，为空表示全渠道',
   sendType             int comment '兑换发送方式,1:直接发放给玩家，2以邮件形式发放给玩家',
   startTime            bigint comment '生效开始时间',
   endTime              bigint comment '生效结束时间',
   minPlayerLevel       int comment '生效最低等级，0表示不限制',
   minVipLevel          int comment '生效最低VIP等级,0表示不限制',
   createFlag           int comment '码生成标志',
   updateTime           bigint(20) default 0 comment '更新时间',
   createTime           bigint(20) default 0 comment '创建时间',
   deleteTime           bigint(20) default 0 comment '删除时间',
   primary key (id)
);

alter table t_redeem comment '兑换码配置表';


create table t_redeem_code
(
   id                   int not null auto_increment comment '兑换码Id',
   redeemCode           varchar(100) comment '兑换码',
   redeemId             int comment '兑换码设置id,来自表t_redeem的主键id',
   useNum               int comment '实际兑换次数',
   updateTime           bigint(20) default 0 comment '更新时间',
   createTime           bigint(20) default 0 comment '创建时间',
   deleteTime           bigint(20) default 0 comment '删除时间',
   primary key (id)
);

alter table t_redeem_code comment '礼包兑换码';

CREATE INDEX IX_t_redeem_code_code ON t_redeem_code(redeemCode);
CREATE INDEX IX_t_redeem_code_pid ON t_redeem_code(redeemId);

create table t_redeem_platform
(
   id                   int not null auto_increment,
   platformId           int,
   redeemId             int comment '兑换码设置id,来自表t_redeem的主键id',
   updateTime           bigint(20) default 0 comment '更新时间',
   createTime           bigint(20) default 0 comment '创建时间',
   deleteTime           bigint(20) default 0 comment '删除时间',
   primary key (id)
);

alter table t_redeem_platform comment '对应中心平台id';

CREATE INDEX IX_t_redeem_platform_redeemId ON t_redeem_platform(redeemId);

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

CREATE VIEW v_player_firstorder AS
SELECT playerId,MIN(updateTime) AS firstOrderTime FROM t_order WHERE status IN (1,2);


-- 开始 2019-01-11 中心用户封禁字段  by cjy
alter table t_user add column  `forbid` int DEFAULT 0 COMMENT "禁号 0正常 1禁号";
alter table t_user add column  `forbidTime` bigint DEFAULT 0 COMMENT "封号时间";
alter table t_user add column  `forbidEndTime` bigint DEFAULT 0 COMMENT "封号结束时间";
alter table t_user add column  `forbidName` varchar(256) DEFAULT "" COMMENT "封号人";
alter table t_user add column  `forbidText` varchar(256) DEFAULT "" COMMENT "禁号原因";

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

-- 开始 2019-02-18 中心服务器表添加归属战区服务器字段  by cjy
alter table t_server add column `parentServerId` int(11) DEFAULT 0 COMMENT "归属服务器,id为服务器表的serverId";

-- 是否提前展示  by zrc
alter table t_server add column `preShow` int(11) DEFAULT 0 COMMENT "提前展示(0:不提前展示 1:提前展示)";

DROP TABLE IF EXISTS `t_platform_marryprice`;
CREATE TABLE `t_platform_marryprice`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platformId` bigint(11) NULL DEFAULT NULL COMMENT '中心平台Id',
  `kindType` int(11) NULL DEFAULT NULL COMMENT '启用类型1现实版，2廉价版',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '结婚平台价格版本配置表' ROW_FORMAT = Dynamic;


-- 客户端版本  by zrc
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

--聊天设置 by cjy 20190527
alter table t_chat_set add column teamVip int(11) DEFAULT 0 COMMENT '组队vip等级';
alter table t_chat_set add column teamPlayerLevel int(11) DEFAULT 0 COMMENT '组队玩家等级';


--聊天设置 by zrc 20190527
ALTER TABLE `t_redeem_record` ADD INDEX redeemCodeIndex (`redeemCode`);
ALTER TABLE `t_redeem_record` ADD INDEX playerIdIndex (`playerId`);

-- 兑换数量  by zrc
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

--中心平台，交易服战区配置 , by cjy 20190627
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

alter table t_server add column jiaoYiZhanQuServerId int(11) DEFAULT 0 COMMENT '交易战区号';
alter table t_server add column pingTaiFuServerId int(11) DEFAULT 0 COMMENT '全平台服Id';

-- 平台聊天
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

-- Table structure for t_arenapvp_bazhu 比武大会历届冠军
-- ----------------------------
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


-- 中心平台，城战战区服归属, by cjy 20190703
alter table t_server add column chengZhanServerId int(11) DEFAULT 0 COMMENT '城战服Id';

-- ----------------------------
-- Table structure for t_feedbackfee_exchange 兑换记录
-- ----------------------------
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
 
-- 兑换码使用次数, by cjy 20190715
alter table t_redeem add column redeemUseNum int(11) DEFAULT 1 COMMENT '兑换码使用次数';

-- ----------------------------
-- Table structure for t_arena_boss 定时boss
-- ----------------------------
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