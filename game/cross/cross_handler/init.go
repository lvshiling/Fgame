package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LOGIN_TYPE), (*crosspb.ISLogin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LOGIN_TYPE), (*crosspb.SILogin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_DATA_TYPE), (*crosspb.SIPlayerData)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_SYSTEM_BATTLE_PROPERTY_CHANGED_TYPE), (*crosspb.SIPlayerSystemBattlePropertyChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_BASIC_PROPERTY_CHANGED_TYPE), (*crosspb.SIPlayerBasicPropertyChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_SHOW_DATA_CHANGED_TYPE), (*crosspb.SIPlayerShowDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_EXIT_CROSS_TYPE), (*crosspb.SIPlayerExitCross)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_EXIT_CROSS_TYPE), (*crosspb.ISPlayerExitCross)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_HEARTBEAT_TYPE), (*crosspb.SIHeartBeat)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_HEARTBEAT_TYPE), (*crosspb.ISHeartBeat)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_BATTLE_DATA_CHANGED_TYPE), (*crosspb.SIPlayerBattleDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_TEAM_CHANGED_TYPE), (*crosspb.SIPlayerTeamSync)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_ALLIANCE_CHANGED_TYPE), (*crosspb.SIPlayerAllianceSync)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_CHANGED_TYPE), (*crosspb.SILingTongDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_INIT_TYPE), (*crosspb.SILingTongDataInit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_REMOVE_TYPE), (*crosspb.SILingTongDataRemove)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_BUFF_REMOVE_TYPE), (*crosspb.SIBuffRemove)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_BUFF_ADD_TYPE), (*crosspb.SIBuffAdd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_BUFF_UPDATE_TYPE), (*crosspb.SIBuffUpdate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_JIEYI_CHANGED_TYPE), (*crosspb.SIPlayerJieYiSync)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_TESHU_SKILL_RESET_TYPE), (*crosspb.SIPlayerTeshuSkillReset)(nil))

}
