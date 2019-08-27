package taifeng

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"

	log "github.com/Sirupsen/logrus"
)

func init() {
	charge.RegisterSDKOrderHandler(logintypes.SDKTypeTaiFeng, charge.GetSDKOrderHandlerFunc(GetSDKOrder))
}

//获取第三方下单
func GetSDKOrder(devicePlatformType logintypes.DevicePlatformType, platformUserId string, productId int32, money int32, roleId int64, roleName string, serverId int32, orderId string) (notifyUrl string, sdkOrderId string, extension string, flag bool) {

	sdkConfig := sdk.GetSdkService().GetSdkConfig(logintypes.SDKTypeTaiFeng)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{
				"platformUserId": platformUserId,
				"prodoctId":      productId,
				"money":          money,
				"roleId":         roleId,
				"roleName":       roleName,
				"serverId":       serverId,
				"orderId":        orderId,
			}).Warn("charge:泰逢下单,sdk配置为空")
		return
	}
	taiFengSDKConfig, ok := sdkConfig.(*sdksdk.TaiFengConfig)
	if !ok {
		log.WithFields(
			log.Fields{
				"platformUserId": platformUserId,
				"prodoctId":      productId,
				"money":          money,
				"roleId":         roleId,
				"roleName":       roleName,
				"serverId":       serverId,
				"orderId":        orderId,
			}).Warn("charge:泰逢下单,sdk配置强制转换失败")
		return
	}

	notifyUrl = taiFengSDKConfig.GetNotifyUrl(devicePlatformType)
	sdkOrderId = ""
	extension = ""
	flag = true
	return
}
