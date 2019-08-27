package marry

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	marryentity "fgame/fgame/game/marry/entity"
)

//喜帖列表排序
type MarryWedCardList []*MarryWedCardObject

func (tdl MarryWedCardList) Len() int {
	return len(tdl)
}

func (tdl MarryWedCardList) Less(i, j int) bool {
	return tdl[i].CreateTime < tdl[j].CreateTime
}

func (tdl MarryWedCardList) Swap(i, j int) {
	tdl[i], tdl[j] = tdl[j], tdl[i]
}

//喜帖数据
type MarryWedCardObject struct {
	Id         int64
	ServerId   int32
	PlayerId   int64
	SpouseId   int64
	PlayerName string
	SpouseName string
	HoldTime   string
	OutOfTime  int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewMarryWedCardObject() *MarryWedCardObject {
	pso := &MarryWedCardObject{}
	return pso
}

func (mwco *MarryWedCardObject) GetDBId() int64 {
	return mwco.Id
}

func (mwco *MarryWedCardObject) ToEntity() (e storage.Entity, err error) {
	pe := &marryentity.MarryWedCardEntity{}
	pe.Id = mwco.Id
	pe.ServerId = mwco.ServerId
	pe.PlayerId = mwco.PlayerId
	pe.SpouseId = mwco.SpouseId
	pe.PlayerName = mwco.PlayerName
	pe.SpouseName = mwco.SpouseName
	pe.HoldTime = mwco.HoldTime
	pe.OutOfTime = mwco.OutOfTime
	pe.UpdateTime = mwco.UpdateTime
	pe.CreateTime = mwco.CreateTime
	pe.DeleteTime = mwco.DeleteTime
	e = pe
	return
}

func (mwco *MarryWedCardObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*marryentity.MarryWedCardEntity)
	mwco.Id = pe.Id
	mwco.ServerId = pe.ServerId
	mwco.PlayerId = pe.PlayerId
	mwco.SpouseId = pe.SpouseId
	mwco.PlayerName = pe.PlayerName
	mwco.SpouseName = pe.SpouseName
	mwco.HoldTime = pe.HoldTime
	mwco.OutOfTime = pe.OutOfTime
	mwco.UpdateTime = pe.UpdateTime
	mwco.CreateTime = pe.CreateTime
	mwco.DeleteTime = pe.DeleteTime
	return
}

func (mwco *MarryWedCardObject) SetModified() {
	e, err := mwco.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
