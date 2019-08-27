package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*XiXiYouConfig)(nil))

}

type XiXiYouSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type XiXiYouConfig struct {
	IOS     *XiXiYouSDKConfig `json:"iOS"`
	Android *XiXiYouSDKConfig `json:"android"`
}

func (c *XiXiYouConfig) FileName() string {
	return "xixiyou.json"
}

func (c *XiXiYouConfig) Platform() types.SDKType {
	return types.SDKTypeXiXiYou
}

func (c *XiXiYouConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *XiXiYouConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
