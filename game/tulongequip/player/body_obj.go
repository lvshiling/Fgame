package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	tulongequipentity "fgame/fgame/game/tulongequip/entity"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	"github.com/pkg/errors"
	"github.com/willf/bitset"
)

//元神装备身体位置包裹
type BodyBag struct {
	p       player.Player
	slotMap map[inventorytypes.BodyPositionType]*PlayerTuLongEquipSlotObject
	//改变的位置
	changedBitset *bitset.BitSet
}

func (bb *BodyBag) GetAll() (slotList []*PlayerTuLongEquipSlotObject) {
	for _, slot := range bb.slotMap {
		slotList = append(slotList, slot)
	}
	return
}

//获取根据位置
func (bb *BodyBag) GetByPosition(pos inventorytypes.BodyPositionType) *PlayerTuLongEquipSlotObject {
	return bb.slotMap[pos]
}

//穿上
func (bb *BodyBag) PutOn(pos inventorytypes.BodyPositionType, itemId int32, bind itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) bool {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if !itemTemplate.IsTuLongEquip() {
		return false
	}

	//位置不存在
	bodySlot := bb.GetByPosition(pos)
	if bodySlot == nil {
		return false
	}

	if bodySlot.IsEmpty() {
		now := global.GetGame().GetTimeService().Now()
		bodySlot.itemId = itemId
		bodySlot.updateTime = now
		bodySlot.bind = bind
		bodySlot.propertyData = propertyData
		bodySlot.SetModified()
		bb.changed(int(pos))
		return true
	}

	return false
}

