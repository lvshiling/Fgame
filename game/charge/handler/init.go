package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHARGE_TYPE), (*uipb.CSCharge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHARGE_TYPE), (*uipb.SCCharge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHARGE_ORDER_TYPE), (*uipb.SCChargeOrder)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FIRST_CHARGE_RECORD_NOTICE_TYPE), (*uipb.SCFirstChargeRecordNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHARGE_GOLD_NOTICE_TYPE), (*uipb.SCChargeGoldNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TODAY_CHARGE_GOLD_TYPE), (*uipb.SCTodayChargeGold)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_NEW_FIRST_CHARGE_RECORD), (*uipb.CSNewFirstChargeRecord)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_NEW_FIRST_CHARGE_RECORD), (*uipb.SCNewFirstChargeRecord)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_NEW_FIRST_CHARGE_RECORD_NOTICE_TYPE), (*uipb.SCNewFirstChargeRecordNotice)(nil))
}
