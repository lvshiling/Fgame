package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHARE_BOSS_LIST_TYPE), (*uipb.CSShareBossList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHARE_BOSS_LIST_TYPE), (*uipb.SCShareBossList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHARE_BOSS_CHALLENGE_TYPE), (*uipb.CSShareBossChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHARE_BOSS_CHALLENGE_TYPE), (*uipb.SCShareBossChallenge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHARE_BOSS_INFO_BROADCAST_TYPE), (*uipb.SCShareBossInfoBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHARE_BOSS_LIST_INFO_NOTICE_TYPE), (*uipb.SCShareBossListInfoNotice)(nil))
}
