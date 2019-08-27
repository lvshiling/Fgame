set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

alter table t_player_item add column `used` int(11) default 0 COMMENT "是否使用";
alter table t_player_item add column `bagType` int(11) default 0 COMMENT "背包类型";

-- ----------------------------
-- Table structure for t_player_wing 玩家战翼数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_wing`;
CREATE TABLE `t_player_wing` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `wingId` int(11) NOT NULL COMMENT "当前战翼id",
  `unrealNum` int(11) NOT NULL COMMENT "战翼幻化丹培养次数",
  `unrealInfo` varchar(50) NOT NULL COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_equipment_slot 玩家装备槽数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_equipment_slot`;
CREATE TABLE `t_player_equipment_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `star` int(11) NOT NULL COMMENT "星级",
  `level` int(11) NOT NULL COMMENT "等级",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_body_shield 玩家护体盾数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_body_shield`;
CREATE TABLE `t_player_body_shield` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `jinjiadanNum` int(11) NOT NULL COMMENT "护体金甲丹培养次数",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_wear_fashion 玩家穿戴时装
-- ----------------------------
DROP TABLE IF EXISTS `t_player_fashion_wear`;
CREATE TABLE `t_player_fashion_wear` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `fashionWear` int(11) NOT NULL COMMENT "穿戴时装id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_fashion 玩家时装数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_fashion`;
CREATE TABLE `t_player_fashion` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `fashionId` int(11) NOT NULL COMMENT "时装id",
  `isExpire` int(11) NOT NULL COMMENT "是否过期",
  `activeTime` bigint(20) NOT NULL COMMENT "激活时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_player_title_wear 玩家穿戴称号
-- ----------------------------
DROP TABLE IF EXISTS `t_player_title_wear`;
CREATE TABLE `t_player_title_wear` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `titleWear` int(11) NOT NULL COMMENT "穿戴称号id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_title 玩家称号数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_title`;
CREATE TABLE `t_player_title` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `titleId` int(11) NOT NULL COMMENT "称号id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_shop 玩家当日商店购买道具(限购使用)
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shop`;
CREATE TABLE `t_player_shop` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `shopId` int(11) NOT NULL COMMENT "道具shopId",
  `dayCount` int(11) NOT NULL COMMENT "购买次数",
  `lastTime` bigint(20) NOT NULL COMMENT "最后一次购买时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_quest 玩家任务数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_quest`;
CREATE TABLE `t_player_quest` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `questId` int(11) NOT NULL COMMENT "任务id",
  `questData` varchar(500) NOT NULL COMMENT "任务数据",
  `collectItemData` varchar(500) NOT NULL COMMENT "收集物品数据",
  `questState` int(11) NOT NULL COMMENT "状态",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_property add column `zhuanSheng` int(11) default 0 COMMENT "转生";


-- ----------------------------
-- Table structure for t_player_weapon_info 玩家兵魂信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_weapon_wear`;
DROP TABLE IF EXISTS `t_player_weapon_info`;
CREATE TABLE `t_player_weapon_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `weaponWear` int(11) NOT NULL COMMENT "穿戴兵魂id",
  `star` int(11) NOT NULL COMMENT "总星数",
  `power` bigint(20) NOT NULL COMMENT "兵魂战斗力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_weapon 玩家兵魂数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_weapon`;
CREATE TABLE `t_player_weapon` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `weaponId` int(11) NOT NULL COMMENT "兵魂id",
  `level` int(11) NOT NULL COMMENT "兵魂星数",
  `culDan` int(11) NOT NULL COMMENT "培养丹数量",
  `state` int(11) NOT NULL COMMENT "觉醒状态 0:未觉醒 1觉醒",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_player_pk 玩家pk值
-- ----------------------------
DROP TABLE IF EXISTS `t_player_pk`;
CREATE TABLE `t_player_pk` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `pkValue` int(11) NOT NULL COMMENT "pk值",
  `killNum` int(11) NOT NULL COMMENT "击杀数量",
  `lastKillTime` bigint(20) DEFAULT 0 COMMENT "上次杀人时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_soul_embed 玩家帝魂镶嵌数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_soul_embed`;
CREATE TABLE `t_player_soul_embed` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `embedInfo` varchar(50) NOT NULL COMMENT "镶嵌帝魂id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_soul 玩家帝魂数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_soul`;
CREATE TABLE `t_player_soul` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `soulTag` int(11) NOT NULL COMMENT "帝魂标签",
  `level` int(11) NOT NULL COMMENT "帝魂等级",
  `experience` int(11) NOT NULL COMMENT "经验",
  `awakenOrder` int(11) NOT NULL  COMMENT "觉醒阶别",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_item_use 玩家物品使用数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_item_use`;
CREATE TABLE `t_player_item_use` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `todayTimes` int(11) NOT NULL COMMENT "今天使用次数",
  `totalTimes` int(11) NOT NULL COMMENT "总共使用次数",
  `lastUseTime` bigint(20) NOT NULL COMMENT "使用时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_mount add column `hidden` int(11) default 0 COMMENT "是否隐藏坐骑";
-- ----------------------------
-- Table structure for t_player_tianjieta 玩家天劫塔数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tianjieta`;
CREATE TABLE `t_player_tianjieta` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `level` int(11) NOT NULL COMMENT "天劫塔等级",
  `usedTime` bigint(20) NOT NULL COMMENT "使用时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_func_open 玩家功能开启数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_func_open`;
CREATE TABLE `t_player_func_open` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `funcOpenList` varchar(5000) NOT NULL COMMENT "功能开启列表",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_juexue_use 玩家绝学使用数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_juexue_use`;
CREATE TABLE `t_player_juexue_use` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "玩家使用绝学类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_juexue 玩家绝学数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_juexue`;
CREATE TABLE `t_player_juexue` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "绝学类型",
  `level` int(11) NOT NULL COMMENT "绝学等级",
  `insight` int(11) NOT NULL COMMENT "是否顿悟",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_xinfa 玩家心法数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_xinfa`;
CREATE TABLE `t_player_xinfa` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "心法类型",
  `level` int(11) NOT NULL COMMENT "心法等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_equipment_slot add column `gemInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "宝石信息";
-- ----------------------------
-- Table structure for t_player_friend 玩家好友数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_friend`;
CREATE TABLE `t_player_friend` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `friendId` bigint(20) NOT NULL COMMENT "好友id",
  `point` int(11) NOT NULL COMMENT "友好度",
  `black` int(11) NOT NULL COMMENT "黑名单",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;




-- ----------------------------
-- Table structure for t_player_cache  玩家缓存表
-- ----------------------------
DROP TABLE IF EXISTS `t_player_cache`;
CREATE TABLE `t_player_cache` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `name` varchar(100) NOT NULL COMMENT "名字",
  `role` int(11) NOT NULL COMMENT "角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `level` int(11) NOT NULL COMMENT "等级",
  `force` bigint(20) NOT NULL COMMENT "战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_mining  玩家矿工挖矿
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mining`;
CREATE TABLE `t_player_mining` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `level` int(11) NOT NULL COMMENT "矿工人数(矿工等级)",
  `storage` int(11) NOT NULL COMMENT "当前库存",
  `stone` bigint(20) NOT NULL COMMENT "玩家原石",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "检验库存时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_gamble  玩家赌石
-- ----------------------------
DROP TABLE IF EXISTS `t_player_gamble`;
CREATE TABLE `t_player_gamble` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "赌石类型 1初级赌石 2高级赌石",
  `num` bigint(20) NOT NULL COMMENT "次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_scene add column `sceneId` bigint(20) NOT NULL COMMENT "当前场景id";
alter table t_player_scene add column `lastSceneId` bigint(20) NOT NULL COMMENT "上一个场景id";


-- ----------------------------
-- Table structure for t_player_tumo  玩家屠魔次数
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tumo`;
CREATE TABLE `t_player_tumo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `num` int(11) NOT NULL COMMENT "今日已屠魔总次数",
  `extraNum` int(11) NOT NULL COMMENT "今日额外总次数",
  `usedNum` int(11) NOT NULL COMMENT "今日屠魔使用默认次数",
  `usedBuyNum` int(11) NOT NULL COMMENT "今日屠魔使用额外次数",
  `buyNum` int(11) NOT NULL COMMENT "今日已购买额外次数",
  `lastTime` bigint(20) NOT NULL COMMENT "最后一次更新时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_email  玩家邮件
