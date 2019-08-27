package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XUECHI_BLOODLINE_TYPE), (*uipb.CSXueChiBloodLine)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUECHI_BLOODLINE_TYPE), (*uipb.SCXueChiBloodLine)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XUECHI_AUTO_BUY_TYPE), (*uipb.CSXueChiAutoBuy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUECHI_AUTO_BUY_TYPE), (*uipb.SCXueChiAutoBuy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUECHI_BLOOD_TYPE), (*uipb.SCXueChiBlood)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XUECHI_GET_TYPE), (*uipb.CSXueChiGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUECHI_GET_TYPE), (*uipb.SCXueChiGet)(nil))
}

func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_XUECHI_BLOODLINE_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_XUECHI_BLOODLINE_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_XUECHI_BLOOD_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_XUECHI_GET_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_XUECHI_GET_TYPE))
}
