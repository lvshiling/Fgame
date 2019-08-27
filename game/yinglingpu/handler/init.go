package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_YLPU_QUERY_TYPE), (*uipb.CSYingLingPuQuery)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YLPU_QUERY_TYPE), (*uipb.SCYingLingPuQuery)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_YLPU_SP_YLP_UP_TYPE), (*uipb.CSYingLingPuUpLevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YLPU_SP_YLP_UP_TYPE), (*uipb.SCYingLingPuUpLevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_YLPU_SP_XIANGQIAN_TYPE), (*uipb.CSYingLingPuSpXiangQian)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YLPU_SP_XIANGQIAN_TYPE), (*uipb.SCYingLingPuSpXiangQian)(nil))
}
