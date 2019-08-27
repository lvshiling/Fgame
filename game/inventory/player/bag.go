package player

import (
	"container/heap"
	"fgame/fgame/game/global"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"

	"github.com/willf/bitset"
)

//空闲格子列表
type slotFreeList []int32

func (sfl *slotFreeList) Len() int {
	return len(*sfl)
}

func (sfl *slotFreeList) Less(i, j int) bool {
	return (*sfl)[i] < (*sfl)[j]
}

func (sfl *slotFreeList) Swap(i, j int) {
	(*sfl)[i], (*sfl)[j] = (*sfl)[j], (*sfl)[i]
}

func (sfl *slotFreeList) Push(h interface{}) {
	*sfl = append(*sfl, h.(int32))
}

func (sfl *slotFreeList) Pop() (v interface{}) {
	*sfl, v = (*sfl)[0:sfl.Len()-1], (*sfl)[sfl.Len()-1]
	return
}

//背包基础
type BagContainer struct {
	//物品列表
	itemList []*PlayerItemObject
	//改变的位置
	changedBitset *bitset.BitSet
	//空闲的索引
	slotFreeList *slotFreeList
}

//创建新背包
func createBagContainer(itemList []*PlayerItemObject) *BagContainer {
	bb := &BagContainer{}
	bb.itemList = itemList
	bb.changedBitset = bitset.New(64)
	bb.slotFreeList = &slotFreeList{}
	return bb
}

//长度
func (bb *BagContainer) Len() int32 {
	return int32(len(bb.itemList))
}

//获取物品列表
func (bb *BagContainer) GetItemList() []*PlayerItemObject {
	return bb.itemList
}

//获取空的位置
func (bb *BagContainer) NumOfEmptySlots() int32 {
	return int32(len(*bb.slotFreeList))
}

//获取位置
func (bb *BagContainer) GetByIndex(index int32) *PlayerItemObject {
	if index < 0 {
		panic(fmt.Errorf("inventory:get by index %d should be no less than 0", index))
	}
	if int(index) >= len(bb.itemList) {
		return nil
	}
	return bb.itemList[index]
}

//获取改变
func (bb *BagContainer) GetChangedSlotAndReset() (itemList []*PlayerItemObject) {
	itemList = make([]*PlayerItemObject, 0, 16)
	for i, valid := bb.changedBitset.NextSet(0); valid; i, valid = bb.changedBitset.NextSet(i + 1) {
		itemList = append(itemList, bb.itemList[i])
	}
	bb.Reset()
	return itemList
}

//---------------------v1.0---------------------------

//获取物品数量
func (bb *BagContainer) NumOfItems(itemId int32) int32 {
	if itemId == 0 {
		panic(fmt.Errorf("inventory:num of items itemId is 0"))
	}
	num := int32(0)
	for _, it := range bb.itemList {
		//不是同一个物品
		if it.ItemId != itemId {
			continue
		}

		num += it.Num
	}
	return num
}

//添加物品
func (bb *BagContainer) AddItem(item *PlayerItemObject) {
	if int(item.Index) != len(bb.itemList) {
		panic(fmt.Errorf("inventory: add item index %d not last %d", item.Index, len(bb.itemList)))
	}
	bb.changed(int(item.Index))
	bb.itemList = append(bb.itemList, item)
	if item.ItemId == 0 {
		heap.Push(bb.slotFreeList, int32(item.Index))
		return
	}
}

//添加到有物品的位置
func (bb *BagContainer) AddItemToOverlapSlot(itemId int32, num int32, bindType itemtypes.ItemBindType) int32 {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		panic(fmt.Errorf("inventory:add item id [%d] invalid", itemId))
	}

	totalRemainNum := num

	if !itemTemplate.CanOverlap() {
		return totalRemainNum
	}

	now := global.GetGame().GetTimeService().Now()
	for index, item := range bb.itemList {
		//不是同一个物品
		if item.ItemId != itemId {
			continue
		}
		//绑定属性不同
		if item.BindType != bindType {
			continue
		}

		//获取剩余可以叠加数量
		remainNum := itemTemplate.MaxOverlap - item.Num
		//没有
		if remainNum <= 0 {
			continue
		}
		//计算叠加数量
		complementNum := remainNum
		if totalRemainNum <= remainNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum
		item.Num += complementNum
		item.UpdateTime = now
		item.SetModified()
		bb.changed(index)
		//放置完成
		if totalRemainNum <= 0 {
			break
		}
	}
	return totalRemainNum
}

