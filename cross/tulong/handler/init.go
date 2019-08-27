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

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_COLLECT_TYPE), (*uipb.CSTuLongCollect)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_TYPE), (*uipb.SCTuLongCollect)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_STOP_TYPE), (*uipb.SCTuLongCollectStop)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_COLLECT_FINISH_TYPE), (*uipb.SCTuLongCollectFinish)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_BOSS_STATUS_TYPE), (*uipb.SCTuLongBossStatus)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_ALLIANCE_BIAOSHI_TYPE), (*uipb.SCTuLongAllianceBiaoShi)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_RESULT_TYPE), (*uipb.SCTuLongResult)(nil))
}
