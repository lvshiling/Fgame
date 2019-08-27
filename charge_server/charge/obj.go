package charge

import "fgame/fgame/charge_server/types"
import logintypes "fgame/fgame/account/login/types"

type OrderObject struct {
	id             int64
	orderId        string
	sdkOrderId     string
	status         types.OrderStatus
	sdkType        logintypes.SDKType
	devicePlatform logintypes.DevicePlatformType
	platformUserId string
	serverId       int32
	userId         int64
	playerId       int64
	playerLevel    int32
	playerName     string
	chargeId       int32
	money          int32
	gold           int32
	receivePayTime int64
	createTime     int64
	updateTime     int64
	deleteTime     int64
}

func (o *OrderObject) GetOrderId() string {
	return o.orderId
}

func (o *OrderObject) GetSdkType() logintypes.SDKType {
	return o.sdkType
}

func (o *OrderObject) GetDevicePlatform() logintypes.DevicePlatformType {
	return o.devicePlatform
}

func (o *OrderObject) GetPlatformUserId() string {
	return o.platformUserId
}

func (o *OrderObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *OrderObject) GetPlayerLevel() int32 {
	return o.playerLevel
}

func (o *OrderObject) GetPlayerName() string {
	return o.playerName
}

func (o *OrderObject) GetChargeId() int32 {
	return o.chargeId
}

func (o *OrderObject) GetServerId() int32 {
	return o.serverId
}

func (o *OrderObject) GetStatus() types.OrderStatus {
	return o.status
}

func (o *OrderObject) SetStatus(status types.OrderStatus) {
	o.status = status
}

func (o *OrderObject) GetMoney() int32 {
	return o.money
}

func (o *OrderObject) GetGold() int32 {
	return o.gold
}

func (o *OrderObject) SetUpdateTime(updateTime int64) {
	o.updateTime = updateTime
}

func (o *OrderObject) FromEntity(e *OrderEntity) *OrderObject {

	o.id = e.Id
	o.orderId = e.OrderId
	o.sdkOrderId = e.SdkOrderId
	o.status = types.OrderStatus(e.Status)
	o.sdkType = logintypes.SDKType(e.SdkType)
	o.serverId = e.ServerId
	o.userId = e.UserId
	o.playerId = e.PlayerId
	o.playerLevel = e.PlayerLevel
	o.playerName = e.PlayerName
	o.chargeId = e.ChargeId
	o.platformUserId = e.PlatformUserId
	o.money = e.Money
	o.gold = e.Gold
	o.receivePayTime = e.ReceivePayTime
	o.createTime = e.CreateTime
	o.updateTime = e.UpdateTime
	o.deleteTime = e.DeleteTime
	return o
}

func (o *OrderObject) ToEntity() *OrderEntity {
	e := &OrderEntity{
		Id:             o.id,
		OrderId:        o.orderId,
		PlatformUserId: o.platformUserId,
		SdkOrderId:     o.sdkOrderId,
		Status:         int32(o.status),
		SdkType:        int32(o.sdkType),
		ServerId:       o.serverId,
		UserId:         o.userId,
		PlayerId:       o.playerId,
		PlayerLevel:    o.playerLevel,
		PlayerName:     o.playerName,
		ChargeId:       o.chargeId,
		Money:          o.money,
		Gold:           o.gold,
		ReceivePayTime: o.receivePayTime,
		CreateTime:     o.createTime,
		UpdateTime:     o.updateTime,
		DeleteTime:     o.deleteTime,
	}
	return e
}

func NewOrderObject() *OrderObject {
	obj := &OrderObject{}
	return obj
}
