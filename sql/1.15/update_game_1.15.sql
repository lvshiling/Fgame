set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;
 


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
CREATE TABLE `t_player_zhenfa` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "阵法类型",
  `level` int(11) NOT NULL COMMENT "阵法等级",
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
CREATE TABLE `t_player_zhenqi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "阵法类型",
  `zhenQiPos` int(11) NOT NULL COMMENT "阵旗部位",
  `number` int(11) NOT NULL COMMENT "阵旗阶数",
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
CREATE TABLE `t_player_zhenqi_xianhuo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `type` int(11) NOT NULL COMMENT "阵法类型",
  `level` int(11) NOT NULL COMMENT "级数",
  `luckyStar` int(11) NOT NULL COMMENT "暴击幸运星",
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


-- create by ylz 2019-03-15
alter table  `t_player_zhenfa` add column `levelNum` int(11) DEFAULT 0 COMMENT "升级次数";  
alter table  `t_player_zhenfa` add column `levelPro` int(11) DEFAULT 0 COMMENT "升级进度值";

alter table  `t_player_zhenqi` add column `numberNum` int(11) DEFAULT 0 COMMENT "升阶次数";  
alter table  `t_player_zhenqi` add column `numberPro` int(11) DEFAULT 0 COMMENT "升阶进度值";

alter table  `t_player_zhenqi_xianhuo` add column `levelNum` int(11) DEFAULT 0 COMMENT "升级次数";  
alter table  `t_player_zhenqi_xianhuo` add column `levelPro` int(11) DEFAULT 0 COMMENT "升级进度值";

