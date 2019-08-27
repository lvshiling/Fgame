package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	towerentity "fgame/fgame/game/tower/entity"

	"github.com/pkg/errors"
)

//打宝塔对象
type PlayerTowerObject struct {
	player        player.Player
	id            int64
	useTime       int64
	extraTime     int64
	lastResetTime int64
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewPlayerTowerObject(pl player.Player) *PlayerTowerObject {
	o := &PlayerTowerObject{
		player: pl,
	}
	return o
}

func convertNewPlayerTowerObjectToEntity(o *PlayerTowerObject) (*towerentity.PlayerTowerEntity, error) {

	e := &towerentity.PlayerTowerEntity{
		Id:            o.id,
		PlayerId:      o.player.GetId(),
		UseTime:       o.useTime,
		ExtralTime:    o.extraTime,
		LastResetTime: o.lastResetTime,
		UpdateTime:    o.updateTime,
		DeleteTime:    o.deleteTime,
		CreateTime:    o.createTime,
	}
	return e, nil
}

func (o *PlayerTowerObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerTowerObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerTowerObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerTowerObjectToEntity(o)
	return e, err
}

func (o *PlayerTowerObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*towerentity.PlayerTowerEntity)

	o.id = pse.Id
	o.useTime = pse.UseTime
	o.extraTime = pse.ExtralTime
	o.lastResetTime = pse.LastResetTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerTowerObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Tower"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
