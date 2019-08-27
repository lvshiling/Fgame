package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGYU_GET_TYPE), (*uipb.CSLingyuGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGYU_GET_TYPE), (*uipb.SCLingyuGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGYU_UNREALDAN_TYPE), (*uipb.CSLingyuUnrealDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGYU_UNREALDAN_TYPE), (*uipb.SCLingyuUnrealDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGYU_UNREAL_TYPE), (*uipb.CSLingyuUnreal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGYU_UNREAL_TYPE), (*uipb.SCLingyuUnreal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGYU_ADVANCED_TYPE), (*uipb.CSLingyuAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGYU_ADVANCED_TYPE), (*uipb.SCLingyuAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGYU_UNLOAD_TYPE), (*uipb.CSLingyuUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGYU_UNLOAD_TYPE), (*uipb.SCLingyuUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGYU_HIDDEN_TYPE), (*uipb.CSLingyuHidden)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGYU_HIDDEN_TYPE), (*uipb.SCLingyuHidden)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGYU_UPSTAR_TYPE), (*uipb.CSLingYuUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGYU_UPSTAR_TYPE), (*uipb.SCLingYuUpstar)(nil))
}
