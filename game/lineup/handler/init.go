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
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINEUP_CANCEL_TYPE), (*uipb.CSLineupCancel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINEUP_CANCEL_TYPE), (*uipb.SCLineupCancel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINEUP_NOTICE_TYPE), (*uipb.SCLineupNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINEUP_SUCCESS_TYPE), (*uipb.SCLineupSuccess)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE), (*uipb.SCLineupSceneFinishToCancel)(nil))
}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_LINEUP_NOTICE_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_LINEUP_SUCCESS_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE))
}
