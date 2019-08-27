package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WING_GET_TYPE), (*uipb.CSWingGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_GET_TYPE), (*uipb.SCWingGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WING_UNREALDAN_TYPE), (*uipb.CSWingUnrealDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_UNREALDAN_TYPE), (*uipb.SCWingUnrealDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WING_UNREAL_TYPE), (*uipb.CSWingUnreal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_UNREAL_TYPE), (*uipb.SCWingUnreal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WING_ADVANCED_TYPE), (*uipb.CSWingAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_ADVANCED_TYPE), (*uipb.SCWingAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WING_UNLOAD_TYPE), (*uipb.CSWingUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_UNLOAD_TYPE), (*uipb.SCWingUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEATHER_GET_TYPE), (*uipb.CSFeatherGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEATHER_GET_TYPE), (*uipb.SCFeatherGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEATHER_ADVANCED_TYPE), (*uipb.CSFeatherAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEATHER_ADVANCED_TYPE), (*uipb.SCFeatherAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WING_HIDDEN_TYPE), (*uipb.CSWingHidden)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_HIDDEN_TYPE), (*uipb.SCWingHidden)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_USE_TRIAL_CARD_TYPE), (*uipb.SCWingTrialCard)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_TRIAL_OVERDUE_TYPE), (*uipb.SCWingTrialOverdue)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WING_UPSTAR_TYPE), (*uipb.CSWingUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WING_UPSTAR_TYPE), (*uipb.SCWingUpstar)(nil))
}
