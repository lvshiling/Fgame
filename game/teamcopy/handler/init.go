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
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAMCOPY_ALL_GET_TYPE), (*uipb.CSTeamCopyAllGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_ALL_GET_TYPE), (*uipb.SCTeamCopyAllGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAMCOPY_START_BATTLE_TYPE), (*uipb.CSTeamCopyStartBattle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_START_BATTLE_TYPE), (*uipb.SCTeamCopyStartBattle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_START_BATTLE_BROADCAST_TYPE), (*uipb.SCTeamCopyStartBattleBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_START_BATTLE_RESULT_BROADCAST_TYPE), (*uipb.SCTeamCopyStartBattleResultBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_SCENE_INFO_TYPE), (*uipb.SCTeamCopySceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_PLAYER_DATA_CHANGED_TYPE), (*uipb.SCTeamCopyPlayerDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_RESULT_TYPE), (*uipb.SCTeamCopyResult)(nil))
}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_PLAYER_DATA_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TEAMCOPY_SCENE_INFO_TYPE))
}
