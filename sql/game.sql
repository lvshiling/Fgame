set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';

DROP DATABASE IF EXISTS `game`;
CREATE DATABASE `game` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `game`;


-- ----------------------------
-- Table structure for t_player 玩家表
-- ----------------------------
DROP TABLE IF EXISTS `t_player`;
CREATE TABLE `t_player` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `userId` bigint(20) NOT NULL COMMENT "用户id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `originServerId` int(11) NOT NULL COMMENT "原始服务器id",
  `name` varchar(100) NOT NULL COMMENT "名字",
  `role` int(11) NOT NULL COMMENT "角色", 
  `sex` int(11) NOT NULL COMMENT "性别",
  `lastLoginTime` bigint(20) NOT NULL COMMENT "上次登陆时间",
  `lastLogoutTime` bigint(20) NOT NULL COMMENT "上次离线时间",
  `onlineTime` bigint(11) NOT NULL COMMENT "在线时间",
  `offlineTime` bigint(11) NOT NULL COMMENT "离线时间",
  `totalOnlineTime` bigint(11) NOT NULL COMMENT "总共在线时间",
  `todayOnlineTime` bigint(11) NOT NULL COMMENT "今日在线时间",
  `forbid` int(11)  NOT NULL COMMENT "禁号 0正常 1禁号",
  `forbidText`  varchar(256) DEFAULT "" COMMENT "禁号原因",
  `forbidTime` bigint(20) NOT NULL COMMENT "封号时间",
  `forbidEndTime` bigint(20) NOT NULL COMMENT "封号结束时间",
  `forbidName` varchar(256) DEFAULT "" COMMENT "封号人",
  `forbidChat` int(11) NOT NULL COMMENT "禁言 0正常 1封禁",
  `forbidChatText` varchar(256) DEFAULT "" COMMENT "禁言原因",
  `forbidChatTime` bigint(20) NOT NULL COMMENT "禁言时间",
  `forbidChatEndTime` bigint(20) NOT NULL COMMENT "禁言结束时间",
  `forbidChatName` varchar(256) DEFAULT "" COMMENT "禁言人",
  `ignoreChat` int(11) NOT NULL COMMENT "禁默 0正常 1封禁",
  `ignoreChatText` varchar(256) DEFAULT "" COMMENT "禁默原因",
  `ignoreChatTime` bigint(20) NOT NULL COMMENT "禁默时间",
  `ignoreChatEndTime` bigint(20) NOT NULL COMMENT "禁默结束时间",
  `ignoreChatName` varchar(256) DEFAULT "" COMMENT "禁默人",
  `isOpenVideo` int(11) NOT NULL COMMENT "是否播放开场动画",
  `privilegeType` int(11) NOT NULL COMMENT "权限0:无 1:普通扶持 2:研发扶持",
  `totalChargeMoney` bigint(20) NOT NULL COMMENT "总共充值金额",
  `totalChargeGold` bigint(20) NOT NULL COMMENT "总共充值元宝",
  `totalPrivilegeChargeGold` bigint(20) NOT NULL COMMENT "总共后台充值元宝",
  `getNewReward` int(11) NOT NULL COMMENT "是否领取新手奖励",
  `systemCompensate` int(11) NOT NULL COMMENT "是否领取进阶系统补偿奖励",
  `online` int(11) NOT NULL COMMENT "在线",
  `sdkType` int(11) NOT NULL COMMENT "sdk类型",
  `ip` varchar(50) NOT NULL COMMENT "ip",
  `todayChargeMoney` bigint(20) DEFAULT 0 COMMENT "今天充值金额",
  `yesterdayChargeMoney` bigint(20) DEFAULT 0 COMMENT "昨天充值金额",
  `chargeTime` bigint(20) DEFAULT 0 COMMENT "充值时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_scene 玩家场景数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_scene`;
CREATE TABLE `t_player_scene` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `mapId` int(11) NOT NULL COMMENT "地图id",
  `sceneId` bigint(20) NOT NULL COMMENT "场景id",
  `posX` float NOT NULL COMMENT "位置x",
  `posY` float NOT NULL COMMENT "位置y",
  `posZ` float NOT NULL COMMENT "位置z",
  `lastMapId` int(11) NOT NULL COMMENT "上一个地图id",
  `lastSceneId` bigint(20) NOT NULL COMMENT "上一个场景id",
  `lastPosX` float NOT NULL COMMENT "上一个地图位置x",
  `lastPosY` float NOT NULL COMMENT "上一个地图位置y",
  `lastPosZ` float NOT NULL COMMENT "上一个地图位置z",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`)  
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_property 玩家基础属性数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_property`;
CREATE TABLE `t_player_property` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `level` int(11) NOT NULL COMMENT "等级",
  `exp` bigint(20) NOT NULL COMMENT "当前经验值",
  `silver` bigint(20) NOT NULL  COMMENT "当前银两",
  `gold` bigint(20) NOT NULL COMMENT "当前元宝",
  `bindGold` bigint(20) NOT NULL COMMENT "绑定元宝",
  `evil` int(11) NOT NULL COMMENT "罪恶值",
  `zhuanSheng` int(11) NOT NULL COMMENT "转生",
  `currentHP` bigint(20) NOT NULL COMMENT "当前血量",
  `currentTP` bigint(20) NOT NULL COMMENT "当前体力",
  `power` bigint(20) NOT NULL COMMENT "战力",
  `charm` int(11) NOT NULL COMMENT "魅力值",
  `goldYuanLevel` int(11) default 0 COMMENT "元神等级",
  `goldYuanExp` bigint(20) default 0 COMMENT "当前元神经验",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_battle_property 玩家战斗属性数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_battle_property`;
CREATE TABLE `t_player_battle_property` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `maxHP` int(11) NOT NULL COMMENT "最大生命",
  `attack` int(11) NOT NULL COMMENT "攻击",
  `defend` int(11) NOT NULL COMMENT "防御",
  `maxTP` int(11) NOT NULL COMMENT "最大体力",
  `moveSpeed` int(11) NOT NULL COMMENT "移动速度",
  `crit` int(11) NOT NULL COMMENT "暴击",
  `tough` int(11) NOT NULL COMMENT "坚韧",
   `block` int(11) NOT NULL COMMENT "格挡",
  `break` int(11) NOT NULL COMMENT "破格",
  `dodge` int(11) NOT NULL COMMENT "闪避",
  `hit` int(11) NOT NULL COMMENT "命中",
  `huanYuanAttack` int(11) NOT NULL COMMENT "混元伤害",
  `huanYuanDef` int(11) NOT NULL COMMENT "混元防御",
  `bindDongRes` int(11) NOT NULL COMMENT "冰冻抗性",
  `poJiaRes` int(11) NOT NULL COMMENT "破解抗性",
  `kuiLeiRes` int(11) NOT NULL COMMENT "傀儡抗性",
  `kuJieRes` int(11) NOT NULL COMMENT "枯竭抗性",
  `shiMingRes` int(11) NOT NULL COMMENT "失明抗性",
  `xuRuoRes` int(11) NOT NULL COMMENT "虚弱抗性",
  `jiaoXieRes` int(11) NOT NULL COMMENT "缴械抗性",
  `zhongDuRes` int(11) NOT NULL COMMENT "中毒抗性",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_skill 玩家技能数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_skill`;
CREATE TABLE `t_player_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `skillId` int(11) NOT NULL COMMENT "技能id",
  `level` int(11) NOT NULL COMMENT "技能等级",
  `tianFuInfo` varchar(256) NOT NULL  COMMENT "天赋信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

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


-- ----------------------------
-- Table structure for t_player_inventory 玩家背包数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_inventory`;
CREATE TABLE `t_player_inventory` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `slotNum` int(11) NOT NULL COMMENT "背包格子数",
  `depotNum` int(11) NOT NULL COMMENT "仓库格子数",
  `miBaoDepotNum` int(11) NOT NULL COMMENT "秘宝仓库格子数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_player_item 玩家物品数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_item`;
CREATE TABLE `t_player_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `bagType` int(11) NOT NULL COMMENT "背包类型",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `index` int(11) NOT NULL COMMENT "索引",
  `num` int(11) NOT NULL COMMENT "叠加次数",
  `used` int(11) NOT NULL DEFAULT 0 COMMENT "是否使用过",
  `lastUseTime` bigint(20) NOT NULL COMMENT "上一次使用时间",
  `itemGetTime` bigint(20) DEFAULT 0 COMMENT "物品获取时间",
  `level` int(11) NOT NULL COMMENT "等级",
  `isDepot` int(11) DEFAULT 0 COMMENT "是否在仓库",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  `gemInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "宝石信息",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_dan 玩家食丹数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_dan`;
CREATE TABLE `t_player_dan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `levelId` int(11) NOT NULL COMMENT "玩家食丹等级",
  `danInfo` varchar(512) NOT NULL DEFAULT "{}" COMMENT "丹药信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_alchemy 玩家炼丹数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_alchemy`;
CREATE TABLE `t_player_alchemy` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `kindId` int(11) NOT NULL COMMENT "模板id",
  `num` int(11) NOT NULL COMMENT "可合成丹药数量",
  `startTime` bigint(20) NOT NULL COMMENT "开始炼丹时间(ms)",
  `state`     int(11) NOT NULL COMMENT "1进行中 2完成  3领取",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_mount 玩家坐骑数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mount`;
CREATE TABLE `t_player_mount` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `mountId` int(11) NOT NULL COMMENT "当前坐骑id",
  `unrealLevel` int(11) NOT NULL COMMENT "食幻化丹等级",
  `unrealNum` int(11) NOT NULL COMMENT "食幻化丹次数",
  `unrealPro` int(11) NOT NULL COMMENT "食幻化丹进度值",
  `culLevel`  int(11) NOT NULL COMMENT "食培养丹等级",
  `culNum` int(11) NOT NULL COMMENT "食培养丹次数",
  `culPro` int(11) NOT NULL COMMENT "食培养丹进度值",
  `unrealInfo` varchar(256) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `hidden` int(11) default 0 COMMENT "是否隐藏坐骑",
  `power` bigint(20) NOT NULL COMMENT "坐骑战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create time:20180402
-- author：ylz
-- Table structure for t_player_wing 玩家战翼数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_wing`;
CREATE TABLE `t_player_wing` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `wingId` int(11) NOT NULL COMMENT "当前战翼id",
  `unrealLevel` int(11) NOT NULL COMMENT "战翼幻化丹食丹等级",
  `unrealNum` int(11) NOT NULL COMMENT "战翼幻化丹次数",
  `unrealPro` int(11) NOT NULL COMMENT "战翼幻化丹培养进度值",
  `unrealInfo` varchar(256) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `featherId` int(11) NOT NULL DEFAULT 1 COMMENT "护体仙羽id",
  `featherNum` int(11) NOT NULL COMMENT "护体仙羽培养次数",
  `featherPro` int(11) NOT NULL COMMENT "护体仙羽培养值",
  `hidden` int(11) default 0 COMMENT "是否隐藏战翼",
  `power` bigint(20) NOT NULL COMMENT "战翼战力",
  `fpower` bigint(20) NOT NULL COMMENT "护体仙羽战力", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create time:20180407
-- author：ylz
-- Table structure for t_player_body_shield 玩家护体盾数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_body_shield`;
CREATE TABLE `t_player_body_shield` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `jinjiadanLevel` int(11) NOT NULL COMMENT "金甲丹食用等级",
  `jinjiadanNum` int(11) NOT NULL COMMENT "护体金甲丹培养次数",
  `jinjiadanPro` int(11) NOT NULL COMMENT "金甲丹培养进度值",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `shieldId` int(11) NOT NULL DEFAULT 1 COMMENT "神盾尖刺id",
  `shieldNum` int(11) NOT NULL COMMENT "神盾尖刺培养次数",
  `shieldPro` int(11) NOT NULL COMMENT "神盾尖刺培养值",
  `power` bigint(20) NOT NULL COMMENT "护体盾战力",
  `spower` bigint(20) NOT NULL COMMENT "神盾尖刺战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_fashion 玩家时装数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_fashion`;
CREATE TABLE `t_player_fashion` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `fashionId` int(11) NOT NULL COMMENT "时装id",
  `star` int(11) NOT NULL COMMENT "时装星数",
  `upStarNum` int(11) NOT NULL COMMENT "时装升星次数",
  `upStarPro` int(11) NOT NULL COMMENT "时装升星进度值",
  `isExpire` int(11) NOT NULL COMMENT "是否过期",
  `activeTime` bigint(20) NOT NULL COMMENT "激活时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_title 玩家称号数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_title`;
CREATE TABLE `t_player_title` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `titleId` int(11) NOT NULL COMMENT "称号id",
  `activeFlag` int(11) NOT NULL COMMENT "是否激活",
  `activeTime` bigint(20) DEFAULT 0 COMMENT "激活时间",
  `validTime` bigint(20) DEFAULT 0 COMMENT "有效时间",
  `starLev` int(11) NOT NULL COMMENT "升星等级",
  `starNum` int(11) NOT NULL COMMENT "升星次数",
  `starBless` int(11) NOT NULL COMMENT "升星祝福值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  `onlineTime` bigint(20) DEFAULT 0 COMMENT "在线时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_weapon_info 玩家兵魂信息
