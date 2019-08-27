package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/equipbaoku/equipbaoku"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
)

func BuildSCEquipBaoKuInfoGet(
	equipObj, materialObj *playerequipbaoku.PlayerEquipBaoKuObject,
	equipLogList, materialLogList []*equipbaoku.EquipBaoKuLogObject,
	buyCountMap map[int32]*playerequipbaoku.PlayerEquipBaoKuShopObject) *uipb.SCEquipbaokuInfoGet {
	scEquipBaoKuInfoGet := &uipb.SCEquipbaokuInfoGet{}
	scEquipBaoKuInfoGet.EquipBaokuInfo = buildEquipBaoKuInfo(equipObj)
	scEquipBaoKuInfoGet.MaterialBaokuInfo = buildEquipBaoKuInfo(materialObj)
	for _, log := range equipLogList {
		scEquipBaoKuInfoGet.EquipLogList = append(scEquipBaoKuInfoGet.EquipLogList, buildEquipBaoKuLog(log))
	}
	for _, log := range materialLogList {
		scEquipBaoKuInfoGet.MaterialLogList = append(scEquipBaoKuInfoGet.MaterialLogList, buildEquipBaoKuLog(log))
	}
	for _, shop := range buyCountMap {
		scEquipBaoKuInfoGet.ShopLimitList = append(scEquipBaoKuInfoGet.ShopLimitList, buildShop(shop))
	}
	return scEquipBaoKuInfoGet
}

func BuildSCEquipBaoKuLogIncr(logList []*equipbaoku.EquipBaoKuLogObject, typ int32) *uipb.SCEquipbaokuLogIncr {
	scEquipBaoKuLogIncr := &uipb.SCEquipbaokuLogIncr{}
	scEquipBaoKuLogIncr.Type = &typ
	for _, log := range logList {
		scEquipBaoKuLogIncr.LogList = append(scEquipBaoKuLogIncr.LogList, buildEquipBaoKuLog(log))
	}
	return scEquipBaoKuLogIncr
}

func buildEquipBaoKuInfo(obj *playerequipbaoku.PlayerEquipBaoKuObject) *uipb.EquipbaokuInfo {
	info := &uipb.EquipbaokuInfo{}
	luckyPoints := obj.GetLuckyPoints()
	info.LuckyPoints = &luckyPoints
	attendPoints := obj.GetAttendPoints()
	info.AttendPoints = &attendPoints
	return info
}

func buildEquipBaoKuLog(log *equipbaoku.EquipBaoKuLogObject) *uipb.EquipbaokuLog {
	info := &uipb.EquipbaokuLog{}
	playerName := log.GetPlayerName()
	itemId := log.GetItemId()
	itemNum := log.GetItemNum()
	time := log.GetUpdateTime()
	info.PlayerName = &playerName
	info.ItemId = &itemId
	info.ItemNum = &itemNum
	info.CreateTime = &time

	return info
}

func BuildSCEquipBaoKuAttend(rewItemList []*droptemplate.DropItemData, logList []*equipbaoku.EquipBaoKuLogObject, luckyPoints int32, attendPoints int32, autoFlag bool, typ int32) *uipb.SCEquipbaokuAttend {
	scEquipBaoKuAttend := &uipb.SCEquipbaokuAttend{}
	scEquipBaoKuAttend.Type = &typ
	for i := int(0); i < len(rewItemList); i++ {
		itemId := rewItemList[i].GetItemId()
		num := rewItemList[i].GetNum()
		level := rewItemList[i].GetLevel()

		scEquipBaoKuAttend.DropInfo = append(scEquipBaoKuAttend.DropInfo, buildDropInfo(itemId, num, level))
	}
	scEquipBaoKuAttend.AutoFlag = &autoFlag
	for _, log := range logList {
		scEquipBaoKuAttend.LogList = append(scEquipBaoKuAttend.LogList, buildEquipBaoKuLog(log))
	}

	scEquipBaoKuAttend.LuckyPoints = &luckyPoints
	scEquipBaoKuAttend.AttendPoints = &attendPoints
	return scEquipBaoKuAttend
}

