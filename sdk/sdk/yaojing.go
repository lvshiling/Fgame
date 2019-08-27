package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*YaoJingConfig)(nil))

}

type YaoJingSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type YaoJingConfig struct {
	IOS     *YaoJingSDKConfig `json:"iOS"`
	Android *YaoJingSDKConfig `json:"android"`
}

func (c *YaoJingConfig) FileName() string {
	return "yaojing.json"
}

func (c *YaoJingConfig) Platform() types.SDKType {
	return types.SDKTypeYaoJing
}

func (c *YaoJingConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *YaoJingConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
