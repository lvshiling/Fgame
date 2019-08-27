package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*HengGeWanConfig)(nil))

}

type HengGenWanSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type HengGeWanConfig struct {
	IOS     *HengGenWanSDKConfig `json:"iOS"`
	Android *HengGenWanSDKConfig `json:"android"`
}

func (c *HengGeWanConfig) FileName() string {
	return "henggewan.json"
}

func (c *HengGeWanConfig) Platform() types.SDKType {
	return types.SDKTypeHengGeWan
}

func (c *HengGeWanConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *HengGeWanConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
