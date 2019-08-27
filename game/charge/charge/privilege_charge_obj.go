package charge

import (
	"fgame/fgame/core/storage"
	chargeentity "fgame/fgame/game/charge/entity"
	"fgame/fgame/game/charge/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

// 后台充值
type PrivilegeChargeObject struct {
	id         int64
	serverId   int32
	status     types.OrderStatus
	playerId   int64
	goldNum    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPrivilegeChargeObject() *PrivilegeChargeObject {
	o := &PrivilegeChargeObject{}
	return o
}

func convertPrivilegeChargeObjectObjectToEntity(o *PrivilegeChargeObject) (e *chargeentity.PrivilegeChargeEntity, err error) {
	e = &chargeentity.PrivilegeChargeEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		Status:     int32(o.status),
		PlayerId:   o.playerId,
		GoldNum:    o.goldNum,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PrivilegeChargeObject) GetPrivilegeCharge() int64 {
	return o.id
}

func (o *PrivilegeChargeObject) GetDBId() int64 {
	return o.id
}

func (o *PrivilegeChargeObject) GetServerId() int32 {
	return o.serverId
}

func (o *PrivilegeChargeObject) GetGoldNum() int64 {
	return o.goldNum
}

func (o *PrivilegeChargeObject) GetStatus() types.OrderStatus {
	return o.status
}

func (o *PrivilegeChargeObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *PrivilegeChargeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPrivilegeChargeObjectObjectToEntity(o)
	return e, err
}

func (o *PrivilegeChargeObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*chargeentity.PrivilegeChargeEntity)
	o.id = te.Id
	o.serverId = te.ServerId
	o.status = types.OrderStatus(te.Status)
	o.playerId = te.PlayerId
	o.goldNum = te.GoldNum
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PrivilegeChargeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "privilegeCharge"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
