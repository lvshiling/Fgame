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
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_LINEUP_TYPE), (*uipb.SCGodSiegeLineUp)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_GET_TYPE), (*uipb.SCGodSiegeGet)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_BOSS_STATUS_TYPE), (*uipb.SCGodSiegeBossStatus)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_RESULT_TYPE), (*uipb.SCGodSiegeResult)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_COLLECT_CHANGED_TYPE), (*uipb.SCGodSiegeCollectChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_COLLECT_NPC_CHANGED_TYPE), (*uipb.SCGodSiegeCollectNpcChanged)(nil))

}
