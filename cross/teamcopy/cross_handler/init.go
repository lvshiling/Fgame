package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()
}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_TEAMCOPY_START_BATTLE_TYPE), (*crosspb.SITeamCopyStartBattle)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_TEAMCOPY_START_BATTLE_TYPE), (*crosspb.ISTeamCopyStartBattle)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_TEAMCOPY_BATTLE_RESULT_TYPE), (*crosspb.ISTeamCopyBattleResult)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_TEAMCOPY_BATTLE_RESULT_TYPE), (*crosspb.SITeamCopyBattleResult)(nil))
}
