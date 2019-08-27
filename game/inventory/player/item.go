package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	gameentity "fgame/fgame/game/inventory/entity"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"

	"github.com/pkg/errors"
)

type itemObjectList []*PlayerItemObject

func (iol itemObjectList) Len() int {
	return len(iol)
}

func (iol itemObjectList) Less(i, j int) bool {

	if iol[i].ItemId == 0 {
		return true
	}
	if iol[j].ItemId == 0 {
		return false
	}
	iItem := item.GetItemService().GetItem(int(iol[i].ItemId))
	jItem := item.GetItemService().GetItem(int(iol[j].ItemId))

	//类型相同
	if iItem.Type == jItem.Type {
		//品质相同
		if iItem.Quality == jItem.Quality {
			//按id升序
			return iol[i].ItemId > iol[j].ItemId
		}
		return iItem.Quality < jItem.Quality
	}
	return iItem.Type > jItem.Type
}

func (iol itemObjectList) Swap(i, j int) {
	iol[i], iol[j] = iol[j], iol[i]
}

//玩家道具数据
type PlayerItemObject struct {
	player       player.Player
	Id           int64
	PlayerId     int64
	BagType      inventorytypes.BagType
	ItemId       int32
	Index        int32
	Num          int32
	Used         int32
	Level        int32
	IsDepot      inventorytypes.IsDepotType
	BindType     itemtypes.ItemBindType
	PropertyData inventorytypes.ItemPropertyData
	ItemGetTime  int64
	LastUseTime  int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

func NewPlayerItemObject(pl player.Player) *PlayerItemObject {
	o := &PlayerItemObject{
		player:   pl,
		PlayerId: pl.GetId(),
	}
	return o
}

func convertPlayerItemObjectToEntity(o *PlayerItemObject) *gameentity.PlayerItemEntity {
	data, err := json.Marshal(o.PropertyData)
	if err != nil {
		return nil
	}

	e := &gameentity.PlayerItemEntity{
		Id:           o.Id,
		PlayerId:     o.PlayerId,
		BagType:      int32(o.BagType),
		ItemId:       o.ItemId,
		Index:        o.Index,
		Num:          o.Num,
		Used:         o.Used,
		Level:        o.Level,
		IsDepot:      int32(o.IsDepot),
		BindType:     int32(o.BindType),
		PropertyData: string(data),
		ItemGetTime:  o.ItemGetTime,
		UpdateTime:   o.UpdateTime,
		CreateTime:   o.CreateTime,
		DeleteTime:   o.DeleteTime,
	}
	return e
}

func (o *PlayerItemObject) GetPlayerId() int64 {
	return o.PlayerId
}

func (o *PlayerItemObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerItemObject) ToEntity() (e storage.Entity, err error) {
	e = convertPlayerItemObjectToEntity(o)
	return
}

func (o *PlayerItemObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*gameentity.PlayerItemEntity)

	o.Id = pse.Id
	o.PlayerId = pse.PlayerId
	o.BagType = inventorytypes.BagType(pse.BagType)
	o.Index = pse.Index
	o.ItemId = pse.ItemId
	o.Num = pse.Num
	o.Level = pse.Level
	o.Used = pse.Used
	o.Level = pse.Level
	o.IsDepot = inventorytypes.IsDepotType(pse.IsDepot)
	o.BindType = itemtypes.ItemBindType(pse.BindType)
	o.ItemGetTime = pse.ItemGetTime
	o.LastUseTime = pse.LastUseTime
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime

	if !o.IsEmpty() {
		itemTemp := item.GetItemService().GetItem(int(pse.ItemId))
		//fmt.Println(pse.ItemId)
		data, err := inventory.CreatePropertyData(itemTemp.GetItemType(), pse.PropertyData)
		if err != nil {
			return err
		}
		o.PropertyData = data
	} else {
		o.PropertyData = inventorytypes.CreateDefaultItemPropertyDataBase()
	}
	return
}

