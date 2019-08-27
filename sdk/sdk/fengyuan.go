package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*FengYuanConfig)(nil))

}

type FengYuanSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type FengYuanConfig struct {
	IOS     *FengYuanSDKConfig `json:"iOS"`
	Android *FengYuanSDKConfig `json:"android"`
}

func (c *FengYuanConfig) FileName() string {
	return "fengyuan.json"
}

func (c *FengYuanConfig) Platform() types.SDKType {
	return types.SDKTypeFengYuan
}

func (c *FengYuanConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *FengYuanConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
