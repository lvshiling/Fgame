package player

import (
	"fgame/fgame/core/storage"
	additionsysentity "fgame/fgame/game/additionsys/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

// 玩家圣痕对象
type PlayerShengHenObject struct {
	player     player.Player
	id         int64
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerShengHenObject(pl player.Player) *PlayerShengHenObject {
	o := &PlayerShengHenObject{
		player: pl,
	}
	return o
}

func createPlayerShengHenObject(p player.Player, now int64) *PlayerShengHenObject {
	obj := NewPlayerShengHenObject(p)
	id, _ := idutil.GetId()
	obj.id = id
	obj.createTime = now
	return obj
}

func convertNewPlayerShengHenObjectToEntity(o *PlayerShengHenObject) (*additionsysentity.PlayerShengHenEntity, error) {

	e := &additionsysentity.PlayerShengHenEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Power:      o.power,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerShengHenObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerShengHenObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerShengHenObject) GetPower() int64 {
	return o.power
}

func (o *PlayerShengHenObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerShengHenObjectToEntity(o)
	return e, err
}

func (o *PlayerShengHenObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*additionsysentity.PlayerShengHenEntity)

	o.id = pse.Id
	o.power = pse.Power
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerShengHenObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerShengHe"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