-- ----------------------------
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_weapon 玩家兵魂数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_weapon`;
CREATE TABLE `t_player_weapon` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `weaponId` int(11) NOT NULL COMMENT "兵魂id",
  `activeFlag` int(11) NOT NULL COMMENT "激活标识",
  `level` int(11) NOT NULL COMMENT "兵魂星数", 
  `upNum` int(11) NOT NULL COMMENT "兵魂升星次数",
  `upPro` int(11) NOT NULL COMMENT "兵魂升星进度值",
  `culLevel` int(11) NOT NULL COMMENT "兵魂食培养丹等级",
  `culNum` int(11) NOT NULL COMMENT "培养丹次数",
  `culPro` int(11) NOT NULL COMMENT "培养丹培养进度值",
  `state` int(11) NOT NULL COMMENT "觉醒状态 0:未觉醒 1觉醒",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;




-- ----------------------------
-- Table structure for t_player_soul_embed 玩家帝魂镶嵌数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_soul_embed`;
CREATE TABLE `t_player_soul_embed` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `embedInfo` varchar(50) NOT NULL COMMENT "镶嵌帝魂",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  `isAwaken` int(11) NOT NULL COMMENT "是否觉醒",
  `awakenOrder` int(11) NOT NULL  COMMENT "觉醒阶别",
  `strengthenLevel` int(11) NOT NULL COMMENT "强化等级",
  `strengthenNum` int(11) NOT NULL COMMENT "强化次数",
  `strengthenPro` int(11) NOT NULL COMMENT "强化值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_tianjieta 玩家天劫塔数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tianjieta`;
CREATE TABLE `t_player_tianjieta` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `playerName` varchar(50) NOT NULL COMMENT "玩家名字",
  `level` int(11) NOT NULL COMMENT "天劫塔等级",
  `usedTime` bigint(20) NOT NULL COMMENT "使用时间",
  `isCheckReissue` int(11) NOT NULL COMMENT "是否检测补发过",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;




-- ----------------------------
-- Table structure for t_player_friend_log 玩家好友日志数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_friend_log`;
CREATE TABLE `t_player_friend_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `friendId` bigint(20) NOT NULL COMMENT "好友id",
  `type` int(11) NOT NULL COMMENT "操作类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


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
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
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
  `title` varchar(500) NOT NULL COMMENT "邮件标题",
  `content` varchar(500) NOT NULL COMMENT "邮件内容",
  `attachementInfo` text(5000) NOT NULL COMMENT "附件信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


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
  `state`     int(11) NOT NULL COMMENT "0未升级 1升级进行中",
  `group` int(11) NOT NULL COMMENT "当前波数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance  仙盟
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance`;
CREATE TABLE `t_alliance` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `originServerId` int(11) NOT NULL COMMENT "原始服务器id",
  `name` varchar(500) NOT NULL COMMENT "名称",
  `notice` varchar(500) NOT NULL COMMENT "公告",
  `level` int(11) NOT NULL COMMENT "等级",
  `jianShe` bigint(20) NOT NULL COMMENT "建设度",
  `huFu` bigint(20) NOT NULL COMMENT "虎符数量",
  `totalForce` bigint(20) NOT NULL COMMENT "总战力",
  `mengzhuId` bigint(20) NOT NULL COMMENT "当前盟主id", 
  `createId` bigint(20) NOT NULL COMMENT "创建人id",
  `transportTimes` int(11) NOT NULL COMMENT "押镖次数",
  `lastTransportRefreshTime` bigint(20) NOT NULL COMMENT "上次押镖次数刷新时间",
  `isAutoAgree` int(11) DEFAULT 0 NOT NULL COMMENT "是否自动同意入盟申请：0否1是",
  `isAutoRemoveDepot` int(11) DEFAULT 0 NOT NULL COMMENT "是否自动销毁仙盟仓库物品：0否1是",
  `maxRemoveZhuanSheng` int(11) DEFAULT 0 NOT NULL COMMENT "自动销毁最大转生条件",
  `maxRemoveQuality` int(11) DEFAULT 0 NOT NULL COMMENT "自动销毁最高品质",
  `lastMergeTime` bigint(20) NOT NULL COMMENT "合帮时间",
  -- `campType` int(11) DEFAULT 0 NOT NULL COMMENT "阵营类型",
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
  `vip` int(11) NOT NULL COMMENT "vip等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`allianceId`), 
  KEY(`memberId`),
   INDEX allianceIdIndex (`allianceId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance_join_apply 仙盟人员申请列表
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
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_emperor 龙椅
-- ----------------------------
DROP TABLE IF EXISTS `t_emperor`;
CREATE TABLE `t_emperor` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `emperorId` bigint(20) NOT NULL COMMENT "帝王id",
  `name` varchar(50) NOT NULL COMMENT "帝王名字",
  `sex` int(11) NOT NULL COMMENT "帝王性别",
  `spouseName` varchar(50) NOT NULL COMMENT "配偶名字", 
  `robNum` bigint(20) NOT NULL COMMENT "第几次争夺",
  `storage` bigint(20) NOT NULL COMMENT "帝王库存",
  `robTime` bigint(20) NOT NULL COMMENT "抢夺时间",
  `lastTime` bigint(20) NOT NULL COMMENT "上次产出时间",
  `boxNum` bigint(20) NOT NULL COMMENT "宝箱库存",
  `boxOutNum` bigint(20) NOT NULL COMMENT "宝箱累计产出",
  `specialBoxLeftNum` int(11) NOT NULL COMMENT "高级宝箱剩余次数",
  `boxLastTime` bigint(20) NOT NULL COMMENT "宝箱上次产出时间",
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
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `type` int(11) NOT NULL COMMENT "1 抢夺记录 2开宝箱",
  `emperorName` varchar(50) NOT NULL COMMENT "帝王名字",
  `robbedName` varchar(50) NOT NULL COMMENT "被抢名字",
  `robTime` bigint(20) NOT NULL COMMENT "操作时间",
  `itemInfo` varchar(256) NOT NULL COMMENT "物品信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance_hegemon 城战
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance_hegemon`;
CREATE TABLE `t_alliance_hegemon` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `allianceId` bigint(20) NOT NULL COMMENT "霸主仙盟id",
  `winNum` int(11) NOT NULL COMMENT "连胜次数",
  `defenceAllianceId` bigint(20) NOT NULL COMMENT "守方仙盟id", 
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
  `allianceLevel` int(11) NOT NULL COMMENT "仙盟等级",
  `lastJuanXuanTime` bigint(20) NOT NULL COMMENT "上次捐献时间",
  `sceneRewardMap` varchar(100) NOT NULL COMMENT "城战奖励数据",
  `warPoint` int(11) DEFAULT 0 COMMENT "城战积分",
  `yaoPai` int(11) NOT NULL COMMENT "腰牌",
  `lastYaoPaiUpdateTime` bigint(20) NOT NULL COMMENT "上次腰牌更新时间",
  `convertTimes` int(11) NOT NULL COMMENT "兑换次数",
  `lastConvertUpdateTime` bigint(20) NOT NULL COMMENT "上次兑换更新时间",
  `lastAllianceSceneEndTime` bigint(20) NOT NULL COMMENT "城战结束时间",
  `reliveTime` int(11) NOT NULL COMMENT "原地复活累计次数",
  `lastReliveTime` bigint(20) DEFAULT 0 COMMENT "原地复活上次时间",
  `lastMemberCallTime` bigint(20) DEFAULT 0 COMMENT "上次仙盟召集时间",
  `lastYuXiMemberCallTime` bigint(20) DEFAULT 0 COMMENT "上次玉玺之战仙盟召集时间", 
  `totalWinTime` int(11) NOT NULL COMMENT "城战胜利次数",
  `depotPoint` int(11) DEFAULT 0 COMMENT "仓库积分",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`allianceId`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_cache  玩家缓存表
-- ----------------------------
DROP TABLE IF EXISTS `t_player_cache`;
CREATE TABLE `t_player_cache` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `name` varchar(100) NOT NULL COMMENT "名字",
  `role` int(11) NOT NULL COMMENT "角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `level` int(11) NOT NULL COMMENT "等级",
  `force` bigint(20) NOT NULL COMMENT "战力",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟",
  `allianceName` varchar(100) DEFAULT NULL COMMENT "仙盟名称",
  `teamId` bigint(20) NOT NULL COMMENT "队伍id",
  `baseProperty` text(2000) NOT NULL COMMENT "基础属性",
  `battleProperty` text(2000) NOT NULL COMMENT "战斗属性",
  `equipmentList` text(2000) NOT NULL COMMENT "装备",
  `mountInfo` text(2000) NOT NULL COMMENT "坐骑",
  `wingInfo` text(2000) NOT NULL COMMENT "战翼",
  `bodyShieldInfo` text(2000) NOT NULL COMMENT "护体盾",
  `allSoulInfo` text(2000) NOT NULL COMMENT "古魂",
  `allWeaponInfo` text(2000) NOT NULL COMMENT "冰魂",
  `fashionId` int(11) NOT NULL COMMENT "时装",
  `marryInfo` text(2000) NOT NULL COMMENT "结婚",
  `goldEquipList` text(2000) NOT NULL COMMENT "元神金装",
  `wushuangListInfo` text(2000) NOT NULL COMMENT "无双神器",
  `shenfaInfo` text(2000) NOT NULL  COMMENT "身法",
  `lingyuInfo` text(2000) NOT NULL COMMENT "领域",
  `shieldInfo` text(2000) NOT NULL COMMENT "神盾尖刺",
  `featherInfo` text(2000) NOT NULL COMMENT "护体仙羽",
  `anqiInfo` text(2000)  NOT NULL  COMMENT "暗器",
  `massacreInfo` text(2000)  NOT NULL  COMMENT "戮仙刃",
  `realmLevel` int(11) NOT NULL COMMENT "天劫塔等级",
  `skillList` text(2000) NOT NULL COMMENT "技能列表",
  `vipInfo` text(2000) NOT NULL COMMENT "vip",
  `fabaoInfo` text(2000) NOT NULL COMMENT "法宝",
  `xuedunInfo` text(2000) NOT NULL COMMENT "血盾",
  `xiantiInfo` text(2000) NOT NULL COMMENT "仙体",
  `baguaInfo` text(2000) NOT NULL COMMENT "八卦秘境",
  `dianxingInfo` text(2000) NOT NULL COMMENT "点星",
  `tianMoTiInfo` text(2000) NOT NULL COMMENT "天魔体",
  `shihunfanInfo` text(2000) NOT NULL COMMENT "噬魂幡",
  `allLingTongDevInfo` text(6000) NOT NULL COMMENT "灵童养成类",
  `lingTongInfo` text(2000) NOT NULL COMMENT "灵童信息",
  `allSystemSkillInfo` text(5000) NOT NULL COMMENT "系统技能",
  `allAdditionSysInfo` text(5000) NOT NULL COMMENT "附加系统类信息",
  `pregnantInfo` text(2000) NOT NULL COMMENT "怀孕信息", 
  `isHuiYuan` int(11) NOT NULL COMMENT "是否永久至尊会员",
  `xianZunCardInfo` text(2000) NOT NULL COMMENT "仙尊特权卡类型",
  `ringInfo` text(2000) NOT NULL COMMENT "特戒",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) ,
   INDEX `nameIndex` (`serverId`, `name`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

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

-- ----------------------------
-- Table structure for t_player_found  资源找回记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_found`;
CREATE TABLE `t_player_found` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `resType` int(11) NOT NULL COMMENT "资源类型",
  `playModeType` int(11) NOT NULL COMMENT "玩法类型：0日常1次数限制活动2无次数限制活动",
  `joinTimes` int(11) NOT NULL COMMENT "参与次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


   -- ----------------------------
-- Table structure for t_player_found_back  资源找回结果
-- ----------------------------
DROP TABLE IF EXISTS `t_player_found_back`;
CREATE TABLE `t_player_found_back` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `resType` int(11) NOT NULL COMMENT "资源类型",
  `resLevel` int(11) NOT NULL COMMENT "资源等级",
  `isReceive` int(11) NOT NULL COMMENT "是否领取0否1是",
  `foundTimes` int(11) DEFAULT 0 COMMENT "找回次数",
  `group` int(11) DEFAULT 0 COMMENT "最大挑战波数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
   
-- ----------------------------
-- Table structure for t_player_secret_card  天机牌
-- ----------------------------
DROP TABLE IF EXISTS `t_player_secret_card`;
CREATE TABLE `t_player_secret_card` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `totalNum` bigint(20) NOT NULL COMMENT "总次数",
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
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

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
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

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
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- Table structure for t_biaoche  镖车
-- ----------------------------
DROP TABLE IF EXISTS `t_biaoche`;
CREATE TABLE `t_biaoche` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `serverId` int(11) COMMENT "服务器id",
  `allianceId`  bigint(20) NOT NULL COMMENT "仙盟id",
  `transportMoveId` int(11) NOT NULL COMMENT "镖车路径模板id",
  `transportType` int(11) NOT NULL COMMENT "镖车类型",
  `state` int(11) NOT NULL COMMENT "镖车状态",
  `owerName` varchar(20) DEFAULT NULL COMMENT "拥有者名字",
  `robName` varchar(20) DEFAULT NULL COMMENT "劫镖人",
  `lastDistressUpdateTime` bigint(20) DEFAULT 0 COMMENT "上一次求救时间",
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
  `mountId` int(11) NOT NULL COMMENT "坐骑皮肤id",
  `level` int(11) NOT NULL COMMENT "升星等级",
  `upNum` int(11) NOT NULL COMMENT "升星次数",
  `upPro` int(11) NOT NULL COMMENT "升星培养值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_wing_other  非进阶战翼
-- ----------------------------
DROP TABLE IF EXISTS `t_player_wing_other`;
CREATE TABLE `t_player_wing_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "战翼类型",
  `wingId` int(11) NOT NULL COMMENT "战翼皮肤id",
  `level` int(11) NOT NULL COMMENT "升星等级",
  `upNum` int(11) NOT NULL COMMENT "升星次数",
  `upPro` int(11) NOT NULL COMMENT "升星培养值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_shenfa 玩家身法数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenfa`;
