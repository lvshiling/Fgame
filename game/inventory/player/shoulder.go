package player

import (
	"fgame/fgame/game/global"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"sort"

	log "github.com/Sirupsen/logrus"
)

//背包
type ShoulderBag struct {
	*BagContainer
	p player.Player
	//位置容量
	capacity int32
	//没有使用过的
	unusedItemList []*PlayerItemObject
}

//添加格子
func (sb *ShoulderBag) AddItem(item *PlayerItemObject) {
	if item.Used == 0 {
		sb.unusedItemList = append(sb.unusedItemList, item)
		return
	}

	sb.addItem(item)
}

func (sb *ShoulderBag) addItem(it *PlayerItemObject) {
	sb.BagContainer.AddItem(it)
}

//获取空闲的格子
func (sb *ShoulderBag) GetEmptySlots() int32 {

	emptyNum := int32(0)

	//获取剩余位置
	remainSlotNum := sb.capacity - sb.Len()
	emptyNum += remainSlotNum
	//使用空的位置
	numOfEmptySlots := sb.NumOfEmptySlots()
	emptyNum += numOfEmptySlots

	return emptyNum
}

// //需要几个格子
// func (sb *ShoulderBag) CountNeedSlot(itemId int32, num int32) (slotNum int32) {
// 	itemTemplate := item.GetItemService().GetItem(int(itemId))
// 	if itemTemplate == nil {
// 		return 0
// 	}
// 	maxNum := int32(0)
// 	if itemTemplate.CanOverlap() {
// 		for _, item := range sb.GetItemList() {
// 			if item.ItemId != itemId {
// 				continue
// 			}
// 			remainNum := itemTemplate.MaxOverlap - item.Num
// 			if remainNum > 0 {
// 				maxNum += remainNum
// 			}
// 		}
// 	}
// 	//不需要额外的空格子
// 	if maxNum >= num {
// 		return
// 	}
// 	remainNum := num - maxNum
// 	maxOverlap := itemTemplate.MaxOverlap
// 	return (remainNum + maxOverlap - 1) / maxOverlap
// }

//-----------------------------v1.0---------------------------------------

func (sb *ShoulderBag) AddItemId(itemId int32, num int32, isDept inventorytypes.IsDepotType, bindType itemtypes.ItemBindType) bool {
	if num <= 0 {
		panic(fmt.Errorf("inventory:add item %d,num %d should more than 0", itemId, num))
	}

	if !sb.HasEnoughSlot(itemId, num) {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	itemTemplate := item.GetItemService().GetItem(int(itemId))

	totalRemainNum := sb.AddItemToOverlapSlot(itemId, num, bindType)
	//填充使用过的空格子
	totalRemainNum = sb.AddItemToEmptySlot(itemId, totalRemainNum, bindType)

	//遍历使用过且整理后空的
	for totalRemainNum > 0 && len(sb.unusedItemList) > 0 {
		//填充没有使用过的格子
		unusedItem := sb.unusedItemList[0]
		sb.unusedItemList = sb.unusedItemList[1:]
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum
		index := int32(sb.Len())
		unusedItem.ItemId = itemId
		unusedItem.Index = index
		unusedItem.UpdateTime = now
		unusedItem.ItemGetTime = now
		unusedItem.Num = complementNum
		unusedItem.Used = 1
		unusedItem.BindType = bindType
		base := inventorytypes.CreateDefaultItemPropertyDataBase()
		unusedItem.PropertyData = inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), base)
		unusedItem.SetModified()
		sb.addItem(unusedItem)
	}

	//填充没有使用过的格子
	for totalRemainNum > 0 {
		nextIndex := int32(sb.Len())
		//超过格子了
		if nextIndex >= sb.capacity {
			break
		}
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum
		item := createItem(sb.p, itemTemplate.GetBagType(), itemId, nextIndex, complementNum, isDept, bindType)
		base := inventorytypes.CreateDefaultItemPropertyDataBase()
		item.PropertyData = inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), base)
		item.SetModified()
		sb.addItem(item)
	}

	if totalRemainNum != 0 {
		panic(fmt.Errorf("total remain num [%d],never reach herer", totalRemainNum))
	}
	return true
}

