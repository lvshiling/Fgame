package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MY_BOSS_CHALLENGE_TYPE), (*uipb.CSMyBossChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MY_BOSS_CHALLENGE_TYPE), (*uipb.SCMyBossChallenge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MY_BOSS_INFO_NOTICE_TYPE), (*uipb.SCMyBossInfoNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MY_BOSS_CHALLENGE_RESULT_TYPE), (*uipb.SCMyBossChallengeResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MY_BOSS_SCENE_INFO_TYPE), (*uipb.SCMyBossSceneInfo)(nil))
}
