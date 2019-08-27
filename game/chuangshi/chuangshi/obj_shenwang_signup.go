package chuangshi

import (
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/game/chuangshi/entity"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type ChuangShiShenWangSignUpObject struct {
	id         int64
	serverId   int32
	playerId   int64
	status     chuangshitypes.ShenWangSignUpType
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewChuangShiShenWangSignUpObject() *ChuangShiShenWangSignUpObject {
	o := &ChuangShiShenWangSignUpObject{}
	return o
}

func convertChuangShiSignUpObjectToEntity(o *ChuangShiShenWangSignUpObject) (*chuangshientity.ChuangShiShenWangSignUpEntity, error) {
	e := &chuangshientity.ChuangShiShenWangSignUpEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		PlayerId:   o.playerId,
		Status:     int32(o.status),
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *ChuangShiShenWangSignUpObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiShenWangSignUpObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiSignUpObjectToEntity(o)
	return e, err
}

func (o *ChuangShiShenWangSignUpObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiShenWangSignUpEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.playerId = pse.PlayerId
	o.status = chuangshitypes.ShenWangSignUpType(pse.Status)
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiShenWangSignUpObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiShenWangSignUp"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

func (o *ChuangShiShenWangSignUpObject) GetStatues() chuangshitypes.ShenWangSignUpType {
	return o.status
}

func (o *ChuangShiShenWangSignUpObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *ChuangShiShenWangSignUpObject) IfSigning() bool {
	return o.status == chuangshitypes.ShenWangSignUpTypeSigning
}
