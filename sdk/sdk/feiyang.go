package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*FeiYangConfig)(nil))

}

type FeiYangSDKConfig struct {
	AppId  string `json:"appId"`
	AppKey string `json:"appKey"`
}

type FeiYangConfig struct {
	IOS     *FeiYangSDKConfig `json:"iOS"`
	Android *FeiYangSDKConfig `json:"android"`
}

func (c *FeiYangConfig) FileName() string {
	return "feiyang.json"
}

func (c *FeiYangConfig) Platform() types.SDKType {
	return types.SDKTypeFeiYang
}

func (c *FeiYangConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *FeiYangConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}
