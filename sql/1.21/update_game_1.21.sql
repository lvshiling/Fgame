set names 'utf8mb4';
set character_set_database = 'utf8mb4';
set character_set_server = 'utf8mb4';


USE `game`;

alter table `t_chat_setting` add column `teamVipLevel` int(11) NOT NULL COMMENT "队伍VIP等级";
alter table `t_chat_setting` add column `teamLevel` int(11) NOT NULL COMMENT "队伍等级";