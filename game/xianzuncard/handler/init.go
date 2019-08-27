package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIAN_ZUN_CARD_INFO_TYPE), (*uipb.CSXianZunCardInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIAN_ZUN_CARD_INFO_TYPE), (*uipb.SCXianZunCardInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIAN_ZUN_CARD_BUY_TYPE), (*uipb.CSXianZunCardBuy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIAN_ZUN_CARD_BUY_TYPE), (*uipb.SCXianZunCardBuy)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RECEIVE_XIAN_ZUN_CARD_REWARD_TYPE), (*uipb.CSReceiveXianZunCardReward)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RECEIVE_XIAN_ZUN_CARD_REWARD_TYPE), (*uipb.SCReceiveXianZunCardReward)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIAN_ZUN_CARD_NOTICE_TYPE), (*uipb.SCXianZunCardNotice)(nil))
}
