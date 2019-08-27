package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FABAO_GET_TYPE), (*uipb.CSFaBaoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FABAO_GET_TYPE), (*uipb.SCFaBaoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FABAO_UNREALDAN_TYPE), (*uipb.CSFaBaoUnrealDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FABAO_UNREALDAN_TYPE), (*uipb.SCFaBaoUnrealDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FABAO_UNREAL_TYPE), (*uipb.CSFaBaoUnreal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FABAO_UNREAL_TYPE), (*uipb.SCFaBaoUnreal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FABAO_ADVANCED_TYPE), (*uipb.CSFaBaoAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FABAO_ADVANCED_TYPE), (*uipb.SCFaBaoAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FABAO_UNLOAD_TYPE), (*uipb.CSFaBaoUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FABAO_UNLOAD_TYPE), (*uipb.SCFaBaoUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FABAO_HIDDEN_TYPE), (*uipb.CSFaBaoHidden)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FABAO_HIDDEN_TYPE), (*uipb.SCFaBaoHidden)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FABAO_UPSTAR_TYPE), (*uipb.CSFaBaoUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FABAO_UPSTAR_TYPE), (*uipb.SCFaBaoUpstar)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FABAO_TONGLING_TYPE), (*uipb.CSFaBaoTongLing)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FABAO_TONGLING_TYPE), (*uipb.SCFaBaoTongLing)(nil))
}
