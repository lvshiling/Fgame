package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*BaJiConfig)(nil))

}

type BaJiSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type BaJiConfig struct {
	IOS     *BaJiSDKConfig `json:"iOS"`
	Android *BaJiSDKConfig `json:"android"`
}

func (c *BaJiConfig) FileName() string {
	return "baji.json"
}

func (c *BaJiConfig) Platform() types.SDKType {
	return types.SDKTypeBaJi
}

func (c *BaJiConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *BaJiConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
