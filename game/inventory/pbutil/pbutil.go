package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	babypbutil "fgame/fgame/game/baby/pbutil"
	babytypes "fgame/fgame/game/baby/types"
	droppbutil "fgame/fgame/game/drop/pbutil"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	ringtypes "fgame/fgame/game/ring/types"
)

func BuildSCInventorySlots(num, depotNum int32) *uipb.SCInventorySlots {
	scInventorySlots := &uipb.SCInventorySlots{}
	scInventorySlots.BagNum = &num
	scInventorySlots.DepotNum = &depotNum
	return scInventorySlots
}

func buildItemList(itemList []*playerinventory.PlayerItemObject) (slotItemList []*uipb.SlotItem) {
	for _, item := range itemList {
		slotItemList = append(slotItemList, buildItem(item))
	}
	return slotItemList
}

func buildItemUseList(itemUseMap map[int32]*playerinventory.PlayerItemUseObject) (slotItemUseList []*uipb.SlotItemUse) {
	for _, itemUse := range itemUseMap {
		slotItemUseList = append(slotItemUseList, buildItemUse(itemUse))
	}
	return slotItemUseList
}
func buildItemUse(itemUse *playerinventory.PlayerItemUseObject) *uipb.SlotItemUse {
	SlotItemUse := &uipb.SlotItemUse{}
	itemId := itemUse.ItemId
	SlotItemUse.ItemId = &itemId
	totalTimes := itemUse.TotalTimes
	SlotItemUse.TotalTimes = &totalTimes
	todayTimes := itemUse.TodayTimes
	SlotItemUse.TodayTimes = &todayTimes
	lastUseTime := itemUse.LastUseTime
	SlotItemUse.LastUseTime = &lastUseTime
	return SlotItemUse
}

func buildItem(item *playerinventory.PlayerItemObject) *uipb.SlotItem {
	slotItem := &uipb.SlotItem{}
	itemId := item.ItemId
	slotItem.ItemId = &itemId
	num := item.Num
	slotItem.Num = &num
	index := item.Index
	slotItem.Index = &index
	bagType := int32(item.BagType)
	slotItem.BagType = &bagType
	level := item.Level
	slotItem.Level = &level
	getTime := item.ItemGetTime
	slotItem.GetTime = &getTime
	bindInt := int32(item.BindType)
	slotItem.BindType = &bindInt
	slotItem.PropertyData = BuildItemPropertyData(item.PropertyData)
	return slotItem
}

func BuildItemPropertyData(data inventorytypes.ItemPropertyData) *uipb.ItemPropertyData {
	info := &uipb.ItemPropertyData{}
	info.Base = buildBaseProperty(int32(data.GetExpireType()), data.GetExpireTimestamp(), data.GetItemGetTime())
	propertyData, ok := data.(*goldequiptypes.GoldEquipPropertyData)
	if ok {
		info.Goldequip = buildGoldEquipProperty(propertyData)
	}

	babyData, ok := data.(*babytypes.BabyPropertyData)
	if ok {
		info.Baby = buildBabyProperty(babyData)
	}

	ringData, ok := data.(*ringtypes.RingPropertyData)
	if ok {
		info.Ring = buildRingProperty(ringData)
	}

	return info
}

func buildGoldEquipProperty(data *goldequiptypes.GoldEquipPropertyData) *uipb.GoldEquipProperty {
	openLightLevel := data.OpenLightLevel
	upstarLevel := data.UpstarLevel
	godcastingTimes := data.GodCastingTimes

	info := &uipb.GoldEquipProperty{}
	info.OpenLightLevel = &openLightLevel
	info.UpstarLevel = &upstarLevel
	info.AttrList = data.AttrList
	info.GodCastingTimes = &godcastingTimes
	return info
}

func buildRingProperty(data *ringtypes.RingPropertyData) *uipb.RingProperty {
	advance := data.Advance
	strengthenLevel := data.StrengthLevel
	jingLingLevel := data.JingLingLevel

	info := &uipb.RingProperty{}
	info.Advance = &advance
	info.StrengthenLevel = &strengthenLevel
	info.JingLingLevel = &jingLingLevel

	return info
}

