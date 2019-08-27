package chuangshi

import (
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/game/chuangshi/entity"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type ChuangShiChengFangJianSheObject struct {
	id          int64
	serverId    int32
	playerId    int64
	cityId      int64
	jianSheType chuangshitypes.ChuangShiCityJianSheType
	num         int32
	status      chuangshitypes.ChengFangStatusType
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func NewChuangShiChengFangJianSheObject() *ChuangShiChengFangJianSheObject {
	o := &ChuangShiChengFangJianSheObject{}
	return o
}

func convertChuangShiChengFangJianSheObjectToEntity(o *ChuangShiChengFangJianSheObject) (*chuangshientity.ChuangShiChengFangJianSheEntity, error) {
	e := &chuangshientity.ChuangShiChengFangJianSheEntity{
		Id:          o.id,
		ServerId:    o.serverId,
		PlayerId:    o.playerId,
		CityId:      o.cityId,
		JianSheType: int32(o.jianSheType),
		Num:         o.num,
		Status:      int32(o.status),
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *ChuangShiChengFangJianSheObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiChengFangJianSheObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiChengFangJianSheObjectToEntity(o)
	return e, err
}

func (o *ChuangShiChengFangJianSheObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiChengFangJianSheEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.playerId = pse.PlayerId
	o.status = chuangshitypes.ChengFangStatusType(pse.Status)
	o.cityId = pse.CityId
	o.jianSheType = chuangshitypes.ChuangShiCityJianSheType(pse.JianSheType)
	o.num = pse.Num
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiChengFangJianSheObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiChengFangJianShe"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

func (o *ChuangShiChengFangJianSheObject) GetStatues() chuangshitypes.ChengFangStatusType {
	return o.status
}

func (o *ChuangShiChengFangJianSheObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *ChuangShiChengFangJianSheObject) IfProgressing() bool {
	return o.status == chuangshitypes.ChengFangStatusTypeProgressing
}
