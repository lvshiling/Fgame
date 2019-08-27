package chuangshi

import (
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/cross/chuangshi/entity"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

type ChuangShiShenWangVoteRecordObject struct {
	id                  int64
	platform            int32
	serverId            int32
	playerServerId      int32
	campType            chuangshitypes.ChuangShiCampType
	playerId            int64
	playerName          string
	houXuanPlatform     int32
	houXuanGameServerId int32
	houXuanPlayerId     int64
	houXuanPlayerName   string
	lastVoteTime        int64
	updateTime          int64
	createTime          int64
	deleteTime          int64
}

func NewChuangShiShenWangVoteRecordObject() *ChuangShiShenWangVoteRecordObject {
	o := &ChuangShiShenWangVoteRecordObject{}
	return o
}

func convertChuangShiVoteRecordObjectToEntity(o *ChuangShiShenWangVoteRecordObject) (*chuangshientity.ChuangShiShenWangVoteRecordEntity, error) {
	e := &chuangshientity.ChuangShiShenWangVoteRecordEntity{
		Id:                  o.id,
		Platform:            o.platform,
		ServerId:            o.serverId,
		PlayerServerId:      o.playerServerId,
		CampType:            int32(o.campType),
		PlayerId:            o.playerId,
		PlayerName:          o.playerName,
		HouXuanPlatform:     o.houXuanPlatform,
		HouXuanGameServerId: o.houXuanGameServerId,
		HouXuanPlayerId:     o.houXuanPlayerId,
		HouXuanPlayerName:   o.houXuanPlayerName,
		LastVoteTime:        o.lastVoteTime,
		UpdateTime:          o.updateTime,
		CreateTime:          o.createTime,
		DeleteTime:          o.deleteTime,
	}
	return e, nil
}

func (o *ChuangShiShenWangVoteRecordObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiShenWangVoteRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiVoteRecordObjectToEntity(o)
	return e, err
}

func (o *ChuangShiShenWangVoteRecordObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiShenWangVoteRecordEntity)

	o.id = pse.Id
	o.platform = pse.Platform
	o.serverId = pse.ServerId
	o.campType = chuangshitypes.ChuangShiCampType(pse.CampType)
	o.playerServerId = pse.PlayerServerId
	o.playerId = pse.PlayerId
	o.playerName = pse.PlayerName
	o.houXuanPlatform = pse.HouXuanPlatform
	o.houXuanGameServerId = pse.HouXuanGameServerId
	o.houXuanPlayerId = pse.HouXuanPlayerId
	o.houXuanPlayerName = pse.HouXuanPlayerName
	o.lastVoteTime = pse.LastVoteTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiShenWangVoteRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiShenWangVote"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

func (o *ChuangShiShenWangVoteRecordObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *ChuangShiShenWangVoteRecordObject) IfCanVote(now int64) bool {
	isSame, _ := timeutils.IsSameDay(now, o.lastVoteTime)
	if isSame {
		return false
	}

	return true
}
