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

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_TULONG_KILL_BOSS_TYPE), (*crosspb.ISTuLongKillBoss)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_TULONG_KILL_BOSS_TYPE), (*crosspb.SITuLongKillBoss)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_TULONG_ATTEND_TYPE), (*crosspb.ISTuLongAttend)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_TULONG_ATTEND_TYPE), (*crosspb.SITuLongAttend)(nil))
}
