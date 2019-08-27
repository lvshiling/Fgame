set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 
 -- create by zrc 2019-06-10
alter table `t_player_shenyu` add column `round` int(11) NOT NULL COMMENT "参赛轮";

 -- create by zrc 2019-06-11
-- ----------------------------
-- Table structure for t_shenmo_rank_time 神魔战场排行榜
-- ----------------------------
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


-- create by cjb 2019-06-11
alter table `t_player_tianjieta` add column `isCheckReissue` int(11) NOT NULL COMMENT "是否检测补发过";
alter table `t_player_gold_equip_slot` add column `gemUnlockInfo` varchar(500) NOT NULL DEFAULT "{}" COMMENT  "解锁宝石孔信息";

-- ----------------------------
-- Table structure for t_player_goldequip_log  元神金装日志
-- ----------------------------
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
CREATE TABLE `t_player_goldequip_setting` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `fenJieIsAuto` int(11) NOT NULL COMMENT "是否自动分解",
  `fenJieQuality` int(11) NOT NULL COMMENT "分解品质",
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

-- create by xzk 2019-06-11
alter table `t_player_title` add column `validTime` bigint(20) DEFAULT 0 COMMENT "有效时间";


-- ----------------------------
-- cjb create by 2019-06-12
-- Table structure for t_player_item_skill 玩家物品技能
-- ----------------------------
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

 -- create by zrc 2019-06-13
alter table `t_player_bagua` add column `isBuChang` int(11) NOT NULL COMMENT "八卦秘境补偿";

-- ----------------------------
-- xzk create by 2019-06-13
-- Table structure for t_player_power_record 玩家战力记录数据
-- ----------------------------
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

-- xzk create by 2019-06-15
alter table `t_trade_recycle` add column `customRecycleGold` bigint(20) NOT NULL COMMENT "自定义回收的元宝";

-- ----------------------------
-- Table structure for t_player_feisheng_receive  玩家飞升次数限制
-- ----------------------------
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
