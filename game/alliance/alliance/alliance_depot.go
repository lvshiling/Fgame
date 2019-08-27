package alliance

import (
	"container/heap"
	"encoding/json"
	"fgame/fgame/core/storage"
	allianceentity "fgame/fgame/game/alliance/entity"
	"fgame/fgame/game/global"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"github.com/willf/bitset"
)

type itemObjectList []*AllianceDepotItemObject

func (iol itemObjectList) Len() int {
	return len(iol)
}

func (iol itemObjectList) Less(i, j int) bool {

	if iol[i].itemId == 0 {
		return true
	}
	if iol[j].itemId == 0 {
		return false
	}
	iItem := item.GetItemService().GetItem(int(iol[i].itemId))
	jItem := item.GetItemService().GetItem(int(iol[j].itemId))

	//类型相同
	if iItem.Type == jItem.Type {
		//品质相同
		if iItem.Quality == jItem.Quality {
			//按id升序
			return iol[i].itemId > iol[j].itemId
		}
		return iItem.Quality < jItem.Quality
	}
	return iItem.Type > jItem.Type
}

func (iol itemObjectList) Swap(i, j int) {
	iol[i], iol[j] = iol[j], iol[i]
}

//仙盟仓库对象
type AllianceDepotItemObject struct {
	id           int64
	allianceId   int64
	itemId       int32
	index        int32
	num          int32
	used         int32
	level        int32
	bindType     itemtypes.ItemBindType
	propertyData inventorytypes.ItemPropertyData
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func createAllianceDepotItemObject() *AllianceDepotItemObject {
	o := &AllianceDepotItemObject{}
	return o
}

func convertAllianceDepotItemObjectToEntity(o *AllianceDepotItemObject) (*allianceentity.AllianceDepotEntity, error) {
	data, err := json.Marshal(o.propertyData)
	if err != nil {
		return nil, err
	}

	e := &allianceentity.AllianceDepotEntity{
		Id:           o.id,
		AllianceId:   o.allianceId,
		ItemId:       o.itemId,
		Index:        o.index,
		Num:          o.num,
		Used:         o.used,
		Level:        o.level,
		BindType:     int32(o.bindType),
		PropertyData: string(data),
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *AllianceDepotItemObject) GetId() int64 {
	return o.id
}

func (o *AllianceDepotItemObject) GetItemId() int32 {
	return o.itemId
}

func (o *AllianceDepotItemObject) GetNum() int32 {
	return o.num
}

func (o *AllianceDepotItemObject) GetIndex() int32 {
	return o.index
}

func (o *AllianceDepotItemObject) GetLevel() int32 {
	return o.level
}

func (o *AllianceDepotItemObject) GetBindType() itemtypes.ItemBindType {
	return o.bindType
}

func (o *AllianceDepotItemObject) GetPropertyData() inventorytypes.ItemPropertyData {
	return o.propertyData
}

func (o *AllianceDepotItemObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceDepotItemObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *AllianceDepotItemObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *AllianceDepotItemObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertAllianceDepotItemObjectToEntity(o)
	return e, err
}

func (o *AllianceDepotItemObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*allianceentity.AllianceDepotEntity)

	o.id = ae.Id
	o.allianceId = ae.AllianceId
	o.itemId = ae.ItemId
	o.index = ae.Index
	o.num = ae.Num
	o.used = ae.Used
	o.level = ae.Level
	o.bindType = itemtypes.ItemBindType(ae.BindType)
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime

	if !o.IsEmpty() {
		itemTemp := item.GetItemService().GetItem(int(ae.ItemId))
		data, err := inventory.CreatePropertyData(itemTemp.GetItemType(), ae.PropertyData)
		if err != nil {
			return err
		}
		o.propertyData = data
	} else {
		o.propertyData = inventorytypes.CreateDefaultItemPropertyDataBase()
	}
	return nil
}

func (o *AllianceDepotItemObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceDepot"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *AllianceDepotItemObject) IsFull() bool {
	if o.itemId == 0 {
		return false
	}

	itemTemplate := item.GetItemService().GetItem(int(o.itemId))
	if !itemTemplate.CanOverlap() {
		return true
	}
	return (itemTemplate.MaxOverlap - o.num) <= 0
}

func (o *AllianceDepotItemObject) IsEmpty() bool {
	return o.itemId == 0
}

// 是否叠加
func (o *AllianceDepotItemObject) IsOverlap(itemId, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, expireTime int64) bool {
	if o.itemId != itemId {
		return false
	}
	if o.level != level {
		return false
	}
	if o.bindType != bindType {
		return false
	}

	if o.propertyData.GetExpireType() != expireType {
		return false
	}
	if o.propertyData.GetExpireType() == inventorytypes.NewItemLimitTimeTypeNone {
		return true
	}
	if o.propertyData.GetExpireTimestamp() != expireTime {
		return false
	}

	return true
}

//空闲格子列表
// 实现“优先队列”接口 container/heap
type slotFreeList []int32

func (sfl *slotFreeList) Len() int {
	return len(*sfl)
}

// 索引大的在队列尾
func (sfl *slotFreeList) Less(i, j int) bool {
	return (*sfl)[i] < (*sfl)[j]
}

func (sfl *slotFreeList) Swap(i, j int) {
	(*sfl)[i], (*sfl)[j] = (*sfl)[j], (*sfl)[i]
}

func (sfl *slotFreeList) Push(h interface{}) {
	*sfl = append(*sfl, h.(int32))
}

// 队列尾删除
func (sfl *slotFreeList) Pop() (v interface{}) {
	*sfl, v = (*sfl)[0:sfl.Len()-1], (*sfl)[sfl.Len()-1]
	return
}

//背包基础
type BagContainer struct {
	//物品列表
	itemList []*AllianceDepotItemObject
	//改变的位置
	changedBitset *bitset.BitSet
	//空闲的索引
	slotFreeList *slotFreeList
}

//创建新背包
func createBagContainer(itemList []*AllianceDepotItemObject) *BagContainer {
	bb := &BagContainer{}
	bb.itemList = itemList
	bb.changedBitset = bitset.New(64)
	bb.slotFreeList = &slotFreeList{}
	return bb
}

func (bb *BagContainer) Reset() {
	bb.changedBitset.ClearAll()
}

//添加物品
// 在未被使用的空格
func (bb *BagContainer) addUseItem(item *AllianceDepotItemObject) {
	if int(item.index) != len(bb.itemList) {
		panic(fmt.Errorf("alliance:allianceId[%d] add item index %d not last %d", item.allianceId, item.index, len(bb.itemList)))
	}
	bb.changed(int(item.index))
	bb.itemList = append(bb.itemList, item)
	if item.itemId == 0 {
		heap.Push(bb.slotFreeList, int32(item.index))
		return
	}
}

//获取位置
func (bb *BagContainer) getByIndex(index int32) *AllianceDepotItemObject {
	if index < 0 {
		panic(fmt.Errorf("alliance:get by index %d should be no less than 0", index))
	}
	if int(index) >= len(bb.itemList) {
		return nil
	}
	return bb.itemList[index]
}

//长度
func (bb *BagContainer) Len() int32 {
	return int32(len(bb.itemList))
}

//获取物品列表
func (bb *BagContainer) GetItemList() []*AllianceDepotItemObject {
	return bb.itemList
}

//获取空的位置
func (bb *BagContainer) NumOfEmptySlots() int32 {
	return int32(len(*bb.slotFreeList))
}

//设置改变
// 哪个位置的物品改变了，置1
func (bb *BagContainer) changed(index int) {
	bb.changedBitset.Set(uint(index))
}

//获取改变
// 遍历，找1
func (bb *BagContainer) getChangedSlotAndReset() (itemList []*AllianceDepotItemObject) {
	itemList = make([]*AllianceDepotItemObject, 0, 16)
	for i, valid := bb.changedBitset.NextSet(0); valid; i, valid = bb.changedBitset.NextSet(i + 1) {
		itemList = append(itemList, bb.itemList[i])
	}
	// 清除所有改变，置0
	bb.Reset()
	return itemList
}

//移除位置的东西
func (bb *BagContainer) removeIndex(index int32, num int32) (empty bool) {
	if num <= 0 {
		panic(fmt.Errorf("alliance: remove index:%d num: %d", index, num))
	}
	it := bb.getByIndex(index)
	if it.num < num {
		panic(fmt.Errorf("alliance: remove num %d,current num %d", num, it.num))
	}
	it.num -= num
	if it.num == 0 {
		it.itemId = 0
		it.level = 0
		it.propertyData = inventorytypes.CreateDefaultItemPropertyDataBase()
		empty = true
		// 移除位置的索引在队列尾添加
		heap.Push(bb.slotFreeList, it.index)
	}
	now := global.GetGame().GetTimeService().Now()
	it.updateTime = now
	it.SetModified()
	bb.changed(int(it.index))
	return
}

// 移除空格（未使用的位置及used=0）添加物品，返回不够空间存放的物品数量
func (bb *BagContainer) addLevelItemWithPropertyDataToEmptySlot(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) int32 {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		panic(fmt.Errorf("alliance:add item id [%d] invalid", itemId))
	}
	now := global.GetGame().GetTimeService().Now()

	totalRemainNum := num
	for totalRemainNum > 0 {
		if bb.NumOfEmptySlots() == 0 {
			return totalRemainNum
		}

		// 移除空格，其实移除的是队列首部的空格
		freeIndex := heap.Pop(bb.slotFreeList).(int32)
		if freeIndex >= int32(len(bb.itemList)) {
			panic(fmt.Errorf("alliance:free index [%d] should in item list", freeIndex))
		}
		item := bb.itemList[freeIndex]
		if item.itemId != 0 {
			panic(fmt.Errorf("alliance:item should be 0"))
		}
		// 判断物品的最大叠加上限
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		// 该空格的物品数量
		totalRemainNum -= complementNum

		item.itemId = itemId
		item.num = complementNum
		item.level = level
		item.bindType = bindType
		item.propertyData = propertyData
		item.updateTime = now
		item.SetModified()
		bb.changed(int(freeIndex))
	}
	// 返回一个空格不能叠加的物品数量
	return totalRemainNum
}

