package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WORLD_BOSS_INFO_BROADCAST_TYPE), (*uipb.SCWorldBossInfoBroadcast)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WORLD_BOSS_LIST_INFO_NOTICE_TYPE), (*uipb.SCWorldBossListInfoNotice)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WORLD_BOSS_RELIVE_TIME_NOTICE_TYPE), (*uipb.SCWorldBossReliveTimeNotice)(nil))
}
