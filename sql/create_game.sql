-- ----------------------------
-- Table structure for t_player 玩家表
-- ----------------------------
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
  `online` int(11) NOT NULL COMMENT "在线",
  `sdkType` int(11) NOT NULL COMMENT "sdk类型",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_player_scene 玩家场景数据
-- ----------------------------
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
CREATE TABLE `t_player_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `skillId` int(11) NOT NULL COMMENT "技能id",
  `level` int(11) NOT NULL COMMENT "技能等级",
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
-- ---------------------------
CREATE TABLE `t_player_inventory` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `slotNum` int(11) NOT NULL COMMENT "背包格子数",
  `depotNum` int(11) NOT NULL COMMENT "仓库格子数",
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
-- ---------------------------
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
-- ----------------------------;
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
CREATE TABLE `t_player_title` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `titleId` int(11) NOT NULL COMMENT "称号id",
  `activeFlag` int(11) NOT NULL COMMENT "是否激活",
  `activeTime` bigint(20) DEFAULT 0 COMMENT "激活时间",
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
CREATE TABLE `t_player_tianjieta` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `playerName` varchar(50) NOT NULL COMMENT "玩家名字",
  `level` int(11) NOT NULL COMMENT "天劫塔等级",
  `usedTime` bigint(20) NOT NULL COMMENT "使用时间",
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
   KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_player_xianfu  秘境仙府
-- ----------------------------
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
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance_member 仙盟成员
-- ----------------------------
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
  KEY(`memberId`),
   INDEX allianceIdIndex (`allianceId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance_join_apply 仙盟人员申请列表
-- ----------------------------
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
-- Table structure for t_alliance_hegemon 
-- ----------------------------
CREATE TABLE `t_alliance_hegemon` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `serverId` int(11) NOT NULL COMMENT "服务器id",
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
  `yaoPai` int(11) NOT NULL COMMENT "腰牌",
  `lastYaoPaiUpdateTime` bigint(20) NOT NULL COMMENT "上次腰牌更新时间",
  `convertTimes` int(11) NOT NULL COMMENT "兑换次数",
  `lastConvertUpdateTime` bigint(20) NOT NULL COMMENT "上次兑换更新时间",
  `lastAllianceSceneEndTime` bigint(20) NOT NULL COMMENT "城战结束时间",
  `reliveTime` int(11) NOT NULL COMMENT "原地复活累计次数",
  `lastReliveTime` bigint(20) DEFAULT 0 COMMENT "原地复活上次时间",
  `lastMemberCallTime` bigint(20) DEFAULT 0 COMMENT "上次仙盟召集时间",
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
   `shenfaInfo` text(2000) NOT NULL  COMMENT "身法",
   `lingyuInfo` text(2000) NOT NULL COMMENT "领域",
  `shieldInfo` text(2000) NOT NULL COMMENT "神盾尖刺",
   `featherInfo` text(2000) NOT NULL COMMENT "护体仙羽",
   `anqiInfo` text(2000)  NOT NULL  COMMENT "暗器",
  `massacreInfo` text(2000)  NOT NULL  COMMENT "戮仙刃",
   `realmLevel` int(11) NOT NULL COMMENT "天劫塔等级",
   `skillList` text(2000) NOT NULL COMMENT "技能列表",
   `vipInfo` text(2000) NOT NULL COMMENT "vip",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_alliance_log  仙盟日志
-- ----------------------------
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
   KEY(`playerId`),
   INDEX playerIdIndex (`playerId`) 
   ) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
   
-- ----------------------------
-- Table structure for t_player_secret_card  天机牌
-- ----------------------------
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
CREATE TABLE `t_biaoche` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `serverId` int(11) COMMENT "服务器id",
  `allianceId`  bigint(20) NOT NULL COMMENT "仙盟id",
  `transportMoveId` int(11) NOT NULL COMMENT "镖车路径模板id",
  `transportType` int(11) NOT NULL COMMENT "镖车类型",
  `state` int(11) NOT NULL COMMENT "镖车状态",
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
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_marry  婚烟表
-- ----------------------------
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
CREATE TABLE `t_player_gold_equip_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `level` int(11) NOT NULL COMMENT "等级",
  `itemId` int(11) NOT NULL COMMENT "物品id",
  `bindType` int(11) NOT NULL COMMENT "绑定类型",
  `porpertyData` varchar(512) DEFAULT "{}" COMMENT "属性数据",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
     INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;

-- Table structure for t_alliance_invitation 仙盟邀请列表
-- ----------------------------
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
CREATE TABLE `t_player_onearena` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `level` int(11) NOT NULL COMMENT "灵池等级",
  `pos` int(11) NOT NULL COMMENT "灵池标识",
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
CREATE TABLE `t_player_arena` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `endTime`  bigint(20) NOT NULL COMMENT "活动结束时间",
  `reliveTime` int(11) NOT NULL COMMENT "复活次数",
  `culRewardTime` int(11) NOT NULL COMMENT "获得奖励次数",
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
CREATE TABLE `t_player_major_num` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `num` int(11) NOT NULL COMMENT "双休次数",
  `lastTime` bigint(20) DEFAULT 0 COMMENT "上次使用时间",
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
-- ---------------------------
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
-- Table structure for t_register_setting_log  //注册设置日志
-- ----------------------------
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
CREATE TABLE `t_player_addition_sys_slot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型",
  `slotId` int(11) NOT NULL COMMENT "装备槽id",
  `level` int(11) NOT NULL COMMENT "等级",
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

