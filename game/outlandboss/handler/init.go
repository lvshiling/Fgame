package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_OUTLAND_BOSS_LIST_TYPE), (*uipb.CSOutlandBossList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OUTLAND_BOSS_LIST_TYPE), (*uipb.SCOutlandBossList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_OUTLAND_BOSS_CHALLENGE_TYPE), (*uipb.CSOutlandBossChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OUTLAND_BOSS_CHALLENGE_TYPE), (*uipb.SCOutlandBossChallenge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OUTLAND_BOSS_INFO_BROADCAST_TYPE), (*uipb.SCOutlandBossInfoBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OUTLAND_BOSS_LIST_INFO_NOTICE_TYPE), (*uipb.SCOutlandBossListInfoNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OUTLAND_BOSS_ENEMIES_NOTICE_TYPE), (*uipb.SCOutlandBossEnemiesNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OUTLAND_BOSS_ZHUOQI_INFO_TYPE), (*uipb.SCOutlandBossZhuoqiInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_OUTLAND_BOSS_DROP_RECORDS_INCR_TYPE), (*uipb.CSOutlandBossDropRecordsIncr)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OUTLAND_BOSS_DROP_RECORDS_INCR_TYPE), (*uipb.SCOutlandBossDropRecordsIncr)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_OUTLAND_BOSS_DROP_RECORDS_GET_TYPE), (*uipb.CSOutlandBossDropRecordsGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OUTLAND_BOSS_DROP_RECORDS_GET_TYPE), (*uipb.SCOutlandBossDropRecordsGet)(nil))
}
