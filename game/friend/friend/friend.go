package friend

import (
	"fgame/fgame/core/storage"
	friendentity "fgame/fgame/game/friend/entity"
	"fgame/fgame/game/global"
)

//好友数据
type FriendObject struct {
	Id         int64
	ServerId   int32
	PlayerId   int64
	FriendId   int64
	Point      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewFriendObject() *FriendObject {
	pso := &FriendObject{}
	return pso
}

func (fo *FriendObject) GetDBId() int64 {
	return fo.Id
}

func (fo *FriendObject) ToEntity() (e storage.Entity, err error) {
	fe := &friendentity.FriendEntity{}
	fe.Id = fo.Id
	fe.ServerId = fo.ServerId
	fe.PlayerId = fo.PlayerId
	fe.FriendId = fo.FriendId
	fe.Point = fo.Point
	fe.UpdateTime = fo.UpdateTime
	fe.CreateTime = fo.CreateTime
	fe.DeleteTime = fo.DeleteTime
	e = fe
	return
}

func (fo *FriendObject) FromEntity(e storage.Entity) (err error) {
	fe, _ := e.(*friendentity.FriendEntity)
	fo.Id = fe.Id
	fo.ServerId = fe.ServerId
	fo.PlayerId = fe.PlayerId
	fo.FriendId = fe.FriendId
	fo.Point = fe.Point
	fo.UpdateTime = fe.UpdateTime
	fo.CreateTime = fe.CreateTime
	fo.DeleteTime = fe.DeleteTime
	return
}

func (fo *FriendObject) SetModified() {
	e, err := fo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
