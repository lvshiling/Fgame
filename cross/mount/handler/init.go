package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLYAER_MOUNT_SYNC_TYPE), (*crosspb.ISPlayerMountSync)(nil))
}

func initProxy() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_HIDDEN_TYPE), (*uipb.CSMountHidden)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_HIDDEN_TYPE), (*uipb.SCMountHidden)(nil))
}
