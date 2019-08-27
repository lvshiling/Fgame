package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	ringentity "fgame/fgame/game/ring/entity"
	ringtypes "fgame/fgame/game/ring/types"

	"github.com/pkg/errors"
)

type PlayerRingObject struct {
	id           int64
	player       player.Player
	typ          ringtypes.RingType
	bindType     itemtypes.ItemBindType
	itemId       int32
	propertyData inventorytypes.ItemPropertyData
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewRingObject(pl player.Player) *PlayerRingObject {
	o := &PlayerRingObject{
		player: pl,
	}
	return o
}

func convertPlayerRingObjectToEntity(o *PlayerRingObject) (*ringentity.PlayerRingEntity, error) {
	data, err := json.Marshal(o.propertyData)
	if err != nil {
		return nil, err
	}

	e := &ringentity.PlayerRingEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		BindType:     int32(o.bindType),
		Typ:          int32(o.typ),
		PropertyData: string(data),
		ItemId:       o.itemId,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *PlayerRingObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerRingObject) GetRingType() ringtypes.RingType {
	return o.typ
}

func (o *PlayerRingObject) GetBindType() itemtypes.ItemBindType {
	return o.bindType
}

func (o *PlayerRingObject) GetPropertyData() inventorytypes.ItemPropertyData {
	return o.propertyData
}

func (o *PlayerRingObject) GetItemId() int32 {
	return o.itemId
}

func (o *PlayerRingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerRingObjectToEntity(o)
	return e, err
}

func (o *PlayerRingObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*ringentity.PlayerRingEntity)

	data, err := inventory.CreatePropertyData(itemtypes.ItemTypeTeRing, pse.PropertyData)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.typ = ringtypes.RingType(pse.Typ)
	o.bindType = itemtypes.ItemBindType(pse.BindType)
	o.itemId = pse.ItemId
	o.propertyData = data
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerRingObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerRing"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
