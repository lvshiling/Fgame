package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	yuxientity "fgame/fgame/game/yuxi/entity"

	"github.com/pkg/errors"
)

//玩家仙盟玉玺对象
type PlayerAllianceYuXiObject struct {
	player     player.Player
	id         int64
	isReceive  int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerAllianceYuXiObject(pl player.Player) *PlayerAllianceYuXiObject {
	o := &PlayerAllianceYuXiObject{
		player: pl,
	}
	return o
}

func convertPlayerAllianceYuXiObjectToEntity(o *PlayerAllianceYuXiObject) (e *yuxientity.PlayerAlliancYuXiEntity, err error) {
	e = &yuxientity.PlayerAlliancYuXiEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		IsReceive:  o.isReceive,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerAllianceYuXiObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerAllianceYuXiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerAllianceYuXiObjectToEntity(o)
	return e, err
}

func (o *PlayerAllianceYuXiObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*yuxientity.PlayerAlliancYuXiEntity)

	o.id = te.Id
	o.isReceive = te.IsReceive
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerAllianceYuXiObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerAllianceYuXi"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerAllianceYuXiObject) GetIsReceive() int32 {
	return o.isReceive
}

func (o *PlayerAllianceYuXiObject) IsReceive() bool {
	return o.isReceive != 0
}
