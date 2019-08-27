package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()

}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_LINEUP_TYPE), (*uipb.SCLianYuLineUp)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_GET_TYPE), (*uipb.SCLianYuGet)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_RANK_CHANGED_TYPE), (*uipb.SCLianYuRankChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_BOSS_STATUS_TYPE), (*uipb.SCLianYuBossStatus)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_RESULT_TYPE), (*uipb.SCLianYuResult)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIANYU_SHAQI_CHANGED_TYPE), (*uipb.SCLianYuShaQiChanged)(nil))
}
