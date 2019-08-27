package player

import (
	"fgame/fgame/core/storage"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	majorentity "fgame/fgame/game/major/entity"

	"github.com/pkg/errors"
)

//玩家双休数对象
type PlayerMajorNumObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	MajorType  majortypes.MajorType
	Num        int32
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerMajorNumObject(pl player.Player) *PlayerMajorNumObject {
	o := &PlayerMajorNumObject{
		player: pl,
	}
	return o
}

func (o *PlayerMajorNumObject) GetPlayerId() int64 {
	return o.PlayerId
}

func (o *PlayerMajorNumObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerMajorNumObject) ToEntity() (e storage.Entity, err error) {
	e = &majorentity.PlayerMajorNumEntity{
		Id:         o.Id,
		PlayerId:   o.PlayerId,
		MajorType:  int32(o.MajorType),
		Num:        o.Num,
		LastTime:   o.LastTime,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerMajorNumObject) FromEntity(e storage.Entity) error {
	pe, _ := e.(*majorentity.PlayerMajorNumEntity)

	o.Id = pe.Id
	o.PlayerId = pe.PlayerId
	o.Num = pe.Num
	o.LastTime = pe.LastTime
	o.MajorType = majortypes.MajorType(pe.MajorType)
	o.UpdateTime = pe.UpdateTime
	o.CreateTime = pe.CreateTime
	o.DeleteTime = pe.DeleteTime
	return nil
}

func (o *PlayerMajorNumObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Major"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
