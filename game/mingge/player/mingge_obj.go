package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/mingge/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//玩家命格补偿对象
type PlayerMingGeObject struct {
	player     player.Player
	id         int64
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerMingGeObject(pl player.Player) *PlayerMingGeObject {
	o := &PlayerMingGeObject{
		player: pl,
	}
	return o
}

func convertMingGeobjectToEntity(o *PlayerMingGeObject) (*entity.PlayerMingGeEntity, error) {

	e := &entity.PlayerMingGeEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Power:      o.power,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerMingGeObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerMingGeObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerMingGeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertMingGeobjectToEntity(o)
	return e, err
}

func (o *PlayerMingGeObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerMingGeEntity)

	o.id = pse.Id
	o.power = pse.Power
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerMingGeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("mingge: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
