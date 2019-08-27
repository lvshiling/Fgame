 set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

 -- create by xubin 2019-07-15
 alter table `t_player_equipbaoku` add column `typ` int(11) DEFAULT 0 COMMENT "宝库类型"

