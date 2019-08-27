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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_QIXUE_DROP_TYPE), (*crosspb.ISQiXueDrop)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_QIXUE_DROP_TYPE), (*crosspb.SIQiXueDrop)(nil))
}
