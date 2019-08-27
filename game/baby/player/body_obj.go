package player

import (
	"fgame/fgame/core/storage"
	babyentity "fgame/fgame/game/baby/entity"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	babytypes "fgame/fgame/game/baby/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	"github.com/pkg/errors"
	"github.com/willf/bitset"
)

//玩具槽位数据
type PlayerBabyToySlotObject struct {
	player     player.Player
	id         int64
	suitType   babytypes.ToySuitType
	slotId     inventorytypes.BodyPositionType
	itemId     int32
	level      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerBabyToySlotObject(pl player.Player) *PlayerBabyToySlotObject {
	o := &PlayerBabyToySlotObject{
		player: pl,
	}
	return o
}

func convertPlayerBabyToySlotObjectToEntity(o *PlayerBabyToySlotObject) (*babyentity.PlayerBabyToySlotEntity, error) {

	e := &babyentity.PlayerBabyToySlotEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		SuitType:   int32(o.suitType),
		SlotId:     int32(o.slotId),
		ItemId:     o.itemId,
		Level:      o.level,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerBabyToySlotObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerBabyToySlotObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerBabyToySlotObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerBabyToySlotObject) GetItemId() int32 {
	return o.itemId
}

func (o *PlayerBabyToySlotObject) GetSlotId() inventorytypes.BodyPositionType {
	return o.slotId
}

func (o *PlayerBabyToySlotObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerBabyToySlotObjectToEntity(o)
	return
}

func (o *PlayerBabyToySlotObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*babyentity.PlayerBabyToySlotEntity)
	suitType := babytypes.ToySuitType(pse.SuitType)

	o.id = pse.Id
	o.suitType = suitType
	o.slotId = inventorytypes.BodyPositionType(pse.SlotId)
	o.itemId = pse.ItemId
	o.level = pse.Level
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return
}

func (o *PlayerBabyToySlotObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "BabyToySlot"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerBabyToySlotObject) IsEmpty() bool {
	return o.itemId == 0
}

func (o *PlayerBabyToySlotObject) IsFull() bool {
	return o.itemId != 0
}

//宝宝玩具位置包裹
type BodyBag struct {
	p       player.Player
	slotMap map[inventorytypes.BodyPositionType]*PlayerBabyToySlotObject
	//改变的位置
	changedBitset *bitset.BitSet
}

func (bb *BodyBag) GetAll() (slotList []*PlayerBabyToySlotObject) {
	for _, slot := range bb.slotMap {
		slotList = append(slotList, slot)
	}
	return
}

//获取根据位置
func (bb *BodyBag) GetByPosition(pos inventorytypes.BodyPositionType) *PlayerBabyToySlotObject {
	return bb.slotMap[pos]
}

//穿上
func (bb *BodyBag) PutOn(itemId int32, pos inventorytypes.BodyPositionType) bool {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	if itemTemplate.GetBabyToyTemplate() == nil {
		return false
	}

	//位置不存在
	bodySlot := bb.GetByPosition(pos)
	if bodySlot == nil {
		return false
	}

	if !bodySlot.IsEmpty() {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	bodySlot.level = 1
	bodySlot.itemId = itemId
	bodySlot.updateTime = now
	bodySlot.SetModified()
	bb.changed(int(pos))

	eventData := babyeventtypes.CreatePlayerBabyToyChangedEventData(bb.CountSuitGroupNum(), itemId)
	gameevent.Emit(babyeventtypes.EventTypeBabyUseToy, bb.p, eventData)
	return true
}

//玩具升级
func (bb *BodyBag) ToyUplevel(pos inventorytypes.BodyPositionType) bool {
	slotIt := bb.GetByPosition(pos)
	if slotIt == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	slotIt.level += 1
	slotIt.updateTime = now
	slotIt.SetModified()
	bb.changed(int(pos))

	gameevent.Emit(babyeventtypes.EventTypeBabyToyUplevel, bb.p, nil)
	return true
}

//玩具回退
func (bb *BodyBag) ToyReturn(pos inventorytypes.BodyPositionType, returnLevel int32) bool {
	slotIt := bb.GetByPosition(pos)
	if slotIt == nil {
		return false
	}

	if slotIt.level < returnLevel {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	slotIt.level = returnLevel
	slotIt.updateTime = now
	slotIt.SetModified()
	bb.changed(int(pos))
	return true
}

func (bb *BodyBag) GetChangedSlotAndReset() []*PlayerBabyToySlotObject {
	itemList := make([]*PlayerBabyToySlotObject, 0, 16)
	for i, valid := bb.changedBitset.NextSet(0); valid; i, valid = bb.changedBitset.NextSet(i + 1) {
		itemList = append(itemList, bb.slotMap[inventorytypes.BodyPositionType(i)])
	}
	bb.Reset()
	return itemList
}

func (bb *BodyBag) CountSuitGroupNum() map[int32]int32 {
	suitGroupMap := make(map[int32]int32)
	for _, slot := range bb.GetAll() {
		if slot.IsEmpty() {
			continue
		}
		itemTemp := item.GetItemService().GetItem(int(slot.itemId))
		groupId := itemTemp.GetBabyToyTemplate().SuitGroup
		if groupId == 0 {
			continue
		}

		suitGroupMap[groupId] += 1
	}
	return suitGroupMap
}

func (bb *BodyBag) Reset() {
	bb.changedBitset.ClearAll()
}

//设置改变
func (bb *BodyBag) changed(index int) {
	bb.changedBitset.Set(uint(index))
}

//创建身体背包
func createBodyBag(p player.Player, suitType babytypes.ToySuitType, slotList []*PlayerBabyToySlotObject) *BodyBag {
	bb := &BodyBag{
		p: p,
	}

	bb.init(suitType, slotList)
	return bb
}

//初始化
func (bb *BodyBag) init(suitType babytypes.ToySuitType, slotList []*PlayerBabyToySlotObject) {
	bb.changedBitset = bitset.New(64)
	bb.slotMap = make(map[inventorytypes.BodyPositionType]*PlayerBabyToySlotObject)
	for _, slot := range slotList {
		bb.slotMap[slot.slotId] = slot
	}

	now := global.GetGame().GetTimeService().Now()
	for slotId := inventorytypes.MinToyPos; slotId <= inventorytypes.MaxToyPos; slotId++ {
		if bb.GetByPosition(slotId) != nil {
			continue
		}
		slot := createBabyToySlotObject(bb.p, suitType, slotId, now)
		slot.SetModified()
		bb.slotMap[slot.slotId] = slot
	}
}

func createBabyToySlotObject(p player.Player, suitType babytypes.ToySuitType, slotId inventorytypes.BodyPositionType, now int64) *PlayerBabyToySlotObject {
	itemObject := NewPlayerBabyToySlotObject(p)
	itemObject.createTime = now
	itemObject.id, _ = idutil.GetId()
	itemObject.level = 0
	itemObject.itemId = 0
	itemObject.slotId = slotId
	itemObject.suitType = suitType
	itemObject.createTime = now
	return itemObject
}
