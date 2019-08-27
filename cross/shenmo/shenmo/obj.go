package shenmo

import (
	"fgame/fgame/core/storage"
	shenmoentity "fgame/fgame/cross/shenmo/entity"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/idutil"
)

//神魔战场排行榜数据
type ShenMoRankObject struct {
	Id           int64
	Platform     int32
	ServerId     int32
	AllianceId   int64
	AllianceName string
	JiFenNum     int32
	LastJiFenNum int32
	LastTime     int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

//本周记录排序
type ThisShenMoRankObjectList []*ShenMoRankObject

func (adl ThisShenMoRankObjectList) Len() int {
	return len(adl)
}

func (adl ThisShenMoRankObjectList) Less(i, j int) bool {
	if adl[i].JiFenNum == adl[j].JiFenNum {
		return adl[i].LastTime < adl[j].LastTime
	}
	return adl[i].JiFenNum < adl[j].JiFenNum
}

func (adl ThisShenMoRankObjectList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

//上周记录排序
type LastShenMoRankObjectList []*ShenMoRankObject

func (adl LastShenMoRankObjectList) Len() int {
	return len(adl)
}

func (adl LastShenMoRankObjectList) Less(i, j int) bool {
	return adl[i].LastJiFenNum < adl[j].LastJiFenNum
}

func (adl LastShenMoRankObjectList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

func initShenMoRankObject(serverId int32,
	allianceId int64,
	allianceName string,
	jiFenNum int32) *ShenMoRankObject {

	platform := global.GetGame().GetPlatform()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()

	o := &ShenMoRankObject{}
	o.Id = id
	o.Platform = platform
	o.ServerId = serverId
	o.AllianceId = allianceId
	o.AllianceName = allianceName
	o.JiFenNum = jiFenNum
	o.LastTime = now
	o.LastJiFenNum = 0
	o.CreateTime = now
	return o
}

func NewShenMoRankObject() *ShenMoRankObject {
	return &ShenMoRankObject{}
}

func (so *ShenMoRankObject) GetDBId() int64 {
	return so.Id
}

func (oo *ShenMoRankObject) ToEntity() (e storage.Entity, err error) {
	oe := &shenmoentity.ShenMoRankEntity{}
	oe.Id = oo.Id
	oe.Platform = oo.Platform
	oe.ServerId = oo.ServerId
	oe.AllianceId = oo.AllianceId
	oe.AllianceName = oo.AllianceName
	oe.JiFenNum = oo.JiFenNum
	oe.LastJiFenNum = oo.LastJiFenNum
	oe.LastTime = oo.LastTime
	oe.UpdateTime = oo.UpdateTime
	oe.CreateTime = oo.CreateTime
	oe.DeleteTime = oo.DeleteTime
	e = oe
	return
}

func (oo *ShenMoRankObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*shenmoentity.ShenMoRankEntity)
	oo.Id = oe.Id
	oo.Platform = oe.Platform
	oo.ServerId = oe.ServerId
	oo.AllianceId = oe.AllianceId
	oo.AllianceName = oe.AllianceName
	oo.JiFenNum = oe.JiFenNum
	oo.LastJiFenNum = oe.LastJiFenNum
	oo.LastTime = oe.LastTime
	oo.UpdateTime = oe.UpdateTime
	oo.CreateTime = oe.CreateTime
	oo.DeleteTime = oe.DeleteTime
	return
}

func (oo *ShenMoRankObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