//添加到有物品的位置
func (bb *BagContainer) addItemLevelToOverlapSlot(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) int32 {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		panic(fmt.Errorf("alliance:add item id [%d] invalid", itemId))
	}

	totalRemainNum := num

	if !itemTemplate.CanOverlap() {
		return totalRemainNum
	}

	now := global.GetGame().GetTimeService().Now()
	for index, item := range bb.itemList {
		if !item.IsOverlap(itemId, level, bindType, propertyData.GetExpireType(), propertyData.GetExpireTimestamp()) {
			continue
		}

		//获取剩余可以叠加数量
		remainNum := itemTemplate.MaxOverlap - item.num
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
		item.num += complementNum
		item.updateTime = now
		item.SetModified()
		// 物品改变标记
		bb.changed(index)
		//放置完成
		if totalRemainNum <= 0 {
			break
		}
	}
	// 返回不能叠加的物品数量
	return totalRemainNum
}

//添加位置的东西
func (bb *BagContainer) addIndex(index int32, itemId int32, num int32) {
	if itemId == 0 || num <= 0 {
		panic(fmt.Errorf("inventory: add item id %d num %d", itemId, num))
	}
	it := bb.getByIndex(index)
	//不能放在空的或不同东西上
	if !it.IsEmpty() && it.itemId != itemId {
		panic(fmt.Errorf("inventory: add item id %d num %d,in different item id %d", itemId, it.itemId))
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	maxNum := int32(1)
	if itemTemplate.CanOverlap() {
		maxNum = itemTemplate.MaxOverlap
	}
	if it.num+num > maxNum {
		panic(fmt.Errorf("inventory: add item id %d currentNum %d num %d,exceed maxNum %d", it.itemId, it.num, num, maxNum))
	}
	it.itemId = itemId
	it.num += num
	now := global.GetGame().GetTimeService().Now()
	it.updateTime = now
	it.SetModified()
	bb.changed(int(it.index))
	return
}

//
// 仙盟仓库
type AllianceDepotBag struct {
	*BagContainer
	//仙盟id
	allianceId int64
	//位置容量
	capacity int32
	//没有使用过的
	unusedItemList []*AllianceDepotItemObject
}

//添加格子
func (sb *AllianceDepotBag) addItem(item *AllianceDepotItemObject) {
	// 是否使用过，有物品的格子都是1
	if item.used == 0 {
		sb.unusedItemList = append(sb.unusedItemList, item)
		return
	}

	sb.BagContainer.addUseItem(item)
}

//获取空闲的格子
func (sb *AllianceDepotBag) GetEmptySlots() int32 {

	emptyNum := int32(0)

	//获取剩余位置
	remainSlotNum := sb.capacity - sb.BagContainer.Len()
	emptyNum += remainSlotNum
	//使用空的位置
	numOfEmptySlots := sb.BagContainer.NumOfEmptySlots()
	emptyNum += numOfEmptySlots

	return emptyNum
}

//需要几个格子
func (sb *AllianceDepotBag) countNeedSlotOfItemLevel(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, expireTime int64) (slotNum int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return 0
	}
	maxNum := int32(0)
	// 计算叠加完剩余的物品数量
	if itemTemplate.CanOverlap() {
		for _, item := range sb.BagContainer.GetItemList() {
			if !item.IsOverlap(itemId, level, bindType, expireType, expireTime) {
				continue
			}

			remainNum := itemTemplate.MaxOverlap - item.num
			if remainNum > 0 {
				maxNum += remainNum
			}
		}
	}
	//不需要额外的空格子
	if maxNum >= num {
		return
	}

	// 计算最终需要几个空格子
	remainNum := num - maxNum
	maxOverlap := itemTemplate.MaxOverlap
	return (remainNum + maxOverlap - 1) / maxOverlap
}

