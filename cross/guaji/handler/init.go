package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()

}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUA_JI_POS_TYPE), (*uipb.SCGuaJiPos)(nil))

}
