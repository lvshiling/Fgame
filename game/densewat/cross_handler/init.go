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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_DENSEWAT_SYNC_TYPE), (*crosspb.SIDenseWatSync)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_DENSEWAT_SYNC_TYPE), (*crosspb.ISDenseWatSync)(nil))

}
