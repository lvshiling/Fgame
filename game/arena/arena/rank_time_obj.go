package arena

import (
	"fgame/fgame/core/storage"
	arenaentity "fgame/fgame/game/arena/entity"
	"fgame/fgame/game/global"
)

//神魔战场时间戳数据
type ArenaRankTimeObject struct {
	Id         int64
	ServerId   int32
	ThisTime   int64
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewArenaRankTimeObject() *ArenaRankTimeObject {
	po := &ArenaRankTimeObject{}
	return po
}

func (so *ArenaRankTimeObject) GetDBId() int64 {
	return so.Id
}

func (o *ArenaRankTimeObject) ToEntity() (e storage.Entity, err error) {
	oe := &arenaentity.ArenaRankTimeEntity{}
	oe.Id = o.Id
	oe.ServerId = o.ServerId
	oe.ThisTime = o.ThisTime
	oe.LastTime = o.LastTime
	oe.UpdateTime = o.UpdateTime
	oe.CreateTime = o.CreateTime
	oe.DeleteTime = o.DeleteTime
	e = oe
	return
}

func (o *ArenaRankTimeObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*arenaentity.ArenaRankTimeEntity)
	o.Id = oe.Id
	o.ServerId = oe.ServerId
	o.ThisTime = oe.ThisTime
	o.LastTime = oe.LastTime
	o.UpdateTime = oe.UpdateTime
	o.CreateTime = oe.CreateTime
	o.DeleteTime = oe.DeleteTime
	return
}

func (o *ArenaRankTimeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