CREATE TABLE `t_player_shenfa` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `shenfaId` int(11) NOT NULL COMMENT "当前身法id",
  `unrealNum` int(11) NOT NULL COMMENT "身法幻化丹次数",
  `unrealInfo` varchar(256) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `power` bigint(20) NOT NULL COMMENT "身法战力",
  `unrealLevel` int(11) NOT NULL NOT NULL COMMENT "身法幻化丹食丹等级",
  `unrealPro` int(11)  NOT NULL COMMENT "身法幻化丹培养进度值",
  `hidden` int(11) DEFAULT 0 NOT NULL COMMENT "是否隐藏身法",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_shenfa_other  非进阶身法皮肤
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenfa_other`;
CREATE TABLE `t_player_shenfa_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "身法类型",
  `shenFaId` int(11) NOT NULL COMMENT "身法皮肤id",
  `level` int(11) NOT NULL COMMENT "升星等级",
  `upNum` int(11) NOT NULL COMMENT "升星次数",
  `upPro` int(11) NOT NULL COMMENT "升星培养值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
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
  `unrealNum` int(11) NOT NULL COMMENT "领域幻化丹次数",
  `unrealInfo` varchar(256) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `power` bigint(20) NOT NULL COMMENT "领域战力",
  `unrealLevel` int(11) NOT NULL NOT NULL COMMENT "领域幻化丹食丹等级",
  `unrealPro` int(11)  NOT NULL COMMENT "领域幻化丹培养进度值",
  `hidden` int(11) DEFAULT 0 NOT NULL COMMENT "是否隐藏领域",
  `chargeVal` bigint(20) NOT NULL COMMENT "模块开启累计充值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_lingyu_other  非进阶领域皮肤
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingyu_other`;
CREATE TABLE `t_player_lingyu_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "领域类型",
  `lingYuId` int(11) NOT NULL COMMENT "领域皮肤id",
  `level` int(11) NOT NULL COMMENT "升星等级",
  `upNum` int(11) NOT NULL COMMENT "升星次数",
  `upPro` int(11) NOT NULL COMMENT "升星培养值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_player_marry  玩家结婚数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_marry`;
CREATE TABLE `t_player_marry` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `spouseId` bigint(20) NOT NULL COMMENT "配偶id",
  `spouseName` varchar(50) NOT NULL COMMENT "配偶名字",
  `status` int(11) NOT NULL COMMENT "婚姻状态 1未婚 2求婚成功 3订婚 4举办过婚礼 5离婚",
  `ring` int(11) NOT NULL COMMENT "婚戒品质",
  `ringLevel` int(11) NOT NULL COMMENT "婚戒等级",
  `ringNum` int(11) NOT NULL COMMENT "婚戒培养次数",
  `ringExp` int(11) NOT NULL COMMENT "婚戒培养进度值",
  `treeLevel` int(11) NOT NULL COMMENT "爱情树等级",
  `treeNum` int(11) NOT  NULL COMMENT "爱情树培养次数",
  `treeExp` int(11) NOT NULL COMMENT "爱情树培养进度值",
  `isProposal` int(11) NOT NULL COMMENT "是否是请求者",
  `wedStatus` int(11) NOT NULL COMMENT "玩家婚宴状态",
  `developExp` int(11) NOT NULL COMMENT "表白经验", 
  `developLevel` int(11) NOT NULL COMMENT "表白等级",
  `coupleDevelopLevel` int(11) NOT NULL COMMENT "配偶表白等级",
  `marryCount` int(11) NOT NULL COMMENT "结婚次数",
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
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `spouseId` bigint(20) NOT NULL COMMENT "配偶id",
  `playerName` varchar(50) NOT NULL COMMENT "玩家名字",
  `spouseName` varchar(50) NOT NULL COMMENT "配偶名字",
  `playerRingLevel` int(11) NOT NULL COMMENT "婚戒等级",
  `spouseRingLevel` int(11) NOT NULL COMMENT "玩家婚戒等级",
  `role` int(11) NOT NULL COMMENT "角色",
  `spouseRole` int(11) NOT NULL COMMENT "配偶角色",
  `sex` int(11) NOT NULL COMMENT "性别",
  `spouseSex` int(11) NOT NULL COMMENT "配偶性别",
  `point` int(11) NOT NULL COMMENT "亲密度",
  `ring` int(11) NOT NULL COMMENT "婚戒类型",
  `status` int(11) NOT NULL COMMENT "婚烟状态 2求婚成功阶段 3订婚 4举办过婚礼",
  `developLevel` int(11) NOT NULL COMMENT "表白等级",
  `spouseDevelopLevel` int(11) NOT NULL COMMENT "配偶表白等级",
  `playerSuit` text NOT NULL COMMENT "玩家定情信物",
  `spouseSuit` text NOT NULL COMMENT "伴侣定情信物",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_marry_divorce_consent  协议离婚成功请求者已下线
-- remark: 拥有协议离婚亲密度扣除使用
-- ----------------------------
DROP TABLE IF EXISTS `t_marry_divorce_consent`;
CREATE TABLE `t_marry_divorce_consent` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
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
  `isFirst` int(11)  DEFAULT 0 COMMENT "是否结婚后第一次",
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
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_wedding_card  喜帖
-- ----------------------------
DROP TABLE IF EXISTS `t_wedding_card`;
CREATE TABLE `t_wedding_card` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
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
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_marry_ring  玩家求婚婚戒(求婚不成功返还)
-- ----------------------------
DROP TABLE IF EXISTS `t_marry_ring`;
CREATE TABLE `t_marry_ring` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
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

-- ----------------------------
-- Table structure for t_player_gold_equip_slot 玩家元神金装装备槽数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_gold_equip_slot`;
CREATE TABLE `t_player_gold_equip_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `level` int(11) NOT NULL COMMENT "等级",
  `newStLevel` int(11) NOT NULL COMMENT "新强化等级",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `gemInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "宝石信息",
  `gemUnlockInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "解锁宝石孔信息",
  `castingSpiritInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "铸灵信息",
  `forgeSoulInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "锻魂信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

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


-- Table structure for t_player_chess 玩家苍龙棋局
-- ----------------------------
DROP TABLE IF EXISTS `t_player_chess`;
CREATE TABLE `t_player_chess` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chessId` int(11) NOT NULL COMMENT "棋局id",
  `attendTimes` int(11) NOT NULL COMMENT "棋局次数",
  `totalAttendTimes` int(11) NOT NULL COMMENT "总破解次数",
  `chessType` int(11) NOT NULL COMMENT "棋局类型",
  `lastSystemRefreshTime` bigint(20) DEFAULT 0 COMMENT "棋局上次自动刷新时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_chess_log  苍龙棋局日志
-- ----------------------------
DROP TABLE IF EXISTS `t_chess_log`;
CREATE TABLE `t_chess_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `playerName` varchar(20) NOT NULL COMMENT "玩家姓名",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;




-- ----------------------------
-- Table structure for t_onearena 灵池信息
-- ----------------------------
DROP TABLE IF EXISTS `t_onearena`;
CREATE TABLE `t_onearena` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
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
  `robTime` bigint(20) DEFAULT 0 COMMENT "抢夺时间",
  `kunSilver` bigint(20) DEFAULT 0 COMMENT "出售鲲总银两",
  `kunBindGold` bigint(20) DEFAULT 0 COMMENT "出售鲲总绑元",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_cross  跨服数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_cross`;
CREATE TABLE `t_player_cross` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `crossType`  int(11) NOT NULL COMMENT "活动类型",
  `crossArgs` varchar(100) NOT NULL DEFAULT "[]" COMMENT "跨服参数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


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
  KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_arena  竞技场数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_arena`;
CREATE TABLE `t_player_arena` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `endTime`  bigint(20) NOT NULL COMMENT "活动结束时间",
  `reliveTime` int(11) NOT NULL COMMENT "复活次数",
  `culRewardTime` int(11) NOT NULL COMMENT "获得奖励次数",
  `totalRewardTime` int(11) NOT NULL COMMENT "累计获胜次数",
  `jiFenCount` int(11) NOT NULL COMMENT "累计积分",
  `jiFenDay` int(11) NOT NULL COMMENT "每日积分",
  `arenaTime` bigint(20) NOT NULL COMMENT "积分更新时间",
  `winCount` int(11) NOT NULL COMMENT "连胜次数",
  `failCount` int(11) NOT NULL COMMENT "连败次数",
  `rankRewTime` bigint(20) NOT NULL COMMENT "上次周榜奖励时间",
  `dayMaxWinCount` int(11) NOT NULL COMMENT "当天最高连胜",
  `dayWinCount` int(11) NOT NULL COMMENT "当天连胜",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_open_activity 开服活动
-- ----------------------------
DROP TABLE IF EXISTS `t_player_open_activity`;
CREATE TABLE `t_player_open_activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType`  int(11) NOT NULL COMMENT "开服活动类型",
  `activitySubType` int(11) NOT NULL COMMENT "开服活动子类型",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `activityData` varchar(1024)  DEFAULT "{}" COMMENT "活动数据",
  `startTime` bigint(20) DEFAULT 0  COMMENT "活动开始时间",
  `endTime` bigint(20) DEFAULT 0  COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_charge  玩家充值记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_charge`;
CREATE TABLE `t_player_charge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chargeType`  int(11) NOT NULL COMMENT "平台类型",
  `chargeId` int(11) NOT NULL COMMENT "充值模板id",
  `chargeNum` bigint(20) NOT NULL COMMENT "元宝数量",
  `orderId` varchar(50) NOT NULL COMMENT "订单id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`), 
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_open_activity_charge  玩家活动充值
-- ----------------------------
DROP TABLE IF EXISTS `t_player_open_activity_charge`;
CREATE TABLE `t_player_open_activity_charge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `goldNum` bigint(20) NOT NULL COMMENT "元宝数量",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime`   bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_open_activity_cost  玩家活动消费
-- ----------------------------
DROP TABLE IF EXISTS `t_player_open_activity_cost`;
CREATE TABLE `t_player_open_activity_cost` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `goldNum` bigint(20) NOT NULL COMMENT "元宝数量",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
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
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
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
  `majorType` int(11) NOT NULL COMMENT "副本类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
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
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


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
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


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
  `lastInterimReceiveTime` bigint(20) DEFAULT 0  COMMENT "上次临时会员奖励领取时间",
  `expireTime` bigint(20) DEFAULT 0 COMMENT "临时会员到期时间",
  `interimBuyTime` bigint(20) DEFAULT 0  COMMENT "临时会员购买时间",
  `plusBuyTime` bigint(20) DEFAULT 0  COMMENT "至尊会员购买时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_open_activity_email_record 开服活动邮件奖励记录
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_email_record`;
CREATE TABLE `t_open_activity_email_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id", 
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


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
  `chargeNum` bigint(20) NOT NULL COMMENT "充值数量",
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


-- Table structure for t_merge 合服标志
-- ----------------------------
DROP TABLE IF EXISTS `t_merge`;
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
  `teamVipLevel` int(11) NOT NULL COMMENT "队伍VIP等级",
  `teamLevel` int(11) NOT NULL COMMENT "队伍等级",
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
  `auto` int(11) NOT NULL COMMENT "自动关闭过",
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
  `playerLevel` int(11) NOT NULL COMMENT "玩家等级",
  `chargeId` int(11) NOT NULL COMMENT "充值档次",
  `money` int(11) NOT NULL COMMENT "钱",
  `gold` int(11) NOT NULL COMMENT "元宝",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


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

-- ----------------------------
-- xzk create by 2018-11-01
-- Table structure for t_alliance_depot 仙盟仓库
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance_depot`;
CREATE TABLE `t_alliance_depot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `index` int(11) NOT NULL COMMENT "索引",
  `num` int(11) NOT NULL COMMENT "叠加次数",
  `used` int(11) NOT NULL DEFAULT 0 COMMENT "是否使用过",
  `level` int(11) NOT NULL COMMENT "等级",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_massacre 玩家戮仙刃数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_massacre`;
CREATE TABLE `t_player_massacre` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `currLevel` int(11) NOT NULL COMMENT "当前阶数",
  `currStar` int(11) NOT NULL COMMENT "当前星数",
  `lastTime` bigint(20) NOT NULL COMMENT "上次被击杀掉落杀气时间",
  `shaQiNum` bigint(20) NOT NULL COMMENT "杀气数量",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数升级次数",
  `power` bigint(20) NOT NULL COMMENT "戮仙刃战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- xzk create by 2018-11-06
-- Table structure for t_player_tower 玩家打宝塔数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tower`;
CREATE TABLE `t_player_tower` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `useTime` bigint(20) NOT NULL COMMENT "已用打宝时间",
  `extraTime` bigint(20) NOT NULL COMMENT "额外打宝时间",
  `lastResetTime` bigint(20) NOT NULL COMMENT "上次打宝时间重置时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
    KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_marry_pre_wed  玩家婚礼预定档次(预定不成功返还)
-- ----------------------------
DROP TABLE IF EXISTS `t_marry_pre_wed`;
CREATE TABLE `t_marry_pre_wed` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `period` int(11) NOT NULL COMMENT "场次",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `playerName` varchar(100) NOT NULL COMMENT "玩家名字",
  `peerId` bigint(20) NOT NULL COMMENT "对方id",
  `grade` int(11) NOT NULL COMMENT "酒席档次",
  `hunCheGrade` int(11) NOT NULL COMMENT "婚车档次",
  `sugarGrade` int(11) NOT NULL COMMENT "喜糖档次",
  `status` int(11) NOT NULL COMMENT "状态 1进行中 2失败",
  `holdTime` bigint(20) NOT NULL COMMENT "举办时间",
  `preWedTime` bigint(20) DEFAULT 0 COMMENT "预定时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- xzk create by 2018-11-09
-- Table structure for t_player_tianshu 玩家天书数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tianshu`;
CREATE TABLE `t_player_tianshu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "天书类型",
  `level` int(11) NOT NULL COMMENT "天书等级",
  `isReceive` int(11) NOT NULL COMMENT "是否领取",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
    KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk create by 2018-11-12
-- Table structure for t_player_myboss 玩家个人BOSS数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_myboss`;
CREATE TABLE `t_player_myboss` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `attendMap` varchar(512) NOT NULL COMMENT "参与次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
    KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_addition_sys_slot 玩家附加系统装备槽数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_addition_sys_slot`;
