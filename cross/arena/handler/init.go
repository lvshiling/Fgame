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

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_MATCH_BROADCAST_TYPE), (*uipb.SCArenaMatchBroadcast)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_MATCH_RESULT_TYPE), (*uipb.SCArenaMatchResult)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_STOP_MATCH_TYPE), (*uipb.CSArenaStopMatch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_STOP_MATCH_TYPE), (*uipb.SCArenaStopMatch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_STOP_MATCH_BROADCAST_TYPE), (*uipb.SCArenaStopMatchBroadcast)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_SELECT_FOUR_GOD_TYPE), (*uipb.CSArenaSelectFourGod)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SELECT_FOUR_GOD_TYPE), (*uipb.SCArenaSelectFourGod)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SELECT_FOUR_GOD_BROADCAST_TYPE), (*uipb.SCArenaSelectFourGodBroadcast)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_QUEUE_TYPE), (*uipb.SCArenaFourGodQueue)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_CANCEL_QUEUE_TYPE), (*uipb.CSArenaFourGodCancelQueue)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_CANCEL_QUEUE_TYPE), (*uipb.SCArenaFourGodCancelQueue)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_CANCEL_QUEUE_BROADCAST_TYPE), (*uipb.SCArenaFourGodCancelQueueBroadcast)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_INFO_TYPE), (*uipb.SCArenaSceneInfo)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_START_TYPE), (*uipb.SCArenaSceneStart)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_END_TYPE), (*uipb.SCArenaSceneEnd)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_INVITE_TYPE), (*uipb.CSArenaInvite)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_INVITE_TYPE), (*uipb.SCArenaInvite)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_INFO_TYPE), (*uipb.SCArenaFourGodSceneInfo)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_NEXT_MATCH_TYPE), (*uipb.CSArenaNextMatch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_NEXT_MATCH_TYPE), (*uipb.SCArenaNextMatch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_NEXT_MATCH_BROADCAST_TYPE), (*uipb.SCArenaNextMatchBroadcast)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_LIST_TYPE), (*uipb.CSArenaFourGodList)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_LIST_TYPE), (*uipb.SCArenaFourGodList)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_QUEUE_CHANGED_TYPE), (*uipb.SCArenaFourGodQueueChanged)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_SCENE_COLLECTING_TYPE), (*uipb.CSArenaFourGodSceneCollecting)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECTING_TYPE), (*uipb.SCArenaFourGodSceneCollecting)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECT_STOP_TYPE), (*uipb.SCArenaFourGodSceneCollectStop)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECT_TYPE), (*uipb.SCArenaFourGodSceneCollect)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_TEAM_INFO_LIST_TYPE), (*uipb.CSArenaFourGodTeamInfoList)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_TEAM_INFO_LIST_TYPE), (*uipb.SCArenaFourGodTeamInfoList)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FINISH_TYPE), (*uipb.SCArenaFinish)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_TEAM_CHANGED_TYPE), (*uipb.SCArenaFourGodTeamChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_PLAYER_DATA_CHANGED_TYPE), (*uipb.SCArenaPlayerDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_PLAYER_EXIT_TYPE), (*uipb.SCArenaPlayerExit)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_DEAD_TYPE), (*uipb.SCFourGodSceneBossDead)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_REBORN_TYPE), (*uipb.SCFourGodSceneBossReborn)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_EXP_TREE_REBORN_TYPE), (*uipb.SCFourGodSceneExpTreeReborn)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_EXP_TREE_DEAD_TYPE), (*uipb.SCFourGodSceneExpTreeDead)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_HP_CHANGED_TYPE), (*uipb.SCFourGodSceneBossHpChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ARENA_DATA_CHANGED_TYPE), (*uipb.SCPlayerArenaDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ARENA_DATA_TYPE), (*uipb.SCPlayerArenaData)(nil))
}