func buildBabyProperty(data *babytypes.BabyPropertyData) *uipb.BabyProperty {
	quality := data.Quality
	sexInt := int32(data.Sex)
	danbei := data.Danbei

	info := &uipb.BabyProperty{}
	info.Quality = &quality
	info.Sex = &sexInt
	info.Danbei = &danbei
	info.TalentList = babypbutil.BuildTalentInfoList(data.TalentList)

	return info
}

func buildBaseProperty(expireType int32, expireTime, itemGetTime int64) *uipb.BaseProperty {
	info := &uipb.BaseProperty{}
	info.ExpireType = &expireType
	info.ExpireTime = &expireTime
	info.ItemGetTime = &itemGetTime
	return info
}

func BuildSCInventoryGetAll(
	itemList, depotList []*playerinventory.PlayerItemObject,
	slotList []*playerinventory.PlayerEquipmentSlotObject,
	itemUseMap map[int32]*playerinventory.PlayerItemUseObject,
	miBaoDepotList, materialDepotList []*playerinventory.PlayerItemObject) *uipb.SCInventoryGet {
	inventoryGet := &uipb.SCInventoryGet{}
	inventoryGet.ItemList = buildItemList(itemList)
	inventoryGet.SlotList = buildEquipmentSlotList(slotList)
	inventoryGet.DepotItemList = buildItemList(depotList)
	inventoryGet.ItemUseList = buildItemUseList(itemUseMap)
	inventoryGet.MiBaoDepotItemList = buildItemList(miBaoDepotList)
	inventoryGet.MaterialDepotItemList = buildItemList(materialDepotList)

	return inventoryGet
}

func BuildSCInventoryItemUseChangedNotice(itemUseMap map[int32]*playerinventory.PlayerItemUseObject) *uipb.SCInventoryItemUseChangedNotice {
	scInventoryItemUseChangedNotice := &uipb.SCInventoryItemUseChangedNotice{}
	scInventoryItemUseChangedNotice.ItemUseList = buildItemUseList(itemUseMap)
	return scInventoryItemUseChangedNotice
}

func BuildSCInventoryBuySlots(slotsNum int32) *uipb.SCInventoryBuySlots {
	inventoryBuySlots := &uipb.SCInventoryBuySlots{}
	inventoryBuySlots.Num = &slotsNum
	return inventoryBuySlots
}

func BuildSCDepotBuySlots(totalNum int32) *uipb.SCDepotBuySlots {
	scDepotBuySlots := &uipb.SCDepotBuySlots{}
	scDepotBuySlots.TotalNum = &totalNum
	return scDepotBuySlots
}

func BuildSCSaveInDepot(itemList []*playerinventory.PlayerItemObject) *uipb.SCSaveInDepot {
	scSaveInDepot := &uipb.SCSaveInDepot{}
	scSaveInDepot.ItemList = buildItemList(itemList)
	return scSaveInDepot
}

func BuildSCDepotTakeOut(itemList []*playerinventory.PlayerItemObject) *uipb.SCDepotTakeOut {
	scDepotTakeOut := &uipb.SCDepotTakeOut{}
	scDepotTakeOut.ItemList = buildItemList(itemList)
	return scDepotTakeOut
}

func BuildSCInventoryItemUse(bagType inventorytypes.BagType, index int32, itemId, num int32) *uipb.SCInventoryItemUse {
	scInventoryItemUse := &uipb.SCInventoryItemUse{}
	bagTypeInt := int32(bagType)
	scInventoryItemUse.BagType = &bagTypeInt
	scInventoryItemUse.Index = &index
	scInventoryItemUse.Num = &num
	scInventoryItemUse.ItemId = &itemId

	return scInventoryItemUse
}

func BuildSCInventoryMerge(bagType inventorytypes.BagType, itemList []*playerinventory.PlayerItemObject) *uipb.SCInventoryMerge {
	inventoryMerge := &uipb.SCInventoryMerge{}
	inventoryMerge.ItemList = buildItemList(itemList)
	bagTypeInt := int32(bagType)
	inventoryMerge.BagType = &bagTypeInt
	return inventoryMerge
}

