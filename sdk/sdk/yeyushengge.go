package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*YeYuShengGeConfig)(nil))

}

type YeYuShengGeSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type YeYuShengGeConfig struct {
	IOS     *YeYuShengGeSDKConfig `json:"iOS"`
	Android *YeYuShengGeSDKConfig `json:"android"`
}

func (c *YeYuShengGeConfig) FileName() string {
	return "yeyushengge.json"
}

func (c *YeYuShengGeConfig) Platform() types.SDKType {
	return types.SDKTypeYeYuShengGe
}

func (c *YeYuShengGeConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *YeYuShengGeConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