CREATE TABLE `t_player_addition_sys_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `level` int(11) NOT NULL COMMENT "等级",
  `shenZhuLev` int(11) NOT NULL COMMENT "神铸等级",
  `shenZhuNum` int(11) NOT NULL COMMENT "神铸次数",
  `shenZhuPro` int(11) NOT NULL COMMENT "神铸进度",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- ylz create by 2018-11-14
-- Table structure for t_player_system_skill 玩家系统技能
-- ----------------------------
DROP TABLE IF EXISTS `t_player_system_skill`;
CREATE TABLE `t_player_system_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "系统技能类型",
  `subType` int(11) NOT NULL COMMENT "技能类型",
  `level` int(11) NOT NULL COMMENT "等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- xzk create by 2018-11-17
-- Table structure for t_player_activity_num_record  玩家活动抽奖次数
-- ----------------------------
DROP TABLE IF EXISTS `t_player_activity_num_record`;
CREATE TABLE `t_player_activity_num_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `times` int(11) NOT NULL COMMENT "次数",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime`   bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- ylz create by 2018-11-19
-- Table structure for t_player_foe  玩家仇人列表
-- ----------------------------
DROP TABLE IF EXISTS `t_player_foe`;
CREATE TABLE `t_player_foe` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `foeId` bigint(20) NOT NULL COMMENT "仇人id",
  `killTime`   bigint(20) DEFAULT 0 COMMENT "击杀时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_open_activity_rewards_limit 活动奖励次数限制数据
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_rewards_limit`;
CREATE TABLE `t_open_activity_rewards_limit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",
  `timesMap` varchar(512) NOT NULL COMMENT "领奖次数map",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime`   bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk create by 2018-11-21
-- Table structure for t_open_activity_discount_limit 折扣商店次数限制数据
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_discount_limit`;
CREATE TABLE `t_open_activity_discount_limit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",
  `discountDay` int(11) NOT NULL COMMENT "折扣日",
  `timesMap` varchar(1024) NOT NULL COMMENT "购买次数map",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime`   bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;




-- ----------------------------
-- ylz create by 2018-11-27
-- Table structure for t_player_fabao_other  非进阶法宝
-- ----------------------------
DROP TABLE IF EXISTS `t_player_fabao_other`;
CREATE TABLE `t_player_fabao_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "法宝类型",
  `faBaoId` int(11) NOT NULL COMMENT "法宝皮肤id",
  `level` int(11) NOT NULL COMMENT "升星等级",
  `upNum` int(11) NOT NULL COMMENT "升星次数",
  `upPro` int(11) NOT NULL COMMENT "升星培养值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- create time:20181127
-- author：ylz
-- Table structure for t_player_fabao 玩家法宝数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_fabao`;
CREATE TABLE `t_player_fabao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `faBaoId` int(11) NOT NULL COMMENT "当前法宝id",
  `unrealLevel` int(11) NOT NULL COMMENT "法宝幻化丹食丹等级",
  `unrealNum` int(11) NOT NULL COMMENT "法宝幻化丹次数",
  `unrealPro` int(11) NOT NULL COMMENT "法宝幻化丹培养进度值",
  `unrealInfo` varchar(256) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `tonglingLevel` int(11) NOT NULL COMMENT "通灵等级",
  `tonglingNum` int(11) NOT NULL COMMENT "通灵次数",
  `tonglingPro` int(11) NOT NULL COMMENT "通灵进度值",
  `hidden` int(11) default 0 COMMENT "是否隐藏法宝",
  `power` bigint(20) NOT NULL COMMENT "法宝战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- xzk create by 2018-11-26
-- Table structure for t_player_unreal_boss 玩家幻境boss数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_unreal_boss`;
CREATE TABLE `t_player_unreal_boss` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `pilaoNum` int(11) NOT NULL COMMENT "疲劳值",
  `buyPiLaoNum` int(11) NOT NULL COMMENT "购买疲劳值",
  `buyPiLaoTimes` int(11) NOT NULL COMMENT "购买次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- cjb create by 2018-11-27
-- Table structure for t_player_addition_sys_level 玩家附加系统等级
-- ----------------------------
DROP TABLE IF EXISTS `t_player_addition_sys_level`;
CREATE TABLE `t_player_addition_sys_level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型",
  `level` int(11) NOT NULL COMMENT "系统等级",
  `upNum` int(11) NOT NULL COMMENT "等级",
  `upPro` int(11) NOT NULL COMMENT "等级",
  `lingLevel` int(11) NOT NULL COMMENT "化灵等级",
  `lingNum` int(11) NOT NULL COMMENT "化灵次数",
  `lingPro` int(11) NOT NULL COMMENT "化灵进度",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- create time:20181128
-- author：ylz
-- Table structure for t_player_xuedun 玩家血盾数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_xuedun`;
CREATE TABLE `t_player_xuedun` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `blood` bigint(20) NOT NULL COMMENT "玩家血炼值",
  `number` int(11) NOT NULL COMMENT "血盾阶别",
  `star` int(11) NOT NULL COMMENT "血盾星级",
  `starNum` int(11) NOT NULL COMMENT "升星次数",
  `starPro` int(11) NOT NULL COMMENT "升星进度值",
  `culLevel` int(11) NOT NULL COMMENT "培养等级",
  `culNum` int(11) NOT NULL COMMENT "培养次数",
  `culPro` int(11) NOT NULL COMMENT "培养进度值",
  `power` bigint(20) NOT NULL COMMENT "血盾战力",
  `isActive` int(11) NOT NULL COMMENT "是否已激活",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk create by 2018-11-28
-- Table structure for t_player_unreal_boss 玩家幻境boss数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_material`;
CREATE TABLE `t_player_material` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `materialType` int(11) NOT NULL COMMENT "材料副本类型",
  `useTimes` int(11) NOT NULL COMMENT "挑战次数",
  `group` int(11) NOT NULL COMMENT "波数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- create time:20181129
-- author：ylz
-- Table structure for t_player_liveness 玩家活跃度数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_liveness`;
CREATE TABLE `t_player_liveness` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `liveness` bigint(20) NOT NULL COMMENT "玩家活跃度",
  `openBoxs` varchar(100) NOT NULL COMMENT "宝箱开启",
  `lastTime` bigint(20) NOT NULL COMMENT "操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create time:20181129
-- author：ylz
-- Table structure for t_player_liveness_quest 玩家活跃度任务数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_liveness_quest`;
CREATE TABLE `t_player_liveness_quest` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `questId` int(11) NOT NULL COMMENT "活跃度任务id",
  `num` int(11) NOT NULL COMMENT "活跃度任务完成次数",
  `lastTime` bigint(20) NOT NULL COMMENT "操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create time:20181129
-- author：cjb
-- Table structure for t_player_xianti 玩家仙体数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_xianti`;
CREATE TABLE `t_player_xianti` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `xianTiId` int(11) NOT NULL COMMENT "当前仙体id",
  `unrealLevel` int(11) NOT NULL COMMENT "食幻化丹等级",
  `unrealNum` int(11) NOT NULL COMMENT "食幻化丹次数",
  `unrealPro` int(11) NOT NULL COMMENT "食幻化丹进度值",
  `unrealInfo` varchar(256) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `hidden` int(11) default 0 COMMENT "是否隐藏仙体",
  `power` bigint(20) NOT NULL COMMENT "仙体战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create time:20181129
-- author：cjb
-- Table structure for t_player_xianti_other  非进阶仙体
-- ----------------------------
DROP TABLE IF EXISTS `t_player_xianti_other`;
CREATE TABLE `t_player_xianti_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "仙体类型",
  `xianTiId` int(11) NOT NULL COMMENT "仙体皮肤id",
  `level` int(11) NOT NULL COMMENT "升星等级",
  `upNum` int(11) NOT NULL COMMENT "升星次数",
  `upPro` int(11) NOT NULL COMMENT "升星培养值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_cycle_charge_record  玩家每日充值记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_cycle_charge_record`;
CREATE TABLE `t_player_cycle_charge_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chargeNum`  bigint(20) NOT NULL COMMENT "充值元宝数",
  `preDayChargeNum` bigint(20) NOT NULL COMMENT "前一天充值数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_foe_protect  玩家仇人反馈保护
-- ----------------------------
DROP TABLE IF EXISTS `t_player_foe_protect`;
CREATE TABLE `t_player_foe_protect` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `expireTime`  bigint(20) NOT NULL COMMENT "保护过期时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk 2018-12-03
-- Table structure for t_player_foe_feedback  玩家仇人反馈
-- ----------------------------
DROP TABLE IF EXISTS `t_player_foe_feedback`;
CREATE TABLE `t_player_foe_feedback` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `isProtected`  int(11) NOT NULL COMMENT "是否保护",
  `feedbackName`  varchar(20) NOT NULL COMMENT "反馈玩家名称",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk 2018-12-04
-- Table structure for t_player_friend_feedback  玩家好友赞赏 
-- ----------------------------
DROP TABLE IF EXISTS `t_player_friend_feedback`;
CREATE TABLE `t_player_friend_feedback` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `friendId` bigint(20) NOT NULL COMMENT "赞赏玩家id",
  `friendName` varchar(20) NOT NULL COMMENT "赞赏玩家名称",
  `noticeType` int(11) NOT NULL COMMENT "消息类型",
  `feedbackType` int(11) NOT NULL COMMENT "赞赏类型",
   `condition` int(11) NOT NULL COMMENT "条件",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xzk 2018-12-04
-- Table structure for t_player_friend_add_rew  玩家添加好友奖励 
-- ----------------------------
DROP TABLE IF EXISTS `t_player_friend_add_rew`;
CREATE TABLE `t_player_friend_add_rew` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `rewRecord` varchar(512) NOT NULL COMMENT "奖励记录",
  `frDummyNum` int(11) NOT NULL COMMENT "虚拟好友数量",
  `lastAddDummyTime` bigint(20) NOT NULL COMMENT "上次增加虚拟好友时间",
  `congratulateTimes` int(11) NOT NULL COMMENT "被祝贺次数",
  `lastCongratulateTime` bigint(20) NOT NULL COMMENT "上次被祝贺时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- ylz create by 2018-12-06
-- Table structure for t_player_daily 玩家日环任务
-- ----------------------------
DROP TABLE IF EXISTS `t_player_daily`;
CREATE TABLE `t_player_daily` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `seqId` int(11) NOT NULL COMMENT "流水id",
  `dailyTag` int(11) NOT NULL COMMENT "日环类型 1日环 2仙盟日常",
  `times` int(11) NOT NULL COMMENT "日环次数",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "操作时间",
  `crossDayTime` bigint(20) DEFAULT 0 COMMENT "跨5点记录时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- ylz create by 2018-12-07
-- Table structure for t_player_bagua 玩家八挂秘境数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_bagua`;
CREATE TABLE `t_player_bagua` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `level` int(11) NOT NULL COMMENT "八卦秘境等级",
  `isBuChang` int(11) NOT NULL COMMENT "八卦秘境补偿",
  `inviteTime` bigint(20) DEFAULT 0 COMMENT "邀请时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- cjb create by 2018-12-10
-- Table structure for t_player_outland_boss 玩家外域boss数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_outland_boss`;
CREATE TABLE `t_player_outland_boss` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `zhuoqiNum` int(11) NOT NULL COMMENT "浊气值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- cjb create by 2018-12-10
-- Table structure for t_outland_boss_drop_records 玩家外域boss掉落记录数据
-- ----------------------------
DROP TABLE IF EXISTS `t_outland_boss_drop_records`;
CREATE TABLE `t_outland_boss_drop_records` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `killerName` varchar(50) NOT NULL COMMENT "击杀者",
  `biologyId` int(11) NOT NULL COMMENT "生物id",
  `mapId` int(11) NOT NULL COMMENT "地图id",
  `dropTime` bigint(20) NOT NULL COMMENT "掉落时间",
  `itemInfo` varchar(256) NOT NULL COMMENT "物品信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xzk 2018-12-10
-- Table structure for t_player_open_activity_mail  玩家运营活动开启邮件数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_open_activity_mail`;
CREATE TABLE `t_player_open_activity_mail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `group` int(11) NOT NULL COMMENT "活动id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ---------------------------
-- create by ylz 2018-12-11
-- Table structure for t_player_yuanbao_songbuting  玩家元宝送不停
-- ----------------------------
DROP TABLE IF EXISTS `t_player_yuanbao_songbuting`;
CREATE TABLE `t_player_yuanbao_songbuting` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `isReceive` int(11) NOT NULL COMMENT "是否能领取 0不能 1可以",
  `times` int(11) NOT NULL COMMENT "领取次数",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_open_activity_laba_log  元宝拉霸日志
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_laba_log`;
CREATE TABLE `t_open_activity_laba_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",
  `playerName` varchar(20) NOT NULL COMMENT "玩家姓名",
  `costGold` int(11) NOT NULL COMMENT "花费元宝",
  `rewGold` int(11) NOT NULL COMMENT "奖励元宝", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- ylz create by 2018-12-12
-- Table structure for t_player_liveness_cross_five 玩家活跃度跨5点记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_liveness_cross_five`;
CREATE TABLE `t_player_liveness_cross_five` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `crossDayTime` bigint(20) DEFAULT 0 COMMENT "跨5点记录时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by xzk 2018-12-15
-- Table structure for t_open_activity_start_mail  运营活动开启邮件通知数据
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_start_mail`;
CREATE TABLE `t_open_activity_start_mail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",  
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- ylz create by 2018-12-16
-- Table structure for t_player_team_copy 玩家组队副本数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_team_copy`;
CREATE TABLE `t_player_team_copy` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `purpose` int(11) DEFAULT 0 COMMENT "组队副本标识",
  `num` int(11) NOT NULL COMMENT "奖励次数", 
  `rewTime` bigint(20) DEFAULT 0 COMMENT "奖励时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- cjb create by 2018-12-21