//脱下
func (bb *BodyBag) TakeOff(pos inventorytypes.BodyPositionType) (flag bool, itemId int32) {
	bodySlot := bb.GetByPosition(pos)
	if bodySlot == nil {
		return
	}
	if bodySlot.IsEmpty() {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	itemId = bodySlot.itemId
	bodySlot.itemId = 0
	bodySlot.propertyData = inventorytypes.CreateDefaultItemPropertyDataBase()
	bodySlot.updateTime = now
	bodySlot.SetModified()
	bb.changed(int(pos))
	flag = true
	return
}

func (bb *BodyBag) GetChangedSlotAndReset() []*PlayerTuLongEquipSlotObject {
	itemList := make([]*PlayerTuLongEquipSlotObject, 0, 16)
	for i, valid := bb.changedBitset.NextSet(0); valid; i, valid = bb.changedBitset.NextSet(i + 1) {
		itemList = append(itemList, bb.slotMap[inventorytypes.BodyPositionType(i)])
	}
	bb.Reset()
	return itemList
}

func (bb *BodyBag) Reset() {
	bb.changedBitset.ClearAll()
}

//设置改变
func (bb *BodyBag) changed(index int) {
	bb.changedBitset.Set(uint(index))
}

// //是否有宝石
// func (bb *BodyBag) IfEmbedGem(pos inventorytypes.BodyPositionType, order int32) (flag bool) {
// 	item := bb.GetByPosition(pos)
// 	if item == nil {
// 		return
// 	}
// 	if item.IsEmpty() {
// 		return
// 	}
// 	//获取装备槽宝石信息
// 	_, exist := item.GemInfo[order]
// 	if !exist {
// 		return
// 	}

// 	return true
// }

// //脱下宝石
// func (bb *BodyBag) TakeOffGem(pos inventorytypes.BodyPositionType, order int32) (itemId int32) {
// 	item := bb.GetByPosition(pos)
// 	if item == nil {
// 		return
// 	}
// 	if item.IsEmpty() {
// 		return
// 	}
// 	//获取装备槽宝石信息
// 	itemId, exist := item.GemInfo[order]
// 	if !exist {
// 		return
// 	}
// 	delete(item.GemInfo, order)
// 	now := global.GetGame().GetTimeService().Now()
// 	item.updateTime = now
// 	item.SetModified()
// 	bb.changed(int(pos))
// 	gameevent.Emit(tulongequipeventtypes.EventTypeTuLongEquipTakeOffGem, bb.p, nil)
// 	return
// }

// //佩戴宝石
// func (bb *BodyBag) PutOnGem(pos inventorytypes.BodyPositionType, order int32, itemId int32) bool {
// 	//物品是不是宝石
// 	itemTemp := item.GetItemService().GetItem(int(itemId))
// 	if !itemTemp.IsGem() {
// 		return false
// 	}

// 	item := bb.GetByPosition(pos)
// 	//位置不存在
// 	if item == nil {
// 		return false
// 	}

// 	//装备不存在
// 	if item.IsEmpty() {
// 		return false
// 	}

// 	//获取装备槽宝石信息
// 	_, exist := item.GemInfo[order]
// 	if exist {
// 		return false
// 	}
// 	item.GemInfo[order] = itemId
// 	now := global.GetGame().GetTimeService().Now()
// 	item.updateTime = now
// 	item.SetModified()
// 	bb.changed(int(pos))

// 	gameevent.Emit(tulongequipeventtypes.EventTypeTuLongEquipEmbedGem, bb.p, nil)
// 	return true
// }

//创建身体背包
func createBodyBag(p player.Player, suitType tulongequiptypes.TuLongSuitType, slotList []*PlayerTuLongEquipSlotObject) *BodyBag {
	bb := &BodyBag{
		p: p,
	}

	bb.init(suitType, slotList)
	return bb
}

//初始化
func (bb *BodyBag) init(suitType tulongequiptypes.TuLongSuitType, slotList []*PlayerTuLongEquipSlotObject) {
	bb.changedBitset = bitset.New(64)
	bb.slotMap = make(map[inventorytypes.BodyPositionType]*PlayerTuLongEquipSlotObject)
	for _, slot := range slotList {
		bb.slotMap[slot.slotId] = slot
	}

	now := global.GetGame().GetTimeService().Now()
	for slotId := inventorytypes.BodyPositionTypeWeapon; slotId <= inventorytypes.BodyPositionTypeRing; slotId++ {
		if bb.GetByPosition(slotId) != nil {
			continue
		}
		slot := createTuLongEquipSlotObject(bb.p, suitType, slotId, now)
		slot.SetModified()
		bb.slotMap[slot.slotId] = slot
	}
}

func createTuLongEquipSlotObject(p player.Player, suitType tulongequiptypes.TuLongSuitType, slotId inventorytypes.BodyPositionType, now int64) *PlayerTuLongEquipSlotObject {
	itemObject := NewPlayerTuLongEquipSlotObject(p)
	itemObject.createTime = now
	itemObject.id, _ = idutil.GetId()
	itemObject.itemId = 0
	itemObject.slotId = slotId
	itemObject.suitType = suitType
	base := inventorytypes.CreateDefaultItemPropertyDataBase()
	itemObject.propertyData = inventory.CreatePropertyDataInterface(itemtypes.ItemTypeTuLongEquip, base)
	itemObject.GemInfo = make(map[int32]int32)
	itemObject.createTime = now
	return itemObject
}

//玩家槽位数据
type PlayerTuLongEquipSlotObject struct {
	player       player.Player
	id           int64
	playerId     int64
	suitType     tulongequiptypes.TuLongSuitType
	slotId       inventorytypes.BodyPositionType
	itemId       int32
	level        int32
	bind         itemtypes.ItemBindType
	propertyData inventorytypes.ItemPropertyData
	GemInfo      map[int32]int32
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerTuLongEquipSlotObject(pl player.Player) *PlayerTuLongEquipSlotObject {
	o := &PlayerTuLongEquipSlotObject{
		player:   pl,
		playerId: pl.GetId(),
	}
	return o
}

func convertPlayerTuLongEquipSlotObjectToEntity(o *PlayerTuLongEquipSlotObject) (*tulongequipentity.PlayerTuLongEquipSlotEntity, error) {
	data, err := json.Marshal(o.propertyData)
	if err != nil {
		return nil, err
	}

	gemInfoBytes, err := json.Marshal(o.GemInfo)
	if err != nil {
		return nil, err
	}

	e := &tulongequipentity.PlayerTuLongEquipSlotEntity{
		Id:           o.id,
		PlayerId:     o.playerId,
		SuitType:     int32(o.suitType),
		ItemId:       o.itemId,
		SlotId:       int32(o.slotId),
		Level:        o.level,
		BindType:     int32(o.bind),
		PropertyData: string(data),
		GemInfo:      string(gemInfoBytes),
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *PlayerTuLongEquipSlotObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *PlayerTuLongEquipSlotObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerTuLongEquipSlotObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerTuLongEquipSlotObject) GetBindType() itemtypes.ItemBindType {
	return o.bind
}

func (o *PlayerTuLongEquipSlotObject) GetItemId() int32 {
	return o.itemId
}

func (o *PlayerTuLongEquipSlotObject) GetSlotId() inventorytypes.BodyPositionType {
	return o.slotId
}

func (o *PlayerTuLongEquipSlotObject) GetPropertyData() inventorytypes.ItemPropertyData {
	return o.propertyData
}

func (o *PlayerTuLongEquipSlotObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerTuLongEquipSlotObjectToEntity(o)
	return
}

func (o *PlayerTuLongEquipSlotObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*tulongequipentity.PlayerTuLongEquipSlotEntity)
	suitType := tulongequiptypes.TuLongSuitType(pse.SuitType)

	data, err := inventory.CreatePropertyData(itemtypes.ItemTypeTuLongEquip, pse.PropertyData)
	if err != nil {
		return
	}

	gemInfo := make(map[int32]int32)
	err = json.Unmarshal([]byte(pse.GemInfo), &gemInfo)
	if err != nil {
		return
	}

	o.id = pse.Id
	o.playerId = pse.PlayerId
	o.suitType = suitType
	o.itemId = pse.ItemId
	o.slotId = inventorytypes.BodyPositionType(pse.SlotId)
	o.level = pse.Level
	o.bind = itemtypes.ItemBindType(pse.BindType)
	o.GemInfo = gemInfo
	o.propertyData = data
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return
}

func (o *PlayerTuLongEquipSlotObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TuLongEquipSlot"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerTuLongEquipSlotObject) IsEmpty() bool {
	return o.itemId == 0
}

func (o *PlayerTuLongEquipSlotObject) IsFull() bool {
	return o.itemId != 0
}
