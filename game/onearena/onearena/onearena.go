package onearena

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	onearenaentity "fgame/fgame/game/onearena/entity"
	onearenatypes "fgame/fgame/game/onearena/types"
)

//灵池数据
type OneArenaObject struct {
	Id         int64
	ServerId   int32
	Level      onearenatypes.OneArenaLevelType
	Pos        int32
	OwnerId    int64
	OwnerName  string
	LastTime   int64
	IsRobbing  bool
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewOneArenaObject() *OneArenaObject {
	poo := &OneArenaObject{}
	return poo
}

func (oo *OneArenaObject) GetDBId() int64 {
	return oo.Id
}

func (oo *OneArenaObject) ToEntity() (e storage.Entity, err error) {
	oe := &onearenaentity.OneArenaEntity{}
	oe.Id = oo.Id
	oe.ServerId = oo.ServerId
	oe.Level = int32(oo.Level)
	oe.Pos = oo.Pos
	oe.OwnerId = oo.OwnerId
	oe.OwnerName = oo.OwnerName
	oe.LastTime = oo.LastTime
	oe.UpdateTime = oo.UpdateTime
	oe.CreateTime = oo.CreateTime
	oe.DeleteTime = oo.DeleteTime
	e = oe
	return
}

func (oo *OneArenaObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*onearenaentity.OneArenaEntity)
	oo.Id = oe.Id
	oo.ServerId = oe.ServerId
	oo.Level = onearenatypes.OneArenaLevelType(oe.Level)
	oo.Pos = oe.Pos
	oo.OwnerId = oe.OwnerId
	oo.OwnerName = oe.OwnerName
	oo.LastTime = oe.LastTime
	oo.UpdateTime = oe.UpdateTime
	oo.CreateTime = oe.CreateTime
	oo.DeleteTime = oe.DeleteTime
	return
}

func (oo *OneArenaObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