-- Table structure for t_quiz 仙尊问答数据
-- ----------------------------
DROP TABLE IF EXISTS `t_quiz`;
CREATE TABLE `t_quiz` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `answerList` varchar(50) NOT NULL COMMENT "最新答案排序",
  `lastQuizId` int(11) NOT NULL COMMENT "最新问题id",
  `lastQuizTime` bigint(20) NOT NULL COMMENT "最新出题时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- ylz create by 2018-12-25
-- Table structure for t_player_dense_wat 玩家金银密窟数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_dense_wat`;
CREATE TABLE `t_player_dense_wat` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `num` int(11) NOT NULL COMMENT "采集次数", 
  `endTime` bigint(20) DEFAULT 0 COMMENT "结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by xzk 2018-12-25
-- Table structure for t_open_activity_drew_log  运营活动-抽奖日志
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_drew_log`;
CREATE TABLE `t_open_activity_drew_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",
  `playerName` varchar(20) NOT NULL COMMENT "玩家姓名",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create time:20181227
-- author：cjb
-- Table structure for t_player_dianxing 玩家点星系统数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_dianxing`;
CREATE TABLE `t_player_dianxing` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `currType` int(11) NOT NULL COMMENT "点星星谱",
  `currLevel` int(11) NOT NULL COMMENT "点星等级",
  `dianXingTimes` int(11) NOT NULL COMMENT "点星升级次数",
  `dianXingBless` int(11) NOT NULL COMMENT "点星升级进度值",
  `dianXingBlessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `xingChenNum` bigint(20) NOT NULL COMMENT "星尘值",
  `jieFengLev` int(11) NOT NULL COMMENT "点星解封等级",
  `jieFengTimes` int(11) NOT NULL COMMENT "点星解封次数",
  `jieFengBless` int(11) NOT NULL COMMENT "点星解封进度值",
  `power` bigint(20) NOT NULL COMMENT "点星战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- ylz create by 2019-01-02
-- Table structure for t_player_wardrobe 玩家衣橱数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_wardrobe`;
CREATE TABLE `t_player_wardrobe` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL  COMMENT "套装类型",
  `subType` int(11) NOT NULL COMMENT "套装子类",
  `activeFlag` int(11) DEFAULT 0 COMMENT "是否激活 0失效 1激活",
  `permanent` int(11) NOT NULL COMMENT "是否永久 0带时效  1永久",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- ylz create by 2019-01-04
-- Table structure for t_player_wardrobe_peiyang 玩家衣橱套装资质丹数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_wardrobe_peiyang`;
CREATE TABLE `t_player_wardrobe_peiyang` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL  COMMENT "套装类型",
  `peiYangNum` int(11) NOT NULL COMMENT "资质丹数量",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by xzk 2019-01-03
-- Table structure for t_player_tianmo 玩家天魔数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tianmo`;
CREATE TABLE `t_player_tianmo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `danLevel` int(11) NOT NULL COMMENT "天魔丹食用等级",
  `danNum` int(11) NOT NULL COMMENT "天魔丹培养次数",
  `danPro` int(11) NOT NULL COMMENT "天魔丹培养进度值",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值", 
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `chargeVal` bigint(20) NOT NULL COMMENT "模块开启累计充值",
  `power` bigint(20) NOT NULL COMMENT "天魔体战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`), 
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by cjb 2019-01-03
-- Table structure for t_player_shihunfan 玩家噬魂幡数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shihunfan`;
CREATE TABLE `t_player_shihunfan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `danLevel` int(11) NOT NULL COMMENT "食用等级",
  `danNum` int(11) NOT NULL COMMENT "培养次数",
  `danPro` int(11) NOT NULL COMMENT "培养进度值",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值", 
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `chargeVal` int(11) NOT NULL COMMENT "模块开启累计充值",
  `power` bigint(20) NOT NULL COMMENT "战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by cjb 2019-01-08
-- Table structure for t_open_activity_crazybox_log  运营活动-疯狂宝箱日志
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_crazybox_log`;
CREATE TABLE `t_open_activity_crazybox_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL COMMENT "活动id",
  `playerName` varchar(20) NOT NULL COMMENT "玩家姓名",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


  -- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_info 玩家灵童信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingtong_info`;
CREATE TABLE `t_player_lingtong_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingTongId` int(11) NOT NULL COMMENT "灵童id",
  `lingTongName` varchar(50) NOT NULL COMMENT "灵童名字",
  `upgradeLevel` int(11) NOT NULL COMMENT "升级等级",
  `upgradeNum` int(11) NOT NULL COMMENT "升级次数",
  `upgradePro` int(11) NOT NULL COMMENT "升级进度值",
  `peiYangLevel` int(11) NOT NULL COMMENT "培养等级",
  `peiYangNum`  int(11) NOT NULL COMMENT "培养次数", 
  `peiYangPro`  bigint(20) NOT NULL COMMENT "培养进度值",
  `starLevel` int(11) NOT NULL COMMENT "升星等级",
  `starNum`  int(11) NOT NULL COMMENT "升星次数", 
  `starPro`  int(11) NOT NULL COMMENT "升星进度值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong 玩家灵童
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingtong`;
CREATE TABLE `t_player_lingtong` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingTongId` int(11) NOT NULL COMMENT "出战灵童id",
  `level` int(11) NOT NULL COMMENT "等级", 
  `basePower`  bigint(20) NOT NULL COMMENT "基础战力",
  `power`  bigint(20) NOT NULL COMMENT "战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_fashion 玩家灵童时装信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingtong_fashion`;
CREATE TABLE `t_player_lingtong_fashion` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingTongId` int(11) NOT NULL COMMENT "灵童id",
  `fashionId` int(11) NOT NULL COMMENT "灵童时装id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_fashion_info 玩家灵童时装信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingtong_fashion_info`;
CREATE TABLE `t_player_lingtong_fashion_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `fashionId` int(11) NOT NULL COMMENT "灵童时装id",
  `upgradeLevel` int(11) NOT NULL COMMENT "升级等级",
  `upgradeNum` int(11) NOT NULL COMMENT "升级次数",
  `upgradePro` int(11) NOT NULL COMMENT "升级进度值",
  `isExpire` int(11) NOT NULL COMMENT "是否失效",     
	`activateTime` bigint(20) DEFAULT 0 COMMENT "激活时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_develop 玩家灵童养成信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingtong_develop`;
CREATE TABLE `t_player_lingtong_develop` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `classType` int(11) NOT NULL COMMENT "养成系统类型",
  `advancedId` int(11) NOT NULL COMMENT "进阶id",
  `seqId` int(11) NOT NULL COMMENT "当前外观id",
  `unrealLevel` int(11) NOT NULL COMMENT "食幻化丹等级",
  `unrealNum` int(11) NOT NULL COMMENT "食幻化丹次数",
  `unrealPro` int(11) NOT NULL COMMENT "食幻化丹进度值",
  `culLevel`  int(11) NOT NULL COMMENT "食培养丹等级",
  `culNum` int(11) NOT NULL COMMENT "食培养丹次数",
  `culPro` int(11) NOT NULL COMMENT "食培养丹进度值",
  `tongLingLevel` int(11) NOT NULL COMMENT "食通灵丹等级",
  `tongLingNum` int(11) NOT NULL COMMENT "食通灵丹次数",
  `tongLingPro` int(11) NOT NULL COMMENT "食通灵丹进度值",
  `unrealInfo` varchar(256) NOT NULL  COMMENT "解锁幻化信息",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数进阶次数",
  `bless`  int(11) NOT NULL COMMENT "当前祝福值",
  `blessTime`  bigint(20) NOT NULL COMMENT "祝福值开始时间",
  `hidden` int(11) default 0 COMMENT "是否隐藏外观",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_other 玩家灵童养成非进阶信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingtong_other`;
CREATE TABLE `t_player_lingtong_other` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `classType` int(11) NOT NULL COMMENT "养成系统类型",
  `type` int(11) NOT NULL COMMENT "皮肤类型",
  `seqId` int(11) NOT NULL COMMENT "皮肤id",
  `level` int(11) NOT NULL COMMENT "升星等级",
  `upNum` int(11) NOT NULL COMMENT "升星次数",
  `upPro` int(11) NOT NULL COMMENT "升星培养值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-07
-- Table structure for t_player_lingtong_power 玩家灵童养成战力信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingtong_power`;
CREATE TABLE `t_player_lingtong_power` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `classType` int(11) NOT NULL COMMENT "养成系统类型",
  `power`  bigint(20) NOT NULL COMMENT "战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-09
-- Table structure for t_player_lingtong_fashion_trial  玩家灵童时装试用卡阶数
-- ----------------------------
DROP TABLE IF EXISTS `t_player_lingtong_fashion_trial`;
CREATE TABLE `t_player_lingtong_fashion_trial` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `trialFashionId`  int(11) NOT NULL COMMENT "时装id",
  `isExpire` int(11) NOT NULL COMMENT "是否失效 0否 1失效", 
  `activateTime` bigint(20) DEFAULT 0 COMMENT "激活时间",
  `durationTime` bigint(20) DEFAULT 0 COMMENT "持续时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by xzk 2019-09-11
-- Table structure for t_open_activity_boss_kill  运营活动BOSS首杀
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_boss_kill`;
CREATE TABLE `t_open_activity_boss_kill` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `bossIdList` varchar(512) NOT NULL COMMENT "bossId列表",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_cycle_cost_record  玩家每日消费记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_cycle_cost_record`;
CREATE TABLE `t_player_cycle_cost_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `costNum`  bigint(20) NOT NULL COMMENT "消费元宝数",
  `preDayCostNum`  bigint(20) NOT NULL COMMENT "上一天消费元宝数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_open_activity_alliance_cheer  城战助威记录
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_alliance_cheer`;
CREATE TABLE `t_open_activity_alliance_cheer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `groupId` int(11) NOT NULL  COMMENT "活动Id",
  `allianceId`  bigint(20) NOT NULL COMMENT "仙盟id",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)  
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_fei_sheng  玩家飞升表
-- ----------------------------
DROP TABLE IF EXISTS `t_player_fei_sheng`;
CREATE TABLE `t_player_fei_sheng` ( 
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `feiLevel`  int(11) NOT NULL COMMENT "飞升等级",
  `addRate`  int(11) NOT NULL COMMENT "增加的成功率",
  `gongDeNum`  bigint(20) NOT NULL COMMENT "功德值",
  `leftPotential`  int(11) NOT NULL COMMENT "剩余潜能点",
  `tiZhi`  int(11) NOT NULL COMMENT "体质点",
  `liDao`  int(11) NOT NULL COMMENT "力道点",
  `jinGu`  int(11) NOT NULL COMMENT "筋骨点",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by ylz 2019-01-22
-- Table structure for t_player_kaifumubiao 玩家开服目标信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_kaifumubiao`;
CREATE TABLE `t_player_kaifumubiao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `kaiFuDay` int(11) NOT NULL COMMENT "开服解锁时间",
  `finishNum` int(11) NOT NULL COMMENT "完成任务数",
  `isReward` int(11) NOT NULL COMMENT "是否领取过组奖励",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


 -- ----------------------------
-- create by ylz 2019-01-22
-- Table structure for t_player_quest_crossday 玩家任务跨天信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_quest_crossday`;
CREATE TABLE `t_player_quest_crossday` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `crossDayTime` bigint(20) DEFAULT 0 COMMENT "跨天时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- create by ylz 2019-01-24
-- Table structure for t_player_shenmo 玩家神魔数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenmo`;
CREATE TABLE `t_player_shenmo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `gongXunNum` int(11) NOT NULL COMMENT "玩家功勋数",
  `killNum` int(11) NOT NULL COMMENT "本次击杀",
  `endTime` bigint(20) DEFAULT 0 COMMENT "结束时间",
  `rewTime` bigint(20) DEFAULT 0 COMMENT "领取奖励的排行榜时间戳",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by cjb 2019-02-11
-- Table structure for t_hongbao 红包数据
-- ----------------------------
DROP TABLE IF EXISTS `t_hongbao`;
CREATE TABLE `t_hongbao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `hongBaoType` int(11) NOT NULL COMMENT "红包类型",
  `sendId` bigint(20) DEFAULT 0 COMMENT "发红包玩家id",
  `awardList` text(2000) NOT NULL COMMENT "红包奖励列表",
  `snatchLog` text(2000) NOT NULL COMMENT "红包领取记录",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by cjb 2019-02-11
-- Table structure for t_player_hongbao 玩家红包数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_hongbao`;
CREATE TABLE `t_player_hongbao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `snatchCount` int(11) NOT NULL COMMENT "抢红包次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by cjb 2019-02-14
-- Table structure for t_player_chat 玩家聊天数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_chat`;
CREATE TABLE `t_player_chat` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chatCount` int(11) NOT NULL COMMENT "聊天次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by zrc 2019-02-19
-- Table structure for t_player_activity_pk 玩家活动pk数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_activity_pk`;
CREATE TABLE `t_player_activity_pk` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType` int(11) NOT NULL COMMENT "活动类型",
  `killedNum` int(11) NOT NULL COMMENT "被杀数",
  `lastKilledTime` bigint(20) NOT NULL COMMENT "上次被杀时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by xzk 2019-02-25
-- Table structure for t_player_friend_add_rew  玩家赞赏 
-- ----------------------------
DROP TABLE IF EXISTS `t_player_friend_admire`;
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


-- ----------------------------
-- create by xzk 2019-02-26
-- Table structure for t_player_open_activity_xun_huan 玩家循环活动数据
-- ----------------------------
DROP TABLE IF EXISTS `t_open_activity_xun_huan`;
CREATE TABLE `t_open_activity_xun_huan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `arrGroup` int(11) NOT NULL COMMENT "随机活动组",
  `activityDay` int(20) NOT NULL COMMENT "当前活动天",
  `startTime` bigint(20) DEFAULT 0 COMMENT "活动开始时间",
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_supreme_title_wear 玩家穿戴至尊称号
-- ----------------------------
DROP TABLE IF EXISTS `t_player_supreme_title_wear`;
CREATE TABLE `t_player_supreme_title_wear` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `titleWear` int(11) NOT NULL COMMENT "穿戴称号id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_supreme_title 玩家至尊称号数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_supreme_title`;
CREATE TABLE `t_player_supreme_title` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `titleId` int(11) NOT NULL COMMENT "称号id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
  

