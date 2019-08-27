package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	inventorypbutil "fgame/fgame/game/inventory/pbutil"
	playertulongequip "fgame/fgame/game/tulongequip/player"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
)

func BuildSCTuLongEquipSlotChanged(slotListMap map[tulongequiptypes.TuLongSuitType][]*playertulongequip.PlayerTuLongEquipSlotObject) *uipb.SCTuLongEquipSlotChanged {
	scMsg := &uipb.SCTuLongEquipSlotChanged{}
	for suitType, slotList := range slotListMap {
		typeInt := int32(suitType)
		info := &uipb.TuLongEquipInfo{}
		info.SuitType = &typeInt
		info.SlotList = buildEquipmentSlotList(slotList)

		scMsg.InfoList = append(scMsg.InfoList, info)
	}
	return scMsg
}

func BuildSCTuLongEquipInfoNotice(slotListMap map[tulongequiptypes.TuLongSuitType][]*playertulongequip.PlayerTuLongEquipSlotObject) *uipb.SCTuLongEquipInfoNotice {
	scMsg := &uipb.SCTuLongEquipInfoNotice{}
	for suitType, slotList := range slotListMap {
		typeInt := int32(suitType)
		info := &uipb.TuLongEquipInfo{}
		info.SuitType = &typeInt
		info.SlotList = buildEquipmentSlotList(slotList)

		scMsg.InfoList = append(scMsg.InfoList, info)
	}
	return scMsg
}

func BuildSCTuLongUseEquip(index, suitType int32) *uipb.SCTuLongUseEquip {
	scMsg := &uipb.SCTuLongUseEquip{}
	scMsg.Index = &index
	scMsg.SuitType = &suitType
	return scMsg
}

func BuildSCTuLongTakeOffEquip(suitType, pos int32) *uipb.SCTuLongTakeOffEquip {
	scMsg := &uipb.SCTuLongTakeOffEquip{}
	scMsg.SlotId = &pos
	scMsg.SuitType = &suitType
	return scMsg
}

func BuildSCTuLongEquipStrengthen(suitType, pos, level int32, success bool) *uipb.SCTuLongEquipStrengthen {
	scMsg := &uipb.SCTuLongEquipStrengthen{}
	scMsg.SlotId = &pos
	scMsg.Level = &level
	scMsg.SuitType = &suitType

	result := int32(0)
	if success {
		result = 1
	}
	scMsg.Result = &result
	return scMsg
}

func BuildSCTuLongEquipRongHe(dropItemDataList []*droptemplate.DropItemData, args int32) *uipb.SCTuLongEquipRongHe {
	scMsg := &uipb.SCTuLongEquipRongHe{}

	var itemIdList []int32
	for _, itemData := range dropItemDataList {
		itemId := itemData.GetItemId()
		itemIdList = append(itemIdList, itemId)
	}
	scMsg.ItemIdList = itemIdList
	scMsg.Args = &args
	return scMsg
}

func BuildSCTuLongEquipZhuanHua(itemId, slot int32) *uipb.SCTuLongEquipZhuanHua {
	scMsg := &uipb.SCTuLongEquipZhuanHua{}
	scMsg.ItemId = &itemId
	scMsg.SlotId = &slot
	return scMsg
}

func BuildSCTuLongEquipSkillUpgrade(suitType, skillLevel int32, success bool) *uipb.SCTuLongEquipSkillUpgrade {
	scMsg := &uipb.SCTuLongEquipSkillUpgrade{}
	scMsg.SuitType = &suitType
	scMsg.IsSuccess = &success
	scMsg.SkillLevel = &skillLevel
	return scMsg
}

func BuildSCTuLongEquipSkillNotice(skillMap map[tulongequiptypes.TuLongSuitType]*playertulongequip.PlayerTuLongSuitSkillObject) *uipb.SCTuLongEquipSkillNotice {
	scMsg := &uipb.SCTuLongEquipSkillNotice{}
	for suitType, skillObj := range skillMap {
		info := buildTuLongEquipSkillInfo(int32(suitType), skillObj.GetLevel())
		scMsg.InfoList = append(scMsg.InfoList, info)
	}
	return scMsg
}

func buildTuLongEquipSkillInfo(suitType, level int32) *uipb.TuLongEquipSkillInfo {
	info := &uipb.TuLongEquipSkillInfo{}
	info.SkillLevel = &level
	info.SuitType = &suitType
	return info
}

func buildEquipmentSlotList(slotList []*playertulongequip.PlayerTuLongEquipSlotObject) (slotItemList []*uipb.TuLongEquipSlot) {
	for _, slot := range slotList {
		slotItemList = append(slotItemList, buildEquipmentSlot(slot))
	}
	return slotItemList
}

func buildEquipmentSlot(slot *playertulongequip.PlayerTuLongEquipSlotObject) *uipb.TuLongEquipSlot {
	slotItem := &uipb.TuLongEquipSlot{}
	slotId := int32(slot.GetSlotId())
	bindInt := int32(slot.GetBindType())
	level := slot.GetLevel()
	itemId := slot.GetItemId()
	slotItem.SlotId = &slotId
	slotItem.Level = &level
	slotItem.ItemId = &itemId
	slotItem.BindType = &bindInt
	slotItem.PropertyData = inventorypbutil.BuildItemPropertyData(slot.GetPropertyData())
	for order, gemId := range slot.GemInfo {
		//注意:range 值拷贝临时变量 导致同一个地址
		tempOrder := order
		tempGemId := gemId
		gem := &uipb.GemSlot{}
		gem.Order = &tempOrder
		gem.ItemId = &tempGemId
		slotItem.Gems = append(slotItem.Gems, gem)
	}
	return slotItem
}

func buildBaseProperty(expireType int32, expireTime, itemGetTime int64) *uipb.BaseProperty {
	info := &uipb.BaseProperty{}
	info.ExpireType = &expireType
	info.ExpireTime = &expireTime
	info.ItemGetTime = &itemGetTime
	return info
}
