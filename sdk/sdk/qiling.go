package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*QiLingConfig)(nil))

}

type QiLingSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type QiLingConfig struct {
	IOS     *QiLingSDKConfig `json:"iOS"`
	Android *QiLingSDKConfig `json:"android"`
}

func (c *QiLingConfig) FileName() string {
	return "qiling.json"
}

func (c *QiLingConfig) Platform() types.SDKType {
	return types.SDKTypeQiLing
}

func (c *QiLingConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *QiLingConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
