package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WELFARE_SCENE_ATTEND_TYPE), (*uipb.CSWelfareSceneAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WELFARE_SCENE_ATTEND_TYPE), (*uipb.SCWelfareSceneAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WELFARE_SCENE_INFO_TYPE), (*uipb.SCWelfareSceneInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WELFARE_SCENE_DATA_CHANGED_NOTICE_TYPE), (*uipb.SCWelfareSceneDataChangedNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WELFARE_SCENE_REFRESH_TYPE), (*uipb.SCWelfareSceneRefersh)(nil))

}
