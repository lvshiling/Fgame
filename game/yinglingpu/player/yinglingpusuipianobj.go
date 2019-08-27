package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	ylentity "fgame/fgame/game/yinglingpu/entity"
	ylptypes "fgame/fgame/game/yinglingpu/types"

	"github.com/pkg/errors"
)

type YingLingPuSuiPianObject struct {
	player     player.Player
	Id         int64
	TuJianId   int32
	TuJianType ylptypes.YingLingPuTuJianType
	SuiPianId  int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewYingLingPuSuiPianObject(pl player.Player) *YingLingPuSuiPianObject {
	rst := &YingLingPuSuiPianObject{
		player: pl,
	}
	return rst
}

func (o *YingLingPuSuiPianObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *YingLingPuSuiPianObject) GetDBId() int64 {
	return o.Id
}

func (o *YingLingPuSuiPianObject) ToEntity() (e storage.Entity, err error) {
	rst := &ylentity.YingLingPuSuiPianEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		TuJianId:   o.TuJianId,
		SuiPianId:  o.SuiPianId,
		TuJianType: int32(o.TuJianType),
		CreateTime: o.CreateTime,
		UpdateTime: o.UpdateTime,
		DeleteTime: o.DeleteTime,
	}
	e = rst
	return e, nil
}

func (o *YingLingPuSuiPianObject) FromEntity(e storage.Entity) error {
	lye, _ := e.(*ylentity.YingLingPuSuiPianEntity)
	o.Id = lye.Id
	o.SuiPianId = lye.SuiPianId
	o.TuJianId = lye.TuJianId
	o.TuJianType = ylptypes.YingLingPuTuJianType(lye.TuJianType)
	o.CreateTime = lye.CreateTime
	o.UpdateTime = lye.UpdateTime
	o.DeleteTime = lye.DeleteTime
	return nil
}

func (o *YingLingPuSuiPianObject) SetModified() {
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