//添加到空的位置
func (bb *BagContainer) AddItemToEmptySlot(itemId int32, num int32, bindType itemtypes.ItemBindType) int32 {

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		panic(fmt.Errorf("inventory:add item id [%d] invalid", itemId))
	}
	now := global.GetGame().GetTimeService().Now()

	totalRemainNum := num
	//填充使用过的空格子
	for totalRemainNum > 0 {
		if bb.NumOfEmptySlots() == 0 {
			return totalRemainNum
		}
		freeIndex := heap.Pop(bb.slotFreeList).(int32)
		if freeIndex >= int32(len(bb.itemList)) {
			panic(fmt.Errorf("inventory:free index [%d] should in item list", freeIndex))
		}
		item := bb.itemList[freeIndex]
		if item.ItemId != 0 {
			panic(fmt.Errorf("inventory:item should be 0"))
		}
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum
		item.ItemId = itemId
		item.BindType = bindType
		item.Num = complementNum
		item.UpdateTime = now
		item.ItemGetTime = now
		base := inventorytypes.CreateDefaultItemPropertyDataBase()
		item.PropertyData = inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), base)
		item.SetModified()
		bb.changed(int(freeIndex))
	}

	return totalRemainNum
}

// 更新物品属性
func (bb *BagContainer) UpdateItem(index int32) bool {
	it := bb.GetByIndex(index)
	if it == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	it.UpdateTime = now
	it.SetModified()
	bb.changed(int(index))

	return true
}

//---------------------v1.1---------------------------

// 更新物品等级
func (bb *BagContainer) UpdateItemLevel(index int32, level int32) bool {
	if level < 0 {
		panic(fmt.Errorf("inventory: update level; index:%d level: %d", index, level))
	}
	it := bb.GetByIndex(index)
	if it == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	it.Level = level
	it.UpdateTime = now
	it.SetModified()
	bb.changed(int(index))

	return true
}

//添加到有物品的位置
func (bb *BagContainer) AddItemLevelToOverlapSlot(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) int32 {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		panic(fmt.Errorf("inventory:add item id [%d] invalid", itemId))
	}

	totalRemainNum := num

	if !itemTemplate.CanOverlap() {
		return totalRemainNum
	}

	now := global.GetGame().GetTimeService().Now()
	for index, item := range bb.itemList {
		if !item.IsOverlap(itemId, level, bindType, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTimestamp()) {
			continue
		}

		//获取剩余可以叠加数量
		remainNum := itemTemplate.MaxOverlap - item.Num
		//没有
		if remainNum <= 0 {
			continue
		}
		//计算叠加数量
		complementNum := remainNum
		if totalRemainNum <= remainNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum
		item.Num += complementNum
		item.UpdateTime = now
		item.SetModified()
		bb.changed(index)
		//放置完成
		if totalRemainNum <= 0 {
			break
		}
	}
	return totalRemainNum
}

func (bb *BagContainer) AddLevelItemToEmptySlot(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType) int32 {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		panic(fmt.Errorf("inventory:add item id [%d] invalid", itemId))
	}
	now := global.GetGame().GetTimeService().Now()

	totalRemainNum := num
	for totalRemainNum > 0 {
		if bb.NumOfEmptySlots() == 0 {
			return totalRemainNum
		}
		freeIndex := heap.Pop(bb.slotFreeList).(int32)
		if freeIndex >= int32(len(bb.itemList)) {
			panic(fmt.Errorf("inventory:free index [%d] should in item list", freeIndex))
		}
		item := bb.itemList[freeIndex]
		if item.ItemId != 0 {
			panic(fmt.Errorf("inventory:item should be 0"))
		}
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum

		item.ItemId = itemId
		item.Num = complementNum
		item.Level = level
		item.BindType = bindType
		item.ItemGetTime = now
		item.UpdateTime = now
		base := inventorytypes.CreateDefaultItemPropertyDataBase()
		item.PropertyData = inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), base)
		item.SetModified()
		bb.changed(int(freeIndex))
	}
	return totalRemainNum
}

//---------------------v1.2---------------------------

