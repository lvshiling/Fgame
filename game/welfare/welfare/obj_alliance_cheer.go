package welfare

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	welfareentity "fgame/fgame/game/welfare/entity"

	"github.com/pkg/errors"
)

//城战助威记录对象
type AllianceCheerObject struct {
	id         int64
	serverId   int32
	groupId    int32
	allianceId int64
	startTime  int64
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newAllianceCheerObject() *AllianceCheerObject {
	o := &AllianceCheerObject{}
	return o
}

func convertNewAllianceCheerObjectToEntity(o *AllianceCheerObject) (*welfareentity.OpenActivityAllianceCheerEntity, error) {
	e := &welfareentity.OpenActivityAllianceCheerEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		GroupId:    o.groupId,
		AllianceId: o.allianceId,
		StartTime:  o.startTime,
		EndTime:    o.endTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *AllianceCheerObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *AllianceCheerObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *AllianceCheerObject) GetGroupId() int32 {
	return o.groupId
}

func (o *AllianceCheerObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceCheerObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewAllianceCheerObjectToEntity(o)
	return e, err
}

func (o *AllianceCheerObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*welfareentity.OpenActivityAllianceCheerEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.groupId = pse.GroupId
	o.allianceId = pse.AllianceId
	o.startTime = pse.StartTime
	o.endTime = pse.EndTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *AllianceCheerObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceCheer"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
