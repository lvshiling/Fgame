package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*QiAConfig)(nil))

}

type QiASDKConfig struct {
	AppId  string `json:"appId"`
	GameId string `json:"gameId"`
	Agent  string `json:"agent"`
	AppKey string `json:"appKey"`
}

type QiAConfig struct {
	IOS     *QiASDKConfig `json:"iOS"`
	Android *QiASDKConfig `json:"android"`
}

func (c *QiAConfig) FileName() string {
	return "qia.json"
}

func (c *QiAConfig) Platform() types.SDKType {
	return types.SDKTypeQiA
}

func (c *QiAConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *QiAConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}

func (c *QiAConfig) GetGameId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GameId
	case types.DevicePlatformTypeIOS:
		return c.IOS.GameId
	default:
		return ""
	}
}
