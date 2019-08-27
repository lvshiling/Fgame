package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*JuDuConfig)(nil))

}

type JuDuSDKConfig struct {
	AppId  string `json:"appId"`
	GameId string `json:"gameId"`
	Agent  string `json:"agent"`
	AppKey string `json:"appKey"`
}

type JuDuConfig struct {
	IOS     *JuDuSDKConfig `json:"iOS"`
	Android *JuDuSDKConfig `json:"android"`
}

func (c *JuDuConfig) FileName() string {
	return "judu.json"
}

func (c *JuDuConfig) Platform() types.SDKType {
	return types.SDKTypeJuDu
}

func (c *JuDuConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *JuDuConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}

func (c *JuDuConfig) GetGameId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GameId
	case types.DevicePlatformTypeIOS:
		return c.IOS.GameId
	default:
		return ""
	}
}