-- Table structure for t_player_equipbaoku 玩家宝库
-- ----------------------------
DROP TABLE IF EXISTS `t_player_equipbaoku`;
CREATE TABLE `t_player_equipbaoku` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `luckyPoints` int(11) NOT NULL COMMENT "幸运值",
  `typ` int(11) NOT NULL COMMENT "宝库类型",
  `attendPoints` int(11) NOT NULL COMMENT "积分",
  `totalAttendTimes` int(11) NOT NULL COMMENT "总参与次数",
  `lastSystemRefreshTime` bigint(20) DEFAULT 0 COMMENT "上次自动刷新时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_equipbaoku_shop 玩家当日宝库商店购买道具(限购使用)
-- ----------------------------
DROP TABLE IF EXISTS `t_player_equipbaoku_shop`;
CREATE TABLE `t_player_equipbaoku_shop` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `shopId` int(11) NOT NULL COMMENT "道具shopId",
  `dayCount` int(11) NOT NULL COMMENT "购买次数",
  `lastTime` bigint(20) NOT NULL COMMENT "最后一次购买时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance_boss  仙盟boss
-- ----------------------------
DROP TABLE IF EXISTS `t_alliance_boss`;
CREATE TABLE `t_alliance_boss` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `allianceId` bigint(20) NOT NULL COMMENT "仙盟id",
  `summonTime` bigint(20) NOT NULL COMMENT "召唤时间",
  `bossLevel` int(11) NOT NULL COMMENT "boss等级",
  `bossExp` int(11) NOT NULL COMMENT "boss经验",
  `isSummon` int(11) NOT NULL COMMENT "当日是否已经召唤过",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_activity_rank 活动排行数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_activity_rank`;
CREATE TABLE `t_player_activity_rank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType` int(11) NOT NULL COMMENT "活动类型",
  `rankMap` varchar(500) NOT NULL COMMENT "排行数据",
  `endTime` bigint(20) NOT NULL COMMENT "结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_mingge_pan 玩家命盘
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mingge_pan`;
CREATE TABLE `t_player_mingge_pan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "命格类型 0 普通 1超级",
  `subType` int(11) NOT NULL COMMENT "命格子类型",
  `itemList` varchar(500) NOT NULL COMMENT "命格信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_mingge_pan_refined 玩家命盘祭炼
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mingge_pan_refined`;
CREATE TABLE `t_player_mingge_pan_refined` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `subType` int(11) NOT NULL COMMENT "命格子类型",
  `number` int(11) NOT NULL COMMENT "阶数",
  `star` int(11) NOT NULL COMMENT "星数",
  `refinedNum` int(11) NOT NULL COMMENT "祭炼次数",
  `refinedPro` int(11) NOT NULL COMMENT "祭炼进度值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_mingge_mingli 玩家命理
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mingge_mingli`;
CREATE TABLE `t_player_mingge_mingli` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "命宫类型",
  `subType` int(11) NOT NULL COMMENT "命理部位",
  `mingLiList` varchar(500) NOT NULL COMMENT "命理信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_tulong_equip_slot 玩家屠龙装备槽数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tulong_equip_slot`;
CREATE TABLE `t_player_tulong_equip_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `suitType` int(11) NOT NULL COMMENT "套装类型",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `level` int(11) NOT NULL COMMENT "等级",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `gemInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "宝石信息",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_tulong_suit_skill 玩家屠龙套装技能数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tulong_suit_skill`;
CREATE TABLE `t_player_tulong_suit_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `suitType` int(11) NOT NULL COMMENT "套装类型",
  `level` int(11) NOT NULL COMMENT "技能等级", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_shenqi_debris 玩家神器碎片
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenqi_debris`;
CREATE TABLE `t_player_shenqi_debris` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `shenQiType` int(11) NOT NULL COMMENT "神器类型",
  `slotId` int(11) NOT NULL COMMENT "部位",
  `level` int(11) NOT NULL COMMENT "等级",
  `upNum` int(11) NOT NULL COMMENT "升级次数",
  `upPro` int(11) NOT NULL COMMENT "升级进度",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_shenqi_smelt 玩家神器淬炼
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenqi_smelt`;
CREATE TABLE `t_player_shenqi_smelt` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `shenQiType` int(11) NOT NULL COMMENT "神器类型",
  `slotId` int(11) NOT NULL COMMENT "部位",
  `level` int(11) NOT NULL COMMENT "等级",
  `upNum` int(11) NOT NULL COMMENT "升级次数",
  `upPro` int(11) NOT NULL COMMENT "升级进度",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_shenqi_qiling 玩家神器器灵
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenqi_qiling`;
CREATE TABLE `t_player_shenqi_qiling` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `shenQiType` int(11) NOT NULL COMMENT "神器类型",
  `qiLingType` int(11) NOT NULL COMMENT "器灵类型",
  `slotId` int(11) NOT NULL COMMENT "部位",
  `level` int(11) NOT NULL COMMENT "等级",
  `upNum` int(11) NOT NULL COMMENT "升级次数",
  `upPro` int(11) NOT NULL COMMENT "升级进度",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_shenqi 玩家神器
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenqi`;
CREATE TABLE `t_player_shenqi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingQiNum` bigint(20) NOT NULL COMMENT "灵气值",
  `power`  bigint(20) NOT NULL COMMENT "战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by xzk 2019-03-08
-- Table structure for t_player_hunt 玩家寻宝数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_hunt`;
CREATE TABLE `t_player_hunt` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `huntType` int(11) NOT NULL COMMENT "寻宝类型",
  `freeHuntCount` int(11) NOT NULL COMMENT "免费寻宝次数", 
  `totalHuntCount` int(11) NOT NULL COMMENT "寻宝总次数", 
  `lastHuntTime` bigint(20) NOT NULL COMMENT "上次寻宝时间",  
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`)  
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_player_zhenfa 玩家阵法
-- ----------------------------
DROP TABLE IF EXISTS `t_player_zhenfa`;
CREATE TABLE `t_player_zhenfa` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "阵法类型",
  `level` int(11) NOT NULL COMMENT "阵法等级",
  `levelNum` int(11) NOT NULL COMMENT "升级次数",
  `levelPro` int(11) NOT NULL COMMENT "升级进度值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_zhenqi 玩家阵旗
-- ----------------------------
DROP TABLE IF EXISTS `t_player_zhenqi`;
CREATE TABLE `t_player_zhenqi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "阵法类型",
  `zhenQiPos` int(11) NOT NULL COMMENT "阵旗部位",
  `number` int(11) NOT NULL COMMENT "阵旗阶数",
  `numberNum` int(11) NOT NULL COMMENT "升阶次数",
  `numberPro` int(11) NOT NULL COMMENT "升阶进度值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_player_zhenqi_xianhuo 玩家阵旗仙火
-- ----------------------------
DROP TABLE IF EXISTS `t_player_zhenqi_xianhuo`;
CREATE TABLE `t_player_zhenqi_xianhuo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "阵法类型",
  `level` int(11) NOT NULL COMMENT "级数",
  `luckyStar` int(11) NOT NULL COMMENT "暴击幸运星",
  `levelNum` int(11) NOT NULL COMMENT "升级次数",
  `levelPro` int(11) NOT NULL COMMENT "升级进度值",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_yinglingpu 玩家英灵谱
-- ----------------------------
DROP TABLE IF EXISTS `t_player_yinglingpu`;
CREATE TABLE `t_player_yinglingpu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `tuJianId` int(11) NOT NULL COMMENT "图鉴id",
  `tuJianType` int(11) NOT NULL COMMENT "图鉴类型",
  `level` int(11) NOT NULL COMMENT "等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_yinglingpu_suipian 英灵谱碎片
-- ----------------------------
DROP TABLE IF EXISTS `t_player_yinglingpu_suipian`;
CREATE TABLE `t_player_yinglingpu_suipian` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `tuJianId` int(11) NOT NULL COMMENT "图鉴id",
  `tuJianType` int(11) NOT NULL COMMENT "图鉴类型",
  `suiPianId` int(11) NOT NULL COMMENT "碎片id",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
--  create by xzk 2019-03-18
-- Table structure for t_player_pregnant 玩家怀孕
-- ----------------------------
DROP TABLE IF EXISTS `t_player_pregnant`;
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
DROP TABLE IF EXISTS `t_player_baby`;
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
  `costItemNum` int(11) NOT NULL COMMENT "洗练消耗道具数量", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_trade_item 交易物品
-- ----------------------------
DROP TABLE IF EXISTS `t_trade_item`;
CREATE TABLE `t_trade_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `originServerId`  int(11) NOT NULL COMMENT "初始服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11)  NOT NULL COMMENT "物品数量",
  `level` int(11) NOT NULL COMMENT "等级",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `gold` bigint(20) NOT NULL COMMENT "价格",
  `status` int(11) NOT NULL COMMENT "状态",
  `system` int(11) NOT NULL COMMENT "系统下架",
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
DROP TABLE IF EXISTS `t_trade_order`;
CREATE TABLE `t_trade_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `buyServerId` int(11) NOT NULL COMMENT "购买服务器id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `playerName` varchar(100)  NOT NULL COMMENT "买家名字", 
  `tradeId` bigint(20) NOT NULL COMMENT "商品id",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11)  NOT NULL COMMENT "物品数量",
  `level` int(11) NOT NULL COMMENT "等级",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `gold` bigint(20) NOT NULL COMMENT "价格",
  `status` int(11) NOT NULL COMMENT "状态0:支付1:发货2:取消",
  `sellPlatform` int(11)  NOT NULL COMMENT "卖家平台",
  `sellServerId` int(11)  NOT NULL COMMENT "卖家服务器id",
  `sellPlayerId` bigint(20)  NOT NULL COMMENT "卖家id",
  `sellPlayerName` varchar(100)  NOT NULL COMMENT "卖家名字", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- cjb create by 2019-3-19
-- Table structure for t_player_activity_add_num  玩家活动增长数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_activity_add_num`;
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


-- ----------------------------
-- Table structure for t_player_baby_toy_slot 玩家宝宝玩具槽数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_baby_toy_slot`;
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




-- ----------------------------
-- Table structure for t_friend_marry_develop_log 全局表白日志记录数据
-- ----------------------------
DROP TABLE IF EXISTS `t_friend_marry_develop_log`;
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
DROP TABLE IF EXISTS `t_player_friend_marry_develop_send_log`;
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
DROP TABLE IF EXISTS `t_player_friend_marry_develop_recv_log`;
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



-- ----------------------------
-- create by xzk 2019-03-26
-- Table structure for t_player_qiyu 玩家奇遇信息
-- ----------------------------
DROP TABLE IF EXISTS `t_player_qiyu`;
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
  `isHadNotice` int(11) NOT NULL COMMENT "是否前置提示过 0否1是",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- create by zrc 2019-04-02
-- Table structure for t_player_marry_jinian 结婚纪念
-- ----------------------------
DROP TABLE IF EXISTS `t_player_marry_jinian`;
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


-- ----------------------------
-- create by zrc 2019-04-02
-- Table structure for t_player_marry_dingqing 结婚定情
-- ----------------------------
DROP TABLE IF EXISTS `t_player_marry_dingqing`;
 CREATE TABLE `t_player_marry_dingqing`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `playerId` bigint(20) NULL DEFAULT NULL COMMENT '玩家Id',
  `suit` text(2000) NOT NULL COMMENT '套装',
  `updateTime` bigint(20) NULL DEFAULT 0 COMMENT '更新时间',
  `createTime` bigint(20) NULL DEFAULT 0 COMMENT '创建时间',
  `deleteTime` bigint(20) NULL DEFAULT 0 COMMENT '删除时间',
PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
)  ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- create by zrc 2019-04-02
-- Table structure for t_player_marry_jinian_sj 纪念时装
-- ----------------------------
DROP TABLE IF EXISTS `t_player_marry_jinian_sj`;
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
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_baby 玩家配偶宝宝数据
-- ----------------------------
DROP TABLE IF EXISTS `t_couple_baby`;
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

-- ----------------------------
--  create by zrc 2019-04-08
-- Table structure for t_player_trade_log 交易日志
-- ----------------------------
DROP TABLE IF EXISTS `t_player_trade_log`;
CREATE TABLE `t_player_trade_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `logType` int(11) NOT NULL COMMENT "日志类型0:出售1:购买",
  `tradeId` bigint(20) NOT NULL COMMENT "交易id",
  `sellServerId` int(11) NOT NULL COMMENT "出售服务器id",
  `sellPlayerId` bigint(20) NOT NULL COMMENT "出售玩家id",
  `sellPlayerName` varchar(100) NOT NULL COMMENT "出售玩家名字",
  `buyServerId` int(11) NOT NULL COMMENT "购买服务器id",
  `buyPlayerId` bigint(20) NOT NULL COMMENT "购买玩家id",
  `buyPlayerName` varchar(100) NOT NULL COMMENT "购买玩家名字",
  `getGold` int(11) NOT NULL COMMENT "获得的元宝",
  `gold` int(11) NOT NULL COMMENT "价格",
  `fee` int(11) NOT NULL COMMENT "手续费",
  `feeRate` int(11) NOT NULL COMMENT "手续费比例",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `itemNum` int(11) NOT NULL COMMENT "物品数量",
  `propertyData` varchar(512) DEFAULT "{}"  COMMENT "属性",
  `level` int(11) NOT NULL COMMENT "等级",
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_shenyu  神域之战
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenyu`;
CREATE TABLE `t_player_shenyu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `keyNum` int(11) NOT NULL COMMENT "钥匙数",
  `round` int(11) NOT NULL COMMENT "参赛轮",
  `exp` bigint(20) NOT NULL COMMENT "获得经验", 
  `itemInfo` varchar(1024) NOT NULL DEFAULT "{}" COMMENT "获得物品",
  `endTime` bigint(20) NOT NULL COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_xian_tao  仙桃大会