func BuildSCEquipBaoKuAttendBatch(rewItemList []*droptemplate.DropItemData, logList []*equipbaoku.EquipBaoKuLogObject, luckyPoints int32, attendPoints int32, autoFlag bool, typ int32) *uipb.SCEquipbaokuAttendBatch {
	scEquipBaoKuAttendBatch := &uipb.SCEquipbaokuAttendBatch{}
	scEquipBaoKuAttendBatch.Type = &typ
	for i := int(0); i < len(rewItemList); i++ {
		itemId := rewItemList[i].GetItemId()
		num := rewItemList[i].GetNum()
		level := rewItemList[i].GetLevel()

		scEquipBaoKuAttendBatch.DropList = append(scEquipBaoKuAttendBatch.DropList, buildDropInfo(itemId, num, level))
	}
	scEquipBaoKuAttendBatch.AutoFlag = &autoFlag
	for _, log := range logList {
		scEquipBaoKuAttendBatch.LogList = append(scEquipBaoKuAttendBatch.LogList, buildEquipBaoKuLog(log))
	}
	scEquipBaoKuAttendBatch.LuckyPoints = &luckyPoints
	scEquipBaoKuAttendBatch.AttendPoints = &attendPoints
	return scEquipBaoKuAttendBatch
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}

func BuildSCEquipBaoKuLuckyBox(rewItemList []*droptemplate.DropItemData, luckyPoints int32, typ int32) *uipb.SCEquipbaokuLuckyBox {
	scEquipBaoKuLuckyBox := &uipb.SCEquipbaokuLuckyBox{}
	scEquipBaoKuLuckyBox.Type = &typ
	for i := int(0); i < len(rewItemList); i++ {
		itemId := rewItemList[i].GetItemId()
		num := rewItemList[i].GetNum()
		level := rewItemList[i].GetLevel()

		scEquipBaoKuLuckyBox.DropList = append(scEquipBaoKuLuckyBox.DropList, buildDropInfo(itemId, num, level))
	}
	scEquipBaoKuLuckyBox.LuckyPoints = &luckyPoints
	return scEquipBaoKuLuckyBox
}

func BuildSCEquipBaoKuPointsExchange(shopId, num, attendPoints, dayCount int32, typ int32) *uipb.SCEquipbaokuPointsExchange {
	scEquipBaoKuPointsExchange := &uipb.SCEquipbaokuPointsExchange{}
	scEquipBaoKuPointsExchange.Type = &typ
	scEquipBaoKuPointsExchange.ShopId = &shopId
	scEquipBaoKuPointsExchange.Num = &num

	scEquipBaoKuPointsExchange.AttendPoints = &attendPoints
	scEquipBaoKuPointsExchange.DayCount = &dayCount
	return scEquipBaoKuPointsExchange
}

func BuildSCEquipBaoKuResolveEquip(level int32, exp int64, returnItemMap map[int32]int32) *uipb.SCEquipbaokuResolveEquip {
	scEquipBaoKuResolveEquip := &uipb.SCEquipbaokuResolveEquip{}
	scEquipBaoKuResolveEquip.GoldYuanLevel = &level
	scEquipBaoKuResolveEquip.GoldYuanExp = &exp
	scEquipBaoKuResolveEquip.DropInfoList = droppbutil.BuildSimpleDropInfoList(returnItemMap)
	return scEquipBaoKuResolveEquip
}

func buildShop(shop *playerequipbaoku.PlayerEquipBaoKuShopObject) *uipb.ShopLimit {
	shopLimit := &uipb.ShopLimit{}
	shopId := int32(shop.ShopId)
	num := int32(shop.DayCount)
	shopLimit.ShopId = &shopId
	shopLimit.DayCount = &num
	return shopLimit
}

func BuildSCEquipBaoKuShopLimit(buyCountMap map[int32]*playerequipbaoku.PlayerEquipBaoKuShopObject) *uipb.SCEquipbaokuShopLimit {
	scShopLimit := &uipb.SCEquipbaokuShopLimit{}
	for _, shop := range buyCountMap {
		scShopLimit.ShopLimitList = append(scShopLimit.ShopLimitList, buildShop(shop))
	}
	return scShopLimit
}
