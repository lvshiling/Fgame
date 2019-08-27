package delete

import (
	model "fgame/fgame/tools/dbdelete/model"
	tool "fgame/fgame/tools/dbdelete/tool"
	"fmt"
)

type serverMySqlDelete struct {
}

func (d *serverMySqlDelete) DeleteTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_mysqlPath string) error {
	whereSql := fmt.Sprintf("serverId=%d", p_serverid)
	return tool.DeleteTable(p_db, p_table, whereSql, p_mysqlPath)
}

func (d *serverMySqlDelete) IsServer() bool {
	return true
}

var (
	serverIdTableArray = []string{
		"t_alliance",
		"t_alliance_hegemon",
		"t_chess_log",
		"t_emperor",
		"t_emperor_records",
		"t_friend",
		"t_marry",
		"t_marry_divorce_consent",
		"t_marry_ring",
		"t_onearena",
		"t_open_activity_email_record",
		"t_player",
		"t_wedding",
		"t_wedding_card",
		"t_merge",
		"t_first_charge",
		"t_compensate",
		"t_chat_setting",
		"t_register_setting",
		"t_register_setting_log",
		"t_order",
		"t_privilege_charge",
		"t_marry_pre_wed",
		"t_open_activity_rewards_limit",
		"t_open_activity_discount_limit",
		"t_outland_boss_drop_records",
		"t_open_activity_laba_log",
		"t_open_activity_start_mail",
		"t_quiz",
		"t_open_activity_drew_log",
		"t_open_activity_boss_kill",
		"t_open_activity_alliance_cheer",
		"t_open_activity_crazybox_log",
		"t_hongbao",
		"t_open_activity_xun_huan",
		"t_trade_item",
		"t_trade_order",
		"t_friend_marry_develop_log",
		"t_couple_baby",
		"t_trade_recycle",
		"t_activity_end_record",
		"t_shenmo_rank_time",
		"t_shenmo_rank",
		"t_chuangshi_yugao",
		"t_feedback_exchange",
		"t_arenapvp_guess_record",
		"t_jieyi",
		"t_jieyi_leave_word",
		"t_jieyi_invite",
		"t_dingshi_boss",
		"t_new_first_charge",
		"t_new_first_charge_log",
	}
)

var (
	serverDel = &serverMySqlDelete{}
)

func init() {
	for _, value := range serverIdTableArray {
		registerDelete(value, serverDel)
	}
}