//有足够的空间
func (sb *ShoulderBag) HasEnoughSlot(itemId int32, num int32) bool {
	remainNum := sb.RemainSlotForItem(itemId)
	if remainNum >= num {
		return true
	}
	return false
}

//剩余填充物品数
func (sb *ShoulderBag) RemainSlotForItem(itemId int32) (maxNum int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return 0
	}
	maxNum = int32(0)
	if itemTemplate.CanOverlap() {
		for _, item := range sb.GetItemList() {
			if item.ItemId != itemId {
				continue
			}
			remainNum := itemTemplate.MaxOverlap - item.Num
			if remainNum > 0 {
				maxNum += remainNum
			}

		}
	}

	//获取剩余位置
	remainSlotNum := sb.capacity - sb.Len()

	//使用空的位置
	numOfEmptySlots := sb.NumOfEmptySlots()
	remainSlotNum += numOfEmptySlots
	if remainSlotNum <= 0 {
		return
	}
	maxNum += itemTemplate.MaxOverlap * int32(remainSlotNum)
	return
}

//---------------------- v1.1------------------------------------

//有足够的空间
func (sb *ShoulderBag) HasEnoughSlotItemLevel(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType) bool {
	now := global.GetGame().GetTimeService().Now()
	return sb.HasEnoughSlotItemLevelWithProperty(itemId, num, level, bindType, inventorytypes.NewItemLimitTimeTypeNone, now, 0)
}

//剩余填充物品数
func (sb *ShoulderBag) RemainSlotForItemLevel(itemId, level int32, bindType itemtypes.ItemBindType) (maxNum int32) {
	now := global.GetGame().GetTimeService().Now()
	return sb.RemainSlotForItemLevelWithProperty(itemId, level, bindType, inventorytypes.NewItemLimitTimeTypeNone, now, 0)
}

//需要几个格子
func (sb *ShoulderBag) CountNeedSlotOfItemLevel(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType) (slotNum int32) {
	now := global.GetGame().GetTimeService().Now()
	return sb.CountNeedSlotOfItemLevelWithProperty(itemId, num, level, bindType, inventorytypes.NewItemLimitTimeTypeNone, now, 0)
}

func (sb *ShoulderBag) AddLevelItem(itemId int32, num int32, level int32, isDept inventorytypes.IsDepotType, bindType itemtypes.ItemBindType) bool {
	if num <= 0 {
		panic(fmt.Errorf("inventory:add item %d,num %d should more than 0", itemId, num))
	}

	if !sb.HasEnoughSlotItemLevel(itemId, num, level, bindType) {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate.GetBindType() == itemtypes.ItemBindTypeBind {
		bindType = itemtypes.ItemBindTypeBind
	}
	base := inventorytypes.CreateDefaultItemPropertyDataBase()
	propertyData := inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), base)
	totalRemainNum := sb.AddItemLevelToOverlapSlot(itemId, num, level, bindType, propertyData)
	//填充使用过的空格子
	totalRemainNum = sb.AddLevelItemToEmptySlot(itemId, totalRemainNum, level, bindType)

	//遍历使用过且整理后空的
	for totalRemainNum > 0 && len(sb.unusedItemList) > 0 {
		//填充没有使用过的格子
		unusedItem := sb.unusedItemList[0]
		sb.unusedItemList = sb.unusedItemList[1:]
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum

		index := int32(sb.Len())
		unusedItem.ItemId = itemId
		unusedItem.Index = index
		unusedItem.UpdateTime = now
		unusedItem.ItemGetTime = now
		unusedItem.Num = complementNum
		unusedItem.Level = level
		unusedItem.Used = 1
		unusedItem.BindType = bindType
		unusedItem.PropertyData = propertyData
		unusedItem.SetModified()
		sb.addItem(unusedItem)
	}

	//填充没有使用过的格子
	for totalRemainNum > 0 {
		nextIndex := int32(sb.Len())
		//超过格子了
		if nextIndex >= sb.capacity {
			log.Warnf("格子不足, itemId:%d, num :%d", itemId, num)
			return false
		}
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum

		item := createItem(sb.p, itemTemplate.GetBagType(), itemId, nextIndex, complementNum, isDept, bindType)
		item.Level = level
		base := inventorytypes.CreateDefaultItemPropertyDataBase()
		item.PropertyData = inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), base)
		item.SetModified()
		sb.addItem(item)
	}
	if totalRemainNum != 0 {
		panic(fmt.Errorf("total remain num [%d],never reach herer", totalRemainNum))
	}
	return true
}

