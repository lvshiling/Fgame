package player

import (
	"fgame/fgame/core/storage"
	compensateentity "fgame/fgame/game/compensate/entity"
	compensatetypes "fgame/fgame/game/compensate/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//玩家补偿对象
type PlayerCompensateObject struct {
	player       player.Player
	id           int64
	compensateId int64
	state        compensatetypes.CompensateRecordSate
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerCompensateObject(pl player.Player) *PlayerCompensateObject {
	o := &PlayerCompensateObject{
		player: pl,
	}
	return o
}

func convertNewPlayerCompensateObjectToEntity(o *PlayerCompensateObject) (*compensateentity.PlayerCompensateEntity, error) {
	e := &compensateentity.PlayerCompensateEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		CompensateId: o.compensateId,
		State:        int32(o.state),
		UpdateTime:   o.updateTime,
		DeleteTime:   o.deleteTime,
		CreateTime:   o.createTime,
	}
	return e, nil
}

func (o *PlayerCompensateObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerCompensateObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerCompensateObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerCompensateObjectToEntity(o)
	return e, err
}

func (o *PlayerCompensateObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*compensateentity.PlayerCompensateEntity)

	o.id = pse.Id
	o.compensateId = pse.CompensateId
	o.state = compensatetypes.CompensateRecordSate(pse.State)
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerCompensateObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerCompensate"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
