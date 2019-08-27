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
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINEUP_NOTICE_TYPE), (*uipb.SCLineupNotice)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINEUP_SUCCESS_TYPE), (*uipb.SCLineupSuccess)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE), (*uipb.SCLineupSceneFinishToCancel)(nil))
}
