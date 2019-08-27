package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_INFO_GET_TYPE), (*uipb.CSEquipbaokuInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EQUIPBAOKU_INFO_GET_TYPE), (*uipb.SCEquipbaokuInfoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_LOG_INCR_TYPE), (*uipb.CSEquipbaokuLogIncr)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EQUIPBAOKU_LOG_INCR_TYPE), (*uipb.SCEquipbaokuLogIncr)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_ATTEND_TYPE), (*uipb.CSEquipbaokuAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EQUIPBAOKU_ATTEND_TYPE), (*uipb.SCEquipbaokuAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_ATTEND_BATCH_TYPE), (*uipb.CSEquipbaokuAttendBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EQUIPBAOKU_ATTEND_BATCH_TYPE), (*uipb.SCEquipbaokuAttendBatch)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_LUCKY_BOX_TYPE), (*uipb.CSEquipbaokuLuckyBox)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EQUIPBAOKU_LUCKY_BOX_TYPE), (*uipb.SCEquipbaokuLuckyBox)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_POINTS_EXCHANGE_TYPE), (*uipb.CSEquipbaokuPointsExchange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EQUIPBAOKU_POINTS_EXCHANGE_TYPE), (*uipb.SCEquipbaokuPointsExchange)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_RESOLVE_EQUIP_TYPE), (*uipb.CSEquipbaokuResolveEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EQUIPBAOKU_RESOLVE_EQUIP_TYPE), (*uipb.SCEquipbaokuResolveEquip)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_SHOP_LIMIT_TYPE), (*uipb.CSEquipbaokuShopLimit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EQUIPBAOKU_SHOP_LIMIT_TYPE), (*uipb.SCEquipbaokuShopLimit)(nil))
}
