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
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_MATCH_TYPE), (*crosspb.SIArenaMatch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_MATCH_TYPE), (*crosspb.ISArenaMatch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_MATCH_RESULT_TYPE), (*crosspb.ISArenaMatchResult)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_STOP_MATCH_TYPE), (*crosspb.SIArenaMatch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_STOP_MATCH_TYPE), (*crosspb.ISArenaStopMatch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_WIN_TYPE), (*crosspb.ISArenaWin)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_WIN_TYPE), (*crosspb.SIArenaWin)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_ARENA_DATA_CHANGED_TYPE), (*crosspb.SIPlayerArenaDataChanged)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_RELIVE_TYPE), (*crosspb.ISArenaRelive)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_RELIVE_TYPE), (*crosspb.SIArenaRelive)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_COLLECT_EXP_TREE_TYPE), (*crosspb.ISArenaCollectExpTree)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_COLLECT_EXP_TREE_TYPE), (*crosspb.SIArenaCollectExpTree)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_COLLECT_BOX_TYPE), (*crosspb.ISArenaCollectBox)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_COLLECT_BOX_TYPE), (*crosspb.SIArenaCollectBox)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_GIVE_UP_TYPE), (*crosspb.ISArenaGiveUp)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_ARENA_GIVE_UP_TYPE), (*crosspb.SIArenaGiveUp)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_ARENA_RESET_RELIVE_TIMES_TYPE), (*crosspb.ISArenaResetReliveTimes)(nil))
}
