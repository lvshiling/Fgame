package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QIXUE_GET_TYPE), (*uipb.CSQiXueGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QIXUE_GET_TYPE), (*uipb.SCQiXueGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QIXUE_ADVANCED_TYPE), (*uipb.CSQiXueAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QIXUE_ADVANCED_TYPE), (*uipb.SCQiXueAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QIXUE_WEAPON_LOSE_INFO_TYPE), (*uipb.SCQiXueWeaponLoseInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QIXUE_SHA_QI_VARY_TYPE), (*uipb.SCQiXueShaQiVary)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QIXUE_SHA_QI_DROP_TYPE), (*uipb.SCQiXueShaQiDrop)(nil))
}
