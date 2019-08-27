package dingshi

import (
	"fgame/fgame/core/storage"
	dingshientity "fgame/fgame/game/dingshi/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type DingShiBossObject struct {
	id           int64
	serverId     int32
	mapId        int32
	bossId       int32
	lastKillTime int64

	updateTime int64
	createTime int64
	deleteTime int64
}

func newDingShiBossObject() *DingShiBossObject {
	o := &DingShiBossObject{}
	return o
}

func convertDingShiBossObjectToEntity(o *DingShiBossObject) (e *dingshientity.DingShiBossEntity, err error) {
	e = &dingshientity.DingShiBossEntity{
		Id:           o.id,
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

func (o *DingShiBossObject) GetDBId() int64 {
	return o.id
}

func (o *DingShiBossObject) GetServerId() int32 {
	return o.serverId
}

func (o *DingShiBossObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertDingShiBossObjectToEntity(o)
	return e, err
}

func (o *DingShiBossObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*dingshientity.DingShiBossEntity)
	o.id = te.Id
	o.serverId = te.ServerId
	o.mapId = te.MapId
	o.bossId = te.BossId
	o.lastKillTime = te.LastKillTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *DingShiBossObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "dingshi"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
