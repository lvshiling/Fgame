package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	tianshuentity "fgame/fgame/game/tianshu/entity"
	tianshutypes "fgame/fgame/game/tianshu/types"

	"github.com/pkg/errors"
)

//天书对象
type PlayerTianShuObject struct {
	player     player.Player
	id         int64
	typ        tianshutypes.TianShuType
	level      int32
	isReceive  int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerTianShuObject(pl player.Player) *PlayerTianShuObject {
	o := &PlayerTianShuObject{
		player: pl,
	}
	return o
}

func convertNewPlayerTianShuObjectToEntity(o *PlayerTianShuObject) (*tianshuentity.PlayerTianShuEntity, error) {

	e := &tianshuentity.PlayerTianShuEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Typ:        int32(o.typ),
		Level:      o.level,
		IsReceive:  o.isReceive,
		UpdateTime: o.updateTime,
		DeleteTime: o.deleteTime,
		CreateTime: o.createTime,
	}
	return e, nil
}

func (o *PlayerTianShuObject) GetPlayerId() int64 {
	return o.player.GetId()
}
func (o *PlayerTianShuObject) GetIsReceive() int32 {
	return o.isReceive
}

func (o *PlayerTianShuObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerTianShuObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerTianShuObject) GetType() tianshutypes.TianShuType {
	return o.typ
}

func (o *PlayerTianShuObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerTianShuObjectToEntity(o)
	return e, err
}

func (o *PlayerTianShuObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*tianshuentity.PlayerTianShuEntity)

	o.id = pse.Id
	o.typ = tianshutypes.TianShuType(pse.Typ)
	o.level = pse.Level
	o.isReceive = pse.IsReceive
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerTianShuObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TianShu"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
