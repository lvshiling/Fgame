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

//用来记录无双神器的历史最高等级
type PlayerWushuangSettingsObject struct {
	player     player.Player
	id         int64
	itemId     int32
	level      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func (o *PlayerWushuangSettingsObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerWushuangSettingsObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerWushuangSettingsObject) ToEntity() (e storage.Entity, err error) {
	e = &wushuangweaponentity.PlayerWushuangSettingsEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		ItemId:     o.itemId,
		Level:      o.level,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerWushuangSettingsObject) FromEntity(e storage.Entity) (err error) {
	te, _ := e.(*wushuangweaponentity.PlayerWushuangSettingsEntity)
	o.id = te.Id
	o.itemId = te.ItemId
	o.level = te.Level
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func NewPlayerWushuangSettingsObject(pl player.Player) *PlayerWushuangSettingsObject {
	obj := &PlayerWushuangSettingsObject{
		player: pl,
	}
	return obj
}

func createPlayerWushuangSettingsObject(pl player.Player, itemId int32, level int32) *PlayerWushuangSettingsObject {
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj := &PlayerWushuangSettingsObject{
		player:     pl,
		id:         id,
		itemId:     itemId,
		level:      level,
		createTime: now,
	}
	return obj
}

func (o *PlayerWushuangSettingsObject) SetModified() {
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
