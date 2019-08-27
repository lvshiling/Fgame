package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENFA_GET_TYPE), (*uipb.CSShenfaGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENFA_GET_TYPE), (*uipb.SCShenfaGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENFA_UNREALDAN_TYPE), (*uipb.CSShenfaUnrealDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENFA_UNREALDAN_TYPE), (*uipb.SCShenfaUnrealDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENFA_UNREAL_TYPE), (*uipb.CSShenfaUnreal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENFA_UNREAL_TYPE), (*uipb.SCShenfaUnreal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENFA_ADVANCED_TYPE), (*uipb.CSShenfaAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENFA_ADVANCED_TYPE), (*uipb.SCShenfaAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENFA_UNLOAD_TYPE), (*uipb.CSShenfaUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENFA_UNLOAD_TYPE), (*uipb.SCShenfaUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENFA_HIDDEN_TYPE), (*uipb.CSShenfaHidden)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENFA_HIDDEN_TYPE), (*uipb.SCShenfaHidden)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENFA_UPSTAR_TYPE), (*uipb.CSShenFaUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENFA_UPSTAR_TYPE), (*uipb.SCShenFaUpstar)(nil))
}
