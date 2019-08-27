package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_GET_TYPE), (*uipb.CSLingTongDevGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_GET_TYPE), (*uipb.SCLingTongDevGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_UNREALDAN_TYPE), (*uipb.CSLingTongDevUnrealDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_UNREALDAN_TYPE), (*uipb.SCLingTongDevUnrealDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_UNREAL_TYPE), (*uipb.CSLingTongDevUnreal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_UNREAL_TYPE), (*uipb.SCLingTongDevUnreal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_ADVANCED_TYPE), (*uipb.CSLingTongDevAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_ADVANCED_TYPE), (*uipb.SCLingTongDevAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_UNLOAD_TYPE), (*uipb.CSLingTongDevUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_UNLOAD_TYPE), (*uipb.SCLingTongDevUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_HIDDEN_TYPE), (*uipb.CSLingTongDevHidden)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_HIDDEN_TYPE), (*uipb.SCLingTongDevHidden)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_UPSTAR_TYPE), (*uipb.CSLingTongDevUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_UPSTAR_TYPE), (*uipb.SCLingTongDevUpstar)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_CULDAN_TYPE), (*uipb.CSLingTongDevCulDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_CULDAN_TYPE), (*uipb.SCLingTongDevCulDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_TONGLING_TYPE), (*uipb.CSLingTongDevTongLing)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONGDEV_TONGLING_TYPE), (*uipb.SCLingTongDevTongLing)(nil))
}
