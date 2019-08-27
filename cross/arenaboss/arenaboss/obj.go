package arenaboss

import (
	"fgame/fgame/core/storage"
	areanbossentity "fgame/fgame/cross/arenaboss/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type ArenaBossObject struct {
	id           int64
	platform     int32
	serverId     int32
	mapId        int32
	bossId       int32
	lastKillTime int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func newArenaBossObject() *ArenaBossObject {
	o := &ArenaBossObject{}
	return o
}

func convertArenaBossObjectToEntity(o *ArenaBossObject) (e *areanbossentity.ArenaBossEntity, err error) {
	e = &areanbossentity.ArenaBossEntity{
		Id:           o.id,
		Platform:     o.platform,
		ServerId:     o.serverId,
		MapId:        o.mapId,
		BossId:       o.bossId,
		LastKillTime: o.lastKillTime,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *ArenaBossObject) GetDBId() int64 {
	return o.id
}

func (o *ArenaBossObject) GetServerId() int32 {
	return o.serverId
}

func (o *ArenaBossObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertArenaBossObjectToEntity(o)
	return e, err
}

func (o *ArenaBossObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*areanbossentity.ArenaBossEntity)
	o.id = te.Id
	o.serverId = te.ServerId
	o.platform = te.Platform
	o.mapId = te.MapId
	o.bossId = te.BossId
	o.lastKillTime = te.LastKillTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *ArenaBossObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "arenaboss"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
