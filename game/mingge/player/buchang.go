package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/mingge/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//玩家命格补偿对象
type PlayerMingGeBuchangObject struct {
	player     player.Player
	id         int64
	buchang    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerMingGeBuchangObject(pl player.Player) *PlayerMingGeBuchangObject {
	o := &PlayerMingGeBuchangObject{
		player: pl,
	}
	return o
}

func convertMingGeBuchangbjectToEntity(o *PlayerMingGeBuchangObject) (*entity.PlayerMingGeBuchangEntity, error) {

	e := &entity.PlayerMingGeBuchangEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Buchang:    o.buchang,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerMingGeBuchangObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerMingGeBuchangObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerMingGeBuchangObject) IsBuchang() bool {
	return o.buchang != 0
}

func (o *PlayerMingGeBuchangObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertMingGeBuchangbjectToEntity(o)
	return e, err
}

func (o *PlayerMingGeBuchangObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerMingGeBuchangEntity)

	o.id = pse.Id
	o.buchang = pse.Buchang
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerMingGeBuchangObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("minggebuchang: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
