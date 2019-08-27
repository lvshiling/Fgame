package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENG_TAN_SCENE_INFO_TYPE), (*uipb.SCShengTanSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENG_TAN_SCENE_BOSS_HP_CHANGED_TYPE), (*uipb.SCShengTanSceneBossHpChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENG_TAN_SCENE_JIU_NIANG_CHANGED_TYPE), (*uipb.SCShengTanSceneJiuNiangChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENG_TAN_SCENE_END_TYPE), (*uipb.SCShengTanSceneEnd)(nil))

}
