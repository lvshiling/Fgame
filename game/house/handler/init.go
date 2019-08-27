package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HOUSE_LIST_GET_TYPE), (*uipb.CSHouseListGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HOUSE_LIST_GET_TYPE), (*uipb.SCHouseListGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HOUSE_ACTIVATE_TYPE), (*uipb.CSHouseActivate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HOUSE_ACTIVATE_TYPE), (*uipb.SCHouseActivate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HOUSE_UPGRADE_TYPE), (*uipb.CSHouseUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HOUSE_UPGRADE_TYPE), (*uipb.SCHouseUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HOUSE_SELL_TYPE), (*uipb.CSHouseSell)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HOUSE_SELL_TYPE), (*uipb.SCHouseSell)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HOUSE_RECEIVE_RENT_TYPE), (*uipb.CSHouseReceiveRent)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HOUSE_RECEIVE_RENT_TYPE), (*uipb.SCHouseReceiveRent)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HOUSE_REPAIR_TYPE), (*uipb.CSHouseRepair)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HOUSE_REPAIR_TYPE), (*uipb.SCHouseRepair)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HOUSE_LOG_INCR_TYPE), (*uipb.CSHouseLogIncr)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HOUSE_LOG_INCR_TYPE), (*uipb.SCHouseLogIncr)(nil))
}
