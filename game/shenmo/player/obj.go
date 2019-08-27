package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shenmoentity "fgame/fgame/game/shenmo/entity"
	"fmt"
)

//神魔对象
type PlayerShenMoObject struct {
	player     player.Player
	id         int64
	gongXunNum int32
	killNum    int32
	endTime    int64
	rewTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerShenMoObject(pl player.Player) *PlayerShenMoObject {
	o := &PlayerShenMoObject{
		player: pl,
	}
	return o
}

func convertNewPlayerShenMoObjectToEntity(o *PlayerShenMoObject) (*shenmoentity.PlayerShenMoEntity, error) {

	e := &shenmoentity.PlayerShenMoEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		GongXunNum: o.gongXunNum,
		KillNum:    o.killNum,
		EndTime:    o.endTime,
		RewTime:    o.rewTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerShenMoObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerShenMoObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerShenMoObject) GetGongXunNum() int32 {
	return o.gongXunNum
}

func (o *PlayerShenMoObject) GetKillNum() int32 {
	return o.killNum
}

func (o *PlayerShenMoObject) GetEndTime() int64 {
	return o.endTime
}

func (o *PlayerShenMoObject) GetRewTime() int64 {
	return o.rewTime
}

func (o *PlayerShenMoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerShenMoObjectToEntity(o)
	return e, err
}

func (o *PlayerShenMoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*shenmoentity.PlayerShenMoEntity)

	o.id = pse.Id
	o.gongXunNum = pse.GongXunNum
	o.killNum = pse.KillNum
	o.endTime = pse.EndTime
	o.rewTime = pse.RewTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerShenMoObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("Shenmo: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
