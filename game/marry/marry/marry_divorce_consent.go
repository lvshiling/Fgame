package marry

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	marryentity "fgame/fgame/game/marry/entity"
)

//协议离婚成功玩家下线数据
type MarryDivorceConsentObject struct {
	Id         int64
	ServerId   int32
	PlayerId   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewMarryDivorceConsentObject() *MarryDivorceConsentObject {
	pso := &MarryDivorceConsentObject{}
	return pso
}

func (mdco *MarryDivorceConsentObject) GetDBId() int64 {
	return mdco.Id
}

func (mdco *MarryDivorceConsentObject) ToEntity() (e storage.Entity, err error) {
	pe := &marryentity.MarryDivorceConsentEntity{}
	pe.Id = mdco.Id
	pe.ServerId = mdco.ServerId
	pe.PlayerId = mdco.PlayerId
	pe.UpdateTime = mdco.UpdateTime
	pe.CreateTime = mdco.CreateTime
	pe.DeleteTime = mdco.DeleteTime
	e = pe
	return
}

func (mdco *MarryDivorceConsentObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*marryentity.MarryDivorceConsentEntity)
	mdco.Id = pe.Id
	mdco.ServerId = pe.ServerId
	mdco.PlayerId = pe.PlayerId
	mdco.UpdateTime = pe.UpdateTime
	mdco.CreateTime = pe.CreateTime
	mdco.DeleteTime = pe.DeleteTime
	return
}

func (mdco *MarryDivorceConsentObject) SetModified() {
	e, err := mdco.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
