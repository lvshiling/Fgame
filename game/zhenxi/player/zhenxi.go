package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/zhenxi/entity"
	"fmt"
)

//玩家珍惜对象
type PlayerZhenXiBossObject struct {
	player     player.Player
	id         int64
	reliveTime int32
	enterTimes int32 //进入次数
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerZhenXiBossObject(pl player.Player) *PlayerZhenXiBossObject {
	pmo := &PlayerZhenXiBossObject{
		player: pl,
	}
	return pmo
}

func convertZhenXiBossObjectToEntity(o *PlayerZhenXiBossObject) (*entity.PlayerZhenXiBossEntity, error) {

	e := &entity.PlayerZhenXiBossEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		ReliveTime: o.reliveTime,
		EnterTimes: o.enterTimes,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerZhenXiBossObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerZhenXiBossObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerZhenXiBossObject) GetReliveTime() int32 {
	return o.reliveTime
}

func (o *PlayerZhenXiBossObject) GetEnterTimes() int32 {
	return o.enterTimes
}

func (o *PlayerZhenXiBossObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertZhenXiBossObjectToEntity(o)
	return e, err
}

func (o *PlayerZhenXiBossObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerZhenXiBossEntity)
	o.id = pse.Id
	o.reliveTime = pse.ReliveTime
	o.enterTimes = pse.EnterTimes
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerZhenXiBossObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("zhenxi: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