//获取物品数量
func (bb *BagContainer) NumOfItemsWithProperty(itemId int32) int32 {
	if itemId == 0 {
		panic(fmt.Errorf("inventory:num of items itemId is 0"))
	}
	num := int32(0)
	for _, it := range bb.itemList {
		//不是同一个物品
		if it.ItemId != itemId {
			continue
		}
		if it.PropertyData.IsExpire() {
			continue
		}
		itemTemplate := item.GetItemService().GetItem(int(itemId))
		if itemTemplate == nil {
			continue
		}
		if it.PropertyData.GetExpireType() == inventorytypes.NewItemLimitTimeTypeNone {
			if itemTemplate.GetLimitTimeType() != inventorytypes.NewItemLimitTimeTypeNone {
				now := global.GetGame().GetTimeService().Now()
				if itemTemplate.GetTrueExpireTime(it.PropertyData.GetItemGetTime()) < now {
					continue
				}
			}
		}

		num += it.Num
	}
	return num
}

func (bb *BagContainer) AddLevelItemWithPropertyDataToEmptySlot(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) int32 {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		panic(fmt.Errorf("inventory:add item id [%d] invalid", itemId))
	}
	now := global.GetGame().GetTimeService().Now()

	totalRemainNum := num
	for totalRemainNum > 0 {
		if bb.NumOfEmptySlots() == 0 {
			return totalRemainNum
		}
		freeIndex := heap.Pop(bb.slotFreeList).(int32)
		if freeIndex >= int32(len(bb.itemList)) {
			panic(fmt.Errorf("inventory:free index [%d] should in item list", freeIndex))
		}
		item := bb.itemList[freeIndex]
		if item.ItemId != 0 {
			panic(fmt.Errorf("inventory:item should be 0"))
		}
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum

		item.ItemId = itemId
		item.Num = complementNum
		item.Level = level
		item.BindType = bindType
		item.PropertyData = propertyData
		item.ItemGetTime = now
		item.UpdateTime = now
		item.SetModified()
		bb.changed(int(freeIndex))
	}
	return totalRemainNum
}

//--------------------分割线-----------------------------

//移除位置的东西
func (bb *BagContainer) RemoveIndex(index int32, num int32) (empty bool) {
	if num <= 0 {
		panic(fmt.Errorf("inventory: remove index:%d num: %d", index, num))
	}
	it := bb.GetByIndex(index)
	if it.Num < num {
		panic(fmt.Errorf("inventory: remove num %d,current num %d", num, it.Num))
	}
	it.Num -= num
	if it.Num == 0 {
		it.ItemId = 0
		it.Level = 0
		it.PropertyData = inventorytypes.CreateDefaultItemPropertyDataBase()
		empty = true
		heap.Push(bb.slotFreeList, it.Index)
	}
	now := global.GetGame().GetTimeService().Now()
	it.UpdateTime = now
	it.LastUseTime = now
	it.SetModified()
	bb.changed(int(it.Index))

	return
}

//添加位置的东西
func (bb *BagContainer) AddIndex(index int32, itemId int32, num int32) {
	if itemId == 0 || num <= 0 {
		panic(fmt.Errorf("inventory: add item id %d num %d", itemId, num))
	}
	it := bb.GetByIndex(index)
	//不能放在空的或不同东西上
	if !it.IsEmpty() && it.ItemId != itemId {
		panic(fmt.Errorf("inventory: add item id %d num %d,in different item id %d", itemId, it.ItemId))
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	maxNum := int32(1)
	if itemTemplate.CanOverlap() {
		maxNum = itemTemplate.MaxOverlap
	}
	if it.Num+num > maxNum {
		panic(fmt.Errorf("inventory: add item id %d currentNum %d num %d,exceed maxNum %d", it.ItemId, it.Num, num, maxNum))
	}
	it.ItemId = itemId
	it.Num += num
	now := global.GetGame().GetTimeService().Now()
	it.UpdateTime = now
	it.SetModified()
	bb.changed(int(it.Index))
	return
}

func (bb *BagContainer) Reset() {
	bb.changedBitset.ClearAll()
}

//设置改变
func (bb *BagContainer) changed(index int) {
	bb.changedBitset.Set(uint(index))
}

//清空
func (bb *BagContainer) ClearAll() {
	for _, item := range bb.itemList {
		if item.Id == 0 {
			continue
		}
		if item.Num == 0 {
			continue
		}
		bb.RemoveIndex(item.Index, item.Num)
	}
}

const (
	pageNum = int32(20)
)
