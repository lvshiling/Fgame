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
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_JIEYI_CHANGED_TYPE), (*crosspb.SIPlayerJieYiSync)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENGWEI_DROP_TYPE), (*crosspb.ISShengWeiDrop)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENGWEI_DROP_TYPE), (*crosspb.SIShengWeiDrop)(nil))
}
