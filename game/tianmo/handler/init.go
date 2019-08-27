package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TIANMOTI_GET_TYPE), (*uipb.CSTianMoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANMOTI_GET_TYPE), (*uipb.SCTianMoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TIANMOTI_EAT_DAN_TYPE), (*uipb.CSTianMoEatDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANMOTI_EAT_DAN_TYPE), (*uipb.SCTianMoEatDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TIANMOTI_ADVANCED_TYPE), (*uipb.CSTianMoAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANMOTI_ADVANCED_TYPE), (*uipb.SCTianMoAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANMOTI_INFO_TYPE), (*uipb.SCTianMoInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANMOTI_CHARGE_VARY_TYPE), (*uipb.SCTianMoChargeGold)(nil))
}
