set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 

-- ----------------------------
--  create by zrc 2019-03-18
-- Table structure for t_trade_item 交易物品
-- ----------------------------
CREATE TABLE `t_trade_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11)  NOT NULL COMMENT "物品数量",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `gold` bigint(20) NOT NULL COMMENT "价格",
  `status` int(11) NOT NULL COMMENT "状态",
  `globalTradeId` bigint(20) NOT NULL COMMENT "全局商品id",
  `buyPlatform` int(11) NOT NULL COMMENT "购买者平台",
  `buyServerId` int(11) NOT NULL COMMENT "购买者服务器",
  `buyPlayerId` bigint(20) NOT NULL COMMENT "购买者玩家id",
  `buyPlayerName` varchar(512) NOT NULL COMMENT "购买者玩家名字",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
--  create by zrc 2019-03-22
-- Table structure for t_trade_order 商品订单
-- ----------------------------
CREATE TABLE `t_trade_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `tradeId` bigint(20) NOT NULL COMMENT "商品id",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11)  NOT NULL COMMENT "物品数量",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `gold` bigint(20) NOT NULL COMMENT "价格",
  `status` int(11) NOT NULL COMMENT "状态0:支付1:发货2:取消",
  `sellPlatform` int(11)  NOT NULL COMMENT "卖家平台",
  `sellServerId` int(11)  NOT NULL COMMENT "卖家服务器id",
  `sellPlayerId` int(11)  NOT NULL COMMENT "卖家id",
  `sellPlayerName` int(11)  NOT NULL COMMENT "卖家名字", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by xzk 2019-03-19
 alter table `t_player_marry` add column `developExp` int(11) NOT NULL COMMENT "表白经验";
 alter table `t_player_marry` add column `developLevel` int(11) NOT NULL COMMENT "表白等级";
 alter table `t_player_marry` add column `coupleDevelopLevel` int(11) NOT NULL COMMENT "配偶表白等级";


 -- ----------------------------
-- cjb create by 2019-3-19
-- Table structure for t_player_activity_add_num  玩家活动增长数据
-- ----------------------------
CREATE TABLE `t_player_activity_add_num` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `addNum` int(11) NOT NULL COMMENT "增长值",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime`   bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xzk 2019-03-20
 alter table `t_marry` add column `developLevel` int(11) NOT NULL COMMENT "表白等级";
 alter table `t_marry` add column `spouseDevelopLevel` int(11) NOT NULL COMMENT "配偶表白等级";


-- ----------------------------
--  create by xzk 2019-03-18
-- Table structure for t_player_pregnant 玩家怀孕
-- ----------------------------
CREATE TABLE `t_player_pregnant` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `tonicPro` int(11) NOT NULL COMMENT "补品进度", 
  `chaoshengNum` int(11) NOT NULL COMMENT "超生数量", 
  `pregnantTime` bigint(20) NOT NULL COMMENT "怀孕时间",  
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  create by xzk 2019-03-18
-- Table structure for t_player_baby 玩家宝宝数据
-- ----------------------------
CREATE TABLE `t_player_baby` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `name` varchar(20) NOT NULL COMMENT "宝宝名称", 
  `sex` int(11) NOT NULL COMMENT "宝宝性别", 
  `quality` int(11) NOT NULL COMMENT "宝宝品质",
  `skillList` varchar(1024) NOT NULL COMMENT "天赋技能", 
  `activateTimes` int(11) NOT NULL COMMENT "激活技能次数", 
  `lockTimes` int(11) NOT NULL COMMENT "锁定技能次数", 
  `refreshTimes` int(11) NOT NULL COMMENT "洗练技能次数", 
  `learnExp` int(11) NOT NULL COMMENT "读书经验", 
  `learnLevel` int(11) NOT NULL COMMENT "读书等级", 
  `attrBeiShu` int(11) NOT NULL COMMENT "属性单倍",  
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by xzk 2019-03-22
-- Table structure for t_player_baby_toy_slot 玩家宝宝玩具槽数据
-- ----------------------------
CREATE TABLE `t_player_baby_toy_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `suitType` int(11) NOT NULL COMMENT "套装类型",
  `itemId` int(11) NOT NULL COMMENT "玩具id",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `level` int(11) NOT NULL COMMENT "玩具等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;




-- create by xzk 2019-03-22
alter table `t_player_major_num` add column `majorType` int(11) NOT NULL COMMENT "副本类型";

