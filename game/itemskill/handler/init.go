package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ITEM_SKILL_ACTIVE_TYPE), (*uipb.SCItemSkillActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ITEM_SKILL_UPGRADE_TYPE), (*uipb.SCItemSkillUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ITEM_SKILL_ALL_GET_TYPE), (*uipb.CSItemSkillAllGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ITEM_SKILL_ALL_GET_TYPE), (*uipb.SCItemSkillAllGet)(nil))

}
