package player

import (
	"fgame/fgame/core/storage"
	outlandbossentity "fgame/fgame/game/outlandboss/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//外域BOSS对象
type PlayerOutlandBossObject struct {
	player     player.Player
	id         int64
	zhuoqiNum  int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerOutlandBossObject(pl player.Player) *PlayerOutlandBossObject {
	o := &PlayerOutlandBossObject{
		player: pl,
	}
	return o
}

func convertNewPlayerOutlandBossObjectToEntity(o *PlayerOutlandBossObject) (*outlandbossentity.PlayerOutlandBossEntity, error) {
	e := &outlandbossentity.PlayerOutlandBossEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		ZhuoQiNum:  o.zhuoqiNum,
		UpdateTime: o.updateTime,
		DeleteTime: o.deleteTime,
		CreateTime: o.createTime,
	}
	return e, nil
}

func (o *PlayerOutlandBossObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerOutlandBossObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerOutlandBossObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerOutlandBossObjectToEntity(o)
	return e, err
}

func (o *PlayerOutlandBossObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*outlandbossentity.PlayerOutlandBossEntity)

	o.id = pse.Id
	o.zhuoqiNum = pse.ZhuoQiNum
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerOutlandBossObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OutlandBoss"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
