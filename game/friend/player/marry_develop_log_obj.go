package player

import (
	"fgame/fgame/core/storage"
	friendentity "fgame/fgame/game/friend/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//我的表白记录数据对象
type PlayerFriendMarryDevelopSendLogObject struct {
	player     player.Player
	Id         int64
	RecvId     int64
	RecvName   string
	ItemId     int32
	ItemNum    int32
	CharmNum   int32
	DevelopExp int32
	ContextStr string
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func newPlayerFriendMarryDevelopSendLogObject(pl player.Player) *PlayerFriendMarryDevelopSendLogObject {
	obj := &PlayerFriendMarryDevelopSendLogObject{
		player: pl,
	}
	return obj
}

func convertPlayerFriendMarryDevelopSendLogObjectToEntity(obj *PlayerFriendMarryDevelopSendLogObject) (e *friendentity.PlayerFriendMarryDevelopSendLogEntity, err error) {

	e = &friendentity.PlayerFriendMarryDevelopSendLogEntity{
		Id:         obj.Id,
		PlayerId:   obj.player.GetId(),
		RecvId:     obj.RecvId,
		RecvName:   obj.RecvName,
		ItemId:     obj.ItemId,
		ItemNum:    obj.ItemNum,
		CharmNum:   obj.CharmNum,
		DevelopExp: obj.DevelopExp,
		ContextStr: obj.ContextStr,
		UpdateTime: obj.UpdateTime,
		CreateTime: obj.CreateTime,
		DeleteTime: obj.DeleteTime,
	}
	return e, nil
}

func (obj *PlayerFriendMarryDevelopSendLogObject) GetPlayerId() int64 {
	return obj.player.GetId()
}

func (obj *PlayerFriendMarryDevelopSendLogObject) GetPlayerName() string {
	return obj.player.GetName()
}

func (obj *PlayerFriendMarryDevelopSendLogObject) GetDBId() int64 {
	return obj.Id
}

func (obj *PlayerFriendMarryDevelopSendLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFriendMarryDevelopSendLogObjectToEntity(obj)
	return e, err
}

func (obj *PlayerFriendMarryDevelopSendLogObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*friendentity.PlayerFriendMarryDevelopSendLogEntity)
	obj.Id = pe.Id
	obj.RecvId = pe.RecvId
	obj.RecvName = pe.RecvName
	obj.ItemId = pe.ItemId
	obj.ItemNum = pe.ItemNum
	obj.CharmNum = pe.CharmNum
	obj.DevelopExp = pe.DevelopExp
	obj.ContextStr = pe.ContextStr
	obj.UpdateTime = pe.UpdateTime
	obj.CreateTime = pe.CreateTime
	obj.DeleteTime = pe.DeleteTime
	return nil
}

func (obj *PlayerFriendMarryDevelopSendLogObject) SetModified() {
	e, err := obj.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Player_Friend_Marry_Develop_Send_Log"))
	}
	pe, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	obj.player.AddChangedObject(pe)
	return
}

//对我的表白记录数据对象
type PlayerFriendMarryDevelopRecvLogObject struct {
	player     player.Player
	Id         int64
	SendId     int64
	SendName   string
	ItemId     int32
	ItemNum    int32
	CharmNum   int32
	DevelopExp int32
	ContextStr string
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func newPlayerFriendMarryDevelopRecvLogObject(pl player.Player) *PlayerFriendMarryDevelopRecvLogObject {
	obj := &PlayerFriendMarryDevelopRecvLogObject{
		player: pl,
	}
	return obj
}

func convertPlayerFriendMarryDevelopRecvLogObjectToEntity(obj *PlayerFriendMarryDevelopRecvLogObject) (e *friendentity.PlayerFriendMarryDevelopRecvLogEntity, err error) {

	e = &friendentity.PlayerFriendMarryDevelopRecvLogEntity{
		Id:         obj.Id,
		PlayerId:   obj.player.GetId(),
		SendId:     obj.SendId,
		SendName:   obj.SendName,
		ItemId:     obj.ItemId,
		ItemNum:    obj.ItemNum,
		CharmNum:   obj.CharmNum,
		DevelopExp: obj.DevelopExp,
		ContextStr: obj.ContextStr,
		UpdateTime: obj.UpdateTime,
		CreateTime: obj.CreateTime,
		DeleteTime: obj.DeleteTime,
	}
	return e, nil
}

func (obj *PlayerFriendMarryDevelopRecvLogObject) GetPlayerId() int64 {
	return obj.player.GetId()
}

func (obj *PlayerFriendMarryDevelopRecvLogObject) GetPlayerName() string {
	return obj.player.GetName()
}

func (obj *PlayerFriendMarryDevelopRecvLogObject) GetDBId() int64 {
	return obj.Id
}

func (obj *PlayerFriendMarryDevelopRecvLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFriendMarryDevelopRecvLogObjectToEntity(obj)
	return e, err
}

func (obj *PlayerFriendMarryDevelopRecvLogObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*friendentity.PlayerFriendMarryDevelopRecvLogEntity)
	obj.Id = pe.Id
	obj.SendId = pe.SendId
	obj.SendName = pe.SendName
	obj.ItemId = pe.ItemId
	obj.ItemNum = pe.ItemNum
	obj.CharmNum = pe.CharmNum
	obj.DevelopExp = pe.DevelopExp
	obj.ContextStr = pe.ContextStr
	obj.UpdateTime = pe.UpdateTime
	obj.CreateTime = pe.CreateTime
	obj.DeleteTime = pe.DeleteTime
	return nil
}

func (obj *PlayerFriendMarryDevelopRecvLogObject) SetModified() {
	e, err := obj.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Player_Friend_Marry_Develop_Recv_Log"))
	}
	pe, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	obj.player.AddChangedObject(pe)
	return
}