func (sb *AllianceDepotBag) addLevelItem(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) bool {
	if num <= 0 {
		panic(fmt.Errorf("alliance:add item %d,num %d should more than 0", itemId, num))
	}

	// 仓库里放不下那么多物品了
	if !sb.hasEnoughSlotItemLevelWithProperty(itemId, num, level, bindType, propertyData.GetExpireType(), propertyData.GetExpireTimestamp()) {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate.GetBindType() == itemtypes.ItemBindTypeBind {
		bindType = itemtypes.ItemBindTypeBind
	}

	// 先计算叠加后剩余物品数量
	totalRemainNum := sb.BagContainer.addItemLevelToOverlapSlot(itemId, num, level, bindType, propertyData)
	//填充使用过的空格子
	totalRemainNum = sb.BagContainer.addLevelItemWithPropertyDataToEmptySlot(itemId, totalRemainNum, level, bindType, propertyData)

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
		unusedItem.itemId = itemId
		unusedItem.index = index
		unusedItem.updateTime = now
		unusedItem.num = complementNum
		unusedItem.level = level
		unusedItem.used = 1
		unusedItem.bindType = bindType
		unusedItem.propertyData = propertyData
		unusedItem.SetModified()
		sb.addItem(unusedItem)
	}

	//填充没有使用过的格子
	for totalRemainNum > 0 {
		nextIndex := int32(sb.Len())
		//超过格子了
		if nextIndex >= sb.capacity {
			return false
		}
		complementNum := itemTemplate.MaxOverlap
		if totalRemainNum <= complementNum {
			complementNum = totalRemainNum
		}
		totalRemainNum -= complementNum

		item := createItem(sb.allianceId, itemId, nextIndex, complementNum, bindType)
		item.level = level
		item.propertyData = propertyData
		item.SetModified()
		sb.addItem(item)
	}
	if totalRemainNum != 0 {
		panic(fmt.Errorf("total remain num [%d],never reach herer", totalRemainNum))
	}
	return true
}

