package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BUY_HUIYUAN_TYPE), (*uipb.CSBuyHuiYuan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BUY_HUIYUAN_TYPE), (*uipb.SCBuyHuiYuan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HUIYUAN_RECEIVE_RWE_TYPE), (*uipb.CSHuiYuanReceiveRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HUIYUAN_RECEIVE_REW_TYPE), (*uipb.SCHuiYuanReceiveRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HUIYUAN_INFO_TYPE), (*uipb.CSHuiYuanInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HUIYUAN_INFO_TYPE), (*uipb.SCHuiYuanInfo)(nil))
}