func BuildSCDepotMerge(itemList []*playerinventory.PlayerItemObject) *uipb.SCDepotMerge {
	scDepotMerge := &uipb.SCDepotMerge{}
	scDepotMerge.ItemList = buildItemList(itemList)
	return scDepotMerge
}

func BuildSCInventoryChanged(itemList []*playerinventory.PlayerItemObject) *uipb.SCInventoryChanged {
	inventoryChanged := &uipb.SCInventoryChanged{}
	inventoryChanged.ItemList = buildItemList(itemList)
	return inventoryChanged

}

func BuildSCDepotChanged(depotItemList []*playerinventory.PlayerItemObject) *uipb.SCDepotChanged {
	scDepotChanged := &uipb.SCDepotChanged{}
	scDepotChanged.DepotItemList = buildItemList(depotItemList)
	return scDepotChanged

}

func BuildSCInventoryItemSell(bagType inventorytypes.BagType, index int32, num int32) *uipb.SCInventoryItemSell {
	scInventoryItemSell := &uipb.SCInventoryItemSell{}
	scInventoryItemSell.Index = &index
	scInventoryItemSell.Num = &num
	bagTypeInt := int32(bagType)
	scInventoryItemSell.BagType = &bagTypeInt
	return scInventoryItemSell
}

func SCInventoryItemSellBatch(silver int64) *uipb.SCInventoryItemSellBatch {
	scInventoryItemSellBatch := &uipb.SCInventoryItemSellBatch{}
	scInventoryItemSellBatch.GainsNum = &silver
	return scInventoryItemSellBatch
}

func BuildSCInventoryEquipmentChanged(slotList []*playerinventory.PlayerEquipmentSlotObject) *uipb.SCInventoryEquipmentSlotChanged {
	inventoryChanged := &uipb.SCInventoryEquipmentSlotChanged{}
	inventoryChanged.SlotList = buildEquipmentSlotList(slotList)
	return inventoryChanged
}

func buildEquipmentSlotList(slotList []*playerinventory.PlayerEquipmentSlotObject) (slotItemList []*uipb.EquipmentSlot) {
	for _, slot := range slotList {
		slotItemList = append(slotItemList, buildEquipmentSlot(slot))
	}
	return slotItemList
}

