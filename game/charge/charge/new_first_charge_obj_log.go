package charge

import (
	"fgame/fgame/core/storage"
	chargeentity "fgame/fgame/game/charge/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//
type NewFirstChargeLogObject struct {
	id       int64
	serverId int32

	updateTime int64
	createTime int64
	deleteTime int64
}

func newNewFirstChargeLogObject() *NewFirstChargeLogObject {
	o := &NewFirstChargeLogObject{}
	return o
}

func convertNewFirstChargeLogObjectToEntity(o *NewFirstChargeLogObject) (e *chargeentity.NewFirstChargeLogEntity, err error) {

	e = &chargeentity.NewFirstChargeLogEntity{
		Id:       o.id,
		ServerId: o.serverId,

		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *NewFirstChargeLogObject) GetDBId() int64 {
	return o.id
}

func (o *NewFirstChargeLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewFirstChargeLogObjectToEntity(o)
	return e, err
}

func (o *NewFirstChargeLogObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*chargeentity.NewFirstChargeEntity)

	o.id = te.Id
	o.serverId = te.ServerId

	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *NewFirstChargeLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "NewFirstChargeLog"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
