package jieyi

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	jieyientity "fgame/fgame/game/jieyi/entity"

	playertypes "fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

type JieYiLeaveWordObject struct {
	id           int64
	serverId     int32
	playerId     int64
	name         string
	level        int32
	role         playertypes.RoleType
	sex          playertypes.SexType
	force        int64
	leaveWord    string
	lastPostTime int64
	onlineStatus playertypes.PlayerOnlineState
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewJieYiLeaveWordObject() *JieYiLeaveWordObject {
	o := &JieYiLeaveWordObject{}
	return o
}

func convertJieYiLeaveWordObjectToEntity(o *JieYiLeaveWordObject) (*jieyientity.JieYiLeaveWordEntity, error) {
	e := &jieyientity.JieYiLeaveWordEntity{
		Id:           o.id,
		ServerId:     o.serverId,
		PlayerId:     o.playerId,
		Name:         o.name,
		Level:        o.level,
		Role:         int32(o.role),
		Sex:          int32(o.sex),
		Force:        o.force,
		LeaveWord:    o.leaveWord,
		LastPostTime: o.lastPostTime,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *JieYiLeaveWordObject) GetDBId() int64 {
	return o.id
}

func (o *JieYiLeaveWordObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *JieYiLeaveWordObject) GetPlayerName() string {
	return o.name
}

func (o *JieYiLeaveWordObject) GetRole() playertypes.RoleType {
	return o.role
}

func (o *JieYiLeaveWordObject) GetSex() playertypes.SexType {
	return o.sex
}

func (o *JieYiLeaveWordObject) GetLevel() int32 {
	return o.level
}

func (o *JieYiLeaveWordObject) GetForce() int64 {
	return o.force
}

func (o *JieYiLeaveWordObject) GetLeaveWord() string {
	return o.leaveWord
}

func (o *JieYiLeaveWordObject) GetLastPostTime() int64 {
	return o.lastPostTime
}

func (o *JieYiLeaveWordObject) GetOnLineState() playertypes.PlayerOnlineState {
	return o.onlineStatus
}

func (o *JieYiLeaveWordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertJieYiLeaveWordObjectToEntity(o)
	return e, err
}

func (o *JieYiLeaveWordObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*jieyientity.JieYiLeaveWordEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.playerId = pse.PlayerId
	o.name = pse.Name
	o.level = pse.Level
	o.role = playertypes.RoleType(pse.Role)
	o.sex = playertypes.SexType(pse.Sex)
	o.force = pse.Force
	o.leaveWord = pse.LeaveWord
	o.lastPostTime = pse.LastPostTime
	o.onlineStatus = playertypes.PlayerOnlineStateOffline
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *JieYiLeaveWordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "JieYiLeaveWord"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

type jieYiLeaveWorldList []*JieYiLeaveWordObject

func (adl jieYiLeaveWorldList) Len() int {
	return len(adl)
}

func (adl jieYiLeaveWorldList) Less(i, j int) bool {
	if adl[i].onlineStatus != adl[j].onlineStatus {
		return adl[i].onlineStatus < adl[j].onlineStatus
	}
	return adl[i].lastPostTime < adl[j].lastPostTime
}

func (adl jieYiLeaveWorldList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}
