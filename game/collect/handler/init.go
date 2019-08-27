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
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_TYPE), (*uipb.CSSceneCollect)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_TYPE), (*uipb.SCSceneCollect)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_STOP_TYPE), (*uipb.SCSceneCollectStop)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_FINISH_TYPE), (*uipb.SCSceneCollectFinish)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_CHOOSE_RESULT_TYPE), (*uipb.CSSceneCollectChooseResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_CHOOSE_RESULT_TYPE), (*uipb.SCSceneCollectChooseResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_OPEN_TYPE), (*uipb.CSSceneCollectMiZangOpen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_MIZANG_OPEN_TYPE), (*uipb.SCSceneCollectMiZangOpen)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_GIVEUP_TYPE), (*uipb.CSSceneCollectMiZangGiveup)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_COLLECT_MIZANG_GIVEUP_TYPE), (*uipb.SCSceneCollectMiZangGiveup)(nil))

}

func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_OPEN_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_GIVEUP_TYPE))
}
