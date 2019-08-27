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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_MATCH_TYPE), (*crosspb.ISArenaMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_MATCH_TYPE), (*crosspb.SIArenaMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_MATCH_RESULT_TYPE), (*crosspb.ISArenaMatchResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_STOP_MATCH_TYPE), (*crosspb.SIArenaStopMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_STOP_MATCH_TYPE), (*crosspb.ISArenaStopMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_WIN_TYPE), (*crosspb.SIArenaWin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_WIN_TYPE), (*crosspb.ISArenaWin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_ARENA_DATA_CHANGED_TYPE), (*crosspb.SIPlayerArenaDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_RELIVE_TYPE), (*crosspb.ISArenaRelive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_RELIVE_TYPE), (*crosspb.SIArenaRelive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_COLLECT_EXP_TREE_TYPE), (*crosspb.ISArenaCollectExpTree)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_COLLECT_EXP_TREE_TYPE), (*crosspb.SIArenaCollectExpTree)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_COLLECT_BOX_TYPE), (*crosspb.ISArenaCollectBox)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_COLLECT_BOX_TYPE), (*crosspb.SIArenaCollectBox)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_GIVE_UP_TYPE), (*crosspb.ISArenaGiveUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_GIVE_UP_TYPE), (*crosspb.SIArenaGiveUp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_RESET_RELIVE_TIMES_TYPE), (*crosspb.ISArenaResetReliveTimes)(nil))
}
