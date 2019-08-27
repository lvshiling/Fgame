package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	transportationentity "fgame/fgame/game/transportation/entity"

	"github.com/pkg/errors"
)

//玩家镖车对象
type PlayerTransportationObject struct {
	id                     int64
	player                 player.Player
	robList                []int64
	personalTransportTimes int32
	updateTime             int64
	createTime             int64
	deleteTime             int64
}

func newPlayerTransportationObject(pl player.Player) *PlayerTransportationObject {
	o := &PlayerTransportationObject{
		player: pl,
	}
	return o
}

func convertPlayerTransportationObjectToEntity(o *PlayerTransportationObject) (e *transportationentity.PlayerTransportationEntity, err error) {
	robList, err := json.Marshal(o.robList)
	if err != nil {
		return
	}

	e = &transportationentity.PlayerTransportationEntity{
		Id:                     o.id,
		PlayerId:               o.player.GetId(),
		RobList:                string(robList),
		PersonalTransportTimes: o.personalTransportTimes,
		UpdateTime:             o.updateTime,
		CreateTime:             o.createTime,
		DeleteTime:             o.deleteTime,
	}
	return e, nil
}

func (o *PlayerTransportationObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerTransportationObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerTransportationObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerTransportationObjectToEntity(o)
	return e, err
}

func (o *PlayerTransportationObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*transportationentity.PlayerTransportationEntity)

	var robList []int64
	json.Unmarshal([]byte(te.RobList), &robList)

	o.id = te.Id
	o.robList = robList
	o.personalTransportTimes = te.PersonalTransportTimes
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerTransportationObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Transporation"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
