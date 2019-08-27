package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XUECHI_BLOODLINE_TYPE), (*uipb.CSXueChiBloodLine)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUECHI_BLOODLINE_TYPE), (*uipb.SCXueChiBloodLine)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUECHI_BLOOD_TYPE), (*uipb.SCXueChiBlood)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XUECHI_GET_TYPE), (*uipb.CSXueChiGet)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUECHI_GET_TYPE), (*uipb.SCXueChiGet)(nil))
}
