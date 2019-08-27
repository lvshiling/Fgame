package charge

import (
	"fgame/fgame/core/storage"
	chargeentity "fgame/fgame/game/charge/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//废弃:zrc
type FirstChargeObject struct {
	Id         int64
	ServerId   int32
	ChargeTime int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func newFirstChargeObject() *FirstChargeObject {
	o := &FirstChargeObject{}
	return o
}

func convertFirstChargeObjectObjectToEntity(o *FirstChargeObject) (e *chargeentity.FirstChargeEntity, err error) {

	e = &chargeentity.FirstChargeEntity{
		Id:         o.Id,
		ServerId:   o.ServerId,
		ChargeTime: o.ChargeTime,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *FirstChargeObject) GetDBId() int64 {
	return o.Id
}

func (o *FirstChargeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertFirstChargeObjectObjectToEntity(o)
	return e, err
}

func (o *FirstChargeObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*chargeentity.FirstChargeEntity)

	o.Id = te.Id
	o.ServerId = te.ServerId
	o.ChargeTime = te.ChargeTime
	o.UpdateTime = te.UpdateTime
	o.CreateTime = te.CreateTime
	o.DeleteTime = te.DeleteTime
	return nil
}

func (o *FirstChargeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "firstCharge"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
