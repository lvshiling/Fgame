package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_USE_ADDITION_SYS_EQUIP_TYPE), (*uipb.CSUseAdditionSysEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_USE_ADDITION_SYS_EQUIP_TYPE), (*uipb.SCUseAdditionSysEquip)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TAKE_OFF_ADDITION_SYS_EQUIP_TYPE), (*uipb.CSTakeOffAdditionSysEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TAKE_OFF_ADDITION_SYS_EQUIP_TYPE), (*uipb.SCTakeOffAdditionSysEquip)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_STRENGTHEN_BODY_TYPE), (*uipb.CSAdditionSysStrengthenBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_STRENGTHEN_BODY_TYPE), (*uipb.SCAdditionSysStrengthenBody)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_SLOT_CHANGED_TYPE), (*uipb.SCAdditionSysSlotChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_SLOT_INFO_TYPE), (*uipb.SCAdditionSysSlotInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_SHENG_JI_TYPE), (*uipb.CSAdditionSysShengJi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_SHENG_JI_TYPE), (*uipb.SCAdditionSysShengJi)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_UPGRADE_TYPE), (*uipb.CSAdditionSysUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_UPGRADE_TYPE), (*uipb.SCAdditionSysUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_HUALING_EAT_TYPE), (*uipb.CSAdditionSysHualingEat)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_HUALING_EAT_TYPE), (*uipb.SCAdditionSysHualingEat)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_SHENZHU_BODY_TYPE), (*uipb.CSAdditionSysShenZhuBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_SHENZHU_BODY_TYPE), (*uipb.SCAdditionSysShenZhuBody)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_AWAKE_EAT_TYPE), (*uipb.CSAdditionSysAwakeEat)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_AWAKE_EAT_TYPE), (*uipb.SCAdditionSysAwakeEat)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_TONGLING_UPGRADE_TYPE), (*uipb.CSAdditionSysTongLingUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_TONGLING_UPGRADE_TYPE), (*uipb.SCAdditionSysTongLingUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_LINGTONG_LINGZHU_INFO_TYPE), (*uipb.CSAdditionSysLingTongLingZhuInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_LINGTONG_LINGZHU_INFO_TYPE), (*uipb.SCAdditionSysLingTongLingZhuInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_LINGTONG_LINGZHU_UPLEVEL_TYPE), (*uipb.CSAdditionSysLingTongLingZhuUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADDITION_SYS_LINGTONG_LINGZHU_UPLEVEL_TYPE), (*uipb.SCAdditionSysLingTongLingZhuUplevel)(nil))
}
