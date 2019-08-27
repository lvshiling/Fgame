package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*XiaoYaoConfig)(nil))

}

type XiaoYaoSDKConfig struct {
	GameId    string `json:"gameId"`
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"SecretKey"`
}

type XiaoYaoConfig struct {
	IOS     *XiaoYaoSDKConfig `json:"iOS"`
	Android *XiaoYaoSDKConfig `json:"android"`
}

func (c *XiaoYaoConfig) FileName() string {
	return "xiaoyao.json"
}

func (c *XiaoYaoConfig) Platform() types.SDKType {
	return types.SDKTypeXiaoYao
}

func (c *XiaoYaoConfig) GetApiKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ApiKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ApiKey
	default:
		return ""
	}
}

func (c *XiaoYaoConfig) GetSecretKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.SecretKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.SecretKey
	default:
		return ""
	}
}
