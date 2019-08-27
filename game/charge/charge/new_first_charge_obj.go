package charge

import (
	"fgame/fgame/core/storage"
	chargeentity "fgame/fgame/game/charge/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//
type NewFirstChargeObject struct {
	id         int64
	serverId   int32
	startTime  int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newNewFirstChargeObject() *NewFirstChargeObject {
	o := &NewFirstChargeObject{}
	return o
}

func convertNewFirstChargeObjectToEntity(o *NewFirstChargeObject) (e *chargeentity.NewFirstChargeEntity, err error) {

	e = &chargeentity.NewFirstChargeEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		StartTime:  o.startTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *NewFirstChargeObject) GetDBId() int64 {
	return o.id
}

func (o *NewFirstChargeObject) GetStartTime() int64 {
	return o.startTime
}

func (o *NewFirstChargeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewFirstChargeObjectToEntity(o)
	return e, err
}

func (o *NewFirstChargeObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*chargeentity.NewFirstChargeEntity)

	o.id = te.Id
	o.serverId = te.ServerId
	o.startTime = te.StartTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *NewFirstChargeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "NewFirstCharge"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
