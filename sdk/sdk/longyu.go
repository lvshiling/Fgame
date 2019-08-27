package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*LongYuConfig)(nil))

}

type LongYuSDKConfig struct {
	AppKey string `json:"appKey"`
}

type LongYuConfig struct {
	IOS     *LongYuSDKConfig `json:"iOS"`
	Android *LongYuSDKConfig `json:"android"`
}

func (c *LongYuConfig) FileName() string {
	return "longyu.json"
}

func (c *LongYuConfig) Platform() types.SDKType {
	return types.SDKTypeLongYu
}

func (c *LongYuConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}
