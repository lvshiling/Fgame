package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANTI_GET_TYPE), (*uipb.CSXiantiGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTI_GET_TYPE), (*uipb.SCXiantiGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANTI_UNREALDAN_TYPE), (*uipb.CSXiantiUnrealDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTI_UNREALDAN_TYPE), (*uipb.SCXiantiUnrealDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANTI_UNREAL_TYPE), (*uipb.CSXiantiUnreal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTI_UNREAL_TYPE), (*uipb.SCXiantiUnreal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANTI_ADVANCED_TYPE), (*uipb.CSXiantiAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTI_ADVANCED_TYPE), (*uipb.SCXiantiAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANTI_UNLOAD_TYPE), (*uipb.CSXiantiUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTI_UNLOAD_TYPE), (*uipb.SCXiantiUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANTI_HIDDEN_TYPE), (*uipb.CSXiantiHidden)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTI_HIDDEN_TYPE), (*uipb.SCXiantiHidden)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANTI_UPSTAR_TYPE), (*uipb.CSXiantiUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTI_UPSTAR_TYPE), (*uipb.SCXiantiUpstar)(nil))
}

func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_XIANTI_HIDDEN_TYPE))
}
