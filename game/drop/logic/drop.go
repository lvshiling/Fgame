package logic

import (
	commonlog "fgame/fgame/common/log"
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/global"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	"fgame/fgame/game/inventory/inventory"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
	"strconv"
	"strings"
)

func SeperateItems(itemMap map[int32]int32) (newItemMap map[int32]int32, resMap map[itemtypes.ItemAutoUseResSubType]int32) {
	if len(itemMap) == 0 {
		panic(fmt.Errorf("drop: 物品不应该为空"))
	}
	newItemMap = make(map[int32]int32)
	resMap = make(map[itemtypes.ItemAutoUseResSubType]int32)
	for itemId, num := range itemMap {
		to := item.GetItemService().GetItem(int(itemId))
		typ := to.GetItemType()
		if typ == itemtypes.ItemTypeAutoUseRes {
			subType := to.GetItemSubType().(itemtypes.ItemAutoUseResSubType)
			resMap[subType] += num
		} else {
			newItemMap[itemId] = num
		}
	}

	return newItemMap, resMap
}

func SeperateItemData(itemData *droptemplate.DropItemData) (newItemData []*droptemplate.DropItemData, resMap map[itemtypes.ItemAutoUseResSubType]int32) {
	if itemData == nil {
		return
	}
	itemId := itemData.GetItemId()
	num := itemData.GetNum()
	to := item.GetItemService().GetItem(int(itemId))
	typ := to.GetItemType()

	resMap = make(map[itemtypes.ItemAutoUseResSubType]int32)
	if typ == itemtypes.ItemTypeAutoUseRes {
		subType := to.GetItemSubType().(itemtypes.ItemAutoUseResSubType)
		resMap[subType] += num
	} else {
		newItemData = append(newItemData, droptemplate.CreateItemDataWithData(itemData))
	}
	return
}

func SeperateItemDatas(itemList []*droptemplate.DropItemData) (newItemList []*droptemplate.DropItemData, resMap map[itemtypes.ItemAutoUseResSubType]int32) {
	if len(itemList) == 0 {
		panic(fmt.Errorf("drop: 物品不应该为空"))
	}

	resMap = make(map[itemtypes.ItemAutoUseResSubType]int32)
	for _, itemData := range itemList {
		if itemData == nil {
			continue
		}

		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		to := item.GetItemService().GetItem(int(itemId))
		typ := to.GetItemType()

		if typ == itemtypes.ItemTypeAutoUseRes {
			subType := to.GetItemSubType().(itemtypes.ItemAutoUseResSubType)
			resMap[subType] += num
		} else {
			isMerge := false
			for _, data := range newItemList {
				if to.IsGoldEquip() {
					continue
				}
				if to.IsBaoBaoCard() {
					continue
				}

				if data.IsMerge(itemData) {
					data.Num += num

					isMerge = true
				}
			}
			if !isMerge {
				newData := droptemplate.CreateItemDataWithData(itemData)
				newItemList = append(newItemList, newData)
			}
		}
	}

	return newItemList, resMap
}

//合并等级物品
func MergeItemLevel(items []*droptemplate.DropItemData) (newItemList []*droptemplate.DropItemData) {
	for _, itemData := range items {
		isMerge := false
		to := item.GetItemService().GetItem(int(itemData.ItemId))
		for _, data := range newItemList {
			if data == nil {
				continue
			}
			if to.IsGoldEquip() {
				continue
			}
			if to.IsBaoBaoCard() {
				continue
			}
			if data.IsMerge(itemData) {
				data.Num += itemData.Num

				isMerge = true
			}
		}
		if !isMerge {
			newData := droptemplate.CreateItemDataWithData(itemData)
			newItemList = append(newItemList, newData)
		}
	}

	return newItemList
}

//TODO 使用addResItem
func AddRes(pl player.Player, resMap map[itemtypes.ItemAutoUseResSubType]int32, goldLog commonlog.GoldLogReason, goldReasonText string, silverLog commonlog.SilverLogReason, silverReasonText string, levelLog commonlog.LevelLogReason, levelReasonText string) error {
	if len(resMap) == 0 {
		panic(fmt.Errorf("drop:资源不应该是空"))
	}

	countResMap := make(map[itemtypes.ItemAutoUseResSubType]int64)
	for resType, num := range resMap {
		_, ok := countResMap[resType]
		if !ok {
			countResMap[resType] = int64(num)
		} else {
			countResMap[resType] += int64(num)
		}
	}

	for resType, num := range countResMap {
		flag := addRes(pl, resType, num, goldLog, goldReasonText, silverLog, silverReasonText, levelLog, levelReasonText)
		if !flag {
			return fmt.Errorf("drop:添加资源失败,resType:%d,resNum:%d", resType, num)
		}
	}

	return nil
}

