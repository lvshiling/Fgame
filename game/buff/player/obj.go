package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	buffentity "fgame/fgame/game/buff/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//buff对象
type PlayerBuffObject struct {
	player     player.Player
	id         int64
	buffMap    map[int32]*buffObject
	updateTime int64
	createTime int64
	deleteTime int64
}

func (o *PlayerBuffObject) GetPlayer() player.Player {
	return o.player
}

func (o *PlayerBuffObject) GetId() int64 {
	return o.id
}

func (o *PlayerBuffObject) GetBuffMap() map[int32]*buffObject {
	return o.buffMap
}

func (o *PlayerBuffObject) GetBuff(groupId int32) *buffObject {
	b, ok := o.buffMap[groupId]
	if !ok {
		return nil
	}
	return b
}

func (o *PlayerBuffObject) AddBuff(bo *buffObject) {
	o.buffMap[bo.GroupId] = bo
}

func (o *PlayerBuffObject) RemoveBuff(groupId int32) {
	delete(o.buffMap, groupId)
}

func (o *PlayerBuffObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *PlayerBuffObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *PlayerBuffObject) GetDeleteTime() int64 {
	return o.deleteTime
}

func newPlayerBuffObject(pl player.Player) *PlayerBuffObject {
	o := &PlayerBuffObject{
		player: pl,
	}
	return o
}

func convertPlayerBuffObjectToEntity(o *PlayerBuffObject) (e *buffentity.PlayerBuffEntity, err error) {
	buffBytes, err := json.Marshal(o.GetBuffMap())
	if err != nil {
		return
	}
	e = &buffentity.PlayerBuffEntity{
		Id:         o.GetId(),
		PlayerId:   o.GetPlayerId(),
		BuffMap:    string(buffBytes),
		UpdateTime: o.GetUpdateTime(),
		CreateTime: o.GetCreateTime(),
		DeleteTime: o.GetDeleteTime(),
	}
	return e, nil
}

func (o *PlayerBuffObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerBuffObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerBuffObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerBuffObjectToEntity(o)
	return e, err
}

func (o *PlayerBuffObject) FromEntity(e storage.Entity) error {
	pe, _ := e.(*buffentity.PlayerBuffEntity)
	o.id = pe.Id
	o.updateTime = pe.UpdateTime
	o.createTime = pe.CreateTime
	o.deleteTime = pe.DeleteTime
	buffMap := make(map[int32]*buffObject)
	err := json.Unmarshal([]byte(pe.BuffMap), &buffMap)
	if err != nil {
		return err
	}
	o.buffMap = buffMap

	return nil
}

func (o *PlayerBuffObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Buff"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
