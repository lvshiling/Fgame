package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WORLD_BOSS_LIST_TYPE), (*uipb.CSWorldBossList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WORLD_BOSS_LIST_TYPE), (*uipb.SCWorldBossList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHALLENGE_WORLD_BOSS_TYPE), (*uipb.CSChallengeWorldBoss)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHALLENGE_WORLD_BOSS_TYPE), (*uipb.SCChallengeWorldBoss)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WORLD_BOSS_INFO_BROADCAST_TYPE), (*uipb.SCWorldBossInfoBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WORLD_BOSS_LIST_INFO_NOTICE_TYPE), (*uipb.SCWorldBossListInfoNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WORLD_BOSS_RELIVE_TIME_NOTICE_TYPE), (*uipb.SCWorldBossReliveTimeNotice)(nil))
}