-- ----------------------------
DROP TABLE IF EXISTS `t_player_email`;
CREATE TABLE `t_player_email` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "邮件id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `isRead` int(11) NOT NULL COMMENT "是否已读",
  `isGetAttachment` int(20) NOT NULL COMMENT "是否已领取附件",
  `title` varchar(30) NOT NULL COMMENT "邮件标题",
  `content` varchar(150) NOT NULL COMMENT "邮件内容",
  `attachementInfo` varchar(512) NOT NULL COMMENT "附件信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_pk add column `onlineTime` bigint(20) default 0 COMMENT "在线时间";
-- ----------------------------
-- Table structure for t_player_xianfu  秘境仙府
-- ----------------------------
DROP TABLE IF EXISTS `t_player_xianfu`;
CREATE TABLE `t_player_xianfu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `xianfuId` int(11) NOT NULL COMMENT "仙府id",
  `xianfuType` int(11) NOT NULL COMMENT "仙府类型",
  `useTimes` int(11) DEFAULT 0 COMMENT "已挑战次数",
  `startTime` bigint(20) DEFAULT 0 COMMENT "开始升级时间(ms)",
  `state`     int(11) NOT NULL COMMENT "0未升级 1升级进行中 ",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_soulruins_num  玩家帝陵遗迹挑战次数
-- ----------------------------
DROP TABLE IF EXISTS `t_player_soulruins_num`;
CREATE TABLE `t_player_soulruins_num` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `num` int(11) NOT NULL COMMENT "今日已挑战总次数",
  `extraBuyNum` int(11) NOT NULL COMMENT "今日额外购买总次数",
  `rewNum` int(11) NOT NULL COMMENT "今日通关奖励次数",
  `usedNum` int(11) NOT NULL COMMENT "今日已使用默认挑战次数",
  `usedBuyNum` int(11) NOT NULL COMMENT "今日已使用购买的挑战次数",
  `usedRewNum` int(11) NOT NULL COMMENT "今日已使用首次通关赠送的挑战次数",
  `buyNum` int(11) NOT NULL COMMENT "今日已购买挑战次数",
  `lastTime` bigint(20) NOT NULL COMMENT "最后一次更新时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_soulruins_rewchapter  玩家帝陵遗迹奖励章节
-- ----------------------------
DROP TABLE IF EXISTS `t_player_soulruins_rewchapter`;
CREATE TABLE `t_player_soulruins_rewchapter` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chapter` int(11) NOT NULL COMMENT "章节数",
  `type` int(11) NOT NULL COMMENT "1普通 2困难",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_soulruins  玩家帝陵遗迹数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_soulruins`;
CREATE TABLE `t_player_soulruins` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chapter` int(11) NOT NULL COMMENT "章节数",
  `type` int(11)  NOT NULL COMMENT "1普通 2困难",
  `level` int(11) NOT NULL COMMENT "关卡",
  `star` int(11) NOT NULL COMMENT "星数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_moonlove  玩家月下情缘活动
-- ----------------------------
DROP TABLE IF EXISTS `t_player_moonlove`;
CREATE TABLE `t_player_moonlove` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `charmNum` int(11) NOT NULL COMMENT "魅力值",
  `generousNum` int(11) NOT NULL COMMENT "豪气值",
  `preActivityTime`  bigint(20) DEFAULT 0 COMMENT  "上次活动时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_buff  玩家buff
