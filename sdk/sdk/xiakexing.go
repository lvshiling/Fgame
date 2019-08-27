package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*XiaKeXingConfig)(nil))

}

type XiaKeXingSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type XiaKeXingConfig struct {
	IOS     *XiaKeXingSDKConfig `json:"iOS"`
	Android *XiaKeXingSDKConfig `json:"android"`
}

func (c *XiaKeXingConfig) FileName() string {
	return "xiakexing.json"
}

func (c *XiaKeXingConfig) Platform() types.SDKType {
	return types.SDKTypeXiaKeXing
}

func (c *XiaKeXingConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *XiaKeXingConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
