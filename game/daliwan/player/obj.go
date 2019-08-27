package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	daliwanentity "fgame/fgame/game/daliwan/entity"

	"github.com/pkg/errors"
)

type DaLiWanObject struct {
	player     player.Player
	id         int64
	typ        int32
	startTime  int64
	duration   int64
	expired    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewDaLiWanObject(pl player.Player) *DaLiWanObject {
	o := &DaLiWanObject{
		player: pl,
	}
	return o
}

func (o *DaLiWanObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *DaLiWanObject) GetDBId() int64 {
	return o.id
}

func (o *DaLiWanObject) ToEntity() (e storage.Entity, err error) {
	rst := &daliwanentity.DaLiWanEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Typ:        o.typ,
		StartTime:  o.startTime,
		Duration:   o.duration,
		Expired:    o.expired,
		CreateTime: o.createTime,
		UpdateTime: o.updateTime,
		DeleteTime: o.deleteTime,
	}
	e = rst
	return e, nil
}

func (o *DaLiWanObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*daliwanentity.DaLiWanEntity)
	o.id = te.Id
	o.typ = te.Typ
	o.startTime = te.StartTime
	o.duration = te.Duration
	o.expired = te.Expired
	o.createTime = te.CreateTime
	o.updateTime = te.UpdateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *DaLiWanObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "daliwan"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *DaLiWanObject) IsExpire(now int64) bool {
	if o.expired != 0 {
		return true
	}
	return (o.startTime + o.duration) <= now
}

func (o *DaLiWanObject) GetTyp() int32 {
	return o.typ
}
