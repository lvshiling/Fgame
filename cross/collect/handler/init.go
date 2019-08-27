package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_TYPE), (*uipb.CSSceneCollect)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_TYPE), (*uipb.SCSceneCollect)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_STOP_TYPE), (*uipb.SCSceneCollectStop)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_FINISH_TYPE), (*uipb.SCSceneCollectFinish)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_OPEN_TYPE), (*uipb.CSSceneCollectMiZangOpen)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_GIVEUP_TYPE), (*uipb.CSSceneCollectMiZangGiveup)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_MIZANG_OPEN_TYPE), (*uipb.SCSceneCollectMiZangOpen)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_MIZANG_GIVEUP_TYPE), (*uipb.SCSceneCollectMiZangGiveup)(nil))

}
