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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_TULONG_KILL_BOSS_TYPE), (*crosspb.ISTuLongKillBoss)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_TULONG_KILL_BOSS_TYPE), (*crosspb.SITuLongKillBoss)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_TULONG_ATTEND_TYPE), (*crosspb.ISTuLongAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_TULONG_ATTEND_TYPE), (*crosspb.SITuLongAttend)(nil))
}
