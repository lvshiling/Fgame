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
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_SCENE_DATA_TYPE), (*uipb.SCArenapvpSceneData)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_BATTLE_START_TYPE), (*uipb.SCArenapvpBattleStart)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_PLAYER_SHOW_DATA_CHANGED_TYPE), (*uipb.SCArenapvpPlayerShowDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_BATTLE_END_TYPE), (*uipb.SCArenapvpBattleEnd)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_FAILED_NOTICE_TYPE), (*uipb.SCArenapvpElectionFailedNotice)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_SCENE_DATA_TYPE), (*uipb.SCArenapvpElectionSceneData)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_SCENE_DATA_CHANGED_TYPE), (*uipb.SCArenapvpElectionSceneDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_END_TYPE), (*uipb.SCArenapvpElectionEnd)(nil))
}
