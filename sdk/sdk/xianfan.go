package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*XianFanConfig)(nil))

}

type XianFanSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type XianFanConfig struct {
	IOS     *XianFanSDKConfig `json:"iOS"`
	Android *XianFanSDKConfig `json:"android"`
}

func (c *XianFanConfig) FileName() string {
	return "xianfan.json"
}

func (c *XianFanConfig) Platform() types.SDKType {
	return types.SDKTypeXianFan
}

func (c *XianFanConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *XianFanConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
