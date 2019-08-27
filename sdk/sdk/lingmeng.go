package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*LingMengConfig)(nil))

}

type LingMengSDKConfig struct {
	AppId  string `json:"appId"`
	AppKey string `json:"appKey"`
}

type LingMengConfig struct {
	IOS     *LingMengSDKConfig `json:"iOS"`
	Android *LingMengSDKConfig `json:"android"`
}

func (c *LingMengConfig) FileName() string {
	return "lingmeng.json"
}

func (c *LingMengConfig) Platform() types.SDKType {
	return types.SDKTypeLingMeng
}

func (c *LingMengConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *LingMengConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}
