package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_DEVOURING_TYPE), (*uipb.CSWushuangWeaponDevouring)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WUSHUANGWEAPON_DEVOURING_TYPE), (*uipb.SCWushuangWeaponDevouring)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_BREAKTHROUGH_TYPE), (*uipb.CSWushuangWeaponBreakthrough)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WUSHUANGWEAPON_BREAKTHROUGH_TYPE), (*uipb.SCWushuangWeaponBreakthrough)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_PUT_ON_TYPE), (*uipb.CSWushuangWeaponPutOn)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WUSHUANGWEAPON_PUT_ON_TYPE), (*uipb.SCWushuangWeaponPutOn)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_TAKE_OFF_TYPE), (*uipb.CSWushuangWeaponTakeOff)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WUSHUANGWEAPON_TAKE_OFF_TYPE), (*uipb.SCWushuangWeaponTakeOff)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_INFO_TYPE), (*uipb.CSWushuangWeaponInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WUSHUANGWEAPON_INFO_TYPE), (*uipb.SCWushuangWeaponInfo)(nil))
}
