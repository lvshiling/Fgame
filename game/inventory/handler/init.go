package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_GET_TYPE), (*uipb.SCInventoryGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_MERGE_TYPE), (*uipb.CSInventoryMerge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_MERGE_TYPE), (*uipb.SCInventoryMerge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_CHANGED_TYPE), (*uipb.SCInventoryChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_ITEM_USE_TYPE), (*uipb.CSInventoryItemUse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_ITEM_USE_TYPE), (*uipb.SCInventoryItemUse)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_ITEM_SELL_TYPE), (*uipb.CSInventoryItemSell)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_ITEM_SELL_TYPE), (*uipb.SCInventoryItemSell)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_ITEM_SELL_BATCH_TYPE), (*uipb.CSInventoryItemSellBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_ITEM_SELL_BATCH_TYPE), (*uipb.SCInventoryItemSellBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_BUY_SLOTS_TYPE), (*uipb.CSInventoryBuySlots)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_BUY_SLOTS_TYPE), (*uipb.SCInventoryBuySlots)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_SLOTS_TYPE), (*uipb.SCInventorySlots)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_EQUIPMENT_CHANGED_TYPE), (*uipb.SCInventoryEquipmentSlotChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_USE_EQUIP_TYPE), (*uipb.CSInventoryUseEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_USE_EQUIP_TYPE), (*uipb.SCInventoryUseEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_TAKE_OFF_EQUIP_TYPE), (*uipb.CSInventoryTakeOffEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_TAKE_OFF_EQUIP_TYPE), (*uipb.SCInventoryTakeOffEquip)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_EQUIP_STRENGTHEN_TYPE), (*uipb.CSInventoryEquipStrengthen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_EQUIP_STRENGTHEN_TYPE), (*uipb.SCInventoryEquipStrengthen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_EQUIP_UPGRADE_TYPE), (*uipb.CSInventoryEquipUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_EQUIP_UPGRADE_TYPE), (*uipb.SCInventoryEquipUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_EQUIP_STRENGTHEN_ALL_TYPE), (*uipb.CSInventoryEquipStrengthenAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_EQUIP_STRENGTHEN_ALL_TYPE), (*uipb.SCInventoryEquipStrengthenAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_USE_GEM_TYPE), (*uipb.CSInventoryUseGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_USE_GEM_TYPE), (*uipb.SCInventoryUseGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_TAKE_OFF_GEM_TYPE), (*uipb.CSInventoryTakeOffGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_TAKE_OFF_GEM_TYPE), (*uipb.SCInventoryTakeOffGem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_USE_GEM_ALL_TYPE), (*uipb.CSInventoryUseGemAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_USE_GEM_ALL_TYPE), (*uipb.SCInventoryUseGemAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_BOX_DROP_INFO_TYPE), (*uipb.SCInventoryBoxDropInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SAVE_IN_DEPOT_TYPE), (*uipb.CSSaveInDepot)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SAVE_IN_DEPOT_TYPE), (*uipb.SCSaveInDepot)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DEPOT_TAKE_OUT_TYPE), (*uipb.CSDepotTakeOut)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DEPOT_TAKE_OUT_TYPE), (*uipb.SCDepotTakeOut)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DEPOT_BUY_SLOTS_TYPE), (*uipb.CSDepotBuySlots)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DEPOT_BUY_SLOTS_TYPE), (*uipb.SCDepotBuySlots)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DEPOT_MERGE_TYPE), (*uipb.CSDepotMerge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DEPOT_MERGE_TYPE), (*uipb.SCDepotMerge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DEPOT_CHANGED_TYPE), (*uipb.SCDepotChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_ITEM_USE_CHANGED_NOTICE_TYPE), (*uipb.SCInventoryItemUseChangedNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MIBAO_DEPOT_CHANGED_TYPE), (*uipb.SCMibaoDepotChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MIBAO_DEPOT_TAKE_OUT_TYPE), (*uipb.CSMibaoDepotTakeOut)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MIBAO_DEPOT_TAKE_OUT_TYPE), (*uipb.SCMibaoDepotTakeOut)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_INVENTORY_ITEM_DECOMPOSE_TYPE), (*uipb.CSInventoryItemDecompose)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_INVENTORY_ITEM_DECOMPOSE_TYPE), (*uipb.SCInventoryItemDecompose)(nil))
}
