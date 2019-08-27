package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*JiuLingConfig)(nil))

}

type JiuLingSDKConfig struct {
	AppKey string `json:"appKey"`
}

type JiuLingConfig struct {
	IOS     *JiuLingSDKConfig `json:"iOS"`
	Android *JiuLingSDKConfig `json:"android"`
}

func (c *JiuLingConfig) FileName() string {
	return "jiuling.json"
}

func (c *JiuLingConfig) Platform() types.SDKType {
	return types.SDKTypeJiuLing
}

func (c *JiuLingConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}
