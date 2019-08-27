package emperor

import (
	"fgame/fgame/core/storage"
	emperorentity "fgame/fgame/game/emperor/entity"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
)

//龙椅数据
type EmperorObject struct {
	Id                int64
	ServerId          int32
	EmperorId         int64
	Name              string
	Sex               playertypes.SexType
	SpouseName        string
	Storage           int64
	RobNum            int64
	RobTime           int64
	LastTime          int64
	BoxNum            int64
	BoxOutNum         int64
	SpecialBoxLeftNum int32
	BoxLastTime       int64
	UpdateTime        int64
	CreateTime        int64
	DeleteTime        int64
}

func NewEmperorObject() *EmperorObject {
	pso := &EmperorObject{}
	return pso
}

func (eo *EmperorObject) GetDBId() int64 {
	return eo.Id
}

func (eo *EmperorObject) ToEntity() (e storage.Entity, err error) {
	pe := &emperorentity.EmperorEntity{}
	pe.Id = eo.Id
	pe.ServerId = eo.ServerId
	pe.EmperorId = eo.EmperorId
	pe.Name = eo.Name
	pe.Sex = int32(eo.Sex)
	pe.SpouseName = eo.SpouseName
	pe.RobNum = eo.RobNum
	pe.Storage = eo.Storage
	pe.RobTime = eo.RobTime
	pe.LastTime = eo.LastTime
	pe.BoxNum = eo.BoxNum
	pe.BoxOutNum = eo.BoxOutNum
	pe.SpecialBoxLeftNum = eo.SpecialBoxLeftNum
	pe.BoxLastTime = eo.BoxLastTime
	pe.UpdateTime = eo.UpdateTime
	pe.CreateTime = eo.CreateTime
	pe.DeleteTime = eo.DeleteTime
	e = pe
	return
}

func (eo *EmperorObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*emperorentity.EmperorEntity)
	eo.Id = pe.Id
	eo.ServerId = pe.ServerId
	eo.EmperorId = pe.EmperorId
	eo.Name = pe.Name
	eo.Sex = playertypes.SexType(pe.Sex)
	eo.SpouseName = pe.SpouseName
	eo.RobNum = pe.RobNum
	eo.Storage = pe.Storage
	eo.RobTime = pe.RobTime
	eo.LastTime = pe.LastTime
	eo.BoxNum = pe.BoxNum
	eo.BoxOutNum = pe.BoxOutNum
	eo.BoxLastTime = pe.BoxLastTime
	eo.SpecialBoxLeftNum = pe.SpecialBoxLeftNum
	eo.UpdateTime = pe.UpdateTime
	eo.CreateTime = pe.CreateTime
	eo.DeleteTime = pe.DeleteTime
	return
}

func (eo *EmperorObject) SetModified() {
	e, err := eo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
