package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*AoTianConfig)(nil))

}

type AoTianSDKConfig struct {
	AppKey string `json:"appKey"`
}

type AoTianConfig struct {
	IOS     *AoTianSDKConfig `json:"iOS"`
	Android *AoTianSDKConfig `json:"android"`
}

func (c *AoTianConfig) FileName() string {
	return "aotian.json"
}

func (c *AoTianConfig) Platform() types.SDKType {
	return types.SDKTypeAoTian
}

func (c *AoTianConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}