func (o *PlayerItemObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Item"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("inventory:set modified never reach here"))
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerItemObject) IsFull() bool {
	if o.ItemId == 0 {
		return false
	}

	if o.PropertyData != nil && o.PropertyData.GetExpireType() != inventorytypes.NewItemLimitTimeTypeNone {
		return true
	}

	itemTemplate := item.GetItemService().GetItem(int(o.ItemId))
	if !itemTemplate.CanOverlap() {
		return true
	}

	return (itemTemplate.MaxOverlap - o.Num) <= 0
}

func (o *PlayerItemObject) IsEmpty() bool {
	return o.ItemId == 0
}

// 是否叠加
func (o *PlayerItemObject) IsOverlap(itemId, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, itemGetTime int64, expireTime int64) bool {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	if o.ItemId != itemId {
		return false
	}
	if o.Level != level {
		return false
	}
	if o.BindType != bindType {
		return false
	}
	if o.PropertyData.GetExpireType() != expireType {
		return false
	}
	if o.PropertyData.GetExpireType() == inventorytypes.NewItemLimitTimeTypeNone {
		if itemTemplate.GetLimitTimeType() == inventorytypes.NewItemLimitTimeTypeNone {
			return true
		}
		// itemGetTime := global.GetGame().GetTimeService().Now()
		if itemTemplate.GetTrueExpireTime(o.PropertyData.GetItemGetTime()) != itemTemplate.GetTrueExpireTime(itemGetTime) {
			return false
		}
		return true
	}
	if o.PropertyData.GetExpireTimestamp() != expireTime {
		return false
	}
	return true
}

// 是否叠加
func (o *PlayerItemObject) IsOverlapExcludeBind(itemId, level int32, expireType inventorytypes.NewItemLimitTimeType, itemGetTime, expireTime int64) bool {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	if o.ItemId != itemId {
		return false
	}
	if o.Level != level {
		return false
	}
	if o.PropertyData.GetExpireType() != expireType {
		return false
	}
	if o.PropertyData.GetExpireType() == inventorytypes.NewItemLimitTimeTypeNone {
		if itemTemplate.GetLimitTimeType() == inventorytypes.NewItemLimitTimeTypeNone {
			return true
		}
		// itemGetTime := global.GetGame().GetTimeService().Now()
		if itemTemplate.GetTrueExpireTime(o.PropertyData.GetItemGetTime()) != itemTemplate.GetTrueExpireTime(itemGetTime) {
			return false
		}
		return true
	}
	if o.PropertyData.GetExpireTimestamp() != expireTime {
		return false
	}
	return true
}

//玩家物品使用数据
type PlayerItemUseObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	ItemId      int32
	TodayTimes  int32
	TotalTimes  int32
	LastUseTime int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerItemUseObject(pl player.Player) *PlayerItemUseObject {
	o := &PlayerItemUseObject{
		player: pl,
	}
	return o
}

func convertPlayerItemUseObjectToEntity(o *PlayerItemUseObject) *gameentity.PlayerItemUseEntity {
	e := &gameentity.PlayerItemUseEntity{
		Id:          o.Id,
		PlayerId:    o.player.GetId(),
		ItemId:      o.ItemId,
		TodayTimes:  o.TodayTimes,
		TotalTimes:  o.TotalTimes,
		LastUseTime: o.LastUseTime,
		UpdateTime:  o.UpdateTime,
		CreateTime:  o.CreateTime,
		DeleteTime:  o.DeleteTime,
	}
	return e
}

func (o *PlayerItemUseObject) GetPlayerId() int64 {
	return o.PlayerId
}

func (o *PlayerItemUseObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerItemUseObject) ToEntity() (e storage.Entity, err error) {
	e = convertPlayerItemUseObjectToEntity(o)
	return
}

func (o *PlayerItemUseObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*gameentity.PlayerItemUseEntity)
	o.Id = pse.Id
	o.PlayerId = pse.PlayerId
	o.ItemId = pse.ItemId
	o.TodayTimes = pse.TodayTimes
	o.TotalTimes = pse.TotalTimes
	o.LastUseTime = pse.LastUseTime
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return
}

func (o *PlayerItemUseObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		return
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("inventory:set modified never reach here"))
	}

	o.player.AddChangedObject(obj)
	return
}
