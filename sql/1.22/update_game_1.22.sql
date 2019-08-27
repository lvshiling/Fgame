set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

-- create by zrc 2019-05-31
alter table `t_player` add column `todayChargeMoney` bigint(20) default 0 COMMENT "今日充值"; 
alter table `t_player` add column `yesterdayChargeMoney` bigint(20) default 0 COMMENT "昨日充值"; 
alter table `t_player` add column `chargeTime` bigint(20) default 0 COMMENT "充值时间"; 
