package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorypbutil "fgame/fgame/game/inventory/pbutil"
	inventorytypes "fgame/fgame/game/inventory/types"
)

func BuildSCUseGoldEquip(index int32) *uipb.SCUseGoldEquip {
	scUseGoldEquip := &uipb.SCUseGoldEquip{}
	scUseGoldEquip.Index = &index
	return scUseGoldEquip
}

func BuildTakeOffGoldEquip(pos inventorytypes.BodyPositionType) *uipb.SCTakeOffGoldEquip {
	scTakeOffGoldEquip := &uipb.SCTakeOffGoldEquip{}
	slotInt := int32(pos)
	scTakeOffGoldEquip.SlotId = &slotInt
	return scTakeOffGoldEquip
}

func BuildSCGoldEquipSlotChanged(slotList []*playergoldequip.PlayerGoldEquipSlotObject) *uipb.SCGoldEquipSlotChanged {
	scGoldEquipSlotChanged := &uipb.SCGoldEquipSlotChanged{}
	scGoldEquipSlotChanged.SlotList = buildEquipmentSlotList(slotList)
	return scGoldEquipSlotChanged
}

func BuildSCGoldEquipStrengthenBag(index int32, result bool) *uipb.SCGoldEquipStrengthenBag {
	scGoldEquipStrengthenBag := &uipb.SCGoldEquipStrengthenBag{}
	scGoldEquipStrengthenBag.Index = &index
	scGoldEquipStrengthenBag.Result = &result
	return scGoldEquipStrengthenBag
}

func BuildSCGoldEquipStrengthenBody(pos inventorytypes.BodyPositionType, result bool) *uipb.SCGoldEquipStrengthenBody {
	scGoldEquipStrengthenBody := &uipb.SCGoldEquipStrengthenBody{}
	slotId := int32(pos)
	scGoldEquipStrengthenBody.SlotId = &slotId
	scGoldEquipStrengthenBody.Result = &result
	return scGoldEquipStrengthenBody
}

func BuildSCGoldEquipChongzhu(itemId int32, level int32) *uipb.SCGoldEquipChongzhu {
	scGoldEquipChongzhu := &uipb.SCGoldEquipChongzhu{}
	scGoldEquipChongzhu.ItemId = &itemId
	scGoldEquipChongzhu.Level = &level
	return scGoldEquipChongzhu
}

func BuildGoldEquipSlotList(slotList []*goldequiptypes.GoldEquipSlotInfo) (es []*uipb.GoldEquipSlot) {
	for _, slot := range slotList {
		es = append(es, buildGoldEquipSlot(slot))
	}
	return
}

func BuildGoldEquipSlotInfoList(slotList []*playergoldequip.PlayerGoldEquipSlotObject) *uipb.SCGoldEquipSlotInfo {
	scGoldEquipSlotInfo := &uipb.SCGoldEquipSlotInfo{}
	scGoldEquipSlotInfo.SlotList = buildEquipmentSlotList(slotList)
	return scGoldEquipSlotInfo
}

func BuildSCGoldEquipZhuanSheng(zhuanShengNum int32) *uipb.SCZhuanSheng {
	zhuanSheng := &uipb.SCZhuanSheng{}
	zhuanSheng.ZhuanShengShu = &zhuanShengNum
	return zhuanSheng
}

func BuildSCEatGoldEquip(level int32, exp int64) *uipb.SCEatGoldEquip {
	scEatGoldEquip := &uipb.SCEatGoldEquip{}
	scEatGoldEquip.GoldYuanLevel = &level
	scEatGoldEquip.GoldYuanExp = &exp
	return scEatGoldEquip
}

func BuildSCGoldEquipOpenLightBody(isSuccess bool, openLevel int32, posType inventorytypes.BodyPositionType) *uipb.SCGoldEquipOpenLightBody {
	scMsg := &uipb.SCGoldEquipOpenLightBody{}
	scMsg.IsSuccess = &isSuccess
	scMsg.OpenLevel = &openLevel
	slotId := int32(posType)
	scMsg.SlotId = &slotId
	return scMsg
}

