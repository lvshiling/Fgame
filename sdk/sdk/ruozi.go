package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*RuoZiConfig)(nil))

}

type RuoZiSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type RuoZiConfig struct {
	IOS     *RuoZiSDKConfig `json:"iOS"`
	Android *RuoZiSDKConfig `json:"android"`
}

func (c *RuoZiConfig) FileName() string {
	return "ruozi.json"
}

func (c *RuoZiConfig) Platform() types.SDKType {
	return types.SDKTypeRuoZi
}

func (c *RuoZiConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *RuoZiConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
