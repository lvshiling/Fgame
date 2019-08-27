package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*TaiFengConfig)(nil))

}

type TaiFengSDKConfig struct {
	AppId     string `json:"appId"`
	AppKey    string `json:"appKey"`
	PayKey    string `json:"payKey"`
	NotifyUrl string `json:"notifyUrl"`
}

type TaiFengConfig struct {
	IOS     *TaiFengSDKConfig `json:"iOS"`
	Android *TaiFengSDKConfig `json:"android"`
}

func (c *TaiFengConfig) FileName() string {
	return "taifeng.json"
}

func (c *TaiFengConfig) Platform() types.SDKType {
	return types.SDKTypeTaiFeng
}

func (c *TaiFengConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *TaiFengConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}

func (c *TaiFengConfig) GetPayKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PayKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PayKey
	default:
		return ""
	}
}
func (c *TaiFengConfig) GetNotifyUrl(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.NotifyUrl
	case types.DevicePlatformTypeIOS:
		return c.IOS.NotifyUrl
	default:
		return ""
	}
}
