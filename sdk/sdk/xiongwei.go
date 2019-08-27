package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*XiongWeiConfig)(nil))

}

type XiongWeiSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type XiongWeiConfig struct {
	IOS     *XiongWeiSDKConfig `json:"iOS"`
	Android *XiongWeiSDKConfig `json:"android"`
}

func (c *XiongWeiConfig) FileName() string {
	return "xiongwei.json"
}

func (c *XiongWeiConfig) Platform() types.SDKType {
	return types.SDKTypeXiongWei
}

func (c *XiongWeiConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *XiongWeiConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
