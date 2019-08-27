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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENGWEI_DROP_TYPE), (*crosspb.ISShengWeiDrop)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENGWEI_DROP_TYPE), (*crosspb.SIShengWeiDrop)(nil))
}
