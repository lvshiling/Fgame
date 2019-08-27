package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_DONGFANG_TYPE), (*uipb.CSBabyDongFang)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_DONGFANG_TYPE), (*uipb.SCBabyDongFang)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_BORN_ACCELERATE_TYPE), (*uipb.CSBabyBornAccelerate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_BORN_ACCELERATE_TYPE), (*uipb.SCBabyBornAccelerate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_EAT_TONIC_TYPE), (*uipb.CSBabyEatTonic)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_EAT_TONIC_TYPE), (*uipb.SCBabyEatTonic)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_BORN_CHAOSHENG_TYPE), (*uipb.CSBabyBornChaoSheng)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_BORN_CHAOSHENG_TYPE), (*uipb.SCBabyBornChaoSheng)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_CHANGE_NAME_TYPE), (*uipb.CSBabyChangeName)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_CHANGE_NAME_TYPE), (*uipb.SCBabyChangeName)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_ACTIVATE_SKILL_TYPE), (*uipb.CSBabyActivateSkill)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_ACTIVATE_SKILL_TYPE), (*uipb.SCBabyActivateSkill)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_LOCK_SKILL_TYPE), (*uipb.CSBabyLockSkill)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_LOCK_SKILL_TYPE), (*uipb.SCBabyLockSkill)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_REFRESH_SKILL_TYPE), (*uipb.CSBabyRefreshSkill)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_REFRESH_SKILL_TYPE), (*uipb.SCBabyRefreshSkill)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_LEARN_UPLEVEL_TYPE), (*uipb.CSBabyLearnUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_LEARN_UPLEVEL_TYPE), (*uipb.SCBabyLearnUplevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_INFO_TYPE), (*uipb.CSBabyInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_INFO_TYPE), (*uipb.SCBabyInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_ZHUAN_SHI_TYPE), (*uipb.CSBabyZhuanShi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_ZHUAN_SHI_TYPE), (*uipb.SCBabyZhuanShi)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_EQUIP_TOY_TYPE), (*uipb.CSBabyEquipToy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_EQUIP_TOY_TYPE), (*uipb.SCBabyEquipToy)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BABY_TOY_UPLEVEL_TYPE), (*uipb.CSBabyToyUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_TOY_UPLEVEL_TYPE), (*uipb.SCBabyToyUplevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_TOY_SLOT_CHANGED_TYPE), (*uipb.SCBabyToySlotChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_BORN_NOTICE_TYPE), (*uipb.SCBabyBornNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_BORN_SPOUSE_NOTICE_TYPE), (*uipb.SCBabyBornSpouseNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_BORN_MESSAGE_NOTICE_TYPE), (*uipb.SCBabyBornMessageNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BABY_POWER_NOTICE_TYPE), (*uipb.SCBabyPowerNotice)(nil))
}
