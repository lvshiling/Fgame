package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ANQI_GET_TYPE), (*uipb.CSAnqiGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ANQI_GET_TYPE), (*uipb.SCAnqiGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ANQI_EAT_DAN_TYPE), (*uipb.CSAnqiEatDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ANQI_EAT_DAN_TYPE), (*uipb.SCAnqiEatDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ANQI_ADVANCED_TYPE), (*uipb.CSAnqiAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ANQI_ADVANCED_TYPE), (*uipb.SCAnqiAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ANQI_INFO_TYPE), (*uipb.SCAnqiInfo)(nil))
}
