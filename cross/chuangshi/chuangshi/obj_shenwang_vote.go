package chuangshi

import (
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/cross/chuangshi/entity"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type ChuangShiShenWangVoteObject struct {
	id             int64
	platform       int32
	serverId       int32
	playerServerId int32
	campType       chuangshitypes.ChuangShiCampType
	playerId       int64
	ticketNum      int32
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewChuangShiShenWangVoteObject() *ChuangShiShenWangVoteObject {
	o := &ChuangShiShenWangVoteObject{}
	return o
}

func convertChuangShiVoteObjectToEntity(o *ChuangShiShenWangVoteObject) (*chuangshientity.ChuangShiShenWangVoteEntity, error) {
	e := &chuangshientity.ChuangShiShenWangVoteEntity{
		Id:             o.id,
		Platform:       o.platform,
		ServerId:       o.serverId,
		PlayerServerId: o.playerServerId,
		CampType:       int32(o.campType),
		PlayerId:       o.playerId,
		TicketNum:      o.ticketNum,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
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
	o.platform = pse.Platform
	o.serverId = pse.ServerId
	o.playerId = pse.PlayerId
	o.campType = chuangshitypes.ChuangShiCampType(pse.CampType)
	o.playerServerId = pse.PlayerServerId
	o.ticketNum = pse.TicketNum
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

func (o *ChuangShiShenWangVoteObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *ChuangShiShenWangVoteObject) GetTicketNum() int32 {
	return o.ticketNum
}
