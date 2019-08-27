package player

import (
	"fgame/fgame/core/storage"
	dianxingentity "fgame/fgame/game/dianxing/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//点星系统对象
type PlayerDianXingObject struct {
	player            player.Player
	Id                int64
	CurrType          int32
	CurrLevel         int32
	DianXingTimes     int32
	DianXingBless     int32
	DianXingBlessTime int64
	XingChenNum       int64
	JieFengLev        int32
	JieFengTimes      int32
	JieFengBless      int32
	Power             int64
	UpdateTime        int64
	CreateTime        int64
	DeleteTime        int64
}

func NewPlayerDianXingObject(pl player.Player) *PlayerDianXingObject {
	o := &PlayerDianXingObject{
		player: pl,
	}
	return o
}

func convertNewPlayerDianXingObjectToEntity(o *PlayerDianXingObject) (*dianxingentity.PlayerDianXingEntity, error) {

	e := &dianxingentity.PlayerDianXingEntity{
		Id:                o.Id,
		PlayerId:          o.player.GetId(),
		CurrType:          o.CurrType,
		CurrLevel:         o.CurrLevel,
		DianXingTimes:     o.DianXingTimes,
		DianXingBless:     o.DianXingBless,
		DianXingBlessTime: o.DianXingBlessTime,
		XingChenNum:       o.XingChenNum,
		JieFengLev:        o.JieFengLev,
		JieFengTimes:      o.JieFengTimes,
		JieFengBless:      o.JieFengBless,
		Power:             o.Power,
		UpdateTime:        o.UpdateTime,
		CreateTime:        o.CreateTime,
		DeleteTime:        o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerDianXingObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerDianXingObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerDianXingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerDianXingObjectToEntity(o)
	return e, err
}

func (o *PlayerDianXingObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*dianxingentity.PlayerDianXingEntity)

	o.Id = pse.Id
	o.CurrType = pse.CurrType
	o.CurrLevel = pse.CurrLevel
	o.DianXingTimes = pse.DianXingTimes
	o.DianXingBless = pse.DianXingBless
	o.DianXingBlessTime = pse.DianXingBlessTime
	o.XingChenNum = pse.XingChenNum
	o.JieFengLev = pse.JieFengLev
	o.JieFengTimes = pse.JieFengTimes
	o.JieFengBless = pse.JieFengBless
	o.Power = pse.Power
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerDianXingObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "DianXing"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
