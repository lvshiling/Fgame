package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*QiLing2Config)(nil))

}

type QiLing2SDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type QiLing2Config struct {
	IOS     *QiLing2SDKConfig `json:"iOS"`
	Android *QiLing2SDKConfig `json:"android"`
}

func (c *QiLing2Config) FileName() string {
	return "qiling2.json"
}

func (c *QiLing2Config) Platform() types.SDKType {
	return types.SDKTypeQiLing2
}

func (c *QiLing2Config) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *QiLing2Config) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