//有足够的空间
// func (sb *AllianceDepotBag) hasEnoughSlotItemLevel(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType) bool {
// 	remainNum := sb.remainSlotForItemLevel(itemId, level, bindType)
// 	if remainNum >= num {
// 		return true
// 	}
// 	return false
// }

//是否有足够的位置存放物品(含时效性)
func (sb *AllianceDepotBag) hasEnoughSlotItemLevelWithProperty(itemId, num, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, expireTime int64) bool {
	remainNum := sb.remainSlotForItemLevelWithProperty(itemId, level, bindType, expireType, expireTime)
	if remainNum >= num {
		return true
	}
	return false
}

//剩余填充物品数
func (sb *AllianceDepotBag) remainSlotForItemLevelWithProperty(itemId, itemLevel int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, expireTime int64) (maxNum int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return 0
	}
	maxNum = int32(0)
	if itemTemplate.CanOverlap() {
		for _, item := range sb.GetItemList() {
			if !item.IsOverlap(itemId, itemLevel, bindType, expireType, expireTime) {
				continue
			}

			remainNum := itemTemplate.MaxOverlap - item.GetNum()
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

//剩余填充物品数
// func (sb *AllianceDepotBag) remainSlotForItemLevel(itemId, itemLevel int32, bindType itemtypes.ItemBindType) (maxNum int32) {
// 	itemTemplate := item.GetItemService().GetItem(int(itemId))
// 	if itemTemplate == nil {
// 		return 0
// 	}
// 	maxNum = int32(0)
// 	if itemTemplate.CanOverlap() {
// 		for _, item := range sb.GetItemList() {
// 			if item.itemId != itemId {
// 				continue
// 			}
// 			if item.level != itemLevel {
// 				continue
// 			}
// 			if item.bindType != bindType {
// 				continue
// 			}

// 			remainNum := itemTemplate.MaxOverlap - item.num
// 			if remainNum > 0 {
// 				maxNum += remainNum
// 			}

// 		}
// 	}

// 	//获取剩余位置
// 	remainSlotNum := sb.capacity - sb.Len()

// 	//使用空的位置
// 	numOfEmptySlots := sb.NumOfEmptySlots()
// 	remainSlotNum += numOfEmptySlots
// 	if remainSlotNum <= 0 {
// 		return
// 	}
// 	maxNum += itemTemplate.MaxOverlap * int32(remainSlotNum)
// 	return
// }

//合并
func (sb *AllianceDepotBag) merge() {
	sb.classifyAndMerge()
	//获取物品列表
	itemList := sb.BagContainer.GetItemList()

	sort.Sort(sort.Reverse(itemObjectList(itemList)))
	//TODO 优化
	sb.BagContainer = createBagContainer(nil)

	unusedItemList := make([]*AllianceDepotItemObject, 0, 16)
	//修改索引
	for index, item := range itemList {
		dirty := false
		//切换索引
		if item.index != int32(index) {
			item.index = int32(index)
			dirty = true
		}

		//设置为没使用
		if item.itemId == 0 {
			item.used = 0
			dirty = true
		}
		//改变过了
		if dirty {
			item.SetModified()
		}

		//使用过
		if item.used != 0 {
			sb.addItem(item)
		} else {
			unusedItemList = append(unusedItemList, item)
		}
	}
	sb.unusedItemList = append(unusedItemList, sb.unusedItemList...)
	//清空标记位
	sb.Reset()
}

//分类合并
func (sb *AllianceDepotBag) classifyAndMerge() {
	classifyMap := make(map[int32][]*AllianceDepotItemObject)
	for _, it := range sb.BagContainer.GetItemList() {
		if it.IsEmpty() {
			continue
		}
		l, exist := classifyMap[it.itemId]
		if !exist {
			l = make([]*AllianceDepotItemObject, 0, 8)
			classifyMap[it.itemId] = l
		}
		classifyMap[it.itemId] = append(l, it)
	}

	for _, classifyList := range classifyMap {
		var previousItem *AllianceDepotItemObject
		for _, it := range classifyList {
			itemTemplate := item.GetItemService().GetItem(int(it.itemId))
			if previousItem == nil {
				//未满
				if !it.IsFull() {
					previousItem = it
				}
				continue
			}
			if !previousItem.IsOverlap(it.itemId, it.level, it.bindType, it.propertyData.GetExpireType(), it.propertyData.GetExpireTimestamp()) {
				continue
			}

			//计算剩余的
			remainNum := itemTemplate.MaxOverlap - previousItem.num
			complementNum := remainNum
			if it.num <= remainNum {
				complementNum = it.num
			}
			//移除当前的
			sb.BagContainer.removeIndex(it.index, complementNum)
			//添加之前的
			sb.BagContainer.addIndex(previousItem.index, previousItem.itemId, complementNum)
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

func createDepotBag(allianceId int64, items []*AllianceDepotItemObject, capacity int32) *AllianceDepotBag {
	sb := &AllianceDepotBag{}
	sb.BagContainer = createBagContainer(nil)
	sb.capacity = capacity
	sb.unusedItemList = make([]*AllianceDepotItemObject, 0, 16)
	sb.allianceId = allianceId

	for _, item := range items {

		sb.addItem(item)
	}

	sb.BagContainer.Reset()
	return sb
}

func createItem(allianceId int64, itemId int32, index int32, num int32, bindType itemtypes.ItemBindType) *AllianceDepotItemObject {
	now := global.GetGame().GetTimeService().Now()
	itemObject := createAllianceDepotItemObject()
	itemObject.createTime = now
	itemObject.id, _ = idutil.GetId()
	itemObject.allianceId = allianceId
	itemObject.itemId = itemId
	itemObject.num = num
	itemObject.index = index
	itemObject.bindType = bindType
	itemObject.used = 1
	return itemObject
}
