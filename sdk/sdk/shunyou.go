package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*ShunYouConfig)(nil))

}

type ShunYouSDKConfig struct {
	AppId      string `json:"appId"`
	AppKey     string `json:"appKey"`
	AppSecret  string `json:"appSecret"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	NotifyUrl  string `json:"notifyUrl"`
}

type ShunYouConfig struct {
	IOS     *ShunYouSDKConfig `json:"iOS"`
	Android *ShunYouSDKConfig `json:"android"`
}

func (c *ShunYouConfig) FileName() string {
	return "shunyou.json"
}

func (c *ShunYouConfig) Platform() types.SDKType {
	return types.SDKTypeShunYou
}

func (c *ShunYouConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *ShunYouConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}

func (c *ShunYouConfig) GetSecretKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppSecret
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppSecret
	default:
		return ""
	}
}

func (c *ShunYouConfig) GetPublicKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PublicKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PublicKey
	default:
		return ""
	}
}

func (c *ShunYouConfig) GetPrivateKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PrivateKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PrivateKey
	default:
		return ""
	}
}

func (c *ShunYouConfig) GetNotifyUrl(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.NotifyUrl
	case types.DevicePlatformTypeIOS:
		return c.IOS.NotifyUrl
	default:
		return ""
	}
}
