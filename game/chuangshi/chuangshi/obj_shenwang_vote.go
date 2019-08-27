package chuangshi

import (
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/game/chuangshi/entity"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type ChuangShiShenWangVoteObject struct {
	id         int64
	serverId   int32
	playerId   int64
	supportId  int64
	status     chuangshitypes.ShenWangVoteType
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewChuangShiShenWangVoteObject() *ChuangShiShenWangVoteObject {
	o := &ChuangShiShenWangVoteObject{}
	return o
}

func convertChuangShiVoteObjectToEntity(o *ChuangShiShenWangVoteObject) (*chuangshientity.ChuangShiShenWangVoteEntity, error) {
	e := &chuangshientity.ChuangShiShenWangVoteEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		PlayerId:   o.playerId,
		SupportId:  o.supportId,
		Status:     int32(o.status),
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *ChuangShiShenWangVoteObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiShenWangVoteObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiVoteObjectToEntity(o)
	return e, err
}

func (o *ChuangShiShenWangVoteObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiShenWangVoteEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.playerId = pse.PlayerId
	o.supportId = pse.SupportId
	o.status = chuangshitypes.ShenWangVoteType(pse.Status)
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiShenWangVoteObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiShenWangVote"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

func (o *ChuangShiShenWangVoteObject) GetStatues() chuangshitypes.ShenWangVoteType {
	return o.status
}

func (o *ChuangShiShenWangVoteObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *ChuangShiShenWangVoteObject) GetSupportId() int64 {
	return o.supportId
}

func (o *ChuangShiShenWangVoteObject) IfVoting() bool {
	return o.status == chuangshitypes.ShenWangVoteTypeVoting
}
