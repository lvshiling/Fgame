package charge

import logintypes "fgame/fgame/account/login/types"

type GetOrderFinishEventData struct {
	sdkType        logintypes.SDKType
	deviceType     logintypes.DevicePlatformType
	notifyUrl      string
	serverId       int32
	platformUserId string
	userId         int64
	playerId       int64
	chargeId       int32
	money          int32
	name           string
	orderId        string
	sdkOrderId     string
	extension      string
}

func (d *GetOrderFinishEventData) GetNotifyUrl() string {
	return d.notifyUrl
}

func (d *GetOrderFinishEventData) GetServerId() int32 {
	return d.serverId
}
func (d *GetOrderFinishEventData) GetPlatformUserId() string {
	return d.platformUserId
}

func (d *GetOrderFinishEventData) GetUserId() int64 {
	return d.userId
}

func (d *GetOrderFinishEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *GetOrderFinishEventData) GetMoney() int32 {
	return d.money
}

func (d *GetOrderFinishEventData) GetChargeId() int32 {
	return d.chargeId
}

func (d *GetOrderFinishEventData) GetName() string {
	return d.name
}

func (d *GetOrderFinishEventData) GetOrderId() string {
	return d.orderId
}

func (d *GetOrderFinishEventData) GetSdkOrderId() string {
	return d.sdkOrderId
}

func (d *GetOrderFinishEventData) GetExtension() string {
	return d.extension
}

func CreateGetOrderFinishEventData(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, notifyUrl string, serverId int32, platformUserId string, userId int64, playerId int64, chargeId int32, money int32, name string, orderId string, sdkOrderId string, extension string) *GetOrderFinishEventData {
	d := &GetOrderFinishEventData{
		sdkType:        sdkType,
		deviceType:     deviceType,
		notifyUrl:      notifyUrl,
		serverId:       serverId,
		platformUserId: platformUserId,
		userId:         userId,
		playerId:       playerId,
		chargeId:       chargeId,
		money:          money,
		name:           name,
		orderId:        orderId,
		sdkOrderId:     sdkOrderId,
		extension:      extension,
	}
	return d
}

type OrderChargeEventData struct {
	orderId  string
	chargeId int32
}

func (d *OrderChargeEventData) GetChargeId() int32 {
	return d.chargeId
}

func (d *OrderChargeEventData) GetOrderId() string {
	return d.orderId
}

func CreateOrderChargeEventData(chargeId int32, orderId string) *OrderChargeEventData {
	eventData := &OrderChargeEventData{
		chargeId: chargeId,
		orderId:  orderId,
	}
	return eventData
}
