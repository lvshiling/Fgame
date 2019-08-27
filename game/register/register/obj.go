package register

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	registerentity "fgame/fgame/game/register/entity"

	"github.com/pkg/errors"
)

//注册设置
type RegisterSettingObject struct {
	id         int64
	serverId   int32
	open       int32
	auto       int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func createRegisterSettingObject() *RegisterSettingObject {
	o := &RegisterSettingObject{}
	return o
}

func convertRegisterSettingObjectToEntity(o *RegisterSettingObject) (*registerentity.RegisterSettingEntity, error) {
	e := &registerentity.RegisterSettingEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		Open:       o.open,
		Auto:       o.auto,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *RegisterSettingObject) GetId() int64 {
	return o.id
}

func (o *RegisterSettingObject) GetDBId() int64 {
	return o.id
}

func (o *RegisterSettingObject) GetServerId() int32 {
	return o.serverId
}

func (o *RegisterSettingObject) GetOpen() int32 {
	return o.open
}

func (o *RegisterSettingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertRegisterSettingObjectToEntity(o)
	return e, err
}

func (o *RegisterSettingObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*registerentity.RegisterSettingEntity)
	o.id = ae.Id
	o.serverId = ae.ServerId
	o.open = ae.Open
	o.auto = ae.Auto
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *RegisterSettingObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "RegisterSetting"))
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

//注册设置
type RegisterSettingLogObject struct {
	id       int64
	serverId int32
	open     int32

	updateTime int64
	createTime int64
	deleteTime int64
}

func createRegisterSettingLogObject() *RegisterSettingLogObject {
	o := &RegisterSettingLogObject{}
	return o
}

func convertRegisterSettingLogObjectToEntity(o *RegisterSettingLogObject) (*registerentity.RegisterSettingLogEntity, error) {
	e := &registerentity.RegisterSettingLogEntity{
		Id:       o.id,
		ServerId: o.serverId,
		Open:     o.open,

		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *RegisterSettingLogObject) GetId() int64 {
	return o.id
}

func (o *RegisterSettingLogObject) GetDBId() int64 {
	return o.id
}

func (o *RegisterSettingLogObject) GetServerId() int32 {
	return o.serverId
}

func (o *RegisterSettingLogObject) GetOpen() int32 {
	return o.open
}

func (o *RegisterSettingLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertRegisterSettingLogObjectToEntity(o)
	return e, err
}

func (o *RegisterSettingLogObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*registerentity.RegisterSettingLogEntity)
	o.id = ae.Id
	o.serverId = ae.ServerId
	o.open = ae.Open

	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *RegisterSettingLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Register"))
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