-- ----------------------------
DROP TABLE IF EXISTS `t_player_xian_tao`;
CREATE TABLE `t_player_xian_tao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `juniorPeachCount` int(11) NOT NULL COMMENT "百年仙桃数量",
  `highPeachCount` int(11) NOT NULL COMMENT "千年仙桃数量",
  `robCount` int(11) NOT NULL COMMENT "劫取次数",
  `beRobCount` int(11) NOT NULL COMMENT "被劫取次数",
  `endTime` bigint(20) NOT NULL COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_house  玩家房子
-- ----------------------------
DROP TABLE IF EXISTS `t_player_house`;
CREATE TABLE `t_player_house` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `houseIndex` int(11) NOT NULL COMMENT "房子序号",
  `houseType` int(11) NOT NULL COMMENT "房子类型", 
  `level` int(11) NOT NULL COMMENT "当前等级",
  `maxLevel` int(11) NOT NULL COMMENT "历史最高等级",
  `dayTimes` int(11) NOT NULL COMMENT "每日维修次数",
  `isBroken` int(11) NOT NULL COMMENT "是否损坏",
  `lastBrokenTime` bigint(20) NOT NULL COMMENT "上次损坏时间",
  `isRent`  int(11) NOT NULL COMMENT "是否领取租金",
  `rentUpdateTime` bigint(20) NOT NULL COMMENT "租金更新时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_trade_recycle  系统回购
-- ----------------------------
DROP TABLE IF EXISTS `t_trade_recycle`;
CREATE TABLE `t_trade_recycle` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `recycleGold` bigint(20) NOT NULL COMMENT "回收的元宝",
  `recycleTime` bigint(20) NOT NULL COMMENT "回收时间", 
  `customRecycleGold` bigint(20) NOT NULL COMMENT "自定义回收的元宝",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_trade_recycle  个人系统回购
-- ----------------------------
DROP TABLE IF EXISTS `t_player_trade_recycle`;
CREATE TABLE `t_player_trade_recycle` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `recycleGold` bigint(20) NOT NULL COMMENT "回收的元宝",
  `recycleTime` bigint(20) NOT NULL COMMENT "回收时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
   PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_activity_collect  个人采集
-- ----------------------------
DROP TABLE IF EXISTS `t_player_activity_collect`;
CREATE TABLE `t_player_activity_collect` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `activityType` int(11) NOT NULL COMMENT "活动类型", 
  `countMap` varchar(512) NOT NULL COMMENT "采集次数map", 
  `endTime` bigint(20) NOT NULL COMMENT "场景结束时间", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",  
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- xb create by 2019-05-07
-- Table structure for t_player_addition_sys_awake  附加系统觉醒
-- ----------------------------
DROP TABLE IF EXISTS `t_player_addition_sys_awake`;
CREATE TABLE `t_player_addition_sys_awake` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型", 
  `isAwake` int(11) NOT NULL COMMENT "是否觉醒", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",  
  PRIMARY KEY (`id`), 
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- cjb create by 2019-05-07
-- Table structure for t_player_addition_sys_tongling 玩家附加系统通灵数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_addition_sys_tongling`;
CREATE TABLE `t_player_addition_sys_tongling` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型",
  `tongLingLev` int(11) NOT NULL COMMENT "通灵等级",
  `tongLingNum` int(11) NOT NULL COMMENT "通灵次数",
  `tongLingPro` int(11) NOT NULL COMMENT "通灵进度",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 
-- ----------------------------
-- Table structure for t_player_alliance_yuxi  玩家仙盟玉玺之战
-- ----------------------------
DROP TABLE IF EXISTS `t_player_alliance_yuxi`;
CREATE TABLE `t_player_alliance_yuxi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `isReceive` int(11) NOT NULL COMMENT "是否领取每日奖励",  
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`),  
   KEY(`playerId`), 
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
 
-- ----------------------------
-- Table structure for t_activity_end_record  日常活动开启记录
-- ----------------------------
DROP TABLE IF EXISTS `t_activity_end_record`;
CREATE TABLE `t_activity_end_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
  `activityType` int(11) NOT NULL COMMENT "活动类型", 
  `endTime` bigint(20) DEFAULT 0 COMMENT "活动结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- create by zrc 2019-06-04
-- ----------------------------
-- Table structure for t_player_week 玩家周卡数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_week`;
CREATE TABLE `t_player_week` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `seniorExpireTime` bigint(20) DEFAULT 0  COMMENT "高级周卡过期时间",
  `seniorLastDayRewTime` bigint(20) DEFAULT 0  COMMENT "高级每日奖励领取时间",
  `seniorCycDay` int(11) NOT NULL COMMENT "高级循环领取天数",
  `juniorExpireTime` bigint(20) DEFAULT 0  COMMENT "初级周卡过期时间",
  `juniorLastDayRewTime` bigint(20) DEFAULT 0  COMMENT "初级每日奖励领取时间",
  `juniorCycDay` int(11) NOT NULL COMMENT "初级循环领取天数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-06-11
-- ----------------------------
-- Table structure for t_player_dushi 玩家八卦符石数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_fushi`;
CREATE TABLE `t_player_fushi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "符石索引",
  `fushiLevel` int(11) NOT NULL COMMENT "符石等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by zrc 2019-06-11
-- ----------------------------
-- Table structure for t_shenmo_rank_time 神魔战场排行榜
-- ----------------------------
DROP TABLE IF EXISTS `t_shenmo_rank_time`;
CREATE TABLE `t_shenmo_rank_time` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "上周时间戳",
  `thisTime` bigint(11) DEFAULT 0 COMMENT "本周时间戳",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

 -- create by zrc 2019-06-11
-- ----------------------------
-- Table structure for t_shenmo_rank 神魔战场排行榜
-- ----------------------------
DROP TABLE IF EXISTS `t_shenmo_rank`;
CREATE TABLE `t_shenmo_rank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
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
-- Table structure for t_player_goldequip_log  元神金装日志
-- ----------------------------
DROP TABLE IF EXISTS `t_player_goldequip_log`;
CREATE TABLE `t_player_goldequip_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `fenJieItemId` varchar(512) NOT NULL COMMENT "分解物品id",
  `rewItemStr` varchar(2014) NOT NULL COMMENT "分解获得物品",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_goldequip_setting  玩家元神金装设置
-- ----------------------------
DROP TABLE IF EXISTS `t_player_goldequip_setting`;
CREATE TABLE `t_player_goldequip_setting` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `fenJieIsAuto` int(11) NOT NULL COMMENT "是否自动分解",
  `fenJieQuality` int(11) NOT NULL COMMENT "分解品质",
  `fenJieZhuanShu` int(11) NOT NULL COMMENT "分解转数",
  `isCheckOldSt` int(11) NOT NULL COMMENT "是否检查过老强化等级", 
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- cjb create by 2019-06-12
-- Table structure for t_player_item_skill 玩家物品技能
-- ----------------------------
DROP TABLE IF EXISTS `t_player_item_skill`;
CREATE TABLE `t_player_item_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "技能类型",
  `level` int(11) NOT NULL COMMENT "等级",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 
-- ----------------------------
-- xzk create by 2019-06-13
-- Table structure for t_player_power_record 玩家战力记录数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_power_record`;
CREATE TABLE `t_player_power_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `todayInitPower` bigint(20) NOT NULL COMMENT "今日初始战力",
  `hisMaxPower` bigint(20) NOT NULL COMMENT "历史最高战力",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- cjb create by 2019-06-13
-- Table structure for t_player_shop_discount 玩家商城促销
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shop_discount`;
CREATE TABLE `t_player_shop_discount` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "特权类型",
  `startTime` bigint(20) DEFAULT 0 COMMENT "开始时间",
  `endTime` bigint(20) DEFAULT 0 COMMENT "结束时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_feisheng_receive  玩家飞升次数限制
-- ----------------------------
DROP TABLE IF EXISTS `t_player_feisheng_receive`;
CREATE TABLE `t_player_feisheng_receive` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `num` int(11) NOT NULL COMMENT "次数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_qixue 玩家泣血枪数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_qixue`;
CREATE TABLE `t_player_qixue` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `currLevel` int(11) NOT NULL COMMENT "当前阶数",
  `currStar` int(11) NOT NULL COMMENT "当前星数",
  `lastTime` bigint(20) NOT NULL COMMENT "上次被击杀掉落杀气时间",
  `shaLuNum` bigint(20) NOT NULL COMMENT "杀戮心数量",
  `timesNum` int(11) NOT NULL COMMENT "当前阶数升级次数",
  `power` bigint(20) NOT NULL COMMENT "戮仙刃战力",
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_chuangshi_yugao 创世之战预告数据
-- ----------------------------
DROP TABLE IF EXISTS `t_chuangshi_yugao`;
CREATE TABLE `t_chuangshi_yugao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` bigint(20) NOT NULL COMMENT "服务器索引",
  `num` int(11) NOT NULL COMMENT "人数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_chuangshi_yugao 玩家创世之战预告数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_chuangshi_yugao`;
