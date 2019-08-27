package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_GET_TYPE), (*uipb.SCSkillGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_ADD_TYPE), (*uipb.SCSkillAdd)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SKILL_UPGRADE_TYPE), (*uipb.CSSkillUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_UPGRADE_TYPE), (*uipb.SCSkillUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SKILL_UPGRADE_ALL_TYPE), (*uipb.CSSkillUpgradeAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_UPGRADE_ALL_TYPE), (*uipb.SCSkillUpgradeAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_REMOVE_TYPE), (*uipb.SCSkillRemove)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_CD_TIME_TYPE), (*uipb.SCSkillCdTime)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_LEARN_TYPE), (*uipb.SCSkillLearn)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SKILL_TIANFU_AWAKEN_TYPE), (*uipb.CSSkillTianFuAwaken)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_TIANFU_AWAKEN_TYPE), (*uipb.SCSkillTianFuAwaken)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SKILL_TIANFU_UPGRADE_TYPE), (*uipb.CSSkillTianFuUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_TIANFU_UPGRADE_TYPE), (*uipb.SCSkillTianFuUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SKILL_TIANFU_GET_TYPE), (*uipb.SCSkillTianFuGet)(nil))
}
