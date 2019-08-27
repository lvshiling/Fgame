package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*JiangHuConfig)(nil))

}

type JiangHuSDKConfig struct {
	AppId  string `json:"appId"`
	GameId string `json:"gameId"`
	Agent  string `json:"agent"`
	AppKey string `json:"appKey"`
}

type JiangHuConfig struct {
	IOS     *JiangHuSDKConfig `json:"iOS"`
	Android *JiangHuSDKConfig `json:"android"`
}

func (c *JiangHuConfig) FileName() string {
	return "jianghu.json"
}

func (c *JiangHuConfig) Platform() types.SDKType {
	return types.SDKTypeJiangHu
}

func (c *JiangHuConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *JiangHuConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}

func (c *JiangHuConfig) GetGameId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GameId
	case types.DevicePlatformTypeIOS:
		return c.IOS.GameId
	default:
		return ""
	}
}
