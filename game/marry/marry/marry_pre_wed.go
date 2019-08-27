package marry

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	marryentity "fgame/fgame/game/marry/entity"
	marrytypes "fgame/fgame/game/marry/types"
)

//婚宴档次预定数据
type MarryPreWedObject struct {
	Id          int64
	ServerId    int32
	Period      int32
	PlayerId    int64
	PlayerName  string
	PeerId      int64
	Grade       marrytypes.MarryBanquetSubTypeWed
	HunCheGrade marrytypes.MarryBanquetSubTypeHunChe
	SugarGrade  marrytypes.MarryBanquetSubTypeSugar
	Status      marrytypes.MarryPreWedStatusType
	HoldTime    int64
	PreWedTime  int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewMarryPreWedObject() *MarryPreWedObject {
	pso := &MarryPreWedObject{}
	return pso
}

func (mpo *MarryPreWedObject) GetDBId() int64 {
	return mpo.Id
}

func (mpo *MarryPreWedObject) ToEntity() (e storage.Entity, err error) {
	pre := &marryentity.MarryPreWedEntity{}
	pre.Id = mpo.Id
	pre.ServerId = mpo.ServerId
	pre.Period = mpo.Period
	pre.PlayerId = mpo.PlayerId
	pre.PlayerName = mpo.PlayerName
	pre.PeerId = mpo.PeerId
	pre.Grade = int32(mpo.Grade)
	pre.HunCheGrade = int32(mpo.HunCheGrade)
	pre.SugarGrade = int32(mpo.SugarGrade)
	pre.Status = int32(mpo.Status)
	pre.HoldTime = mpo.HoldTime
	pre.PreWedTime = mpo.PreWedTime
	pre.UpdateTime = mpo.UpdateTime
	pre.CreateTime = mpo.CreateTime
	pre.DeleteTime = mpo.DeleteTime
	e = pre
	return
}

func (mpo *MarryPreWedObject) FromEntity(e storage.Entity) (err error) {
	pre, _ := e.(*marryentity.MarryPreWedEntity)
	mpo.Id = pre.Id
	mpo.ServerId = pre.ServerId
	mpo.Period = pre.Period
	mpo.PlayerId = pre.PlayerId
	mpo.PlayerName = pre.PlayerName
	mpo.PeerId = pre.PeerId
	mpo.Grade = marrytypes.MarryBanquetSubTypeWed(pre.Grade)
	mpo.HunCheGrade = marrytypes.MarryBanquetSubTypeHunChe(pre.HunCheGrade)
	mpo.SugarGrade = marrytypes.MarryBanquetSubTypeSugar(pre.SugarGrade)
	mpo.Status = marrytypes.MarryPreWedStatusType(pre.Status)
	mpo.HoldTime = pre.HoldTime
	mpo.PreWedTime = pre.PreWedTime
	mpo.UpdateTime = pre.UpdateTime
	mpo.CreateTime = pre.CreateTime
	mpo.DeleteTime = pre.DeleteTime
	return
}

func (mpo *MarryPreWedObject) SetModified() {
	e, err := mpo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
