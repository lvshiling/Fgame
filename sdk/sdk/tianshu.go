package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*TianShuConfig)(nil))

}

type TianShuSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type TianShuConfig struct {
	IOS     *TianShuSDKConfig `json:"iOS"`
	Android *TianShuSDKConfig `json:"android"`
}

func (c *TianShuConfig) FileName() string {
	return "tianshu.json"
}

func (c *TianShuConfig) Platform() types.SDKType {
	return types.SDKTypeSanJie
}

func (c *TianShuConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *TianShuConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