func AddItem(pl player.Player,
	itemData *droptemplate.DropItemData,
	goldLog commonlog.GoldLogReason,
	goldReasonText string,
	silverLog commonlog.SilverLogReason,
	silverReasonText string,
	inventoryLog commonlog.InventoryLogReason,
	inventoryReasonText string,
	levelLog commonlog.LevelLogReason,
	levelReasonText string) (flag bool, err error) {

	if itemData.Num <= 0 {
		panic(fmt.Errorf("drop:添加物品[%d]应该不能小于0", itemData.Num))
	}
	to := item.GetItemService().GetItem(int(itemData.ItemId))
	typ := to.GetItemType()
	if typ == itemtypes.ItemTypeAutoUseRes {
		flag = addResItem(pl, itemData.ItemId, itemData.Num, goldLog, goldReasonText, silverLog, silverReasonText, levelLog, levelReasonText)
		return
	}

	//添加物品
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotItemLevel(itemData.ItemId, itemData.Num, itemData.Level, itemData.BindType) {
		return
	}

	flag = inventoryManager.AddItemLevel(itemData, inventoryLog, inventoryReasonText)
	if !flag {
		panic(fmt.Errorf("drop:添加物品应该成功"))
	}
	return
}

func AddItemWithProperty(pl player.Player,
	itemId int32,
	num int32,
	level int32,
	upstar int32,
	attrList []int32,
	bind itemtypes.ItemBindType,
	goldLog commonlog.GoldLogReason,
	goldReasonText string,
	silverLog commonlog.SilverLogReason,
	silverReasonText string,
	inventoryLog commonlog.InventoryLogReason,
	inventoryReasonText string,
	levelLog commonlog.LevelLogReason,
	levelReasonText string) (flag bool, err error) {

	if num <= 0 {
		panic(fmt.Errorf("drop:添加物品[%d]应该不能小于0", num))
	}
	to := item.GetItemService().GetItem(int(itemId))
	typ := to.GetItemType()
	if typ == itemtypes.ItemTypeAutoUseRes {
		flag = addResItem(pl, itemId, num, goldLog, goldReasonText, silverLog, silverReasonText, levelLog, levelReasonText)
		return
	}

	//添加物品
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotItemLevel(itemId, num, level, bind) {
		return
	}

	base := inventorytypes.CreateDefaultItemPropertyDataBase()
	propertyData := inventory.CreatePropertyDataInterface(to.GetItemType(), base)
	if to.IsGoldEquip() {
		data := propertyData.(*goldequiptypes.GoldEquipPropertyData)
		data.UpstarLevel = upstar
		data.AttrList = attrList
		data.IsHadCountAttr = true
	}

	flag = inventoryManager.AddItemLevelWithPropertyData(itemId, num, level, bind, propertyData, inventoryLog, inventoryReasonText)
	if !flag {
		panic(fmt.Errorf("drop:添加物品应该成功"))
	}
	return
}

//添加资源物品
func addResItem(pl player.Player, itemId int32, num int32, goldLog commonlog.GoldLogReason, goldReasonText string, silverLog commonlog.SilverLogReason, silverReasonText string, levelLog commonlog.LevelLogReason, levelReasonText string) (flag bool) {
	if num <= 0 {
		panic(fmt.Errorf("drop:添加物品[%d]应该不能小于0", num))
	}

	to := item.GetItemService().GetItem(int(itemId))
	typ := to.GetItemType()
	if typ != itemtypes.ItemTypeAutoUseRes {
		panic(fmt.Errorf("drop:添加物品[%d]应该是资源物品", itemId))
	}
	resType, ok := to.GetItemSubType().(itemtypes.ItemAutoUseResSubType)
	if !ok {
		return
	}

	flag = addRes(pl, resType, int64(num), goldLog, goldReasonText, silverLog, silverReasonText, levelLog, levelReasonText)
	return
}

