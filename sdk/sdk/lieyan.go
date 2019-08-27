package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*LieYanConfig)(nil))

}

type LieYanSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type LieYanConfig struct {
	IOS     *LieYanSDKConfig `json:"iOS"`
	Android *LieYanSDKConfig `json:"android"`
}

func (c *LieYanConfig) FileName() string {
	return "lieyan.json"
}

func (c *LieYanConfig) Platform() types.SDKType {
	return types.SDKTypeLieYan
}

func (c *LieYanConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *LieYanConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
