package marry

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	marryentity "fgame/fgame/game/marry/entity"
	marrytypes "fgame/fgame/game/marry/types"
)

//婚期数据
type MarryWedObject struct {
	Id          int64
	ServerId    int32
	Period      int32
	Grade       marrytypes.MarryBanquetSubTypeWed
	HunCheGrade marrytypes.MarryBanquetSubTypeHunChe
	SugarGrade  marrytypes.MarryBanquetSubTypeSugar
	Status      marrytypes.MarryWedStatusType
	PlayerId    int64
	SpouseId    int64
	Name        string
	SpouseName  string
	HTime       int64
	LastTime    int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
	IsFirst     bool
}

func NewMarryWedObject() *MarryWedObject {
	pso := &MarryWedObject{}
	return pso
}

func (mwo *MarryWedObject) GetDBId() int64 {
	return mwo.Id
}

func (mwo *MarryWedObject) ToEntity() (e storage.Entity, err error) {
	pe := &marryentity.MarryWedEntity{}
	pe.Id = mwo.Id
	pe.ServerId = mwo.ServerId
	pe.Period = mwo.Period
	pe.Grade = int32(mwo.Grade)
	pe.HunCheGrade = int32(mwo.HunCheGrade)
	pe.SugarGrade = int32(mwo.SugarGrade)
	pe.Status = int32(mwo.Status)
	pe.PlayerId = mwo.PlayerId
	pe.SpouseId = mwo.SpouseId
	pe.Name = mwo.Name
	pe.SpouseName = mwo.SpouseName
	pe.HTime = mwo.HTime
	pe.LastTime = mwo.LastTime
	pe.UpdateTime = mwo.UpdateTime
	pe.CreateTime = mwo.CreateTime
	pe.DeleteTime = mwo.DeleteTime
	if mwo.IsFirst {
		pe.IsFirst = 1
	} else {
		pe.IsFirst = 0
	}
	e = pe
	return
}

func (mwo *MarryWedObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*marryentity.MarryWedEntity)
	mwo.Id = pe.Id
	mwo.ServerId = pe.ServerId
	mwo.Period = pe.Period
	mwo.Grade = marrytypes.MarryBanquetSubTypeWed(pe.Grade)
	mwo.HunCheGrade = marrytypes.MarryBanquetSubTypeHunChe(pe.HunCheGrade)
	mwo.SugarGrade = marrytypes.MarryBanquetSubTypeSugar(pe.SugarGrade)
	mwo.Status = marrytypes.MarryWedStatusType(pe.Status)
	mwo.PlayerId = pe.PlayerId
	mwo.SpouseId = pe.SpouseId
	mwo.Name = pe.Name
	mwo.SpouseName = pe.SpouseName
	mwo.HTime = pe.HTime
	mwo.LastTime = pe.LastTime
	mwo.UpdateTime = pe.UpdateTime
	mwo.CreateTime = pe.CreateTime
	mwo.DeleteTime = pe.DeleteTime
	if pe.IsFirst > 0 {
		mwo.IsFirst = true
	} else {
		mwo.IsFirst = false
	}
	return
}

func (mwo *MarryWedObject) SetModified() {
	e, err := mwo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
