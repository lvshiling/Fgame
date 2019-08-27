set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;





  -- create by xzk 2019-06-29
 -- ----------------------------
 -- Table structure for t_player_feedbackfee 玩家逆付费数据
 -- ----------------------------
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

 -- ----------------------------
 -- Table structure for t_player_arenapvp 玩家比武大会竞猜数据
 -- ----------------------------
 CREATE TABLE `t_player_arenapvp` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT "id",
   `playerId` bigint(20) NOT NULL COMMENT "玩家id",
   `reliveTimes` int(11) NOT NULL COMMENT "复活次数",
   `outStatus` int(11) NOT NULL COMMENT "是否淘汰：0否1是",
   `jiFen` int(11) NOT NULL COMMENT "积分", 
   `guessNotice` int(11) NOT NULL COMMENT "竞猜提醒设置",  
   `pvpRecord` int(11) NOT NULL COMMENT "pvp成绩",  
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



 -- create by xzk 2019-07-8
 -- Table structure for t_player_daliwan 大力丸
 -- ----------------------------
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