CREATE TABLE `t_player_chuangshi_yugao` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `isJoin` int(11) NOT NULL COMMENT "是否参加",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`), 
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for t_player_mingge_buchang 玩家命格补偿数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mingge_buchang`;
CREATE TABLE `t_player_mingge_buchang` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `buchang` int(11) NOT NULL COMMENT "补偿",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_charm_add_log 魅力增加日志
-- ----------------------------
 DROP TABLE IF EXISTS `t_player_charm_add_log`;
 CREATE TABLE `t_player_charm_add_log` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `sendId` bigint(20) NOT NULL COMMENT "赠送者id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `charm` int(11) NOT NULL COMMENT "魅力值",
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
   PRIMARY KEY (`id`),
   KEY(`playerId`), 
     INDEX playerIdIndex (`playerId`) 
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;



 -- ----------------------------
 -- Table structure for t_player_feedbackfee 玩家逆付费数据
 -- ----------------------------
 DROP TABLE IF EXISTS `t_player_feedbackfee`;
 CREATE TABLE `t_player_feedbackfee` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `money` int(11) NOT NULL COMMENT "库存金额",
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
   PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- ----------------------------
 -- Table structure for t_player_arenapvp 玩家比武大会竞猜数据
 -- ----------------------------
 DROP TABLE IF EXISTS `t_player_arenapvp`;
 CREATE TABLE `t_player_arenapvp` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `reliveTimes` int(11) NOT NULL COMMENT "复活次数", 
   `outStatus` int(11) NOT NULL COMMENT "是否淘汰：0否1是",
   `jiFen` int(11) NOT NULL COMMENT "积分", 
   `guessNotice` int(11) NOT NULL COMMENT "竞猜提醒设置",  
   `pvpRecord` int(11) NOT NULL COMMENT "pvp成绩",  
   `ticketFlag` int(11) NOT NULL COMMENT "是否购买门票",  
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
   PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- ----------------------------
 -- Table structure for t_player_arenapvp_guess_log 玩家比武大会竞猜日志数据
 -- ----------------------------
 DROP TABLE IF EXISTS `t_player_arenapvp_guess_log`;
 CREATE TABLE `t_player_arenapvp_guess_log` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `raceNum` int(11) NOT NULL COMMENT "届数",
   `guessId` bigint(20) NOT NULL COMMENT "竞猜玩家id",
   `guessType` int(11) NOT NULL COMMENT "竞猜类型", 
   `winnerId` bigint(20) NOT NULL COMMENT "获胜玩家id",  
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",  
   PRIMARY KEY (`id`),  
   KEY(`playerId`), 
      INDEX playerIdIndex (`playerId`) 
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4; 

 -- create by xzk 2019-07-8
 -- Table structure for t_arenapvp_guess_record 竞猜记录
 -- ----------------------------
 DROP TABLE IF EXISTS `t_arenapvp_guess_record`;
 CREATE TABLE `t_arenapvp_guess_record` ( 
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `serverId` int(11) NOT NULL COMMENT "服务器id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id", 
   `raceNumber` int(11) NOT NULL COMMENT "届数",
   `guessType` int(11) NOT NULL COMMENT "竞猜类型",
   `guessId` bigint(20) NOT NULL COMMENT "竞猜玩家id",   
   `winnerId` bigint(20) NOT NULL COMMENT "获胜玩家id",   
   `status` int(11) NOT NULL COMMENT "状态",  
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",  
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
   PRIMARY KEY (`id`)
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4; 
 

-- create by xzk 2019-06-29
-- ----------------------------
-- Table structure for t_player_feedbackfee 玩家逆付费数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_feedbackfee`;
CREATE TABLE `t_player_feedbackfee` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `totalGetMoney` bigint(20) NOT NULL COMMENT "总共获取库存",
  `money` int(11) NOT NULL COMMENT "库存金额",
  `todayUseNum` int(11) NOT NULL COMMENT "今天使用数量",
  `useTime` bigint(20) NOT NULL COMMENT "使用时间",
  `cashMoney` bigint(20) NOT NULL COMMENT "现金兑换",
  `goldMoney` bigint(20) NOT NULL COMMENT "元宝兑换",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4; 


  -- create by zrc 2019-07-4
 -- ----------------------------
 -- Table structure for t_player_feedback_record 玩家逆付费记录数据
 -- ----------------------------
 DROP TABLE IF EXISTS `t_player_feedbackfee_record`;
  CREATE TABLE `t_player_feedbackfee_record` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `money` int(11) NOT NULL COMMENT "金额",
   `code` varchar(50) NOT NULL COMMENT "兑换码",
   `status` int(11) NOT NULL COMMENT "状态",
   `type` int(11) NOT NULL COMMENT "类型0:现金兑换1:元宝兑换",
   `expiredTime`  bigint(20) NOT NULL COMMENT "过期时间",
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
   PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

  -- create by zrc 2019-07-4
 -- ----------------------------
 -- Table structure for t_feedback_exchange 兑换记录
 -- ----------------------------
  DROP TABLE IF EXISTS `t_feedback_exchange`;
  CREATE TABLE `t_feedback_exchange` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `serverId` int(11) NOT NULL COMMENT "服务器id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `exchangeId` bigint(20) NOT NULL COMMENT "兑换id",
   `code` varchar(50) NOT NULL COMMENT "兑换码",
   `status` int(11) NOT NULL COMMENT "状态",
   `money` int(11) NOT NULL COMMENT "金额",
   `expiredTime` bigint(11) NOT NULL COMMENT "过期时间", 
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
   PRIMARY KEY (`id`)
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4; 


 
 -- create by xzk 2019-07-12
 -- Table structure for t_player_daliwan 大力丸
 -- ----------------------------
DROP TABLE IF EXISTS `t_player_daliwan`;
 CREATE TABLE `t_player_daliwan` ( 
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id", 
   `typ` int(11) NOT NULL COMMENT "类型",
   `startTime` bigint(20) NOT NULL COMMENT "使用时间",
   `duration` bigint(20) NOT NULL COMMENT "持续时间",
   `expired` int(11) NOT NULL COMMENT "过期",
   `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",  
   `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间", 
   `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间", 
    PRIMARY KEY (`id`),  
   KEY(`playerId`), 
      INDEX playerIdIndex (`playerId`) 
 ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4; 

 -- create by jzy 2019-07-15
 -- Table structure for t_player_wushuangweapon_slot 无双神器位置
 -- ----------------------------
DROP TABLE IF EXISTS `t_player_wushuangweapon_slot`;
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

  -- ----------------------------
  -- Table structure for t_player_jieyi 玩家结义数据
  -- ----------------------------
  DROP TABLE IF EXISTS `t_player_jieyi`;
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

  -- ----------------------------
  -- Table structure for t_jieyi 结义数据 
  -- ----------------------------
   DROP TABLE IF EXISTS `t_jieyi`;
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

  -- ----------------------------
  -- Table structure for t_jieyi_member 结义成员数据
  -- ----------------------------
   DROP TABLE IF EXISTS `t_jieyi_member`;
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

  -- ----------------------------
  -- Table structure for t_jieyi_leave_word 结义留言数据
  -- ----------------------------
   DROP TABLE IF EXISTS `t_jieyi_leave_word`;
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
   DROP TABLE IF EXISTS `t_jieyi_invite`;
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

   
     -- ----------------------------
  -- Table structure for t_dingshi_boss 定时boss
  -- ----------------------------
  DROP TABLE IF EXISTS `t_dingshi_boss`;
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
DROP TABLE IF EXISTS `t_player_zhenxi_boss`;
CREATE TABLE `t_player_zhenxi_boss` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `reliveTime` int(11) NOT NULL COMMENT "复活次数",
    `enterTimes` int(11) NOT NULL COMMENT "进入次数",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_new_first_charge_record 玩家新首充记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_new_first_charge_record`;
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

-- ----------------------------
-- Table structure for t_new_first_charge 新首充活动信息
-- ----------------------------
DROP TABLE IF EXISTS `t_new_first_charge`;
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
DROP TABLE IF EXISTS `t_player_boss_relive`;
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


-- ----------------------------
-- Table structure for t_new_first_charge_log 新首充活动记录信息
-- ----------------------------
DROP TABLE IF EXISTS `t_new_first_charge_log`;
CREATE TABLE `t_new_first_charge_log` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `serverId` bigint(11) NOT NULL COMMENT "服务器id",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by jzy 2019-08-02
-- ----------------------------
-- Table structure for t_player_wushuang_settings 无双神器配置
DROP TABLE IF EXISTS `t_player_wushuang_settings`;
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


/*
  -- ----------------------------
  -- Table structure for t_chuangshi_chengfang_jianshe 城防建设记录
  -- ----------------------------
   DROP TABLE IF EXISTS `t_chuangshi_chengfang_jianshe`;
   CREATE TABLE `t_chuangshi_chengfang_jianshe` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
     `serverId` bigint(11) NOT NULL COMMENT "服务器id",
     `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `cityId` bigint(20) NOT NULL COMMENT "城池id",
     `jianSheType` int(11) NOT NULL COMMENT "建设类型",
     `num` int(11) NOT NULL COMMENT "数量", 
     `status` int(11) NOT NULL COMMENT "状态",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`),
     KEY(`playerId`), 
       INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

  -- ----------------------------
  -- Table structure for t_chuangshi_shenwang_signup 神王报名记录
  -- ----------------------------
   DROP TABLE IF EXISTS `t_chuangshi_shenwang_signup`;
   CREATE TABLE `t_chuangshi_shenwang_signup` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
     `serverId` bigint(11) NOT NULL COMMENT "服务器id",
     `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `status` int(11) NOT NULL COMMENT "状态",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`),
     KEY(`playerId`), 
       INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

  -- ----------------------------
  -- Table structure for t_chuangshi_shenwang_vote 神王选举投票记录
  -- ----------------------------
   DROP TABLE IF EXISTS `t_chuangshi_shenwang_vote`;
   CREATE TABLE `t_chuangshi_shenwang_vote` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
     `serverId` bigint(11) NOT NULL COMMENT "服务器id",
     `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `supportId` bigint(20) NOT NULL COMMENT "支持玩家id",
     `status` int(11) NOT NULL COMMENT "状态",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`),
     KEY(`playerId`), 
       INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


  -- ----------------------------
  -- Table structure for t_player_chuangshi_guanzhi 玩家创世官职数据
  -- ----------------------------
  DROP TABLE IF EXISTS `t_player_chuangshi_guanzhi`;
  CREATE TABLE `t_player_chuangshi_guanzhi` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
     `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `receiveRewLevel` int(11) NOT NULL COMMENT "当前领取奖励等级",
     `level` int(11) NOT NULL COMMENT "等级",
     `times` int(11) NOT NULL COMMENT "次数",
     `weiWang` int(11) NOT NULL COMMENT "威望值",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`),
     KEY(`playerId`),
        INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

  -- ----------------------------
  -- Table structure for t_player_chuangshi_sign 玩家创世神王报名数据
  -- ----------------------------
  DROP TABLE IF EXISTS `t_player_chuangshi_sign`;
  CREATE TABLE `t_player_chuangshi_sign` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
     `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `status` int(11) NOT NULL COMMENT "状态",
     `lastSignTime` bigint(20) DEFAULT 0 COMMENT "上次报名时间",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间", 
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`),
     KEY(`playerId`),
        INDEX playerIdIndex (`playerId`)  
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

  -- ----------------------------
  -- Table structure for t_player_chuangshi_vote 玩家创世神王投票数据
  -- ----------------------------
  DROP TABLE IF EXISTS `t_player_chuangshi_vote`;
  CREATE TABLE `t_player_chuangshi_vote` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
     `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `status` int(11) NOT NULL COMMENT "状态",
     `lastVoteTime` bigint(20) DEFAULT 0 COMMENT "上次投票时间", 
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",  
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`),
     KEY(`playerId`),
        INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

  -- ----------------------------
  -- Table structure for t_player_chuangshi 玩家创世信息数据
  -- ----------------------------
  DROP TABLE IF EXISTS `t_player_chuangshi`;
  CREATE TABLE `t_player_chuangshi` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
     `playerId` bigint(20) NOT NULL COMMENT "玩家id",
     `campType` int(11) NOT NULL COMMENT "阵营类型",
     `pos` int(11) NOT NULL COMMENT "阵营官职",
     `jifen` bigint(20) NOT NULL COMMENT "积分",
     `diamonds` bigint(20) NOT NULL COMMENT "钻石",
     `weiWang` bigint(20) NOT NULL COMMENT "威望值",
     `lastMyPayTime` bigint(20) NOT NULL COMMENT "上次工资时间",
     `joinCampTime` bigint(20) NOT NULL COMMENT "加入阵营时间",
     `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
     `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
     `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
     PRIMARY KEY (`id`),
     KEY(`playerId`),
        INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

   */



-- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_goldequip 玩家元神金装数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_goldequip`;
CREATE TABLE `t_player_goldequip` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_mingge 玩家命格数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mingge`;
CREATE TABLE `t_player_mingge` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_shenghen 玩家圣痕数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shenghen`;
CREATE TABLE `t_player_shenghen` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_zhenfa_power 玩家阵法数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_zhenfa_power`;
CREATE TABLE `t_player_zhenfa_power` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_tulong_equip 玩家屠龙装数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_tulong_equip`;
CREATE TABLE `t_player_tulong_equip` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xubin 2019-08-06
-- ----------------------------
-- Table structure for t_player_baby_power 玩家宝宝数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_baby_power`;
CREATE TABLE `t_player_baby_power` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `power`  bigint(20) NOT NULL COMMENT "战力",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


 -- create by jzy 2019-08-08
-- ----------------------------
-- Table structure for t_player_addition_sys_lingzhu 附加系统五行灵珠
-- ----------------------------
DROP TABLE IF EXISTS `t_player_addition_sys_lingzhu`;
CREATE TABLE `t_player_addition_sys_lingzhu` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `sysType` bigint(11) NOT NULL COMMENT "附加系统类型",
    `lingZhuId` bigint(11) NOT NULL COMMENT "灵珠类型",
    `level` bigint(11) NOT NULL COMMENT "等级",
    `times` bigint(11) NOT NULL COMMENT "次数",
    `bless` bigint(20) NOT NULL COMMENT "祝福值",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

 -- create by xubin 2019-08-13
-- ----------------------------
-- Table structure for t_player_xianzun_card 玩家仙尊特权卡数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_xianzun_card`;
CREATE TABLE `t_player_xianzun_card` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `typ` bigint(11) NOT NULL COMMENT "仙尊特权卡类型",
    `IsActivite` bigint(11) NOT NULL COMMENT "是否激活",
    `isReceive` bigint(11) NOT NULL COMMENT "是否领取",
    `activiteTime` bigint(20) NOT NULL COMMENT "激活时间",
    `receiveTime` bigint(20) NOT NULL COMMENT "领取时间",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`),
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


 -- create by jzy 2019-08-14
-- ----------------------------
-- Table structure for t_player_wushuang_buchang 玩家无双神器补偿邮件数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_wushuang_buchang`;
CREATE TABLE `t_player_wushuang_buchang` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id", 
    `playerId` bigint(20) NOT NULL COMMENT "玩家id",
    `isSendEmail` bigint(11) NOT NULL COMMENT "是否发送邮件",
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`), 
    KEY(`playerId`), 
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_privilege_charge  玩家扶持充值记录
-- ----------------------------
DROP TABLE IF EXISTS `t_player_privilege_charge`;
CREATE TABLE `t_player_privilege_charge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chargeType`  int(11) NOT NULL COMMENT "平台类型",
  `chargeId` int(11) NOT NULL COMMENT "充值模板id",
  `chargeNum` bigint(20) NOT NULL COMMENT "元宝数量",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`), 
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- create by xzk 2019-08-13
-- ----------------------------
-- Table structure for t_arena_rank 3v3排行榜
-- ----------------------------
DROP TABLE IF EXISTS `t_arena_rank`;
CREATE TABLE `t_arena_rank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `playerId` bigint(20) DEFAULT 0 COMMENT "玩家id",
  `playerName` varchar(100) DEFAULT "" COMMENT "玩家名字",
  `curWinCount` int(11) DEFAULT 0 COMMENT "本周连胜",
  `winCount` int(11) DEFAULT 0 COMMENT "本周最高连胜", 
  `lastWinCount` int(11) DEFAULT 0 COMMENT "上周连胜",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "最后操作时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_arena_rank_time 3v3排行榜
-- ----------------------------
DROP TABLE IF EXISTS `t_arena_rank_time`;
CREATE TABLE `t_arena_rank_time` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) DEFAULT 0 COMMENT "服务器id",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "上周时间戳",
  `thisTime` bigint(11) DEFAULT 0 COMMENT "本周时间戳",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_shangguzhiling 玩家上古之灵数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_shangguzhiling`;
CREATE TABLE `t_player_shangguzhiling` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `lingShouType` int(11) NOT NULL COMMENT "灵兽类型",
  `level` int(11) NOT NULL COMMENT "等级",
  `experience` bigint(20) NOT NULL COMMENT "经验",
  `lingwen` varchar(512) DEFAULT "{}" COMMENT "灵纹",
  `uprankLevel` int(11) NOT NULL COMMENT "阶级",
  `uprankBless` bigint(20) NOT NULL COMMENT "祝福值",
  `uprankTimes` int(11) NOT NULL COMMENT "本阶尝试进阶次数",
  `linglian` varchar(512) DEFAULT "{}" COMMENT "灵炼",
  `linglianTimes` int(11) NOT NULL COMMENT "灵炼次数",
  `receiveTime` bigint(20) DEFAULT 0 COMMENT "上一次领取时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_ring 玩家特戒数据
-- ----------------------------
DROP TABLE IF EXISTS `t_player_ring`;
CREATE TABLE `t_player_ring` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) DEFAULT 0 COMMENT "玩家id",
  `typ` bigint(11) DEFAULT 0 COMMENT "特戒类型",
  `bindType` bigint(11) DEFAULT 0 COMMENT "绑定类型",
  `itemId` bigint(11) DEFAULT 0 COMMENT "特戒物品id",
  `propertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`), 
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_ring_baoku 玩家特戒宝库
-- ----------------------------
DROP TABLE IF EXISTS `t_player_ring_baoku`;
CREATE TABLE `t_player_ring_baoku` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `typ` int(11) NOT NULL COMMENT "宝库类型",
  `luckyPoints` int(11) NOT NULL COMMENT "幸运值",
  `attendPoints` int(11) NOT NULL COMMENT "积分",
  `totalAttendTimes` int(11) NOT NULL COMMENT "总参与次数",
  `lastSystemRefreshTime` bigint(20) DEFAULT 0 COMMENT "上次自动刷新时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;