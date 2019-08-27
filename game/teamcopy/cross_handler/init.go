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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_TEAMCOPY_START_BATTLE_TYPE), (*crosspb.SITeamCopyStartBattle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_TEAMCOPY_START_BATTLE_TYPE), (*crosspb.ISTeamCopyStartBattle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_TEAMCOPY_BATTLE_RESULT_TYPE), (*crosspb.ISTeamCopyBattleResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_TEAMCOPY_BATTLE_RESULT_TYPE), (*crosspb.SITeamCopyBattleResult)(nil))
}
