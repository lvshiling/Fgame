package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEEK_INFO_TYPE), (*uipb.CSWeekInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEEK_INFO_TYPE), (*uipb.SCWeekInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEEK_BUY_TYPE), (*uipb.CSWeekBuy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEEK_BUY_TYPE), (*uipb.SCWeekBuy)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WEEK_RECEIVE_REW_TYPE), (*uipb.CSWeekReceiveRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WEEK_RECEIVE_REW_TYPE), (*uipb.SCWeekReceiveRew)(nil))
}
