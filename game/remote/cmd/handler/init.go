package handler

import (
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
)

func init() {
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_KICK_OUT_PLAYER_TYPE), (*cmdpb.CmdKickoutPlayer)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_FORBID_PLAYER_TYPE), (*cmdpb.CmdForbidPlayer)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_UNFORBID_PLAYER_TYPE), (*cmdpb.CmdUnforbidPlayer)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_FORBID_PLAYER_CHAT_TYPE), (*cmdpb.CmdForbidPlayerChat)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_UNFORBID_PLAYER_CHAT_TYPE), (*cmdpb.CmdUnforbidPlayerChat)(nil))

	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_IGNORE_PLAYER_CHAT_TYPE), (*cmdpb.CmdIgnorePlayerChat)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_UNIGNORE_PLAYER_CHAT_TYPE), (*cmdpb.CmdUnignorePlayerChat)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_CHAT_SET_TYPE), (*cmdpb.CmdChatSet)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_REGISTER_SET_TYPE), (*cmdpb.CmdRegisterSet)(nil))

	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_SEND_SERVER_COMPENSATE_TYPE), (*cmdpb.CmdSendServerCompensate)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_SEND_PLAYER_COMPENSATE_TYPE), (*cmdpb.CmdSendPlayerCompensate)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_PRIVILEGE_SET_TYPE), (*cmdpb.CmdPrivilegeSet)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_PRIVILEGE_CHARGE_TYPE), (*cmdpb.CmdPrivilegeCharge)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_CHARGE_TYPE), (*cmdpb.CmdCharge)(nil))

	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_BROADCAST_NOTICE), (*cmdpb.CmdBroadcastNotice)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_TRADE_SELL_TYPE), (*cmdpb.CmdTradeSell)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_MARRY_SET_TYPE), (*cmdpb.CmdMarrySet)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_CREATE_ROLE_TYPE), (*cmdpb.CmdCreateRole)(nil))

	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_ALLIANCE_NOTICE_TYPE), (*cmdpb.CmdAllainceNotice)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_CUSTOM_TRADE_RECYCLE_GOLD_TYPE), (*cmdpb.CmdCustomRecycleGold)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_PING_TYPE), (*cmdpb.CmdPing)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_EXCHANGE_TYPE), (*cmdpb.CmdExchange)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_FIRST_CHARGE_RESET_TYPE), (*cmdpb.CmdFirstChargeReset)(nil))
	cmd.RegisterCmd(cmd.CmdType(cmdpb.CmdType_CMD_ALLIANCE_DISMISS_TYPE), (*cmdpb.CmdAllianceDismiss)(nil))

}
