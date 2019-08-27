package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_USE_EQUIP_TYPE), (*uipb.CSTuLongUseEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_USE_EQUIP_TYPE), (*uipb.SCTuLongUseEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_TAKE_OFF_EQUIP_TYPE), (*uipb.CSTuLongTakeOffEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_TAKE_OFF_EQUIP_TYPE), (*uipb.SCTuLongTakeOffEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_EQUIP_STRENGTHEN_TYPE), (*uipb.CSTuLongEquipStrengthen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_EQUIP_STRENGTHEN_TYPE), (*uipb.SCTuLongEquipStrengthen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_EQUIP_RONGHE_TYPE), (*uipb.CSTuLongEquipRongHe)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_EQUIP_RONGHE_TYPE), (*uipb.SCTuLongEquipRongHe)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_EQUIP_ZHUANHUA_TYPE), (*uipb.CSTuLongEquipZhuanHua)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_EQUIP_ZHUANHUA_TYPE), (*uipb.SCTuLongEquipZhuanHua)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TULONG_EQUIP_SKILL_UPGRADE_TYPE), (*uipb.CSTuLongEquipSkillUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_EQUIP_SKILL_UPGRADE_TYPE), (*uipb.SCTuLongEquipSkillUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_EQUIP_SKILL_NOTICE_TYPE), (*uipb.SCTuLongEquipSkillNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_EQUIP_SLOT_CHANGED_TYPE), (*uipb.SCTuLongEquipSlotChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TULONG_EQUIP_INFO_NOTICE_TYPE), (*uipb.SCTuLongEquipInfoNotice)(nil))
}
