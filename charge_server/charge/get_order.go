package charge

import (
	logintypes "fgame/fgame/account/login/types"
	"fmt"
)

type GetSDKOrderHandler interface {
	GetSDKOrder(devicePlatformType logintypes.DevicePlatformType, platformUserId string, productId int32, money int32, roleId int64, roleName string, serverId int32, orderId string) (notifyUrl string, thirdOrderId string, extension string, flag bool)
}

type GetSDKOrderHandlerFunc func(devicePlatformType logintypes.DevicePlatformType, platformUserId string, productId int32, money int32, roleId int64, roleName string, serverId int32, orderId string) (notifyUrl string, thirdOrderId string, extension string, flag bool)

func (f GetSDKOrderHandlerFunc) GetSDKOrder(devicePlatformType logintypes.DevicePlatformType, platformUserId string, productId int32, money int32, roleId int64, roleName string, serverId int32, orderId string) (notifyUrl string, thirdOrderId string, extension string, flag bool) {
	return f(devicePlatformType, platformUserId, productId, money, roleId, roleName, serverId, orderId)
}

var (
	getSDKOrderHandlerMap = make(map[logintypes.SDKType]GetSDKOrderHandler)
)

func RegisterSDKOrderHandler(sdkType logintypes.SDKType, h GetSDKOrderHandler) {
	_, ok := getSDKOrderHandlerMap[sdkType]

	if ok {
		panic(fmt.Errorf("重复注册第三方下单[%s],[%s]", sdkType.String()))
	}
	getSDKOrderHandlerMap[sdkType] = h
}

func GetGetSDKOrderHandler(sdkType logintypes.SDKType) GetSDKOrderHandler {
	h, ok := getSDKOrderHandlerMap[sdkType]
	if !ok {
		return nil
	}
	return h
}
