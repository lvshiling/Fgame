package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*TianShenConfig)(nil))

}

type TianShenSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type TianShenConfig struct {
	IOS     *TianShenSDKConfig `json:"iOS"`
	Android *TianShenSDKConfig `json:"android"`
}

func (c *TianShenConfig) FileName() string {
	return "tianshen.json"
}

func (c *TianShenConfig) Platform() types.SDKType {
	return types.SDKTypeTianShen
}

func (c *TianShenConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *TianShenConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}