func BuildSCGoldEquipOpenLightBag(isSuccess bool, openLevel int32, index int32) *uipb.SCGoldEquipOpenLightBag {
	scMsg := &uipb.SCGoldEquipOpenLightBag{}
	scMsg.IsSuccess = &isSuccess
	scMsg.OpenLevel = &openLevel
	scMsg.Index = &index
	return scMsg
}

func BuildSCGoldEquipUpstarBody(result goldequiptypes.UpstarResultType, starLevel int32, posType inventorytypes.BodyPositionType, isAvoid bool) *uipb.SCGoldEquipUpstarBody {
	scMsg := &uipb.SCGoldEquipUpstarBody{}
	resultInt := int32(result)
	scMsg.Result = &resultInt
	scMsg.StarLevel = &starLevel
	slotId := int32(posType)
	scMsg.SlotId = &slotId
	scMsg.IsAvoid = &isAvoid
	return scMsg
}

func BuildSCGoldEquipUpstarBag(result goldequiptypes.UpstarResultType, starLevel int32, index int32, isAvoid bool) *uipb.SCGoldEquipUpstarBag {
	scMsg := &uipb.SCGoldEquipUpstarBag{}
	resultInt := int32(result)
	scMsg.Result = &resultInt
	scMsg.StarLevel = &starLevel
	scMsg.Index = &index
	scMsg.IsAvoid = &isAvoid
	return scMsg
}

func BuildSCGoldEquipUseGemAll() *uipb.SCGoldEquipUseGemAll {
	scMsg := &uipb.SCGoldEquipUseGemAll{}
	return scMsg
}

func BuildSCGoldEquipTakeOffGem(slot, order int32) *uipb.SCGoldEquipTakeOffGem {
	scMsg := &uipb.SCGoldEquipTakeOffGem{}
	scMsg.SlotId = &slot
	scMsg.Order = &order
	return scMsg
}

func BuildSCGoldEquipUseGem(index, slot, order int32) *uipb.SCGoldEquipUseGem {
	scMsg := &uipb.SCGoldEquipUseGem{}
	scMsg.Index = &index
	scMsg.SlotId = &slot
	scMsg.Order = &order
	return scMsg
}

func BuildSCGoldEquipExtendBag(targetIndex, useIndex, targetLevel int32) *uipb.SCGoldEquipExtendBag {
	scMsg := &uipb.SCGoldEquipExtendBag{}
	scMsg.ItemIndex = &targetIndex
	scMsg.UseItemIndex = &useIndex
	scMsg.ItemUpstarLevel = &targetLevel

	return scMsg
}

func BuildSCGoldEquipExtendBody(slotId, useIndex, targetLevel int32) *uipb.SCGoldEquipExtendBody {
	scMsg := &uipb.SCGoldEquipExtendBody{}
	scMsg.SlotId = &slotId
	scMsg.UseItemIndex = &useIndex
	scMsg.ItemUpstarLevel = &targetLevel

	return scMsg
}

func BuildSCGoldEquipLog(logList []*playergoldequip.PlayerGoldEquipLogObject) *uipb.SCGoldEquipLog {
	scMsg := &uipb.SCGoldEquipLog{}
	scMsg.LogList = buildGoldEquipLogList(logList)
	return scMsg
}

func BuildSCGoldEquipAutoFenJie(isAuto, quality, zhuanShu int32) *uipb.SCGoldEquipAutoFenJie {
	scMsg := &uipb.SCGoldEquipAutoFenJie{}
	scMsg.IsAuto = &isAuto
	scMsg.MaxQuality = &quality
	scMsg.ZhuanShu = &zhuanShu
	return scMsg
}

