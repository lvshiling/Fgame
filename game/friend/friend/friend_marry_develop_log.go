package friend

import (
	"fgame/fgame/core/storage"
	friendentity "fgame/fgame/game/friend/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//所有表白记录数据
type FriendMarryDevelopLogObject struct {
	Id         int64
	ServerId   int32
	SendId     int64
	RecvId     int64
	SendName   string
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

func NewFriendMarryDevelopLogObject() *FriendMarryDevelopLogObject {
	o := &FriendMarryDevelopLogObject{}
	return o
}

func convertNewFriendMarryDevelopLogObjectToEntity(o *FriendMarryDevelopLogObject) (*friendentity.FriendMarryDevelopLogEntity, error) {
	e := &friendentity.FriendMarryDevelopLogEntity{
		Id:         o.Id,
		ServerId:   o.ServerId,
		SendId:     o.SendId,
		RecvId:     o.RecvId,
		SendName:   o.SendName,
		RecvName:   o.RecvName,
		ItemId:     o.ItemId,
		ItemNum:    o.ItemNum,
		CharmNum:   o.CharmNum,
		DevelopExp: o.DevelopExp,
		ContextStr: o.ContextStr,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *FriendMarryDevelopLogObject) GetDBId() int64 {
	return o.Id
}

func (o *FriendMarryDevelopLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewFriendMarryDevelopLogObjectToEntity(o)
	return e, err
}

func (o *FriendMarryDevelopLogObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*friendentity.FriendMarryDevelopLogEntity)

	o.Id = pse.Id
	o.ServerId = pse.ServerId
	o.SendId = pse.SendId
	o.RecvId = pse.RecvId
	o.SendName = pse.SendName
	o.RecvName = pse.RecvName
	o.ItemId = pse.ItemId
	o.ItemNum = pse.ItemNum
	o.CharmNum = pse.CharmNum
	o.DevelopExp = pse.DevelopExp
	o.ContextStr = pse.ContextStr
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *FriendMarryDevelopLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FriendMarryDevelopLog"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
