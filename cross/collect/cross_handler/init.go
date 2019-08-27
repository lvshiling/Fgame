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

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_COLLECT_FINISH_TYPE), (*crosspb.ISCollectFinish)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_COLLECT_FINISH_TYPE), (*crosspb.SICollectFinish)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_COLLECT_MIZANG_FINISH_TYPE), (*crosspb.SICollectMiZangFinish)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_COLLECT_MIZANG_FINISH_TYPE), (*crosspb.ISCollectMiZangFinish)(nil))
}
