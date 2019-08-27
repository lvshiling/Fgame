package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_CHANGED_TYPE), (*crosspb.SILingTongDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_INIT_TYPE), (*crosspb.SILingTongDataInit)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_REMOVE_TYPE), (*crosspb.SILingTongDataRemove)(nil))
}

func initProxy() {

}
