package player

import (
	"fgame/fgame/core/storage"
	itemtypes "fgame/fgame/game/item/types"
	luckyentity "fgame/fgame/game/lucky/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//合成
type PlayerLuckyObject struct {
	player     player.Player
	id         int64
	typ        itemtypes.ItemType
	subType    itemtypes.ItemSubType
	itemId     int32
	expireTime int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerLuckyObject(pl player.Player) *PlayerLuckyObject {
	o := &PlayerLuckyObject{
		player: pl,
	}
	return o
}

func convertNewPlayerLuckyObjectToEntity(o *PlayerLuckyObject) (*luckyentity.PlayerLuckyEntity, error) {
	e := &luckyentity.PlayerLuckyEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Typ:        int32(o.typ),
		SubType:    o.subType.SubType(),
		ItemId:     o.itemId,
		ExpireTime: o.expireTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerLuckyObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerLuckyObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerLuckyObject) GetExpireTime() int64 {
	return o.expireTime
}

func (o *PlayerLuckyObject) GetItemId() int32 {
	return o.itemId
}

func (o *PlayerLuckyObject) GetType() itemtypes.ItemType {
	return o.typ
}

func (o *PlayerLuckyObject) GetSubType() itemtypes.ItemSubType {
	return o.subType
}

func (o *PlayerLuckyObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerLuckyObjectToEntity(o)
	return e, err
}

func (o *PlayerLuckyObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*luckyentity.PlayerLuckyEntity)

	o.id = pse.Id
	o.typ = itemtypes.ItemType(pse.Typ)
	o.subType = itemtypes.CreateItemSubType(o.typ, pse.SubType)
	o.expireTime = pse.ExpireTime
	o.itemId = pse.ItemId
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerLuckyObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("ShenFa: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