func addRes(pl player.Player, resType itemtypes.ItemAutoUseResSubType, resNum int64,
	goldLog commonlog.GoldLogReason, goldReasonText string,
	silverLog commonlog.SilverLogReason, silverReasonText string,
	levelLog commonlog.LevelLogReason, levelReasonText string) bool {

	var reason commonlogtypes.LogReason
	reasonText := ""
	h := drop.GetDropResHandler(resType)
	if h == nil {
		return false
	}

	switch resType {
	case itemtypes.ItemAutoUseResSubTypeGold,
		itemtypes.ItemAutoUseResSubTypeBindGold:
		{
			reason = goldLog
			reasonText = goldReasonText
		}
	case itemtypes.ItemAutoUseResSubTypeSilver:
		{
			reason = silverLog
			reasonText = silverReasonText
		}
	case itemtypes.ItemAutoUseResSubTypeYaoPai:
		{
			reason = commonlog.YaoPaiLogReasonPlayerKilled
			reasonText = commonlog.YaoPaiLogReasonPlayerKilled.String()
		}
	case itemtypes.ItemAutoUseResSubTypeExp:
		{
			reason = levelLog
			reasonText = levelReasonText
		}
	}

	return h.AddRes(pl, resNum, reason, reasonText)
}

func GetAddAutoRes(subType itemtypes.ItemAutoUseResSubType, num int32) (silver, gold, bindGold, exp, stone, shaqiNum, xingchenNum, gongdeNum, equipBaoKuPoint, lingqiNum, shenYuKey, qXianTao, bXianTao, shaLuXin, arenaPoint, shengWei, arenapvpJiFen, materialJiFen int32) {
	switch subType {
	case itemtypes.ItemAutoUseResSubTypeBindGold:
		{
			bindGold = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeGold:
		{
			gold = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeSilver:
		{
			silver = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeExp:
		{
			exp = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeStorage:
		{
			stone = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeShaQi:
		{
			shaqiNum = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeXingChen:
		{
			xingchenNum = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeGongDe:
		{
			gongdeNum = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeEquipBaoKuAttendPoints:
		{
			equipBaoKuPoint = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeLingQi:
		{
			lingqiNum = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeShenYuKey:
		{
			shenYuKey = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeQianNianXianTao:
		{
			qXianTao = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeBaiNianXianTao:
		{
			bXianTao = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeShaLuXin:
		{
			shaLuXin = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeArenaPoint:
		{
			arenaPoint = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeShengWei:
		{
			shengWei = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeArenapvpJiFen:
		{
			arenapvpJiFen = num
			break
		}
	case itemtypes.ItemAutoUseResSubTypeMaterialBaoKuAttendPoints:
		{
			materialJiFen = num
			break
		}
	}
	return
}

func ParseAttachmentList(itemStr string) (itemDataList []*droptemplate.DropItemData, err error) {
	itemArr := strings.Split(itemStr, ",")
	for _, tempItem := range itemArr {
		itemNumArr := strings.Split(tempItem, ":")
		if len(itemNumArr) != 2 {
			return nil, fmt.Errorf("格式不对[%s]", tempItem)
		}
		itemId, err := strconv.ParseInt(itemNumArr[0], 10, 64)
		if err != nil {
			return nil, err
		}
		itemIdInt := int32(itemId)

		itemTemplate := item.GetItemService().GetItem(int(itemIdInt))
		if itemTemplate == nil {
			return nil, fmt.Errorf("物品不存在[%d]", itemIdInt)
		}

		itemNum, err := strconv.ParseInt(itemNumArr[1], 10, 64)
		if err != nil {
			return nil, err
		}
		itemNumInt := int32(itemNum)
		itemData := droptemplate.CreateItemData(int32(itemId), itemNumInt, 0, itemtypes.ItemBindTypeUnBind)
		itemDataList = append(itemDataList, itemData)
	}
	return
}

func SceneDropConvertToDropItemData(dropItem scene.DropItem) *droptemplate.DropItemData {
	itemId := dropItem.GetItemId()
	num := dropItem.GetItemNum()
	level := dropItem.GetLevel()
	bind := dropItem.GetBindType()
	upstar := dropItem.GetUpstar()
	attrList := dropItem.GetAttrList()
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := global.GetGame().GetTimeService().Now()
	return droptemplate.CreateItemDataWithPropertyData(itemId, num, level, bind, upstar, attrList, true, expireType, expireTime, itemGetTime)
}
