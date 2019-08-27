package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TOWER_ENTER_TYPE), (*uipb.CSTowerEnter)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TOWER_ENTER_TYPE), (*uipb.SCTowerEnter)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TOWER_BOSS_INFO_NOTICE_TYPE), (*uipb.SCTowerBossInfoNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TOWER_TIME_NOTICE_TYPE), (*uipb.SCTowerTimeNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TOWER_FLOOR_LIST_TYPE), (*uipb.CSTowerFloorList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TOWER_FLOOR_LIST_TYPE), (*uipb.SCTowerFloorList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TOWER_LOG_INCR_TYPE), (*uipb.CSTowerLogIncr)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TOWER_LOG_INCR_TYPE), (*uipb.SCTowerLogIncr)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TOWER_DA_BAO_TYPE), (*uipb.CSTowerDaBao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TOWER_DA_BAO_TYPE), (*uipb.SCTowerDaBao)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TOWER_RESULT_NOTICE_TYPE), (*uipb.SCTowerResultNotice)(nil))
}
