package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*YouMengConfig)(nil))

}

type YouMengSDKConfig struct {
	AppId     string `json:"appId"`
	AppKey    string `json:"appKey"`
	ChargeKey string `json:"chargeKey"`
}

type YouMengConfig struct {
	IOS     *YouMengSDKConfig `json:"iOS"`
	Android *YouMengSDKConfig `json:"android"`
}

func (c *YouMengConfig) FileName() string {
	return "youmeng.json"
}

func (c *YouMengConfig) Platform() types.SDKType {
	return types.SDKTypeYouMeng
}

func (c *YouMengConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *YouMengConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}

func (c *YouMengConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
