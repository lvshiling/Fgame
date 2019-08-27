package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*TianXingConfig)(nil))

}

type TianXingSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type TianXingConfig struct {
	IOS     *TianXingSDKConfig `json:"iOS"`
	Android *TianXingSDKConfig `json:"android"`
}

func (c *TianXingConfig) FileName() string {
	return "tianxing.json"
}

func (c *TianXingConfig) Platform() types.SDKType {
	return types.SDKTypeTianXing
}

func (c *TianXingConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *TianXingConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