-- ----------------------------
-- Table structure for t_friend_marry_develop_log 全局表白日志记录数据
-- ----------------------------
CREATE TABLE `t_friend_marry_develop_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `sendId` bigint(20) NOT NULL COMMENT "发送玩家id",
  `recvId` bigint(20) NOT NULL COMMENT "接收玩家id",
  `sendName` varchar(100) NOT NULL COMMENT "发送玩家名字",
  `recvName` varchar(100) NOT NULL COMMENT "接收玩家名字",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量",
  `charmNum` int(11) NOT NULL COMMENT "魅力值",
  `developExp` int(11) NOT NULL COMMENT "表白经验",
  `contextStr` varchar(100) NOT NULL COMMENT "留言",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_friend_marry_develop_send_log 玩家表白日志记录数据
-- ----------------------------
CREATE TABLE `t_player_friend_marry_develop_send_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `recvId` bigint(20) NOT NULL COMMENT "接收玩家id",
  `recvName` varchar(100) NOT NULL COMMENT "接收玩家名字",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量",
  `charmNum` int(11) NOT NULL COMMENT "魅力值",
  `developExp` int(11) NOT NULL COMMENT "表白经验",
  `contextStr` varchar(100) NOT NULL COMMENT "留言",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_friend_marry_develop_recv_log 玩家被表白日志记录数据
-- ----------------------------
CREATE TABLE `t_player_friend_marry_develop_recv_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sendId` bigint(20) NOT NULL COMMENT "发送玩家id",
  `sendName` varchar(100) NOT NULL COMMENT "发送玩家名字",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量",
  `charmNum` int(11) NOT NULL COMMENT "魅力值",
  `developExp` int(11) NOT NULL COMMENT "表白经验",
  `contextStr` varchar(100) NOT NULL COMMENT "留言",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by zrc 2019-04-04
CREATE TABLE `t_player_marry_jinian`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `playerId` bigint(20) NULL DEFAULT NULL COMMENT '玩家Id',
  `jiNianType` int(11) NULL DEFAULT NULL COMMENT '纪念类型,1普通，2中等，3高级',
  `jiNianCount` int(11) NULL DEFAULT NULL COMMENT '举行的数量',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  `sendFlag` tinyint(4) NULL DEFAULT 0 COMMENT '是否发送0否1是',
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by zrc 2019-04-04
alter table `t_marry` add column `playerSuit` text NOT NULL COMMENT "玩家定情信物";
update t_marry set `playerSuit`="{}";

-- create by zrc 2019-04-04
alter table `t_marry` add column `spouseSuit` text NOT NULL COMMENT "伴侣定情信物";
update t_marry set `spouseSuit`="{}";

-- create by zrc 2019-04-04
CREATE TABLE `t_player_marry_dingqing`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `playerId` bigint(20) NULL DEFAULT NULL COMMENT '玩家Id',
  `suit` text DEFAULT NULL COMMENT '套装',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
)  ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by zrc 2019-04-04
CREATE TABLE `t_player_marry_jinian_sj`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NULL DEFAULT NULL COMMENT '玩家Id',
  `sjGetFlag` int(11) NULL DEFAULT 0 COMMENT '时装是否获取',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
  )  ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by xzk 2019-03-26
-- Table structure for t_player_qiyu 玩家奇遇信息
-- ----------------------------
CREATE TABLE `t_player_qiyu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `qiyuId` int(11) NOT NULL COMMENT "奇遇Id",
  `level` int(11) NOT NULL COMMENT "等级",
  `zhuan` int(11) NOT NULL COMMENT "转生",
  `fei` int(11) NOT NULL COMMENT "飞升",
  `endTime` bigint(20) NOT NULL COMMENT "任务结束时间",
  `isFinish` int(11) NOT NULL COMMENT "是否完成0否1是",
  `isReceive` int(11) NOT NULL COMMENT "是否领取0否1是",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table `t_player_marry` add column `marryCount` int(11) NOT NULL COMMENT "结婚次数";
-- 不能使用这个有一些已经离婚过的
-- UPDATE t_player_marry SET marryCount=marryCount+(SELECT COUNT(1) FROM t_marry WHERE t_player_marry.playerId=t_marry.playerId);
-- UPDATE t_player_marry SET marryCount=marryCount+(SELECT COUNT(1) FROM t_marry WHERE t_player_marry.playerId=t_marry.spouseId);

alter table `t_wedding` add column `isFirst` int(11) DEFAULT 0 NOT NULL COMMENT "是否结婚后第一次";


-- create by xzk 2019-03-28 
alter table `t_player_baby` add column `costItemNum` int(11) NOT NULL COMMENT "洗练消耗道具数量";


-- ----------------------------
-- Table structure for t_baby 玩家配偶宝宝数据
-- ----------------------------
CREATE TABLE `t_couple_baby` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `babyList` text(5000) NOT NULL COMMENT "宝宝列表", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;