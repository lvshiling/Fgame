package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LIANYU_CANCLE_LINEUP_TYPE), (*uipb.CSLianYuCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_CANCLE_LINEUP_TYPE), (*uipb.SCLianYuCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_LINEUP_TYPE), (*uipb.SCLianYuLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_LINEUP_SUCCESS_TYPE), (*uipb.SCLianYuLineUpSuccess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_GET_TYPE), (*uipb.SCLianYuGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_BOSS_STATUS_TYPE), (*uipb.SCLianYuBossStatus)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_RANK_CHANGED_TYPE), (*uipb.SCLianYuRankChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_RESULT_TYPE), (*uipb.SCLianYuResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_FINISH_TO_LINEUP_TYPE), (*uipb.SCLianYuFinishToLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_SHAQI_CHANGED_TYPE), (*uipb.SCLianYuShaQiChanged)(nil))
}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_LIANYU_GET_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_LIANYU_BOSS_STATUS_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_LIANYU_RANK_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_LIANYU_RESULT_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_LIANYU_SHAQI_CHANGED_TYPE))
}
