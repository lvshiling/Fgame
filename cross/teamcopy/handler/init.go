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
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_SCENE_INFO_TYPE), (*uipb.SCTeamCopySceneInfo)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_PLAYER_DATA_CHANGED_TYPE), (*uipb.SCTeamCopyPlayerDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_RESULT_TYPE), (*uipb.SCTeamCopyResult)(nil))
}
