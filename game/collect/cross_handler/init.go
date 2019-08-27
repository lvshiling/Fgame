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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_COLLECT_FINISH_TYPE), (*crosspb.ISCollectFinish)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_COLLECT_FINISH_TYPE), (*crosspb.SICollectFinish)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_COLLECT_MIZANG_FINISH_TYPE), (*crosspb.ISCollectMiZangFinish)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_COLLECT_MIZANG_FINISH_TYPE), (*crosspb.SICollectMiZangFinish)(nil))
}
