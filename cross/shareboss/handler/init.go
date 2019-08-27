package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHARE_BOSS_LIST_TYPE), (*uipb.CSShareBossList)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHARE_BOSS_LIST_TYPE), (*uipb.SCShareBossList)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHARE_BOSS_CHALLENGE_TYPE), (*uipb.CSShareBossChallenge)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHARE_BOSS_CHALLENGE_TYPE), (*uipb.SCShareBossChallenge)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHARE_BOSS_INFO_BROADCAST_TYPE), (*uipb.SCShareBossInfoBroadcast)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHARE_BOSS_LIST_INFO_NOTICE_TYPE), (*uipb.SCShareBossListInfoNotice)(nil))

}
