package player

import (
	"fgame/fgame/core/storage"
	baguaentity "fgame/fgame/game/bagua/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//八卦对象
type PlayerBaGuaObject struct {
	player     player.Player
	id         int64
	playerId   int64
	level      int32
	isBuChang  int32
	inviteTime int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerBaGuaObject(pl player.Player) *PlayerBaGuaObject {
	o := &PlayerBaGuaObject{
		player: pl,
	}
	return o
}

func (o *PlayerBaGuaObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *PlayerBaGuaObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerBaGuaObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerBaGuaObject) GetInviteTime() int64 {
	return o.inviteTime
}

func (o *PlayerBaGuaObject) ToEntity() (e storage.Entity, err error) {
	e = &baguaentity.PlayerBaGuaEntity{
		Id:         o.id,
		PlayerId:   o.playerId,
		Level:      o.level,
		IsBuChang:  o.isBuChang,
		InviteTime: o.inviteTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerBaGuaObject) FromEntity(e storage.Entity) error {
	pe, _ := e.(*baguaentity.PlayerBaGuaEntity)

	o.id = pe.Id
	o.playerId = pe.PlayerId
	o.level = pe.Level
	o.isBuChang = pe.IsBuChang
	o.inviteTime = pe.InviteTime
	o.updateTime = pe.UpdateTime
	o.createTime = pe.CreateTime
	o.deleteTime = pe.DeleteTime
	return nil
}

func (o *PlayerBaGuaObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "bagua"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
