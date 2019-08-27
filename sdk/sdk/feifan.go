package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*FeiFanConfig)(nil))

}

type FeiFanSDKConfig struct {
	AppId      string `json:"appId"`
	AppKey     string `json:"appKey"`
	AppSecret  string `json:"appSecret"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	NotifyUrl  string `json:"notifyUrl"`
}

type FeiFanConfig struct {
	IOS     *FeiFanSDKConfig `json:"iOS"`
	Android *FeiFanSDKConfig `json:"android"`
}

func (c *FeiFanConfig) FileName() string {
	return "feifan.json"
}

func (c *FeiFanConfig) Platform() types.SDKType {
	return types.SDKTypeFeiFan
}

func (c *FeiFanConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *FeiFanConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}

func (c *FeiFanConfig) GetSecretKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppSecret
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppSecret
	default:
		return ""
	}
}

func (c *FeiFanConfig) GetPublicKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PublicKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PublicKey
	default:
		return ""
	}
}

func (c *FeiFanConfig) GetPrivateKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PrivateKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PrivateKey
	default:
		return ""
	}
}

func (c *FeiFanConfig) GetNotifyUrl(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.NotifyUrl
	case types.DevicePlatformTypeIOS:
		return c.IOS.NotifyUrl
	default:
		return ""
	}
}
