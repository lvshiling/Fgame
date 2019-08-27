package player

import (
	"fgame/fgame/core/storage"
	chargeentity "fgame/fgame/game/charge/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//充值对象
type PlayerPrivilegeChargeObject struct {
	player     player.Player
	id         int64
	chargeType int32
	chargeId   int32
	chargeNum  int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerPrivilegeChargeObject(pl player.Player) *PlayerPrivilegeChargeObject {
	o := &PlayerPrivilegeChargeObject{
		player: pl,
	}
	return o
}

func convertNewPlayerPrivilegeChargeObjectToEntity(o *PlayerPrivilegeChargeObject) (*chargeentity.PlayerPrivilegeChargeEntity, error) {
	e := &chargeentity.PlayerPrivilegeChargeEntity{
		Id:         o.id,
		PlayerId:   o.GetPlayerId(),
		ChargeType: o.chargeType,
		ChargeId:   o.chargeId,
		ChargeNum:  o.chargeNum,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerPrivilegeChargeObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerPrivilegeChargeObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerPrivilegeChargeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerPrivilegeChargeObjectToEntity(o)
	return e, err
}

func (o *PlayerPrivilegeChargeObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chargeentity.PlayerPrivilegeChargeEntity)

	o.id = pse.Id
	o.chargeType = pse.ChargeType
	o.chargeNum = pse.ChargeNum
	o.chargeId = pse.ChargeId
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerPrivilegeChargeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PrivilegeCharge"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerPrivilegeChargeObject) GetChargeNum() int32 {

	return o.chargeNum
}
