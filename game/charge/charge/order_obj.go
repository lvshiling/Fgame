package charge

import (
	"fgame/fgame/core/storage"
	chargeentity "fgame/fgame/game/charge/entity"
	"fgame/fgame/game/charge/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type OrderObject struct {
	id          int64
	serverId    int32
	orderId     string
	playerId    int64
	playerLevel int32
	chargeId    int32
	money       int32
	gold        int32
	orderStatus types.OrderStatus
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func newOrderObject() *OrderObject {
	o := &OrderObject{}
	return o
}

func convertOrderObjectObjectToEntity(o *OrderObject) (e *chargeentity.OrderEntity, err error) {
	e = &chargeentity.OrderEntity{
		Id:          o.id,
		ServerId:    o.serverId,
		OrderId:     o.orderId,
		OrderStatus: int32(o.orderStatus),
		PlayerId:    o.playerId,
		PlayerLevel: o.playerLevel,
		ChargeId:    o.chargeId,
		Money:       o.money,
		Gold:        o.gold,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *OrderObject) GetDBId() int64 {
	return o.id
}

func (o *OrderObject) GetServerId() int32 {
	return o.serverId
}
func (o *OrderObject) GetOrderId() string {
	return o.orderId
}

func (o *OrderObject) GetOrderStatus() types.OrderStatus {
	return o.orderStatus
}

func (o *OrderObject) GetChargeId() int32 {
	return o.chargeId
}

func (o *OrderObject) GetMoney() int32 {
	return o.money
}

func (o *OrderObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *OrderObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertOrderObjectObjectToEntity(o)
	return e, err
}

func (o *OrderObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*chargeentity.OrderEntity)
	o.id = te.Id
	o.serverId = te.ServerId
	o.orderId = te.OrderId
	o.orderStatus = types.OrderStatus(te.OrderStatus)
	o.chargeId = te.ChargeId
	o.playerId = te.PlayerId
	o.playerLevel = te.PlayerLevel
	o.money = te.Money
	o.gold = te.Gold
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *OrderObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "order"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
