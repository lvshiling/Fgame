set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


-- xzk create by 2018-11-26
-- Table structure for t_player_unreal_boss 玩家幻境boss数据
-- ----------------------------
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
CREATE TABLE `t_player_addition_sys_level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `sysType` int(11) NOT NULL COMMENT "系统类型",
  `level` int(11) NOT NULL COMMENT "系统等级",
  `upNum` int(11) NOT NULL COMMENT "等级",
  `upPro` int(11) NOT NULL COMMENT "等级",
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

-- ----------------------------
-- create time:20181129
-- author：ylz
-- Table structure for t_player_liveness 玩家活跃度数据
-- ----------------------------
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
-- ylz create by 2018-11-27
-- Table structure for t_player_fabao_other  非进阶法宝
-- ----------------------------
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


-- ylz 2018-11-28
alter table `t_player_cache` add column `fabaoInfo` text(2000) NOT NULL COMMENT "法宝";
UPDATE `t_player_cache` SET `fabaoInfo`= "{}";






-- xzk create by 2018-11-28
-- Table structure for t_player_unreal_boss 玩家幻境boss数据
-- ----------------------------
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



-- ylz 2018-11-29
alter table `t_player_cache` add column `xuedunInfo` text(2000) NOT NULL COMMENT "血盾";
UPDATE `t_player_cache` SET `xuedunInfo`= "{}";






-- ----------------------------
-- create time:20181129
-- author：cjb
-- Table structure for t_player_xianti 玩家仙体数据
-- ----------------------------
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

-- cjb 2018-11-30
alter table `t_player_cache` add column `xiantiInfo` text(2000) NOT NULL COMMENT "仙体";
UPDATE `t_player_cache` SET `xiantiInfo`= "{}";



-- xzk 2018-12-03
-- Table structure for t_player_cycle_charge_record  玩家每日充值记录
-- ----------------------------
CREATE TABLE `t_player_cycle_charge_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `chargeNum`  bigint(20) NOT NULL COMMENT "充值元宝数",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
  KEY(`playerId`),
  INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;


-- xzk 2018-12-03
-- Table structure for t_player_foe_protect  玩家仇人反馈保护
-- ----------------------------
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
CREATE TABLE `t_player_friend_feedback` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `friendId` bigint(20) NOT NULL COMMENT "赞赏玩家id",
  `friendName` varchar(20) NOT NULL COMMENT "赞赏玩家名称",
  `noticeType` int(11) NOT NULL COMMENT "消息类型",
  `feedbackType` int(11) NOT NULL COMMENT "赞赏类型",
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
CREATE TABLE `t_player_friend_add_rew` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
  `playerId` bigint(20) NOT NULL COMMENT "玩家id",
  `rewRecord` varchar(512) NOT NULL COMMENT "奖励记录",
  `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
  `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
  `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
  PRIMARY KEY (`id`),
   KEY(`playerId`),
      INDEX playerIdIndex (`playerId`) 
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
