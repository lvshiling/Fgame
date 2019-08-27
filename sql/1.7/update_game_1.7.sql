set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;


-- ----------------------------
-- cjb create by 2018-12-21
-- Table structure for t_quiz 仙尊问答数据
-- ----------------------------
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
-- create by xzk 2018-12-25
-- Table structure for t_open_activity_drew_log  运营活动-抽奖日志
-- ----------------------------
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


-- 2018-12-26 create by xzk
alter table `t_alliance` add column `isAutoRemoveDepot` int(11) DEFAULT 0 NOT NULL COMMENT "是否自动销毁仙盟仓库物品：0否1是";
alter table `t_alliance` add column `maxRemoveZhuanSheng` int(11) DEFAULT 0 NOT NULL COMMENT "自动销毁最大转生条件";
alter table `t_alliance` add column `maxRemoveQuality` int(11) DEFAULT 0 NOT NULL COMMENT "自动销毁最高品质";


-- ----------------------------
-- ylz create by 2018-12-25
-- Table structure for t_player_dense_wat 玩家金银密窟数据
-- ----------------------------
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

-- xzk create by 2018-12-27
alter table `t_biaoche` add column `owerName` varchar(20) DEFAULT NULL COMMENT "拥有者名字";



-- ----------------------------
-- create time:20181227
-- author：cjb
-- Table structure for t_player_dianxing 玩家点星系统数据
-- ----------------------------
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


-- xzk create by 2018-12-28
alter table `t_player_friend_add_rew` add column `congratulateTimes` int(11) NOT NULL COMMENT "被祝贺次数";
alter table `t_player_friend_add_rew` add column `lastCongratulateTime` bigint(20) NOT NULL COMMENT "上次被祝贺时间";

-- cjb 2018-12-28
alter table `t_player_cache` add column `dianxingInfo` text(2000) NOT NULL COMMENT "点星";
UPDATE `t_player_cache` SET `dianxingInfo`= "{}";

-- xzk create by 2018-12-29 
alter table `t_player_friend_feedback` add column `condition` int(11) NOT NULL COMMENT "条件";
alter table `t_player_email` modify column `title` varchar(500)  NOT NULL COMMENT "邮件标题";
alter table `t_player_email` modify column `content` varchar(500) NOT NULL COMMENT "邮件内容";