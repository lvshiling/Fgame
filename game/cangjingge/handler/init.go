package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CANGJINGGE_BOSS_LIST_TYPE), (*uipb.CSCangjinggeBossList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CANGJINGGE_BOSS_LIST_TYPE), (*uipb.SCCangjinggeBossList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CANGJINGGE_BOSS_CHALLENGE_TYPE), (*uipb.CSCangjinggeBossChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CANGJINGGE_BOSS_CHALLENGE_TYPE), (*uipb.SCCangjinggeBossChallenge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CANGJINGGE_BOSS_INFO_BROADCAST_TYPE), (*uipb.SCCangjinggeBossInfoBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CANGJINGGE_BOSS_LIST_INFO_NOTICE_TYPE), (*uipb.SCCangjinggeBossListInfoNotice)(nil))
}
