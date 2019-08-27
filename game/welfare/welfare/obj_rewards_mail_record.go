package welfare

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	welfareentity "fgame/fgame/game/welfare/entity"

	"github.com/pkg/errors"
)

//排行榜邮件奖励记录对象
type OpenActivityEmailRecordObject struct {
	id         int64
	serverId   int32
	endTime    int64
	groupId    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newOpenActivityEmailRecordObject() *OpenActivityEmailRecordObject {
	o := &OpenActivityEmailRecordObject{}
	return o
}

func convertOpenActivityEmailRecordObjectToEntity(o *OpenActivityEmailRecordObject) (e *welfareentity.OpenActivityEmailRecordEntity, err error) {

	e = &welfareentity.OpenActivityEmailRecordEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		EndTime:    o.endTime,
		GroupId:    o.groupId,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *OpenActivityEmailRecordObject) GetDBId() int64 {
	return o.id
}

func (o *OpenActivityEmailRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertOpenActivityEmailRecordObjectToEntity(o)
	return e, err
}

func (o *OpenActivityEmailRecordObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.OpenActivityEmailRecordEntity)

	o.id = te.Id
	o.serverId = te.ServerId
	o.endTime = te.EndTime
	o.groupId = te.GroupId
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *OpenActivityEmailRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityEmailRecord"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

//排行榜邮件奖励记录对象
type OpenActivityStartEmailObject struct {
	id         int64
	serverId   int32
	endTime    int64
	groupId    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newOpenActivityStartEmailObject() *OpenActivityStartEmailObject {
	o := &OpenActivityStartEmailObject{}
	return o
}

func convertOpenActivityStartEmailObjectToEntity(o *OpenActivityStartEmailObject) (e *welfareentity.OpenActivityStartEmailEntity, err error) {

	e = &welfareentity.OpenActivityStartEmailEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		EndTime:    o.endTime,
		GroupId:    o.groupId,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *OpenActivityStartEmailObject) GetDBId() int64 {
	return o.id
}

func (o *OpenActivityStartEmailObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertOpenActivityStartEmailObjectToEntity(o)
	return e, err
}

func (o *OpenActivityStartEmailObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.OpenActivityStartEmailEntity)

	o.id = te.Id
	o.serverId = te.ServerId
	o.endTime = te.EndTime
	o.groupId = te.GroupId
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *OpenActivityStartEmailObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityStartEmail"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