//---------------------- v1.2------------------------------------

//有足够的空间
func (sb *ShoulderBag) HasEnoughSlotItemLevelWithProperty(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, itemGetTime int64, expireTime int64) bool {
	remainNum := sb.RemainSlotForItemLevelWithProperty(itemId, level, bindType, expireType, itemGetTime, expireTime)
	if remainNum >= num {
		return true
	}
	return false
}

//剩余填充物品数
func (sb *ShoulderBag) RemainSlotForItemLevelWithProperty(itemId, itemLevel int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, itemGetTime int64, expireTime int64) (maxNum int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return 0
	}
	maxNum = int32(0)
	if itemTemplate.CanOverlap() {
		for _, item := range sb.GetItemList() {
			if !item.IsOverlap(itemId, itemLevel, bindType, expireType, itemGetTime, expireTime) {
				continue
			}

			remainNum := itemTemplate.MaxOverlap - item.Num
			if remainNum > 0 {
				maxNum += remainNum
			}

		}
	}

	//获取剩余位置
	remainSlotNum := sb.capacity - sb.Len()

	//使用空的位置
	numOfEmptySlots := sb.NumOfEmptySlots()
	remainSlotNum += numOfEmptySlots
	if remainSlotNum <= 0 {
		return
	}
	maxNum += itemTemplate.MaxOverlap * int32(remainSlotNum)
	return
}

//需要几个格子
func (sb *ShoulderBag) CountNeedSlotOfItemLevelWithProperty(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, itemGetTime int64, expireTime int64) (slotNum int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return 0
	}
	maxNum := int32(0)
	if itemTemplate.CanOverlap() {
		for _, item := range sb.GetItemList() {
			if !item.IsOverlap(itemId, level, bindType, expireType, itemGetTime, expireTime) {
				continue
			}

			remainNum := itemTemplate.MaxOverlap - item.Num
			if remainNum > 0 {
				maxNum += remainNum
			}
		}
	}
	//不需要额外的空格子
	if maxNum >= num {
		return
	}

	remainNum := num - maxNum
	maxOverlap := itemTemplate.MaxOverlap
	return (remainNum + maxOverlap - 1) / maxOverlap
}

