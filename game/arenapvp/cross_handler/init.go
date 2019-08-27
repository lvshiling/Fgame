package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	initCodec()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_ATTEND_TYPE), (*crosspb.ISArenapvpAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENAPVP_ATTEND_TYPE), (*crosspb.SIArenapvpAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_RELIVE_TYPE), (*crosspb.ISArenapvpRelive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENAPVP_RELIVE_TYPE), (*crosspb.SIArenapvpRelive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_RESET_RELIVETIMES_TYPE), (*crosspb.ISArenapvpResetReliveTimes)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_ATTEND_SUCCESS_TYPE), (*crosspb.ISArenapvpAttendSuccess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_RESULT_ELECTION_TYPE), (*crosspb.ISArenapvpResultElection)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_RESULT_BATTLE_TYPE), (*crosspb.ISArenapvpResultBattle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENAPVP_PLAYER_DATA_CHANGED_TYPE), (*crosspb.SIPlayerArenapvpDataChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_ELECTION_LUCKY_REW_TYPE), (*crosspb.ISAreanapvpElectionLuckyRew)(nil))
}
