package shenmo

import (
	"fgame/fgame/core/storage"
	shenmoentity "fgame/fgame/cross/shenmo/entity"
	"fgame/fgame/game/global"
)

//神魔战场时间戳数据
type ShenMoRankTimeObject struct {
	Id         int64
	Platform   int32
	ThisTime   int64
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewShenMoRankTimeObject() *ShenMoRankTimeObject {
	poo := &ShenMoRankTimeObject{}
	return poo
}

func (so *ShenMoRankTimeObject) GetDBId() int64 {
	return so.Id
}

func (oo *ShenMoRankTimeObject) ToEntity() (e storage.Entity, err error) {
	oe := &shenmoentity.ShenMoRankTimeEntity{}
	oe.Id = oo.Id
	oe.Platform = oo.Platform
	oe.ThisTime = oo.ThisTime
	oe.LastTime = oo.LastTime
	oe.UpdateTime = oo.UpdateTime
	oe.CreateTime = oo.CreateTime
	oe.DeleteTime = oo.DeleteTime
	e = oe
	return
}

func (oo *ShenMoRankTimeObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*shenmoentity.ShenMoRankTimeEntity)
	oo.Id = oe.Id
	oo.Platform = oe.Platform
	oo.ThisTime = oe.ThisTime
	oo.LastTime = oe.LastTime
	oo.UpdateTime = oe.UpdateTime
	oo.CreateTime = oe.CreateTime
	oo.DeleteTime = oe.DeleteTime
	return
}

func (oo *ShenMoRankTimeObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
