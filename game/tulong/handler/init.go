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
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_RANK_TYPE), (*uipb.CSTuLongRank)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_RANK_TYPE), (*uipb.SCTuLongRank)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_COLLECT_TYPE), (*uipb.CSTuLongCollect)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_TYPE), (*uipb.SCTuLongCollect)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_STOP_TYPE), (*uipb.SCTuLongCollectStop)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_FINISH_TYPE), (*uipb.SCTuLongCollectFinish)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_BOSS_STATUS_TYPE), (*uipb.SCTuLongBossStatus)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_START_TYPE), (*uipb.SCTuLongStart)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_ALLIANCE_BIAOSHI_TYPE), (*uipb.SCTuLongAllianceBiaoShi)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_RESULT_TYPE), (*uipb.SCTuLongResult)(nil))

}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_TULONG_COLLECT_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_STOP_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_FINISH_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TULONG_BOSS_STATUS_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TULONG_ALLIANCE_BIAOSHI_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TULONG_RESULT_TYPE))
}
