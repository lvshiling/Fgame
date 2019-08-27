package friend

import (
	"fgame/fgame/core/storage"
	friendentity "fgame/fgame/game/friend/entity"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/global"
)

//好友系统操作数据日志
type PlayerFriendLogObject struct {
	Id         int64
	PlayerId   int64
	FrinedId   int64
	Type       friendtypes.FriendLogType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerFriendLogObject() *PlayerFriendLogObject {
	pso := &PlayerFriendLogObject{}
	return pso
}

func (pflo *PlayerFriendLogObject) GetDBId() int64 {
	return pflo.Id
}

func (pflo *PlayerFriendLogObject) ToEntity() (e storage.Entity, err error) {
	pfle := &friendentity.PlayerFriendLogEntity{}
	pfle.Id = pflo.Id
	pfle.PlayerId = pflo.PlayerId
	pfle.FriendId = pflo.FrinedId
	pfle.Type = int32(pflo.Type)
	pfle.UpdateTime = pflo.UpdateTime
	pfle.CreateTime = pflo.CreateTime
	pfle.DeleteTime = pflo.DeleteTime
	e = pfle
	return
}

func (pflo *PlayerFriendLogObject) FromEntity(e storage.Entity) (err error) {
	pfle, _ := e.(*friendentity.PlayerFriendLogEntity)
	pflo.Id = pfle.Id
	pflo.PlayerId = pfle.PlayerId
	pflo.FrinedId = pfle.FriendId
	pflo.Type = friendtypes.FriendLogType(pfle.Type)
	pflo.UpdateTime = pfle.UpdateTime
	pflo.CreateTime = pfle.CreateTime
	pflo.DeleteTime = pfle.DeleteTime
	return
}

func (pflo *PlayerFriendLogObject) SetModified() {
	e, err := pflo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
