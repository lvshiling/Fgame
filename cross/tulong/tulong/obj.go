package tulong

import (
	"fgame/fgame/core/storage"
	tulongentity "fgame/fgame/cross/tulong/entity"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/idutil"
)

//跨服屠龙排行榜数据
type TuLongRankObject struct {
	Id           int64
	Platform     int32
	AreaId       int32
	ServerId     int32
	AllianceId   int64
	AllianceName string
	KillNum      int32
	LastTime     int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

//记录排序
type TuLongRankObjectList []*TuLongRankObject

func (adl TuLongRankObjectList) Len() int {
	return len(adl)
}

func (adl TuLongRankObjectList) Less(i, j int) bool {
	if adl[i].KillNum == adl[j].KillNum {
		return adl[i].LastTime < adl[j].LastTime
	}
	return adl[i].KillNum < adl[j].KillNum
}

func (adl TuLongRankObjectList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

func NewTuLongRankObject() *TuLongRankObject {
	poo := &TuLongRankObject{}
	return poo
}

func initNewTuLongRankObject(playerServerId int32, allianceId int64, allianceName string) *TuLongRankObject {
	now := global.GetGame().GetTimeService().Now()
	poo := NewTuLongRankObject()
	id, _ := idutil.GetId()
	poo.Id = id
	poo.Platform = global.GetGame().GetPlatform()
	poo.AreaId = global.GetGame().GetServerIndex()
	poo.ServerId = playerServerId
	poo.AllianceId = allianceId
	poo.AllianceName = allianceName
	poo.KillNum = 1
	poo.LastTime = now
	poo.CreateTime = now
	poo.SetModified()
	return poo
}

func (tlro *TuLongRankObject) GetDBId() int64 {
	return tlro.Id
}

func (oo *TuLongRankObject) ToEntity() (e storage.Entity, err error) {
	oe := &tulongentity.TuLongRankEntity{}
	oe.Id = oo.Id
	oe.Platform = oo.Platform
	oe.AreaId = oo.AreaId
	oe.ServerId = oo.ServerId
	oe.AllianceId = oo.AllianceId
	oe.AllianceName = oo.AllianceName
	oe.KillNum = oo.KillNum
	oe.LastTime = oo.LastTime
	oe.UpdateTime = oo.UpdateTime
	oe.CreateTime = oo.CreateTime
	oe.DeleteTime = oo.DeleteTime
	e = oe
	return
}

func (oo *TuLongRankObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*tulongentity.TuLongRankEntity)
	oo.Id = oe.Id
	oe.Platform = oe.Platform
	oo.AreaId = oe.AreaId
	oo.ServerId = oe.ServerId
	oo.AllianceId = oe.AllianceId
	oo.AllianceName = oe.AllianceName
	oo.KillNum = oe.KillNum
	oo.LastTime = oe.LastTime
	oo.UpdateTime = oe.UpdateTime
	oo.CreateTime = oe.CreateTime
	oo.DeleteTime = oe.DeleteTime
	return
}

func (oo *TuLongRankObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
