package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*JiuTianConfig)(nil))

}

type JiuTianSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type JiuTianConfig struct {
	IOS     *JiuTianSDKConfig `json:"iOS"`
	Android *JiuTianSDKConfig `json:"android"`
}

func (c *JiuTianConfig) FileName() string {
	return "jiutian.json"
}

func (c *JiuTianConfig) Platform() types.SDKType {
	return types.SDKTypeJiuTian
}

func (c *JiuTianConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *JiuTianConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
