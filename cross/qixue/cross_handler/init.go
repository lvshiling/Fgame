package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()
}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_QIXUE_DROP_TYPE), (*crosspb.ISQiXueDrop)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_QIXUE_DROP_TYPE), (*crosspb.SIQiXueDrop)(nil))
}
