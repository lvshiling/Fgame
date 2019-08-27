package merge

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	mergeentity "fgame/fgame/game/merge/entity"

	"github.com/pkg/errors"
)

type MergeObject struct {
	id         int64
	serverId   int32
	merge      int32
	mergeTime  int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func (o *MergeObject) GetId() int64 {
	return o.id
}

func (o *MergeObject) GetServerId() int32 {
	return o.serverId
}

func (o *MergeObject) GetMerge() int32 {
	return o.merge
}

func (o *MergeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertMergeObjectToEntity(o)
	return e, err
}

func (o *MergeObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*mergeentity.MergeEntity)
	o.id = ae.Id
	o.serverId = ae.ServerId
	o.merge = ae.Merge
	o.mergeTime = ae.MergeTime
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *MergeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Merge"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func newMergeObject() *MergeObject {
	o := &MergeObject{}
	return o
}

func convertMergeObjectToEntity(o *MergeObject) (*mergeentity.MergeEntity, error) {
	e := &mergeentity.MergeEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		Merge:      o.merge,
		MergeTime:  o.mergeTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}
