package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_USE_GOLD_EQUIP_TYPE), (*uipb.CSUseGoldEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_USE_GOLD_EQUIP_TYPE), (*uipb.SCUseGoldEquip)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TAKE_OFF_GOLD_EQUIP_TYPE), (*uipb.CSTakeOffGoldEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TAKE_OFF_GOLD_EQUIP_TYPE), (*uipb.SCTakeOffGoldEquip)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_STRENGTHEN_BAG_TYPE), (*uipb.CSGoldEquipStrengthenBag)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_STRENGTHEN_BAG_TYPE), (*uipb.SCGoldEquipStrengthenBag)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_STRENGTHEN_BODY_TYPE), (*uipb.CSGoldEquipStrengthenBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_STRENGTHEN_BODY_TYPE), (*uipb.SCGoldEquipStrengthenBody)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_CHONGZHU_TYPE), (*uipb.CSGoldEquipChongzhu)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_CHONGZHU_TYPE), (*uipb.SCGoldEquipChongzhu)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_SLOT_CHANGED_TYPE), (*uipb.SCGoldEquipSlotChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_SLOT_INFO_TYPE), (*uipb.SCGoldEquipSlotInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ZHUANSHENG_TYPE), (*uipb.CSZhuanSheng)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ZHUANSHENG_TYPE), (*uipb.SCZhuanSheng)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EAT_GOLD_EQUIP_TYP), (*uipb.CSEatGoldEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EAT_GOLD_EQUIP_TYPE), (*uipb.SCEatGoldEquip)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_OPENLIGHT_BODY_TYPE), (*uipb.CSGoldEquipOpenLightBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_OPENLIGHT_BODY_TYPE), (*uipb.SCGoldEquipOpenLightBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_OPENLIGHT_BAG_TYPE), (*uipb.CSGoldEquipOpenLightBag)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_OPENLIGHT_BAG_TYPE), (*uipb.SCGoldEquipOpenLightBag)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_UPSTAR_BODY_TYPE), (*uipb.CSGoldEquipUpstarBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_UPSTAR_BODY_TYPE), (*uipb.SCGoldEquipUpstarBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_UPSTAR_BAG_TYPE), (*uipb.CSGoldEquipUpstarBag)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_UPSTAR_BAG_TYPE), (*uipb.SCGoldEquipUpstarBag)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_USE_GEM_TYPE), (*uipb.CSGoldEquipUseGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_USE_GEM_TYPE), (*uipb.SCGoldEquipUseGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_TAKE_OFF_GEM_TYPE), (*uipb.CSGoldEquipTakeOffGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_TAKE_OFF_GEM_TYPE), (*uipb.SCGoldEquipTakeOffGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_USE_GEM_ALL_TYPE), (*uipb.CSGoldEquipUseGemAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_USE_GEM_ALL_TYPE), (*uipb.SCGoldEquipUseGemAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_EXTEND_BODY_TYPE), (*uipb.CSGoldEquipExtendBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_EXTEND_BODY_TYPE), (*uipb.SCGoldEquipExtendBody)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_EXTEND_BAG_TYPE), (*uipb.CSGoldEquipExtendBag)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_EXTEND_BAG_TYPE), (*uipb.SCGoldEquipExtendBag)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_UNLOCK_GEM_TYPE), (*uipb.CSGoldEquipUnlockGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_UNLOCK_GEM_TYPE), (*uipb.SCGoldEquipUnlockGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_LOG_TYPE), (*uipb.CSGoldEquipLog)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_LOG_TYPE), (*uipb.SCGoldEquipLog)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_AUTO_FENJIE_TYPE), (*uipb.CSGoldEquipAutoFenJie)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_AUTO_FENJIE_TYPE), (*uipb.SCGoldEquipAutoFenJie)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_STRENGTHEN_BUWEI), (*uipb.CSGoldEquipStrengthenBuwei)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_STRENGTHEN_BUWEI), (*uipb.SCGoldEquipStrengthenBuwei)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GOLD_EQUIP_USE_ITEM_WITH_GROW_UP_TYPE), (*uipb.SCGoldEquipUseItemWithGrowUp)(nil))

	//神铸
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GODCASTING_EQUIP_UPLEVEL_TYPE), (*uipb.CSGodCastingEquipUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODCASTING_EQUIP_UPLEVEL_TYPE), (*uipb.SCGodCastingEquipUplevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GODCASTING_CASTINGSPIRIT_UPLEVEL_TYPE), (*uipb.CSGodCastingCastingSpiritUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODCASTING_CASTINGSPIRIT_UPLEVEL_TYPE), (*uipb.SCGodCastingCastingSpiritUplevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GODCASTING_FORGESOUL_UPLEVEL_TYPE), (*uipb.CSGodCastingForgeSoulUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODCASTING_FORGESOUL_UPLEVEL_TYPE), (*uipb.SCGodCastingForgeSoulUplevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GODCASTING_EQUIP_INHERIT_TYPE), (*uipb.CSGodCastingEquipInherit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GODCASTING_EQUIP_INHERIT_TYPE), (*uipb.SCGodCastingEquipInherit)(nil))
}
