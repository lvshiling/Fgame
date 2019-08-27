package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YUXI_POS_BROADCAST_TYPE), (*uipb.SCYuXiPosBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YUXI_COLLECT_INFO_BROADCAST_TYPE), (*uipb.SCYuXiCollectInfoBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_YUXI_RECEIVE_DAY_REW_TYPE), (*uipb.CSYuXiReceiveDayRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YUXI_RECEIVE_DAY_REW_TYPE), (*uipb.SCYuXiReceiveDayRew)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_YUXI_GET_IFNO_TYPE), (*uipb.CSYuXiGetInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YUXI_GET_IFNO_TYPE), (*uipb.SCYuXiGetInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YUXI_WINNER_BROADCAST_TYPE), (*uipb.SCYuXiWinnerBroadcast)(nil))
}
