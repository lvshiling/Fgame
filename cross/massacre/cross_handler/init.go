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

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_MASSACRE_DROP_TYPE), (*crosspb.ISMassacreDrop)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_MASSACRE_DROP_TYPE), (*crosspb.SIMassacreDrop)(nil))
}
