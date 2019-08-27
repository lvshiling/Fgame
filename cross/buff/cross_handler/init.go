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
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_BUFF_ADD_TYPE), (*crosspb.SIBuffAdd)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_BUFF_REMOVE_TYPE), (*crosspb.SIBuffRemove)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_BUFF_UPDATE_TYPE), (*crosspb.SIBuffUpdate)(nil))

}
