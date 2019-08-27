package chuangshi

import (
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/game/chuangshi/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type ChuangShiYuGaoObject struct {
	id         int64
	serverId   int32
	num        int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewChuangShiYuGaoObject() *ChuangShiYuGaoObject {
	o := &ChuangShiYuGaoObject{}
	return o
}

func convertChuangShiYuGaoObjectToEntity(o *ChuangShiYuGaoObject) (*chuangshientity.ChuangShiYuGaoEntity, error) {
	e := &chuangshientity.ChuangShiYuGaoEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		Num:        o.num,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *ChuangShiYuGaoObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiYuGaoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiYuGaoObjectToEntity(o)
	return e, err
}

func (o *ChuangShiYuGaoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiYuGaoEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.num = pse.Num
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiYuGaoObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShi"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
