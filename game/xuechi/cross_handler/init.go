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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_XUE_CHI_ADD_TYPE), (*crosspb.SIXueChiAdd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_XUE_CHI_SYNC_TYPE), (*crosspb.SIXueChiSync)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_XUE_CHI_SYNC_TYPE), (*crosspb.ISXueChiSync)(nil))

}
