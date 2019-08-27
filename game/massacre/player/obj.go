package player

import (
	"fgame/fgame/core/storage"
	massacreentity "fgame/fgame/game/massacre/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//戮仙刃对象
type PlayerMassacreObject struct {
	player     player.Player
	Id         int64
	AdvanceId  int
	CurrLevel  int32
	CurrStar   int32
	TimesNum   int32
	LastTime   int64
	ShaQiNum   int64
	Power      int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerMassacreObject(pl player.Player) *PlayerMassacreObject {
	o := &PlayerMassacreObject{
		player: pl,
	}
	return o
}

func convertNewPlayerMassacreObjectToEntity(o *PlayerMassacreObject) (*massacreentity.PlayerMassacreEntity, error) {

	e := &massacreentity.PlayerMassacreEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		AdvancedId: o.AdvanceId,
		CurrLevel:  o.CurrLevel,
		CurrStar:   o.CurrStar,
		TimesNum:   o.TimesNum,
		LastTime:   o.LastTime,
		ShaQiNum:   o.ShaQiNum,
		Power:      o.Power,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerMassacreObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerMassacreObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerMassacreObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerMassacreObjectToEntity(o)
	return e, err
}

func (o *PlayerMassacreObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*massacreentity.PlayerMassacreEntity)

	o.Id = pse.Id
	o.AdvanceId = pse.AdvancedId
	o.CurrLevel = pse.CurrLevel
	o.CurrStar = pse.CurrStar
	o.TimesNum = pse.TimesNum
	o.LastTime = pse.LastTime
	o.ShaQiNum = pse.ShaQiNum
	o.Power = pse.Power
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerMassacreObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Massacre"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
