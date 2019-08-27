package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TRADE_INFO_LIST_TYPE), (*uipb.CSTradeInfoList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TRADE_INFO_LIST_TYPE), (*uipb.SCTradeInfoList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TRADE_UPLOAD_ITEM_TYPE), (*uipb.CSTradeUploadItem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TRADE_UPLOAD_ITEM_TYPE), (*uipb.SCTradeUploadItem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TRADE_WITHDRAW_ITEM_TYPE), (*uipb.CSTradeWithDrawItem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TRADE_WITHDRAW_ITEM_TYPE), (*uipb.SCTradeWithDrawItem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TRADE_ITEM_TYPE), (*uipb.CSTradeItem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TRADE_ITEM_TYPE), (*uipb.SCTradeItem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SELF_TRADE_INFO_LIST_TYPE), (*uipb.CSSelfTradeInfoList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SELF_TRADE_INFO_LIST_TYPE), (*uipb.SCSelfTradeInfoList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TRADE_LOG_LIST_TYPE), (*uipb.CSTradeLogList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TRADE_LOG_LIST_TYPE), (*uipb.SCTradeLogList)(nil))
}