-- ----------------------------
DROP TABLE IF EXISTS `t_player_buff`;
CREATE TABLE `t_player_buff` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `buffMap` varchar(5000) NOT NULL COMMENT "buff列表",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_activity  玩家活动
-- ----------------------------
DROP TABLE IF EXISTS `t_player_activity`;
CREATE TABLE `t_player_activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType` int(11) NOT NULL COMMENT "活动类型",
  `attendTimes` int(11)  NOT NULL COMMENT "已参与次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance  仙盟
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance`;
CREATE TABLE `t_alliance` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `name` varchar(500) NOT NULL COMMENT "名称",
  `notice` varchar(500) NOT NULL COMMENT "公告",
  `level` int(11) NOT NULL COMMENT "等级",
  `jianShe` bigint(20) NOT NULL COMMENT "建设度",
  `huFu` bigint(20) NOT NULL COMMENT "虎符数量",
  `totalForce` bigint(20) NOT NULL COMMENT "总战力",
  `mengzhuId` bigint(20) NOT NULL COMMENT "当前盟主id",
  `createId` bigint(20) NOT NULL COMMENT "创建人id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance_member 仙盟成员
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance_member`;
CREATE TABLE `t_alliance_member` (

  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",

  `memberId` bigint(20) NOT NULL COMMENT "成员id",

  `name` varchar(50) NOT NULL COMMENT "成员名字",
  `force` bigint(20) NOT NULL COMMENT "成员战力",
  `level` int(11) NOT NULL COMMENT "等级",
  `role` int(11) NOT NULL COMMENT "角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `zhuanSheng` int(11) NOT NULL COMMENT "转生",
  `position` int(11) NOT NULL COMMENT "职位",
  `gongXian` bigint(20) NOT NULL COMMENT "贡献",
  `joinTime` bigint(20) NOT NULL COMMENT "加入时间",
  `lastLogoutTime` bigint(20) NOT NULL COMMENT "上次退出时间",
  `lingyuId` int(11) default 0 COMMENT "领域id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`allianceId`),
  KEY(`memberId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance_join_apply 申请列表
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance_join_apply`;
CREATE TABLE `t_alliance_join_apply` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",
  `joinId` bigint(20) NOT NULL COMMENT "申请人id",
  `level` int(11) NOT NULL COMMENT "等级",
    `role` int(11) NOT NULL COMMENT "角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `name` varchar(50) NOT NULL COMMENT "申请人名字",
  `force` bigint(20) NOT NULL COMMENT "申请人战斗力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`allianceId`),
  KEY(`joinId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_emperor_worship 玩家膜拜次数
-- ----------------------------
DROP TABLE IF EXISTS `t_player_emperor_worship`;
CREATE TABLE `t_player_emperor_worship` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `num` int(11) NOT NULL COMMENT "膜拜次数",
  `lastTime` bigint(20) NOT NULL COMMENT "最后一次膜拜时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_emperor 龙椅
-- ----------------------------
DROP TABLE IF EXISTS `t_emperor`;
CREATE TABLE `t_emperor` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `emperorId` bigint(20) NOT NULL COMMENT "帝王id",
  `name` varchar(50) NOT NULL COMMENT "帝王名字",
  `spouseName` varchar(50) NOT NULL COMMENT "配偶名字", 
  `robNum` bigint(20) NOT NULL COMMENT "第几次争夺",
  `storage` bigint(20) NOT NULL COMMENT "帝王库存",
  `robTime` bigint(20) NOT NULL COMMENT "抢夺时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_emperor_record 龙椅抢夺记录
-- ----------------------------
DROP TABLE IF EXISTS `t_emperor_records`;
CREATE TABLE `t_emperor_records` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `emperorName` varchar(50) NOT NULL COMMENT "帝王名字",
  `robbedName` varchar(50) NOT NULL COMMENT "被抢名字",
  `robTime` bigint(20) NOT NULL COMMENT "抢夺时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_cache  玩家缓存表
-- ----------------------------
DROP TABLE IF EXISTS `t_player_cache`;
CREATE TABLE `t_player_cache` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `name` varchar(100) NOT NULL COMMENT "名字",
  `role` int(11) NOT NULL COMMENT "角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `level` int(11) NOT NULL COMMENT "等级",
  `force` bigint(20) NOT NULL COMMENT "战力",
  `baseProperty` varchar(100) NOT NULL COMMENT "基础属性",
  `battleProperty` varchar(1000) NOT NULL COMMENT "战斗属性",
  `equipmentList` varchar(1000) NOT NULL COMMENT "装备",
  `mountInfo` varchar(1000) NOT NULL COMMENT "坐骑",
  `wingInfo` varchar(1000) NOT NULL COMMENT "战翼",
  `bodyShieldInfo` varchar(1000) NOT NULL COMMENT "护体盾",
  `allSoulInfo` varchar(1000) NOT NULL COMMENT "古魂",
   `allWeaponInfo` varchar(1000) NOT NULL COMMENT "冰魂",
   `marryInfo` varchar(1000) NOT NULL COMMENT "结婚",
   `shieldInfo` varchar(1000) NOT NULL COMMENT "神盾尖刺",
   `featherInfo` varchar(1000) NOT NULL COMMENT "护体仙羽",
   `anqiInfo` varchar(1000) NOT NULL COMMENT "暗器",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- Table structure for t_alliance_hegemon 
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance_hegemon`;
CREATE TABLE `t_alliance_hegemon` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",
  `winNum` int(11) NOT NULL COMMENT "连胜次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`allianceId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_alliance 
-- ----------------------------
DROP TABLE IF EXISTS `t_player_alliance`;
CREATE TABLE `t_player_alliance` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",
  `allianceName` varchar(50) NOT NULL COMMENT "仙盟名字",
  `donateMap` varchar(500) NOT NULL COMMENT "捐献次数",
  `currentGongXian` bigint(20) NOT NULL COMMENT "当前捐献",
  `lastJuanXuanTime` bigint(20) NOT NULL COMMENT "上次捐献时间",
  `sceneRewardMap` varchar(100) NOT NULL COMMENT "城战奖励数据",
  `lastAllianceSceneEndTime` bigint(20) NOT NULL COMMENT "城战结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
     KEY(`allianceId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_alliance_skill  仙盟仙术
-- ----------------------------
DROP TABLE IF EXISTS `t_player_alliance_skill`;
CREATE TABLE `t_player_alliance_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `level` int(11) NOT NULL COMMENT "仙术等级",
  `skillType` int(11) NOT NULL COMMENT "仙术类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_cache add column `fashionId` int(11) default 0 COMMENT "时装";

-- ----------------------------
-- Table structure for t_alliance_log  仙盟日志
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance_log`;
CREATE TABLE `t_alliance_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",
  `content` varchar(150) NOT NULL COMMENT "内容",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`allianceId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_alliance add column `yaoPai` int(11) default 0 COMMENT "腰牌";
alter table t_player_alliance add column `lastYaoPaiUpdateTime` bigint(20) default 0 COMMENT "上次腰牌更新时间";
alter table t_player_alliance add column `convertTimes` int(11) default 0 COMMENT "兑换次数";
alter table t_player_alliance add column `lastConvertUpdateTime` bigint(20) default 0 COMMENT "上次兑换更新时间";

-- ----------------------------
-- Table structure for t_player_secret_card  天机牌
-- ----------------------------
DROP TABLE IF EXISTS `t_player_secret_card`;
CREATE TABLE `t_player_secret_card` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `num` int(11) NOT NULL COMMENT "天机牌",
  `totalStar` int(11) NOT NULL COMMENT "总星数",
  `openBoxs` varchar(20) NOT NULL COMMENT "运势箱开启",
  `cardId` int(11) NOT NULL COMMENT "接取天机牌",
  `star` int(11) NOT NULL COMMENT "接取天机牌星数",
  `cards` varchar(100) NOT NULL COMMENT "下发天机",
  `usedCards` varchar(1000) NOT NULL COMMENT "已使用天机牌id",
  `lastTime` bigint(20) NOT NULL COMMENT "记录时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
-- ----------------------------
-- Table structure for t_player_found  资源找回记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_found`;
CREATE TABLE `t_player_found` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `resType` int(11) NOT NULL COMMENT "资源类型",
  `playModeType` int(11) NOT NULL COMMENT "玩法类型：0日常1次数限制活动2无次数限制活动",
  `joinTimes` int(11) DEFAULT 0 COMMENT "参与次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_found_back  资源找回结果
-- ----------------------------
DROP TABLE IF EXISTS `t_player_found_back`;
CREATE TABLE `t_player_found_back` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `resType` int(11) NOT NULL COMMENT "资源类型",
  `isReceive` int(11) NOT NULL COMMENT "是否领取0否1是",
  `foundTimes` int(11) DEFAULT 0 COMMENT "找回次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_alliance add column `allianceLevel` int(11) NOT NULL COMMENT "仙盟等级";
   
-- ----------------------------
-- Table structure for t_player_dragon  神龙现世
-- ----------------------------
DROP TABLE IF EXISTS `t_player_dragon`;
CREATE TABLE `t_player_dragon` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `stageId` int(11) NOT NULL COMMENT "阶段id",
  `itemInfo` varchar(512) NOT NULL COMMENT "道具信息",
  `status` int(11) NOT NULL COMMENT "激活状态 0未激活 1激活神龙",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_cache add column `allianceId` bigint(20) default 0 COMMENT "仙盟";
alter table t_player_cache add column `teamId` bigint(20) default 0 COMMENT "队伍";

-- ----------------------------
-- Table structure for t_player_biaoche  玩家镖车信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_biaoche`;
CREATE TABLE `t_player_biaoche` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `robList` varchar(100) DEFAULT "[]" COMMENT "个人劫镖次数",
  `personalTransportTimes` int(11) NOT NULL COMMENT "个人押镖次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_biaoche  镖车
-- ----------------------------
DROP TABLE IF EXISTS `t_biaoche`;
CREATE TABLE `t_biaoche` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `transportType` int(11) NOT NULL COMMENT "镖车类型",
  `state` int(11) NOT NULL COMMENT "镖车状态",
  `robName` varchar(20) DEFAULT NULL COMMENT "劫镖人",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_four_god  四神遗迹
-- ----------------------------
DROP TABLE IF EXISTS `t_player_four_god`;
CREATE TABLE `t_player_four_god` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `keyNum` int(11) NOT NULL COMMENT "钥匙数",
  `exp` bigint(20) NOT NULL COMMENT "获得经验",
  `itemInfo` varchar(1024) NOT NULL DEFAULT "{}" COMMENT "获得物品",
  `endTime` bigint(20) NOT NULL COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_mount_other  非进阶坐骑
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mount_other`;
CREATE TABLE `t_player_mount_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "坐骑类型",
  `mountInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "激活的非进阶坐骑",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_wing_other  非进阶战翼
-- ----------------------------
DROP TABLE IF EXISTS `t_player_wing_other`;
CREATE TABLE `t_player_wing_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "战翼类型",
  `wingInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "激活的非进阶战翼",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_marry  玩家结婚数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_marry`;
CREATE TABLE `t_player_marry` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `spouseId` bigint(20) NOT NULL COMMENT "配偶id",
  `status` int(11) NOT NULL COMMENT "婚姻状态 1未婚 2求婚成功 3订婚 4举办过婚礼 5离婚",
  `ring` int(11) NOT NULL COMMENT "婚戒品质",
  `ringLevel` int(11) NOT NULL COMMENT "婚戒等级",
  `ringNum` int(11) NOT NULL COMMENT "婚戒培养次数",
  `ringExp` int(11) NOT NULL COMMENT "婚戒培养进度值",
  `treeLevel` int(11) NOT NULL COMMENT "爱情树等级",
  `treeNum` int(11) NOT  NULL COMMENT "爱情树培养次数",
  `treeExp` int(11) NOT NULL COMMENT "爱情树培养进度值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_marry  婚烟表
-- ----------------------------
DROP TABLE IF EXISTS `t_marry`;
CREATE TABLE `t_marry` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `spouseId` bigint(20) NOT NULL COMMENT "配偶id",
  `ring` int(11) NOT NULL COMMENT "婚戒类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_wedding  婚期安排
-- ----------------------------
DROP TABLE IF EXISTS `t_wedding`;
CREATE TABLE `t_wedding` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `period` int(11) NOT NULL COMMENT "预定场次",
  `grade` int(11) NOT NULL COMMENT "酒席档次",
  `hunCheGrade` int(11) NOT NULL COMMENT "婚车档次",
  `sugarGrade` int(11) NOT NULL COMMENT "喜糖档次",
  `status` int(11) NOT NULL COMMENT "1 未开始 2取消 3进行中 4举办过",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `spouseId` bigint(20) NOT NULL COMMENT "配偶id",
  `name` varchar(50) NOT NULL DEFAULT "" COMMENT  "名字",
  `spouseName` varchar(50) NOT NULL DEFAULT "" COMMENT "配偶名字",
  `hTime` bigint(20) DEFAULT 0 COMMENT "举办时间",
  `lastTime` bigint(20)  DEFAULT 0 COMMENT "最后操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`period`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_view_wedcard  玩家查看过喜帖
-- ----------------------------
DROP TABLE IF EXISTS `t_player_view_wedcard`;
CREATE TABLE `t_player_view_wedcard` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `cardId` bigint(20) NOT NULL COMMENT "喜帖id",
  `viewTime` bigint(20) NOT NULL COMMENT "查看时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_wedding_card  喜帖
-- ----------------------------
DROP TABLE IF EXISTS `t_wedding_card`;
CREATE TABLE `t_wedding_card` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `spouseId` bigint(20) NOT NULL COMMENT "配偶id",
  `playerName` varchar(50) NOT NULL COMMENT "玩家名字",
  `spouseName` varchar(50) NOT NULL COMMENT "配偶名字",
  `holdTime` varchar(100)  DEFAULT 0 COMMENT "举办时间",
  `outOfTime` bigint(20) DEFAULT 0 COMMENT "失效时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_marry_heroism  玩家豪气值
-- ----------------------------
DROP TABLE IF EXISTS `t_player_marry_heroism`;
CREATE TABLE `t_player_marry_heroism` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `heroism`  int(11) NOT NULL COMMENT "玩家豪气值",
  `outOfTime` bigint(20) DEFAULT 0 COMMENT "失效时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



alter table t_player_biaoche add column `robIdList` varchar(50) default 0 COMMENT "队伍";


alter table t_player_wing add column `featherId` int(11) default 1 COMMENT "护体仙羽id";
alter table t_player_wing add column `featherNum` int(11) default 0 COMMENT "护体仙羽培养次数";
alter table t_player_wing add column `featherPro` int(11) default 0 COMMENT "护体仙羽培养值";

alter table t_player_body_shield add column `shieldId` int(11) default 1 COMMENT "神盾尖刺id";
alter table t_player_body_shield add column `shieldNum` int(11) default 0 COMMENT "神盾尖刺培养次数";
alter table t_player_body_shield add column `shieldPro` int(11) default 0 COMMENT "神盾尖刺培养值";

alter table t_alliance_member add column `lingyuId` int(11) default 0 COMMENT "领域id";

-- ----------------------------
-- Table structure for t_player_shenfa 玩家身法数据

-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenfa`;
CREATE TABLE `t_player_shenfa` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `shenfaId` int(11) NOT NULL COMMENT "当前身法id",
  `unrealNum` int(11) NOT NULL COMMENT "身法幻化丹",
  `unrealInfo` varchar(50) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `power` bigint(20) NOT NULL COMMENT "身法战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_shenfa_other  非进阶身法皮肤
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenfa_other`;
CREATE TABLE `t_player_shenfa_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "战翼类型",
  `shenfaInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "激活的非进阶身法",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_lingyu 玩家领域数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingyu`;
CREATE TABLE `t_player_lingyu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `lingyuId` int(11) NOT NULL COMMENT "当前领域id",
  `unrealNum` int(11) NOT NULL COMMENT "领域幻化丹",
  `unrealInfo` varchar(50) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `power` bigint(20) NOT NULL COMMENT "领域战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_lingyu_other  非进阶领域皮肤
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingyu_other`;
CREATE TABLE `t_player_lingyu_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "战翼类型",
  `lingyuInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "激活的非进阶领域",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_cache add column `shenfaInfo` varchar(1000) DEFAULT "{}" COMMENT "身法";
alter table t_player_cache add column `lingyuInfo` varchar(1000) DEFAULT "{}" COMMENT "领域";

alter table t_player_cache add column `marryInfo` varchar(1000) DEFAULT "{}" COMMENT "结婚";

alter table t_player_item add column `level` int(11) DEFAULT 0 COMMENT "等级";


-- ----------------------------
-- Table structure for t_player_gold_equip_slot 玩家元神金装装备槽数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_gold_equip_slot`;
CREATE TABLE `t_player_gold_equip_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `level` int(11) NOT NULL COMMENT "等级",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_cache add column `goldEquipList` varchar(1000) DEFAULT "[]" COMMENT "元神金装";
alter table t_marry add column `status` int(11) NOT NULL COMMENT "婚烟状态 2求婚成功阶段 3订婚 4举办过婚礼";


-- ----------------------------
-- Table structure for t_alliance_invitation 仙盟邀请列表
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance_invitation`;
CREATE TABLE `t_alliance_invitation` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",
  `invitationId` bigint(20) NOT NULL COMMENT "邀请对象id",
  `level` int(11) NOT NULL COMMENT "等级",
  `role` int(11) NOT NULL COMMENT "角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `name` varchar(50) NOT NULL COMMENT "申请人名字",
  `force` bigint(20) NOT NULL COMMENT "申请人战斗力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`allianceId`),
  KEY(`invitationId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_secret_card add column  `totalNum` bigint(20) NOT NULL COMMENT "总次数";


alter table t_player_soul add column  `isAwaken` int(11) NOT NULL COMMENT "是否觉醒";
alter table t_player_soul add column  `strengthenLevel` int(11) NOT NULL DEFAULT 1  COMMENT "强化等级";
alter table t_player_soul add column  `strengthenNum` int(11) NOT NULL COMMENT "强化次数";
alter table t_player_soul add column  `strengthenPro` int(11) NOT NULL COMMENT "强化值";


alter table t_player_mount add column  `unrealLevel` int(11) DEFAULT 0 NOT NULL COMMENT "食幻化丹等级";
alter table t_player_mount add column  `unrealPro` int(11) NOT NULL COMMENT "食幻化丹进度值";
alter table t_player_mount add column  `culLevel` int(11) DEFAULT 0 NOT NULL COMMENT "食培养丹等级";
alter table t_player_mount add column  `culPro` int(11) NOT NULL COMMENT "食培养丹进度值";

alter table t_player_wing add column  `unrealLevel` int(11) DEFAULT 0 NOT NULL COMMENT "战翼幻化丹食丹等级";
alter table t_player_wing add column  `unrealPro` int(11)  NOT NULL COMMENT "战翼幻化丹培养进度值";
alter table t_player_wing add column  `hidden` int(11) DEFAULT 0 NOT NULL COMMENT "是否隐藏战翼";

alter table t_player_body_shield add column  `jinjiadanLevel` int(11) DEFAULT 0  NOT NULL COMMENT "金甲丹食用等级";
alter table t_player_body_shield add column  `jinjiadanNum` int(11) DEFAULT 0  NOT NULL COMMENT "护体金甲丹培养次数";
alter table t_player_body_shield add column  `jinjiadanPro` int(11) DEFAULT 0  NOT NULL COMMENT "金甲丹培养进度值";

alter table t_player_weapon add column  `culLevel` int(11) DEFAULT 0  NOT NULL COMMENT "兵魂食培养丹等级";
alter table t_player_weapon add column  `culNum` int(11) DEFAULT 0  NOT NULL COMMENT "培养丹次数";
alter table t_player_weapon add column  `culPro` int(11) DEFAULT 0  NOT NULL COMMENT "培养丹培养进度值";

alter table t_player_shenfa add column  `unrealLevel` int(11) DEFAULT 0 NOT NULL COMMENT "身法幻化丹食丹等级";
alter table t_player_shenfa add column  `unrealPro` int(11)  NOT NULL COMMENT "身法幻化丹培养进度值";
alter table t_player_shenfa add column  `hidden` int(11) DEFAULT 0 NOT NULL COMMENT "是否隐藏身法";

alter table t_player_lingyu add column  `unrealLevel` int(11) DEFAULT 0 NOT NULL COMMENT "领域幻化丹食丹等级";
alter table t_player_lingyu add column  `unrealPro` int(11)  NOT NULL COMMENT "领域幻化丹培养进度值";
alter table t_player_lingyu add column  `hidden` int(11) DEFAULT 0 NOT NULL COMMENT "是否隐藏领域";

-- ----------------------------
-- Table structure for t_onearena 灵池信息
-- ----------------------------
DROP TABLE IF EXISTS `t_onearena`;
CREATE TABLE `t_onearena` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `level` int(11) NOT NULL COMMENT "灵池等级",
  `pos` int(11) NOT NULL COMMENT "灵池位置",
  `ownerId` bigint(20) NOT NULL COMMENT "占领者id",
  `ownerName` varchar(50) NOT NULL COMMENT "占领者名字",
  `lastTime` bigint(20) NOT NULL COMMENT "上次产出时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_onearena  玩家灵池争夺数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_onearena`;
CREATE TABLE `t_player_onearena` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `level` int(11) NOT NULL COMMENT "灵池等级",
  `pos` int(11) NOT NULL COMMENT "灵池标识",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_cross  跨服数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_cross`;
CREATE TABLE `t_player_cross` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType`  int(11) NOT NULL COMMENT "活动类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_onearena_record  玩家灵池争夺时间数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_onearena_record`;
CREATE TABLE `t_player_onearena_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `level` int(11) NOT NULL COMMENT "灵池等级",
  `pos` int(11) NOT NULL COMMENT "灵池标识",
  `robTime`bigint(20) NOT NULL COMMENT "灵池抢夺时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_onearena_robbed 玩家灵池被抢记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_onearena_robbed`;
CREATE TABLE `t_player_onearena_robbed` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `robName` varchar(50) NOT NULL COMMENT "抢夺名字",
  `robTime` bigint(20) NOT NULL COMMENT "抢夺时间",
  `status` int(11) NOT NULL COMMENT "结果 1成功 2失败",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_onearena_kun 玩家下线后灵池获取的鲲
-- ----------------------------
DROP TABLE IF EXISTS `t_player_onearena_kun`;
CREATE TABLE `t_player_onearena_kun` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `kunInfo` varchar(1000) NOT NULL DEFAULT "{}" COMMENT "鲲信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;






-- Table structure for t_player_chess 玩家苍龙棋局
-- ----------------------------
DROP TABLE IF EXISTS `t_player_chess`;
CREATE TABLE `t_player_chess` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chessDropId` int(11) NOT NULL COMMENT "棋局掉落id",
  `attendTimes` int(11) NOT NULL COMMENT "棋局次数",
  `totalAttendTimes` int(11) NOT NULL COMMENT "总破解次数",
  `chessType` int(11) NOT NULL COMMENT "棋局类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_chess_log  苍龙棋局日志
-- ----------------------------
DROP TABLE IF EXISTS `t_chess_log`;
CREATE TABLE `t_chess_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerName` varchar(20) NOT NULL COMMENT "玩家姓名",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_cache add column `allianceName` varchar(100) DEFAULT NULL COMMENT "仙盟名称";
alter table t_player_item add column `isDepot` int(11) DEFAULT 0 COMMENT "是否在仓库";
alter table t_player_inventory add column `depotNum` int(11) NOT NULL COMMENT "仓库格子数";
alter table t_player_chess add column `lastSystemRefreshTime` bigint(20) DEFAULT 0 COMMENT "棋局上次自动刷新时间";

alter table t_alliance add column `transportTimes` int(11) NOT NULL COMMENT "押镖次数";


alter table t_player_title add column `activeTime` bigint(20) DEFAULT 0 NOT NULL  COMMENT "激活时间";
alter table t_player_title add column `activeFlag` int(11) DEFAULT 1 NOT NULL  COMMENT "是否激活";



-- ----------------------------
-- Table structure for t_marry_divorce_consent  协议离婚成功请求者已下线
-- remark: 拥有协议离婚亲密度扣除使用
-- ----------------------------
DROP TABLE IF EXISTS `t_marry_divorce_consent`;
CREATE TABLE `t_marry_divorce_consent` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_marry add column `isProposal` int(11) DEFAULT 0 NOT NULL  COMMENT "是否是求婚者";

alter table t_player_weapon add column `upNum` int(11) DEFAULT 0 NOT NULL  COMMENT "兵魂升星次数";
alter table t_player_weapon add column `upPro` int(11) DEFAULT 0 NOT NULL  COMMENT "兵魂升星进度值";

alter table t_marry add column `playerName` varchar(50) DEFAULT "" NOT NULL  COMMENT "玩家名字";
alter table t_marry add column `spouseName` varchar(50) DEFAULT "" NOT NULL  COMMENT "配偶名字";
alter table t_marry add column `point` int(11) DEFAULT 0 NOT NULL  COMMENT "亲密度";




-- ----------------------------
-- Table structure for t_player_relive  复活数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_relive`;
CREATE TABLE `t_player_relive` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `culTime`  int(11) NOT NULL COMMENT "累计复活次数",
  `lastReliveTime` bigint(20) NOT NULL COMMENT "上次复活时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_marry add column `spouseName` varchar(50) DEFAULT "" NOT NULL COMMENT "配偶名字";

alter table t_marry add column  `playerRingLevel`  int(11) DEFAULT 1 NOT NULL COMMENT "婚戒等级";
alter table t_marry add column  `spouseRingLevel` int(11) DEFAULT 1 NOT NULL COMMENT "玩家婚戒等级";

-- ----------------------------
-- Table structure for t_marry_ring  玩家求婚婚戒(求婚不成功返还)
-- ----------------------------
DROP TABLE IF EXISTS `t_marry_ring`;
CREATE TABLE `t_marry_ring` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `peerId` bigint(20) NOT NULL COMMENT "对方id",
  `peerName` varchar(50) NOT NULL COMMENT "对方名字",
  `ring` int(11) NOT NULL COMMENT "婚戒类型",
  `status` int(11) NOT NULL COMMENT "状态 1进行中 2失败",
  `proposalTime` bigint(20) DEFAULT 0 COMMENT "求婚时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_friend add column `isFriend` int(11) DEFAULT 0  NOT NULL COMMENT "玩家添加好友";
alter table t_player_friend add column `isReverseFriend` int(11) DEFAULT 0  NOT NULL COMMENT "好友添加玩家";


-- ----------------------------
-- Table structure for t_player_friend_log 玩家好友日志数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_friend_log`;
CREATE TABLE `t_player_friend_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `friendId` bigint(20) NOT NULL COMMENT "好友id",
  `type` int(11) NOT NULL COMMENT "操作类型 1添加  2删除 3加入黑名单",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_cache add column `shieldInfo` varchar(1000) DEFAULT "{}" COMMENT "神盾尖刺";
alter table t_player_cache add column `featherInfo` varchar(1000) DEFAULT "{}" COMMENT "护体仙羽";

alter table t_emperor add column `sex` int(11) NOT NULL DEFAULT 0 COMMENT "帝王性别";

ALTER TABLE t_player_cache MODIFY baseProperty VARCHAR(1000) default "{}" COMMENT "基础属性";

ALTER TABLE t_player_biaoche DROP COLUMN `allianceTransportTimes`;

ALTER TABLE t_player_tianjieta add column `playerName` varchar(50) DEFAULT "" COMMENT "玩家名字";

ALTER TABLE t_player_marry add column `wedStatus`  int(11) NOT NULL DEFAULT  0 COMMENT "玩家婚宴状态";

ALTER TABLE t_player_lingyu DROP COLUMN `skillId`;

alter table t_player_friend add column `isReverseBlack` int(11) DEFAULT 0  NOT NULL COMMENT "对方把玩家拉黑";

ALTER TABLE t_player_chess DROP COLUMN `chessDropId`;
ALTER TABLE t_player_chess add COLUMN  `chessId` int(11) NOT NULL COMMENT "棋局id";

alter table t_player_cache add column `realmLevel` int(11) DEFAULT 0 COMMENT "天劫塔等级";


-- ----------------------------
-- Table structure for t_player_arena  竞技场数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_arena`;
CREATE TABLE `t_player_arena` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `endTime`  bigint(20) NOT NULL COMMENT "活动结束时间",
    `culRewardTime` int(11) NOT NULL COMMENT "获得奖励次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_wing_trial  玩家战翼试用卡阶数
-- ----------------------------
DROP TABLE IF EXISTS `t_player_wing_trial`;
CREATE TABLE `t_player_wing_trial` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `trialOrderId`  int(11) NOT NULL COMMENT "试用卡获得阶数",
  `activeTime` bigint(20) NOT NULL COMMENT "激活时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_open_activity 福利厅
-- ----------------------------
DROP TABLE IF EXISTS `t_player_open_activity`;
CREATE TABLE `t_player_open_activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType`  int(11) NOT NULL COMMENT "开服活动类型：0投资计划1返利2排行3福利礼包",
  `activitySubType` int(11) NOT NULL COMMENT "开服活动子类型 type=0:0：初级投资 1：高级投资 2：七日投资;type=1:0：充值返利 1：消费返利;type=2: 0：充值排行1：消费排行2：坐骑排行3：战翼排行4：护盾排行5：领域排行6：身法排行7：护体仙域排行8：盾刺排行;type=3: 0:登录奖励1:升级奖励3:在线奖励",
  `activityData` varchar(1024)  DEFAULT "{}" COMMENT "活动数据",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player add column `todayOnlineTime` bigint(20) default 0 COMMENT "今天在线时间";

alter table t_player_chess add column `totalAttendTimes` int(11) NOT NULL COMMENT "总破解次数";

alter table t_player_wing add column `fpower` bigint(20) default 0 COMMENT "护体仙羽战力";
alter table t_player_body_shield add column `spower` bigint(20) default 0 COMMENT "神盾尖刺战力";
alter table t_player_arena add column `reliveTime` int(11) NOT NULL COMMENT "复活次数";
-- ----------------------------
-- Table structure for t_player_charge  玩家充值记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_charge`;
CREATE TABLE `t_player_charge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chargeType`  int(11) NOT NULL COMMENT "平台类型",
  `chargeNum` bigint(20) NOT NULL COMMENT "元宝数量",
  `chargeId` int(11) NOT NULL COMMENT "充值模板id",
  `orderId` varchar(50) NOT NULL COMMENT "订单id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_open_activity_charge  玩家活动充值
-- ----------------------------
DROP TABLE IF EXISTS `t_player_open_activity_charge`;
CREATE TABLE `t_player_open_activity_charge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `goldNum` bigint(20) NOT NULL COMMENT "元宝数量",
    `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime`   bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_open_activity_cost  玩家活动消费
-- ----------------------------
DROP TABLE IF EXISTS `t_player_open_activity_cost`;
CREATE TABLE `t_player_open_activity_cost` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `goldNum` bigint(20) NOT NULL COMMENT "元宝数量",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime`   bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_first_charge  玩家首充
-- ----------------------------
DROP TABLE IF EXISTS `t_player_first_charge`;
CREATE TABLE `t_player_first_charge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `isReceive` int(11) NOT NULL COMMENT "是否领取",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_major_num  玩家双修数
-- ----------------------------
DROP TABLE IF EXISTS `t_player_major_num`;
CREATE TABLE `t_player_major_num` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `num` int(11) NOT NULL COMMENT "双休次数",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "上次使用时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_hidden_weapon 玩家暗器数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_anqi`;
CREATE TABLE `t_player_anqi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `anqiDanLevel` int(11) NOT NULL COMMENT "暗器丹食用等级",
  `anqiDanNum` int(11) NOT NULL COMMENT "暗器丹培养次数",
  `anqiDanPro` int(11) NOT NULL COMMENT "暗器丹培养进度值",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `power` bigint(20) NOT NULL COMMENT "暗器战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_cache add column `anqiInfo` varchar(1000) DEFAULT "{}" COMMENT "暗器";

alter table t_player_fashion add column `star` int(11) DEFAULT 0  COMMENT "时装星数";
alter table t_player_charge add column `chargeId` int(11) NOT NULL COMMENT "充值模板id";

alter table t_alliance_member add column `zhuanSheng` int(11) NOT NULL COMMENT "转生";

alter table t_player_fashion add column `upStarNum` int(11) DEFAULT 0  COMMENT "时装升星次数";
alter table t_player_fashion add column `upStarPro` int(11) DEFAULT 0  COMMENT "时装升星进度值";


alter table t_player_open_activity add column `groupId` int(11) NOT NULL  COMMENT "活动Id";
alter table t_player_open_activity_charge add column `groupId` int(11) NOT NULL COMMENT "活动Id";
alter table t_player_open_activity_cost add column `groupId` int(11) NOT NULL  COMMENT "活动Id";

-- ----------------------------
-- Table structure for t_player_xuechi 玩家血池数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_xuechi`;
CREATE TABLE `t_player_xuechi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `bloodLine` int(11) NOT NULL COMMENT "补血生命线",
  `blood` bigint(20) NOT NULL COMMENT "血池剩余血量",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "上次补血时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_cross drop `activityType`;

alter table t_player_cross add `crossType` int(11);

alter table t_player_open_activity add column `startTime` bigint(20) DEFAULT 0  COMMENT "活动开始时间";
alter table t_player_open_activity add column `endTime` bigint(20) DEFAULT 0  COMMENT "活动结束时间";


alter table t_player_cross add column `crossArgs` varchar(100) DEFAULT "[]" COMMENT "跨服参数";

alter table t_player add column `serverId` int(11)  COMMENT "服务器id";

-- ----------------------------
-- Table structure for t_player_huiyuan 玩家会员数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_huiyuan`;
CREATE TABLE `t_player_huiyuan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "会员类型",
  `level` int(11) NOT NULL COMMENT "会员等级",
  `lastReceiveTime` bigint(20) DEFAULT 0 COMMENT "上次至尊会员领奖时间",
  `expireTime` bigint(20) DEFAULT 0 COMMENT "到期时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4; 


alter table t_player_open_activity_charge add column `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间";
alter table t_player_open_activity_cost add column `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间";

alter table t_player_huiyuan add column `level` int(11) NOT NULL COMMENT "会员等级"; 
alter table t_player_huiyuan add column `lastInterimReceiveTime` bigint(20) DEFAULT 0  COMMENT "上次临时会员奖励领取时间"; 
alter table t_player_huiyuan add column `interimBuyTime` bigint(20) DEFAULT 0  COMMENT "临时会员购买时间"; 
alter table t_player_huiyuan add column `plusBuyTime` bigint(20) DEFAULT 0  COMMENT "至尊会员购买时间"; 

alter table t_player_item add column `itemGetTime` bigint(20) DEFAULT 0 COMMENT "物品获取时间"; 


alter table t_player_cache add column `serverId` int(11)  COMMENT "服务器id";

alter table t_player_cache add column `skillList` text(2000)   COMMENT "技能列表";
update t_player_cache set `skillList`="[]";

alter table t_emperor add column `serverId` int(11) COMMENT "服务器id";
alter table t_emperor_records add column `serverId` int(11) COMMENT "服务器id";
alter table t_marry add column `serverId` int(11) COMMENT "服务器id";
alter table t_marry_divorce_consent add column `serverId` int(11) COMMENT "服务器id";
alter table t_wedding add column `serverId` int(11) COMMENT "服务器id";
alter table t_wedding_card add column `serverId` int(11) COMMENT "服务器id";
alter table t_marry_ring add column `serverId` int(11) COMMENT "服务器id";
alter table t_onearena add column `serverId` int(11) COMMENT "服务器id";

alter table t_player_property add column `charm` int(11) NOT NULL COMMENT "魅力值";

alter table t_player add column `forbid` int(11) DEFAULT 0 COMMENT "禁号 0正常 1禁号";
alter table t_player add column `forbidText` varchar(256) DEFAULT "" COMMENT "禁号原因";


-- ----------------------------
-- Table structure for t_player_skill_cd 玩家技能cd时间
-- ----------------------------
DROP TABLE IF EXISTS `t_player_skill_cd`;
CREATE TABLE `t_player_skill_cd` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `skillId` int(11) NOT NULL COMMENT "玩家技能",
  `lastTime` bigint(20) NOT NULL COMMENT "上次使用时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_skill drop `lastTime`;


alter table t_player_onearena add column `kunSilver` bigint(20) DEFAULT 0 COMMENT "出售鲲总银两";
alter table t_player_onearena add column `kunBindGold` bigint(20) DEFAULT 0 COMMENT "出售鲲总绑元";



alter table t_player add column `forbidTime` bigint(20) NOT NULL COMMENT "封号时间";
alter table t_player add column `forbidName` varchar(256) DEFAULT "" COMMENT "封号人";

alter table t_player add column `forbidChat` int(11) NOT NULL COMMENT "禁言 0正常 1封禁";
alter table t_player add column `forbidChatText` varchar(256) DEFAULT "" COMMENT "禁言原因";
alter table t_player add column `forbidChatTime` bigint(20) NOT NULL COMMENT "禁言时间";
alter table t_player add column `forbidChatName` varchar(256) DEFAULT "" COMMENT "禁言人";


alter table t_player_alliance add column `reliveTime` int(11) NOT NULL COMMENT "原地复活累计次数";
alter table t_player_alliance add column `lastReliveTime` bigint(20) DEFAULT 0 COMMENT "原地复活上次时间";

alter table t_alliance add column `serverId` int(11) COMMENT "服务器id";
alter table t_alliance_hegemon add column `serverId` int(11) COMMENT "服务器id";
alter table t_chess_log add column `serverId` int(11) COMMENT "服务器id";
alter table t_open_activity_email_record add column `serverId` int(11) COMMENT "服务器id";
alter table t_open_activity_email_record add column `endTime` bigint(20) COMMENT "活动结束时间";



alter table t_player_mount_other drop `mountInfo`;
alter table t_player_mount_other add column `mountId` int(11) DEFAULT 0 COMMENT "坐骑皮肤id";
alter table t_player_mount_other add column `level` int(11) DEFAULT 0 COMMENT "升星等级";
alter table t_player_mount_other add column `upNum` int(11) DEFAULT 0 COMMENT "升星次数";
alter table t_player_mount_other add column `upPro` int(11) DEFAULT 0 COMMENT "升星培养值";



alter table t_player_wing_other drop `wingInfo`;
alter table t_player_wing_other add column `wingId` int(11) DEFAULT 0 COMMENT "战翼皮肤id";
alter table t_player_wing_other add column `level` int(11) DEFAULT 0 COMMENT "升星等级";
alter table t_player_wing_other add column `upNum` int(11) DEFAULT 0 COMMENT "升星次数";
alter table t_player_wing_other add column `upPro` int(11) DEFAULT 0 COMMENT "升星培养值";

alter table t_player_item add column `bindType` int(11) default 0 COMMENT "绑定类型";
alter table t_player_gold_equip_slot add column `bindType` int(11) default 0 COMMENT "绑定类型";

alter table t_player_wing modify column `unrealInfo` varchar(256) ;
alter table t_player_mount modify column `unrealInfo` varchar(256) ;
alter table t_player_lingyu modify column `unrealInfo` varchar(256) ;
alter table t_player_shenfa modify column `unrealInfo` varchar(256) ;

alter table t_emperor add column `lastTime` bigint(20)  DEFAULT 0 COMMENT "上次产出时间";

-- ----------------------------
-- Table structure for t_player_first_charge_record  玩家档次首充记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_first_charge_record`;
CREATE TABLE `t_player_first_charge_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chargeType`  int(11) NOT NULL COMMENT "平台类型",
  `chargeId` int(11) NOT NULL COMMENT "充值模板id", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player_property add column `goldYuanLevel` int(11) default 0 COMMENT "元神等级";
alter table t_player_property add column `goldYuanExp` bigint(20) default 0 COMMENT "当前元神经验";



-- ----------------------------
-- Table structure for t_player_vip  玩家vip
-- ----------------------------
DROP TABLE IF EXISTS `t_player_vip`;
CREATE TABLE `t_player_vip` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `vipLevel`  int(11) NOT NULL COMMENT "vip等级",
  `vipStar` int(11) NOT NULL COMMENT "vip星级", 
  `consumeLevel` int(11) NOT NULL COMMENT "消费等级", 
  `chargeNum` int(11) NOT NULL COMMENT "充值数量",
  `freeGiftMap` varchar(512) DEFAULT "{}" COMMENT "免费礼包领取记录",
  `discountMap` varchar(512) DEFAULT "{}" COMMENT "礼包购买记录",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_friend 好友系统
-- ----------------------------
DROP TABLE IF EXISTS `t_friend`;
CREATE TABLE `t_friend` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `friendId` bigint(20) NOT NULL COMMENT "好友id",
  `point` int(11) NOT NULL COMMENT "友好度",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`serverId`,`playerId`,`friendId`),
    INDEX friendIdIndex (`serverId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_friend_black 玩家黑名单
-- ----------------------------
DROP TABLE IF EXISTS `t_player_friend_black`;
CREATE TABLE `t_player_friend_black` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `friendId` bigint(20) NOT NULL COMMENT "好友id",
  `black` int(11) NOT NULL COMMENT "拉黑对方",
  `revBlack` int(11) NOT NULL COMMENT "被对方拉黑",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_friend_invite 玩家收到的添加好友的邀请
-- ----------------------------
DROP TABLE IF EXISTS `t_player_friend_invite`;
CREATE TABLE `t_player_friend_invite` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `inviteId` bigint(20) NOT NULL COMMENT "邀请者id",
  `level` int(11) NOT NULL COMMENT "等级",
  `role` int(11) NOT NULL COMMENT "角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `name` varchar(50) NOT NULL COMMENT "邀请者名字",
  `force` bigint(20) NOT NULL COMMENT "邀请者战斗力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `t_player_friend`;

-- Table structure for t_player_lucky 幸运符
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lucky`;
CREATE TABLE `t_player_lucky` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "物品类型",
  `subType` int(11) NOT NULL COMMENT "物品子类型",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `expireTime` bigint(20) NOT NULL COMMENT "过期时间", 
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
    KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player add column `isOpenVideo` int(11) NOT NULL COMMENT "是否播放开场动画";

alter table t_chess_log drop `content`;
alter table t_chess_log add column `playerName` varchar(20) NOT NULL COMMENT "玩家姓名";
alter table t_chess_log add column `itemId` int(11) NOT NULL COMMENT "物品id";
alter table t_chess_log add column `itemNum` int(11) NOT NULL COMMENT "物品数量";


alter table t_player add column `originServerId` int(11) NOT NULL COMMENT "起始服务器id";
update t_player set originServerId=serverId;

alter table t_biaoche add column `serverId` int(11) COMMENT "服务器id";

alter table t_player_property modify column `gold` bigint(20) NOT NULL COMMENT "当前元宝";
alter table t_player_property modify column `bindGold` bigint(20) NOT NULL COMMENT "当前绑元";
alter table t_player_vip modify column `costNum` bigint(20) NOT NULL COMMENT "消费金额";
alter table t_player_vip modify column `chargeNum` bigint(20) NOT NULL COMMENT "充值数量";

CREATE TABLE `t_merge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` bigint(20) NOT NULL COMMENT "服务器id",
  `merge` int(11) DEFAULT 0 COMMENT "合服",
  `mergeTime` bigint(20) DEFAULT 0 COMMENT "合服时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
    KEY(`serverId`),
    INDEX serverIdIndex (`serverId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_emperor_records add column `type` int(11) NOT NULL default 1 COMMENT "1 抢夺记录 2开宝箱";
alter table t_emperor_records add column `itemInfo` varchar(256)  NOT NULL default "[]" COMMENT "物品信息";

alter table t_emperor add column `boxNum` bigint(20)  NOT NULL default 0 COMMENT "宝箱库存";
alter table t_emperor add column `specialBoxLeftNum` int(11)  NOT NULL default 0 COMMENT "高级宝箱剩余次数";
alter table t_emperor add column `boxLastTime` bigint(20)  NOT NULL default 0 COMMENT "宝箱上次产出时间";

alter table t_player_vip drop column `costNum`; 
alter table t_player_vip add column  `freeGiftMap` varchar(512) DEFAULT "{}" COMMENT "免费礼包领取记录";

-- ----------------------------
-- Table structure for t_player_fashion_trial  玩家时装试用卡阶数
-- ----------------------------
DROP TABLE IF EXISTS `t_player_fashion_trial`;
CREATE TABLE `t_player_fashion_trial` (  
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `trialFashionId`  int(11) NOT NULL COMMENT "时装id",
  `expireTime` bigint(20) NOT NULL COMMENT "过期时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`), 
   KEY(`playerId`),  
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



alter table t_player_lingyu_other drop `lingyuInfo`;
alter table t_player_lingyu_other add column `lingYuId` int(11) DEFAULT 0 COMMENT "领域皮肤id";
alter table t_player_lingyu_other add column `level` int(11) DEFAULT 0 COMMENT "升星等级";
alter table t_player_lingyu_other add column `upNum` int(11) DEFAULT 0 COMMENT "升星次数";
alter table t_player_lingyu_other add column `upPro` int(11) DEFAULT 0 COMMENT "升星培养值";



alter table t_player_shenfa_other drop `shenfaInfo`;
alter table t_player_shenfa_other add column `shenFaId` int(11) DEFAULT 0 COMMENT "身法皮肤id";
alter table t_player_shenfa_other add column `level` int(11) DEFAULT 0 COMMENT "升星等级";
alter table t_player_shenfa_other add column `upNum` int(11) DEFAULT 0 COMMENT "升星次数";
alter table t_player_shenfa_other add column `upPro` int(11) DEFAULT 0 COMMENT "升星培养值";

alter table t_player_charge add column `orderId` varchar(50) NOT NULL COMMENT "订单id";
alter table t_player_charge add column `state` int(11) NOT NULL COMMENT "订单状态";

-- ----------------------------
-- Table structure for t_first_charge  //首次重置时间
-- ----------------------------
DROP TABLE IF EXISTS `t_first_charge`;
CREATE TABLE `t_first_charge` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `chargeTime` bigint(20) NOT NULL COMMENT "首冲时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_pushwed_record  //推送婚礼按钮记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_pushwed_record`;
CREATE TABLE `t_player_pushwed_record` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `wedId` bigint(20) NOT NULL COMMENT "婚礼id",
  `hunCheTime` bigint(20) NOT NULL COMMENT "推送巡游时间",
  `banquetTime` bigint(20) NOT NULL COMMENT "推送酒席时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player add column `forbidEndTime` bigint(20) NOT NULL COMMENT "封禁结束时间";
alter table t_player add column `forbidChatEndTime` bigint(20) NOT NULL COMMENT "禁言结束时间";

alter table t_player add column `ignoreChat` int(11) NOT NULL COMMENT "禁默 0正常 1封禁";
alter table t_player add column `ignoreChatText` varchar(256) NOT NULL COMMENT "禁默原因";
alter table t_player add column `ignoreChatTime` bigint(20) NOT NULL COMMENT "禁默时间";
alter table t_player add column `ignoreChatEndTime` bigint(20) NOT NULL COMMENT "禁默结束时间";
alter table t_player add column `ignoreChatName` varchar(256) DEFAULT "" COMMENT "禁默人";


alter table t_emperor add column `boxOutNum` bigint(20) NOT NULL  DEFAULT 0 COMMENT "宝箱累计产出";


alter table t_player_equipment_slot add column `bindType` int(11) NOT NULL COMMENT "绑定类型";

-- ----------------------------
-- Table structure for t_chat_setting  //聊天设置
-- ----------------------------
DROP TABLE IF EXISTS `t_chat_setting`;
CREATE TABLE `t_chat_setting` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` bigint(20) NOT NULL COMMENT "服务器id",
  `worldVipLevel` int(11) NOT NULL COMMENT "世界频道vip等级",
  `worldLevel` int(11) NOT NULL COMMENT "世界频道等级",
  `allianceVipLevel` int(11) NOT NULL COMMENT "公会频道vip等级",
  `allianceLevel` int(11) NOT NULL COMMENT "公会频道等级",
  `privateVipLevel` int(11) NOT NULL COMMENT "私聊VIP等级",
  `privateLevel` int(11) NOT NULL COMMENT "私聊等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_register_setting  //注册设置
-- ----------------------------
DROP TABLE IF EXISTS `t_register_setting`;
CREATE TABLE `t_register_setting` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `open` int(11) NOT NULL COMMENT "0:关闭1:开放",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_compensate  //玩家补偿数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_compensate`;
CREATE TABLE `t_player_compensate` (  
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `compensateId` bigint(20) NOT NULL COMMENT "补偿id",
  `state` int(11) NOT NULL COMMENT "状态",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`), 
   KEY(`playerId`),  
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_compensate  //全服补偿信息
-- ----------------------------
DROP TABLE IF EXISTS `t_compensate`;
CREATE TABLE `t_compensate` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId`  int(11) NOT NULL COMMENT "服务器id", 
  `titlte` varchar(256) NOT NULL COMMENT "补偿标题",
  `content` varchar(256) NOT NULL COMMENT "补偿内容",
  `attachment` varchar(512) DEFAULT "{}" COMMENT "附件",
  `roleLevel` int(11) NOT NULL COMMENT "角色等级",
  `roleCreateTime` bigint(20) NOT NULL COMMENT "角色创建时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_register_setting_log  //注册设置日志
-- ----------------------------
DROP TABLE IF EXISTS `t_register_setting_log`;
CREATE TABLE `t_register_setting_log` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `open` int(11) NOT NULL COMMENT "0:关闭1:开放",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


alter table t_player add column `privilegeType` int(11) NOT NULL COMMENT "权限0:无 1:普通扶持 2:研发扶持";
alter table t_player add column `totalChargeMoney` bigint(20) NOT NULL COMMENT "总共充值金额";
alter table t_player add column `totalChargeGold` bigint(20) NOT NULL COMMENT "总共充值元宝";
alter table t_player add column `totalPrivilegeChargeGold` bigint(20) NOT NULL COMMENT "总共后台充值元宝";

-- ----------------------------
-- Table structure for t_order  //订单号
-- ----------------------------
DROP TABLE IF EXISTS `t_order`;
CREATE TABLE `t_order` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId`int(11) NOT NULL COMMENT "服务器id",
  `orderId` varchar(256) NOT NULL COMMENT "订单号",
  `orderStatus` int(11) NOT NULL COMMENT "0:充值成功1:发货成功",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chargeId` int(11) NOT NULL COMMENT "充值档次",
  `money` int(11) NOT NULL COMMENT "钱",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_player_charge drop column `state`;


alter table t_player_cache add column `vipInfo` varchar(1000) DEFAULT "{}" COMMENT "vip";


-- ----------------------------
-- Table structure for t_privilege_charge  //后台充值
-- ----------------------------
DROP TABLE IF EXISTS `t_privilege_charge`;
CREATE TABLE `t_privilege_charge` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId`int(11) NOT NULL COMMENT "服务器id",
  `status` int(11) NOT NULL COMMENT "0:充值成功 1:发货成功",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `goldNum` bigint(20) NOT NULL COMMENT "元宝数量",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

alter table t_order add column `playerLevel` int(11) NOT NULL COMMENT "玩家等级";
alter table t_order add column `gold` int(11) NOT NULL COMMENT "元宝";

alter table t_player add column `online` int(11) DEFAULT 0 COMMENT "在线状态";

alter table t_player add column `getNewReward` int(11) DEFAULT 0 COMMENT "是否领取新手礼包";

alter table t_player add column `sdkType` int(11) DEFAULT 0 COMMENT "sdk类型";

