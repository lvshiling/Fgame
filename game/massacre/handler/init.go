package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MASSACRE_GET_TYPE), (*uipb.CSMassacreGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MASSACRE_GET_TYPE), (*uipb.SCMassacreGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MASSACRE_ADVANCED_TYPE), (*uipb.CSMassacreAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MASSACRE_ADVANCED_TYPE), (*uipb.SCMassacreAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MASSACRE_INFO_TYPE), (*uipb.SCMassacreInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MASSACRE_WEAPON_LOSE_INFO_TYPE), (*uipb.SCMassacreWeaponLoseInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MASSACRE_SHA_QI_VARY_TYPE), (*uipb.SCMassacreShaQiVary)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MASSACRE_SHA_QI_DROP_TYPE), (*uipb.SCMassacreShaQiDrop)(nil))
}
