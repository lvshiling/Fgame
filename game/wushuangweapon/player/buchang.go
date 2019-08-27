package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wushuangweaponentity "fgame/fgame/game/wushuangweapon/entity"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

type PlayerWushuangBuchangObject struct {
	player      player.Player
	id          int64
	isSendEmail int32
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func (o *PlayerWushuangBuchangObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerWushuangBuchangObject) ToEntity() (e storage.Entity, err error) {
	e = &wushuangweaponentity.PlayerWushuangBuchangEntity{
		Id:          o.id,
		PlayerId:    o.player.GetId(),
		IsSendEmail: o.isSendEmail,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *PlayerWushuangBuchangObject) FromEntity(e storage.Entity) (err error) {
	te, _ := e.(*wushuangweaponentity.PlayerWushuangBuchangEntity)
	o.id = te.Id
	o.isSendEmail = te.IsSendEmail
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func NewPlayerWushuangBuchangObject(pl player.Player) *PlayerWushuangBuchangObject {
	obj := &PlayerWushuangBuchangObject{
		player: pl,
	}
	return obj
}

func createPlayerWushuangBuchangObject(pl player.Player) *PlayerWushuangBuchangObject {
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj := &PlayerWushuangBuchangObject{
		player:      pl,
		id:          id,
		isSendEmail: int32(0),
		createTime:  now,
	}
	return obj
}

func (o *PlayerWushuangBuchangObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "WushuangWeapon"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