func buildGoldEquipSlot(slot *goldequiptypes.GoldEquipSlotInfo) *uipb.GoldEquipSlot {
	slotItem := &uipb.GoldEquipSlot{}
	slotId := slot.SlotId
	level := slot.Level
	newStLevel := slot.NewStLevel
	itemId := slot.ItemId
	slotItem.SlotId = &slotId
	slotItem.Level = &level
	slotItem.NewStLevel = &newStLevel
	slotItem.ItemId = &itemId
	slotItem.PropertyData = inventorypbutil.BuildItemPropertyData(slot.PropertyData)
	for order, gemId := range slot.Gems {
		//注意:range 值拷贝临时变量 导致同一个地址
		tempOrder := order
		tempGemId := gemId
		gem := &uipb.GemSlot{}
		gem.Order = &tempOrder
		gem.ItemId = &tempGemId
		slotItem.Gems = append(slotItem.Gems, gem)
	}
	for order := range slot.GemUnlockInfo {
		slotItem.GemsUnlockList = append(slotItem.GemsUnlockList, order)
	}
	for typ, info := range slot.CastingSpiritInfo {
		slotItem.SpiritInfo = append(slotItem.SpiritInfo, buildCastingSpiritInfo(typ, info))
	}
	for typ, info := range slot.ForgeSoulInfo {
		slotItem.SoulInfo = append(slotItem.SoulInfo, buildForgeSoulInfo(typ, info))
	}
	return slotItem
}

func buildCastingSpiritInfo(spiritType goldequiptypes.SpiritType, spiritInfo *goldequiptypes.CastingSpiritInfo) *uipb.CastingSpiritInfo {
	info := &uipb.CastingSpiritInfo{}
	level := spiritInfo.Level
	times := spiritInfo.Times
	bless := spiritInfo.Bless
	tpy := int32(spiritType)
	info.SpiritType = &tpy
	info.Level = &level
	info.Times = &times
	info.Bless = &bless
	return info
}

func buildForgeSoulInfo(soulType goldequiptypes.ForgeSoulType, soulInfo *goldequiptypes.ForgeSoulInfo) *uipb.ForgeSoulInfo {
	info := &uipb.ForgeSoulInfo{}
	level := soulInfo.Level
	times := soulInfo.Times
	tpy := int32(soulType)
	info.SoulType = &tpy
	info.Level = &level
	info.Times = &times
	return info
}

func BuildSCGodCastingCastingSpiritUpLevel(bodyPos inventorytypes.BodyPositionType, spiritType goldequiptypes.SpiritType, spiritInfo *goldequiptypes.CastingSpiritInfo, isSuccess int32) *uipb.SCGodCastingCastingSpiritUplevel {
	info := &uipb.SCGodCastingCastingSpiritUplevel{}
	bless := spiritInfo.Bless
	times := spiritInfo.Times
	level := spiritInfo.Level
	bp := int32(bodyPos)
	st := int32(spiritType)
	info.BodyPos = &bp
	info.SpiritType = &st
	info.Bless = &bless
	info.Times = &times
	info.Level = &level
	info.IsSuccess = &isSuccess
	return info
}

func BuildSCGodCastingForgeSoulUpLevel(bodyPos inventorytypes.BodyPositionType, soulType goldequiptypes.ForgeSoulType, soulInfo *goldequiptypes.ForgeSoulInfo, isSuccess int32) *uipb.SCGodCastingForgeSoulUplevel {
	scMsg := &uipb.SCGodCastingForgeSoulUplevel{}
	bp := int32(bodyPos)
	st := int32(soulType)
	level := soulInfo.Level
	times := soulInfo.Times
	scMsg.BodyPos = &bp
	scMsg.SoulType = &st
	scMsg.Times = &times
	scMsg.Level = &level
	scMsg.IsSuccess = &isSuccess
	return scMsg
}

func BuildSCGodCastingEquipUplevel(bodyPos inventorytypes.BodyPositionType, times int32, itemId int32, isSuccess int32) *uipb.SCGodCastingEquipUplevel {
	scMsg := &uipb.SCGodCastingEquipUplevel{}
	bp := int32(bodyPos)
	scMsg.BodyPos = &bp
	scMsg.Times = &times
	scMsg.ItemId = &itemId
	scMsg.IsSuccess = &isSuccess
	return scMsg
}

