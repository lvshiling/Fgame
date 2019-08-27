package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	qixueentity "fgame/fgame/game/qixue/entity"

	"github.com/pkg/errors"
)

//泣血枪对象
type PlayerQiXueObject struct {
	player     player.Player
	id         int64
	currLevel  int32
	currStar   int32
	timesNum   int32
	lastTime   int64
	shaLuNum   int64
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerQiXueObject(pl player.Player) *PlayerQiXueObject {
	o := &PlayerQiXueObject{
		player: pl,
	}
	return o
}

func convertNewPlayerQiXueObjectToEntity(o *PlayerQiXueObject) (*qixueentity.PlayerQiXueEntity, error) {

	e := &qixueentity.PlayerQiXueEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		CurrLevel:  o.currLevel,
		CurrStar:   o.currStar,
		TimesNum:   o.timesNum,
		LastTime:   o.lastTime,
		ShaLuNum:   o.shaLuNum,
		Power:      o.power,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerQiXueObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerQiXueObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerQiXueObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerQiXueObjectToEntity(o)
	return e, err
}

func (o *PlayerQiXueObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*qixueentity.PlayerQiXueEntity)

	o.id = pse.Id
	o.currLevel = pse.CurrLevel
	o.currStar = pse.CurrStar
	o.timesNum = pse.TimesNum
	o.lastTime = pse.LastTime
	o.shaLuNum = pse.ShaLuNum
	o.power = pse.Power
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerQiXueObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "QiXue"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerQiXueObject) GetShaLuNum() int64 {
	return o.shaLuNum
}

func (o *PlayerQiXueObject) GetLevel() int32 {
	return o.currLevel
}

func (o *PlayerQiXueObject) GetStar() int32 {
	return o.currStar
}

func (o *PlayerQiXueObject) GetTimesNum() int32 {
	return o.timesNum
}

func (o *PlayerQiXueObject) IfEnoughShaLuNum() bool {
	return o.shaLuNum > 0
}
