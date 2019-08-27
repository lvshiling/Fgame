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
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_MATCH_TYPE), (*uipb.CSArenaMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_MATCH_TYPE), (*uipb.SCArenaMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_MATCH_BROADCAST_TYPE), (*uipb.SCArenaMatchBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_MATCH_RESULT_TYPE), (*uipb.SCArenaMatchResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_STOP_MATCH_TYPE), (*uipb.CSArenaStopMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_STOP_MATCH_TYPE), (*uipb.SCArenaStopMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_STOP_MATCH_BROADCAST_TYPE), (*uipb.SCArenaStopMatchBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_SELECT_FOUR_GOD_TYPE), (*uipb.CSArenaSelectFourGod)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SELECT_FOUR_GOD_TYPE), (*uipb.SCArenaSelectFourGod)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SELECT_FOUR_GOD_BROADCAST_TYPE), (*uipb.SCArenaSelectFourGodBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_QUEUE_TYPE), (*uipb.SCArenaFourGodQueue)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_CANCEL_QUEUE_TYPE), (*uipb.CSArenaFourGodCancelQueue)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_CANCEL_QUEUE_TYPE), (*uipb.SCArenaFourGodCancelQueue)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_CANCEL_QUEUE_BROADCAST_TYPE), (*uipb.SCArenaFourGodCancelQueueBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_INFO_TYPE), (*uipb.SCArenaSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_START_TYPE), (*uipb.SCArenaSceneStart)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_END_TYPE), (*uipb.SCArenaSceneEnd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_INVITE_TYPE), (*uipb.CSArenaInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_INVITE_TYPE), (*uipb.SCArenaInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_INFO_TYPE), (*uipb.SCArenaFourGodSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_NEXT_MATCH_TYPE), (*uipb.CSArenaNextMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_NEXT_MATCH_TYPE), (*uipb.SCArenaNextMatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_NEXT_MATCH_BROADCAST_TYPE), (*uipb.SCArenaNextMatchBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_LIST_TYPE), (*uipb.CSArenaFourGodList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_LIST_TYPE), (*uipb.SCArenaFourGodList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_QUEUE_CHANGED_TYPE), (*uipb.SCArenaFourGodQueueChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_SCENE_COLLECTING_TYPE), (*uipb.CSArenaFourGodSceneCollecting)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECTING_TYPE), (*uipb.SCArenaFourGodSceneCollecting)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECT_STOP_TYPE), (*uipb.SCArenaFourGodSceneCollectStop)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECT_TYPE), (*uipb.SCArenaFourGodSceneCollect)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_TEAM_INFO_LIST_TYPE), (*uipb.CSArenaFourGodTeamInfoList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_TEAM_INFO_LIST_TYPE), (*uipb.SCArenaFourGodTeamInfoList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FINISH_TYPE), (*uipb.SCArenaFinish)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_TEAM_CHANGED_TYPE), (*uipb.SCArenaFourGodTeamChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_PLAYER_DATA_CHANGED_TYPE), (*uipb.SCArenaPlayerDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_PLAYER_EXIT_TYPE), (*uipb.SCArenaPlayerExit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_DEAD_TYPE), (*uipb.SCFourGodSceneBossDead)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_REBORN_TYPE), (*uipb.SCFourGodSceneBossReborn)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_EXP_TREE_REBORN_TYPE), (*uipb.SCFourGodSceneExpTreeReborn)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_EXP_TREE_DEAD_TYPE), (*uipb.SCFourGodSceneExpTreeDead)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_HP_CHANGED_TYPE), (*uipb.SCFourGodSceneBossHpChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ARENA_DATA_CHANGED_TYPE), (*uipb.SCPlayerArenaDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ARENA_DATA_TYPE), (*uipb.SCPlayerArenaData)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PLAYER_ARENA_INFO_TYPE), (*uipb.CSPlayerArenaInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ARENA_INFO_TYPE), (*uipb.SCPlayerArenaInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_GET_REWARD_TYPE), (*uipb.CSArenaGetReward)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_GET_REWARD_TYPE), (*uipb.SCArenaGetReward)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_MY_RANK_TYPE), (*uipb.CSArenaMyRank)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_MY_RANK_TYPE), (*uipb.SCArenaMyRank)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ARENA_RANK_GET_TYPE), (*uipb.CSArenaRankGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ARENA_RANK_GET_TYPE), (*uipb.SCArenaRankGet)(nil))
}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_ARENA_SELECT_FOUR_GOD_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_SELECT_FOUR_GOD_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_SELECT_FOUR_GOD_BROADCAST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_QUEUE_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_CANCEL_QUEUE_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_CANCEL_QUEUE_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_CANCEL_QUEUE_BROADCAST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_INFO_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_START_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_SCENE_END_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_INFO_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_ARENA_NEXT_MATCH_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_NEXT_MATCH_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_NEXT_MATCH_BROADCAST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_LIST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_LIST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_QUEUE_CHANGED_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_SCENE_COLLECTING_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECTING_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECT_STOP_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_SCENE_COLLECT_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_TEAM_INFO_LIST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_TEAM_INFO_LIST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FINISH_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_FOUR_GOD_TEAM_CHANGED_TYPE))

	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_PLAYER_DATA_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ARENA_PLAYER_EXIT_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_DEAD_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_REBORN_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_EXP_TREE_REBORN_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_EXP_TREE_DEAD_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_FOUR_GOD_SCENE_BOSS_HP_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_PLAYER_ARENA_DATA_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_PLAYER_ARENA_DATA_TYPE))

}
