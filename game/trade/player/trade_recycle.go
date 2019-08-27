package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	tradeentity "fgame/fgame/game/trade/entity"

	"github.com/pkg/errors"
)

type PlayerTradeRecycleObject struct {
	player      player.Player
	id          int64
	recycleGold int64
	recycleTime int64

	updateTime int64
	createTime int64
	deleteTime int64
}

func createPlayerTradeRecycleObject(p player.Player) *PlayerTradeRecycleObject {
	o := &PlayerTradeRecycleObject{}
	o.player = p
	return o
}

func convertPlayerTradeRecyleToEntity(o *PlayerTradeRecycleObject) (*tradeentity.PlayerTradeRecycleEntity, error) {
	e := &tradeentity.PlayerTradeRecycleEntity{
		Id:          o.id,
		PlayerId:    o.player.GetId(),
		RecycleGold: o.recycleGold,
		RecycleTime: o.recycleTime,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *PlayerTradeRecycleObject) GetId() int64 {
	return o.id
}

func (o *PlayerTradeRecycleObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerTradeRecycleObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerTradeRecycleObject) GetRecycleGold() int64 {
	return o.recycleGold
}

func (o *PlayerTradeRecycleObject) GetRecycleTime() int64 {
	return o.recycleTime
}

func (o *PlayerTradeRecycleObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerTradeRecyleToEntity(o)
	return e, err
}

func (o *PlayerTradeRecycleObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*tradeentity.PlayerTradeRecycleEntity)
	o.id = ae.Id
	o.recycleGold = ae.RecycleGold
	o.recycleTime = ae.RecycleTime
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *PlayerTradeRecycleObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerTradeRecycle"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}
	o.player.AddChangedObject(obj)
	return
}
