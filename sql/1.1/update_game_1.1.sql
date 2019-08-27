set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

-- xzk create by 2018-11-01
alter table t_player_alliance add column `depotPoint` int(11) DEFAULT 0 COMMENT "仓库积分";

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


-- cjb create by 2018-11-06
alter table t_player_massacre add column `currLevel` int(11) DEFAULT 0 COMMENT "当前阶数";
alter table t_player_massacre add column `currStar` int(11) DEFAULT 0 COMMENT "当前星数";

-- ----------------------------
-- xzk create by 2018-11-06
-- Table structure for t_player_tower 玩家打宝塔数据
-- ----------------------------
CREATE TABLE `t_player_tower` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `useTime` bigint(20) NOT NULL COMMENT "已用打宝时间",
  `extraTime` bigint(20) NOT NULL COMMENT "额外打宝时间",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
    KEY(`playerId`),
    INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- ylz create by 2018-11-08
alter table `t_player_weapon` add column `activeFlag` int(11) DEFAULT 1 COMMENT "激活标识";


-- zrc create by 2018-11-14
alter table `t_alliance` add column `originServerId` int(11) DEFAULT 0 COMMENT "原始服务器";
update t_alliance set originServerId=serverId;