func BuildSCGodCastingEquipInherit() *uipb.SCGodCastingEquipInherit {
	scMsg := &uipb.SCGodCastingEquipInherit{}
	return scMsg
}

func buildEquipmentSlot(slot *playergoldequip.PlayerGoldEquipSlotObject) *uipb.GoldEquipSlot {
	slotItem := &uipb.GoldEquipSlot{}
	slotId := int32(slot.GetSlotId())
	bindInt := int32(slot.GetBindType())
	level := slot.GetLevel()
	newStLevel := slot.GetNewStLevel()
	itemId := slot.GetItemId()
	slotItem.SlotId = &slotId
	slotItem.Level = &level
	slotItem.NewStLevel = &newStLevel
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
	for order := range slot.GemUnlockInfo {
		slotItem.GemsUnlockList = append(slotItem.GemsUnlockList, order)
	}
	for key, value := range slot.CastingSpiritInfo {
		info := buildCastingSpiritInfo(key, value)
		slotItem.SpiritInfo = append(slotItem.SpiritInfo, info)
	}
	for key, value := range slot.ForgeSoulInfo {
		info := buildForgeSoulInfo(key, value)
		slotItem.SoulInfo = append(slotItem.SoulInfo, info)
	}
	return slotItem
}

func buildEquipmentSlotList(slotList []*playergoldequip.PlayerGoldEquipSlotObject) (slotItemList []*uipb.GoldEquipSlot) {
	for _, slot := range slotList {
		slotItemList = append(slotItemList, buildEquipmentSlot(slot))
	}
	return slotItemList
}

func buildGoldEquipLogList(logObjList []*playergoldequip.PlayerGoldEquipLogObject) (logList []*uipb.GoldEquipLog) {
	for _, logObj := range logObjList {
		createTime := logObj.GetUpdateTime()
		itemIdList := logObj.GetFenJieItemIdList()
		rewItemStr := logObj.GetRewItemStr()

		log := &uipb.GoldEquipLog{}
		log.CreateTime = &createTime
		log.ItemIdList = itemIdList
		log.RewItemStr = &rewItemStr

		logList = append(logList, log)
	}
	return logList
}

func buildBaseProperty(expireType int32, expireTime, itemGetTime int64) *uipb.BaseProperty {
	info := &uipb.BaseProperty{}
	info.ExpireType = &expireType
	info.ExpireTime = &expireTime
	info.ItemGetTime = &itemGetTime
	return info
}

func BuildSCGoldEquipUnlockGem(slot, order int32) *uipb.SCGoldEquipUnlockGem {
	scMsg := &uipb.SCGoldEquipUnlockGem{}
	scMsg.SlotId = &slot
	scMsg.Order = &order
	return scMsg
}

func BuildSCGoldEquipStrengthenBuwei(result goldequiptypes.UpstarResultType, level int32, posType inventorytypes.BodyPositionType, isAvoid bool) *uipb.SCGoldEquipStrengthenBuwei {
	scMsg := &uipb.SCGoldEquipStrengthenBuwei{}
	resultInt := int32(result)
	scMsg.Result = &resultInt
	scMsg.Level = &level
	slotId := int32(posType)
	scMsg.SlotId = &slotId
	scMsg.IsAvoid = &isAvoid
	return scMsg
}

func BuildSCGoldEquipUseItemWithGrowUp(itemId int32, posType inventorytypes.BodyPositionType, level int32) *uipb.SCGoldEquipUseItemWithGrowUp {
	scMsg := &uipb.SCGoldEquipUseItemWithGrowUp{}
	slotId := int32(posType)
	scMsg.ItemId = &itemId
	scMsg.SlotId = &slotId
	scMsg.Level = &level
	return scMsg
}
