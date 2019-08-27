package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	ylentity "fgame/fgame/game/yinglingpu/entity"

	ylptypes "fgame/fgame/game/yinglingpu/types"

	"github.com/pkg/errors"
)

type YingLingPuObject struct {
	player     player.Player
	Id         int64
	TuJianId   int32
	TuJianType ylptypes.YingLingPuTuJianType
	Level      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewYingLingPuObject(pl player.Player) *YingLingPuObject {
	rst := &YingLingPuObject{
		player: pl,
	}
	return rst
}

func (o *YingLingPuObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *YingLingPuObject) GetDBId() int64 {
	return o.Id
}

func (o *YingLingPuObject) ToEntity() (e storage.Entity, err error) {
	rst := &ylentity.YingLingPuEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		TuJianId:   o.TuJianId,
		Level:      o.Level,
		TuJianType: int32(o.TuJianType),
		CreateTime: o.CreateTime,
		UpdateTime: o.UpdateTime,
		DeleteTime: o.DeleteTime,
	}
	e = rst
	return e, nil
}

func (o *YingLingPuObject) FromEntity(e storage.Entity) error {
	lye, _ := e.(*ylentity.YingLingPuEntity)
	o.Id = lye.Id
	o.Level = lye.Level
	o.TuJianId = lye.TuJianId
	o.TuJianType = ylptypes.YingLingPuTuJianType(lye.TuJianType)
	o.CreateTime = lye.CreateTime
	o.UpdateTime = lye.UpdateTime
	o.DeleteTime = lye.DeleteTime
	return nil
}

func (o *YingLingPuObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "YingLingPu"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
