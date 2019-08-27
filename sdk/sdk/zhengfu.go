package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*ZhengFuConfig)(nil))

}

type ZhengFuSDKConfig struct {
	AppId      string `json:"appId"`
	AppKey     string `json:"appKey"`
	AppSecret  string `json:"appSecret"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	NotifyUrl  string `json:"notifyUrl"`
}

type ZhengFuConfig struct {
	IOS     *ZhengFuSDKConfig `json:"iOS"`
	Android *ZhengFuSDKConfig `json:"android"`
}

func (c *ZhengFuConfig) FileName() string {
	return "zhengfu.json"
}

func (c *ZhengFuConfig) Platform() types.SDKType {
	return types.SDKTypeZhengFu
}

func (c *ZhengFuConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *ZhengFuConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}

func (c *ZhengFuConfig) GetSecretKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppSecret
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppSecret
	default:
		return ""
	}
}

func (c *ZhengFuConfig) GetPublicKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PublicKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PublicKey
	default:
		return ""
	}
}

func (c *ZhengFuConfig) GetPrivateKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PrivateKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PrivateKey
	default:
		return ""
	}
}

func (c *ZhengFuConfig) GetNotifyUrl(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.NotifyUrl
	case types.DevicePlatformTypeIOS:
		return c.IOS.NotifyUrl
	default:
		return ""
	}
}
