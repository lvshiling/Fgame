package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*TaiGuConfig)(nil))

}

type TaiGuSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type TaiGuConfig struct {
	IOS     *TaiGuSDKConfig `json:"iOS"`
	Android *TaiGuSDKConfig `json:"android"`
}

func (c *TaiGuConfig) FileName() string {
	return "taigu.json"
}

func (c *TaiGuConfig) Platform() types.SDKType {
	return types.SDKTypeTaiGu
}

func (c *TaiGuConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *TaiGuConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
