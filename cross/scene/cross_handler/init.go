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

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_KILL_BIOLOGY_TYPE), (*crosspb.ISPlayerKillBiology)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLYAER_KILL_BIOLOGY_TYPE), (*crosspb.SIPlayerKillBiology)(nil))
}
