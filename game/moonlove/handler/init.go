package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_GENEROUS_CHANGED_TYPE), (*uipb.SCMoonloveGenerousChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_CHARM_CHANGED_TYPE), (*uipb.SCMoonloveCharmChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOONLOVE_VIEW_DOUBLE_TYPE), (*uipb.CSMoonloveViewDouble)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_VIEW_DOUBLE_TYPE), (*uipb.SCMoonloveViewDouble)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOONLOVE_VIEW_DOUBLE_STATE_TYPE), (*uipb.CSMoonloveViewDoubleState)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_VIEW_DOUBLE_STATE_TYPE), (*uipb.SCMoonloveViewDoubleState)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_VIEW_DOUBLE_RELEASE_TYPE), (*uipb.SCMoonloveViewDoubleRelease)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_SCENE_INFO_TYPE), (*uipb.SCMoonloveSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_SCENE_RESULT_TYPE), (*uipb.SCMoonloveSceneResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_RANK_REWARDS_TYPE), (*uipb.SCMoonloveRankRewards)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_PUSH_CHARM_RANK_TYPE), (*uipb.SCMoonlovePushCharmRank)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_PUSH_GENEROUS_RANK_TYPE), (*uipb.SCMoonlovePushGenerousRank)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOONLOVE_PLAYER_LIST_TYPE), (*uipb.CSMoonlovePlayerList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_PLAYER_LIST_TYPE), (*uipb.SCMoonlovePlayerList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_GIFT_NOTICE_TYPE), (*uipb.SCMoonloveGiftNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOONLOVE_EXP_COUNT_NOTICE_TYPE), (*uipb.SCMoonloveExpCountNotice)(nil))
}
