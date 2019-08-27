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

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GODSIEGE_CANCLE_LINEUP_TYPE), (*uipb.CSGodSiegeCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_CANCLE_LINEUP_TYPE), (*uipb.SCGodSiegeCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_LINEUP_TYPE), (*uipb.SCGodSiegeLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_LINEUP_SUCCESS_TYPE), (*uipb.SCGodSiegeLineUpSuccess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_GET_TYPE), (*uipb.SCGodSiegeGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_BOSS_STATUS_TYPE), (*uipb.SCGodSiegeBossStatus)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_RESULT_TYPE), (*uipb.SCGodSiegeResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_FINISH_TO_LINEUP_TYPE), (*uipb.SCGodSiegeFinishToLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_COLLECT_CHANGED_TYPE), (*uipb.SCGodSiegeCollectChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODSIEGE_COLLECT_NPC_CHANGED_TYPE), (*uipb.SCGodSiegeCollectNpcChanged)(nil))
}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_GODSIEGE_GET_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_GODSIEGE_BOSS_STATUS_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_GODSIEGE_RESULT_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_GODSIEGE_COLLECT_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_GODSIEGE_COLLECT_NPC_CHANGED_TYPE))
}
