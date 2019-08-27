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
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENAPVP_RACE_INFO_TYPE), (*uipb.CSArenapvpRaceInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_RACE_INFO_TYPE), (*uipb.SCArenapvpRaceInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENAPVP_BAZHU_INFO_TYPE), (*uipb.CSArenapvpBaZhuInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_BAZHU_INFO_TYPE), (*uipb.SCArenapvpBaZhuInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENAPVP_GUESS_INFO_TYPE), (*uipb.CSArenapvpGuessInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_GUESS_INFO_TYPE), (*uipb.SCArenapvpGuessInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENAPVP_GUESS_TYPE), (*uipb.CSArenapvpGuess)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_GUESS_TYPE), (*uipb.SCArenapvpGuess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_SCENE_DATA_TYPE), (*uipb.SCArenapvpSceneData)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_BATTLE_START_TYPE), (*uipb.SCArenapvpBattleStart)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_PLAYER_SHOW_DATA_CHANGED_TYPE), (*uipb.SCArenapvpPlayerShowDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_GUESS_BEGIN_NOTICE_TYPE), (*uipb.SCArenapvpGuessBeginNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENAPVP_INFO_TYPE), (*uipb.CSArenapvpInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_INFO_TYPE), (*uipb.SCArenapvpInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_BATTLE_END_TYPE), (*uipb.SCArenapvpBattleEnd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_FAILED_NOTICE_TYPE), (*uipb.SCArenapvpElectionFailedNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_SCENE_DATA_TYPE), (*uipb.SCArenapvpElectionSceneData)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_SCENE_DATA_CHANGED_TYPE), (*uipb.SCArenapvpElectionSceneDataChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENAPVP_ELECTION_RACE_INFO_TYPE), (*uipb.CSArenapvpElectionRaceInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_RACE_INFO_TYPE), (*uipb.SCArenapvpElectionRaceInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENAPVP_GUESS_NOTICE_SETTING_TYPE), (*uipb.CSArenapvpGuessNoticeSetting)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_GUESS_NOTICE_SETTING_TYPE), (*uipb.SCArenapvpGuessNoticeSetting)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_JIFEN_NOTICE_TYPE), (*uipb.SCArenapvpJiFenChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_END_TYPE), (*uipb.SCArenapvpElectionEnd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENAPVP_GUESS_INFO_PUSH_TYPE), (*uipb.SCArenapvpGuessInfoPush)(nil))
}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENAPVP_SCENE_DATA_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENAPVP_BATTLE_START_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENAPVP_PLAYER_SHOW_DATA_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENAPVP_BATTLE_END_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_FAILED_NOTICE_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_SCENE_DATA_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_SCENE_DATA_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENAPVP_ELECTION_END_TYPE))
}