func (sb *ShoulderBag) AddLevelItemWithPropertyData(itemId int32, num int32, level int32, propertyData inventorytypes.ItemPropertyData, isDept inventorytypes.IsDepotType, bindType itemtypes.ItemBindType) bool {
	if num <= 0 {
		panic(fmt.Errorf("inventory:add item %d,num %d should more than 0", itemId, num))
	}

	if !sb.HasEnoughSlotItemLevelWithProperty(itemId, num, level, bindType, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTimestamp()) {
		return false
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate.IsGoldEquip() && int(num) != 1 {
		return false
	}

	if itemTemplate.GetBindType() == itemtypes.ItemBindTypeBind {
		bindType = itemtypes.ItemBindTypeBind
	}
	if itemTemplate.IsGoldEquip() {
		data := propertyData.(*goldequiptypes.GoldEquipPropertyData)
		if !data.IsHadCountAttr {
			data.AttrList = itemTemplate.GetGoldEquipTemplate().RandomGoldEquipAttr()
			data.IsHadCountAttr = true
		}
	}

	totalRemainNum := sb.AddItemLevelToOverlapSlot(itemId, num, level, bindType, propertyData)
	//填充使用过的空格子

	totalRemainNum = sb.AddLevelItemWithPropertyDataToEmptySlot(itemId, totalRemainNum, level, bindType, propertyData)
	//遍历使用过且整理后空的
	now := global.GetGame().GetTimeService().Now()
	for totalRemainNum > 0 && len(sb.unusedItemList) > 0 {
		//填充没有使用过的格子
		unusedItem := sb.unusedItemList[0]
		sb.unusedItemList = sb.unusedItemList[1:]
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum

		index := int32(sb.Len())
		unusedItem.ItemId = itemId
		unusedItem.Index = index
		unusedItem.UpdateTime = now
		unusedItem.ItemGetTime = now
		unusedItem.Num = complementNum
		unusedItem.Level = level
		unusedItem.Used = 1
		unusedItem.BindType = bindType
		unusedItem.PropertyData = propertyData
		unusedItem.SetModified()
		sb.addItem(unusedItem)
	}

	//填充没有使用过的格子
	for totalRemainNum > 0 {
		nextIndex := int32(sb.Len())
		//超过格子了
		if nextIndex >= sb.capacity {
			log.Warnf("格子不足, itemId:%d, num :%d", itemId, num)
			return false
		}
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum

		item := createItem(sb.p, itemTemplate.GetBagType(), itemId, nextIndex, complementNum, isDept, bindType)
		item.Level = level
		item.PropertyData = propertyData
		item.SetModified()
		sb.addItem(item)
	}
	if totalRemainNum != 0 {
		panic(fmt.Errorf("total remain num [%d],never reach herer", totalRemainNum))
	}
	return true
}

//-----------------------分割------------------------------------

//移除物品
func (sb *ShoulderBag) RemoveItem(itemId int32, num int32) bool {
	if num <= 0 {
		panic(fmt.Errorf("inventory:remove item %d,num %d", itemId, num))
	}

	if sb.NumOfItems(itemId) < num {
		return false
	}

	totalRemoveNum := num
	// 时效物品
	for _, item := range sb.GetItemList() {
		if item.ItemId != itemId {
			continue
		}
		if item.PropertyData.GetExpireType() == inventorytypes.NewItemLimitTimeTypeNone {
			continue
		}

		if item.PropertyData.IsExpire() {
			continue
		}

		//移除的数量
		removeNum := item.Num
		if totalRemoveNum < removeNum {
			removeNum = totalRemoveNum
		}
		totalRemoveNum -= removeNum

		sb.RemoveIndex(item.Index, removeNum)
		if totalRemoveNum <= 0 {
			return true
		}
	}

	// 非时效物品
	for _, item := range sb.GetItemList() {
		if item.ItemId != itemId {
			continue
		}
		if item.PropertyData.GetExpireType() != inventorytypes.NewItemLimitTimeTypeNone {
			continue
		}

		//移除的数量
		removeNum := item.Num
		if totalRemoveNum < removeNum {
			removeNum = totalRemoveNum
		}
		totalRemoveNum -= removeNum

		sb.RemoveIndex(item.Index, removeNum)
		if totalRemoveNum <= 0 {
			return true
		}
	}

	panic(fmt.Errorf("inventory:remove item %d,num %d", itemId, num))
}

//分类合并
func (sb *ShoulderBag) classifyAndMerge() {
	classifyMap := make(map[int32][]*PlayerItemObject)
	for _, it := range sb.GetItemList() {
		if it.IsEmpty() {
			continue
		}
		l, exist := classifyMap[it.ItemId]
		if !exist {
			l = make([]*PlayerItemObject, 0, 8)
			classifyMap[it.ItemId] = l
		}
		classifyMap[it.ItemId] = append(l, it)
	}

	for _, classifyList := range classifyMap {
		var previousItem *PlayerItemObject
		for _, it := range classifyList {
			itemTemplate := item.GetItemService().GetItem(int(it.ItemId))
			if previousItem == nil {
				//未满
				if !it.IsFull() {
					previousItem = it
				}
				continue
			}
			if !previousItem.IsOverlapExcludeBind(it.ItemId, it.Level, it.PropertyData.GetExpireType(), it.PropertyData.GetItemGetTime(), it.PropertyData.GetExpireTimestamp()) {
				continue
			}

			// 合并后都为绑定类型
			if previousItem.BindType != it.BindType {
				previousItem.BindType = itemtypes.ItemBindTypeBind
				it.BindType = itemtypes.ItemBindTypeBind
			}

			//TODO:xzk:修改合并时间和等级
			//计算剩余的
			remainNum := itemTemplate.MaxOverlap - previousItem.Num
			complementNum := remainNum
			if it.Num <= remainNum {
				complementNum = it.Num
			}
			//移除当前的
			sb.RemoveIndex(it.Index, complementNum)
			//添加之前的
			sb.AddIndex(previousItem.Index, previousItem.ItemId, complementNum)
			//当前为空
			if it.IsEmpty() {
				//先前位置填满了
				if previousItem.IsFull() {
					previousItem = nil
				}
				continue
			}
			previousItem = it
		}
	}
}

//合并
func (sb *ShoulderBag) Merge() {
	sb.classifyAndMerge()
	//获取物品列表
	itemList := sb.GetItemList()

	sort.Sort(sort.Reverse(itemObjectList(itemList)))
	//TODO 优化
	sb.BagContainer = createBagContainer(nil)

	unusedItemList := make([]*PlayerItemObject, 0, 16)
	//修改索引
	for index, item := range itemList {
		dirty := false
		//切换索引
		if item.Index != int32(index) {
			item.Index = int32(index)
			dirty = true
		}

		//设置为没使用
		if item.ItemId == 0 {
			item.Used = 0
			dirty = true
		}
		//改变过了
		if dirty {
			item.SetModified()
		}

		//使用过
		if item.Used != 0 {
			sb.addItem(item)
		} else {
			unusedItemList = append(unusedItemList, item)
		}
	}
	sb.unusedItemList = append(unusedItemList, sb.unusedItemList...)
	//清空标记位
	sb.Reset()
}

func (sb *ShoulderBag) ResetCapacity(capacity int32) {
	sb.capacity = capacity
}

func createShoulderBag(p player.Player, items []*PlayerItemObject, capacity int32) *ShoulderBag {
	sb := &ShoulderBag{
		p: p,
	}
	sb.BagContainer = createBagContainer(nil)
	sb.capacity = capacity
	sb.unusedItemList = make([]*PlayerItemObject, 0, 16)

	for _, item := range items {
		sb.AddItem(item)
	}

	sb.Reset()
	return sb
}

func createItem(p player.Player, bagType inventorytypes.BagType, itemId int32, index int32, num int32, isDept inventorytypes.IsDepotType, bindType itemtypes.ItemBindType) *PlayerItemObject {
	now := global.GetGame().GetTimeService().Now()
	itemObject := NewPlayerItemObject(p)
	itemObject.CreateTime = now
	itemObject.ItemGetTime = now
	itemObject.Id, _ = idutil.GetId()
	itemObject.ItemId = itemId
	itemObject.Num = num
	itemObject.Index = index
	itemObject.BagType = bagType
	itemObject.IsDepot = isDept
	itemObject.BindType = bindType
	itemObject.Used = 1
	itemObject.PropertyData = inventorytypes.CreateDefaultItemPropertyDataBase()
	return itemObject
}
