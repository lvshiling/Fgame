package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEAPON_GET_TYPE), (*uipb.CSWeaponGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEAPON_GET_TYPE), (*uipb.SCWeaponGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEAPON_ACTIVE_TYPE), (*uipb.CSWeaponActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEAPON_ACTIVE_TYPE), (*uipb.SCWeaponActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEAPON_EATDAN_TYPE), (*uipb.CSWeaponEatDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEAPON_EATDAN_TYPE), (*uipb.SCWeaponEatDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEAPON_UPSTAR_TYPE), (*uipb.CSWeaponUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEAPON_UPSTAR_TYPE), (*uipb.SCWeaponUpstar)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEAPON_AWAKEN_TYPE), (*uipb.CSWeaponAwaken)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEAPON_AWAKEN_TYPE), (*uipb.SCWeaponAwaken)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEAPON_WEAR_TYPE), (*uipb.CSWeaponWear)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEAPON_WEAR_TYPE), (*uipb.SCWeaponWear)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEAPON_UNLOAD_TYPE), (*uipb.CSWeaponUnLoad)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEAPON_UNLOAD_TYPE), (*uipb.SCWeaponUnLoad)(nil))
}
