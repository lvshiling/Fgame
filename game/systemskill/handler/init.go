package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SYSTEM_SKILL_GET_TYPE), (*uipb.CSSystemSkillGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYSTEM_SKILL_GET_TYPE), (*uipb.SCSystemSkillGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SYSTEM_SKILL_ACTIVE_TYPE), (*uipb.CSSystemSkillActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYSTEM_SKILL_ACTIVE_TYPE), (*uipb.SCSystemSkillActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SYSTEM_SKILL_UPGRADE_TYPE), (*uipb.CSSystemSkillUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYSTEM_SKILL_UPGRADE_TYPE), (*uipb.SCSystemSkillUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SYSTEM_SKILL_ALL_GET_TYPE), (*uipb.CSSystemSkillAllGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYSTEM_SKILL_ALL_GET_TYPE), (*uipb.SCSystemSkillAllGet)(nil))

}
