package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_UNREAL_BOSS_LIST_TYPE), (*uipb.CSUnrealBossList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_UNREAL_BOSS_LIST_TYPE), (*uipb.SCUnrealBossList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_UNREAL_BOSS_CHALLENGE_TYPE), (*uipb.CSUnrealBossChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_UNREAL_BOSS_CHALLENGE_TYPE), (*uipb.SCUnrealBossChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_UNREAL_BOSS_BUY_PILAO_TYPE), (*uipb.CSUnrealBossBuyPilaoNum)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_UNREAL_BOSS_BUY_PILAO_TYPE), (*uipb.SCUnrealBossBuyPilaoNum)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_UNREAL_BOSS_INFO_BROADCAST_TYPE), (*uipb.SCUnrealBossInfoBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_UNREAL_BOSS_LIST_INFO_NOTICE_TYPE), (*uipb.SCUnrealBossListInfoNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_UNREAL_BOSS_ENEMIES_NOTICE_TYPE), (*uipb.SCUnrealBossEnemiesNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_UNREAL_BOSS_PILAO_INFO_TYPE), (*uipb.SCUnrealBossPilaoInfo)(nil))

}
