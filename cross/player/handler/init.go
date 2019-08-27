package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_SYSTEM_BATTLE_PROPERTY_CHANGED_TYPE), (*crosspb.SIPlayerSystemBattlePropertyChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_BASIC_PROPERTY_CHANGED_TYPE), (*crosspb.SIPlayerBasicPropertyChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_SHOW_DATA_CHANGED_TYPE), (*crosspb.SIPlayerShowDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_EXIT_CROSS_TYPE), (*crosspb.SIPlayerExitCross)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_EXIT_CROSS_TYPE), (*crosspb.ISPlayerExitCross)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_BATTLE_DATA_CHANGED_TYPE), (*crosspb.SIPlayerBattleDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_TESHU_SKILL_RESET_TYPE), (*crosspb.SIPlayerTeshuSkillReset)(nil))
}

func initProxy() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_PROPERTY_TYPE), (*uipb.SCPlayerProperty)(nil))
}