func buildEquipmentSlot(slot *playerinventory.PlayerEquipmentSlotObject) *uipb.EquipmentSlot {
	slotItem := &uipb.EquipmentSlot{}
	slotId := int32(slot.SlotId)
	slotItem.SlotId = &slotId
	level := slot.Level
	slotItem.Level = &level
	star := slot.Star
	slotItem.Star = &star
	itemId := slot.ItemId
	slotItem.ItemId = &itemId
	bindInt := int32(slot.BindType)
	slotItem.BindType = &bindInt
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

func BuildSCInventoryTakeOffEquip(pos inventorytypes.BodyPositionType) *uipb.SCInventoryTakeOffEquip {
	scInventoryTakeOffEquip := &uipb.SCInventoryTakeOffEquip{}
	slotInt := int32(pos)
	scInventoryTakeOffEquip.SlotId = &slotInt
	return scInventoryTakeOffEquip
}

func BuildSCInventoryUseEquip(index int32) *uipb.SCInventoryUseEquip {
	scInventoryUseEquip := &uipb.SCInventoryUseEquip{}
	scInventoryUseEquip.Index = &index
	return scInventoryUseEquip
}

func BuildSCInventoryUseGemAll() *uipb.SCInventoryUseGemAll {
	scInventoryUseGemAll := &uipb.SCInventoryUseGemAll{}
	return scInventoryUseGemAll
}

func BuildSCInventoryEquipmentSlotStrength(pos inventorytypes.BodyPositionType, typ inventorytypes.EquipmentStrengthenType, result inventorytypes.EquipmentStrengthenResultType, auto bool) *uipb.SCInventoryEquipStrengthen {
	scInventoryEquipStrengthen := &uipb.SCInventoryEquipStrengthen{}
	slotInt := int32(pos)
	scInventoryEquipStrengthen.SlotId = &slotInt
	typInt := int32(typ)
	scInventoryEquipStrengthen.Typ = &typInt
	resultInt := int32(result)
	scInventoryEquipStrengthen.Result = &resultInt
	scInventoryEquipStrengthen.Auto = &auto
	return scInventoryEquipStrengthen
}

func BuildSCInventoryEquipmentUpgrade(pos inventorytypes.BodyPositionType, result inventorytypes.EquipmentStrengthenResultType) *uipb.SCInventoryEquipUpgrade {
	scInventoryEquipUpgrade := &uipb.SCInventoryEquipUpgrade{}
	slotInt := int32(pos)
	scInventoryEquipUpgrade.SlotId = &slotInt
	resultInt := int32(result)
	scInventoryEquipUpgrade.Result = &resultInt
	return scInventoryEquipUpgrade
}

func BuildSCInventoryEquipmentSlotStrengthAll(resultList []*inventorytypes.StrengthenResult) *uipb.SCInventoryEquipStrengthenAll {
	scInventoryEquipStrengthenAll := &uipb.SCInventoryEquipStrengthenAll{}
	for _, result := range resultList {
		scInventoryEquipStrengthenAll.ResultList = append(scInventoryEquipStrengthenAll.ResultList, BuildEquipStrengthenResult(result))
	}

	return scInventoryEquipStrengthenAll
}

func BuildEquipStrengthenResult(result *inventorytypes.StrengthenResult) *uipb.EquipStrengthenResult {
	equipStrengthenResult := &uipb.EquipStrengthenResult{}
	slotInt := int32(result.Pos)
	equipStrengthenResult.SlotId = &slotInt
	resultInt := int32(result.Result)
	equipStrengthenResult.Result = &resultInt
	return equipStrengthenResult
}

func BuildEquipmentSlotInfoList(slotList []*inventorytypes.EquipmentSlotInfo) (es []*uipb.EquipmentSlot) {
	for _, slot := range slotList {
		es = append(es, BuildEquipmentSlotInfo(slot))
	}
	return
}

func BuildEquipmentSlotInfo(slot *inventorytypes.EquipmentSlotInfo) (slotItem *uipb.EquipmentSlot) {
	slotItem = &uipb.EquipmentSlot{}
	slotId := int32(slot.SlotId)
	slotItem.SlotId = &slotId
	level := slot.Level
	slotItem.Level = &level
	star := slot.Star
	slotItem.Star = &star
	itemId := slot.ItemId
	slotItem.ItemId = &itemId
	for order, gemId := range slot.Gems {
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

func BuildSCMiBaoDepotChanged(depotItemList []*playerinventory.PlayerItemObject, typ int32) *uipb.SCMibaoDepotChanged {
	scDepotChanged := &uipb.SCMibaoDepotChanged{}
	scDepotChanged.Type = &typ
	scDepotChanged.DepotItemList = buildItemList(depotItemList)
	return scDepotChanged
}

func BuildSCMiBaoDepotTakeOut(itemList []*playerinventory.PlayerItemObject, typ int32, isBatch bool, index int32) *uipb.SCMibaoDepotTakeOut {
	scDepotTakeOut := &uipb.SCMibaoDepotTakeOut{}
	scDepotTakeOut.ItemList = buildItemList(itemList)
	scDepotTakeOut.Type = &typ
	scDepotTakeOut.IsBatch = &isBatch
	scDepotTakeOut.Index = &index
	return scDepotTakeOut
}

func BuildSCInventoryItemDecompose(bagType inventorytypes.BagType, index int32, num int32, rewItemMap map[int32]int32) *uipb.SCInventoryItemDecompose {
	scMsg := &uipb.SCInventoryItemDecompose{}
	bagTypeInt := int32(bagType)
	scMsg.BagType = &bagTypeInt
	scMsg.Index = &index
	scMsg.Num = &num
	scMsg.DropInfoList = droppbutil.BuildSimpleDropInfoList(rewItemMap)

	return scMsg
}
