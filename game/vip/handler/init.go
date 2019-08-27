package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_VIP_INFO_NOTICE_TYPE), (*uipb.SCVipInfoNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_VIP_GIFT_BUY_TYPE), (*uipb.CSVipGiftBuy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_VIP_GIFT_BUY_TYPE), (*uipb.SCVipGiftBuy)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RECEIVE_FREE_GIFT_TYPE), (*uipb.CSReceiveFreeGift)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RECEIVE_FREE_GIFT_TYPE), (*uipb.SCReceiveFreeGift)(nil))
}
