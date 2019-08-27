package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*QiLing3Config)(nil))

}

type QiLing3SDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type QiLing3Config struct {
	IOS     *QiLing3SDKConfig `json:"iOS"`
	Android *QiLing3SDKConfig `json:"android"`
}

func (c *QiLing3Config) FileName() string {
	return "qiling3.json"
}

func (c *QiLing3Config) Platform() types.SDKType {
	return types.SDKTypeQiLing3
}

func (c *QiLing3Config) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *QiLing3Config) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
