package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*JiuMengConfig)(nil))

}

type JiuMengSDKConfig struct {
	GameId string `json:"gameId"`
	AppId  string `json:"appId"`
	AppKey string `json:"appKey"`
}

type JiuMengConfig struct {
	IOS     *JiuMengSDKConfig `json:"iOS"`
	Android *JiuMengSDKConfig `json:"android"`
}

func (c *JiuMengConfig) FileName() string {
	return "jiumeng.json"
}

func (c *JiuMengConfig) Platform() types.SDKType {
	return types.SDKTypeJiuMeng
}

func (c *JiuMengConfig) GetGameId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GameId
	case types.DevicePlatformTypeIOS:
		return c.IOS.GameId
	default:
		return ""
	}
}

func (c *JiuMengConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *JiuMengConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}
