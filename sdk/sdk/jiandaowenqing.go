package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*JianDaoConfig)(nil))
}

type JianDaoSDKConfig struct {
	AppId  string `json:"appId"`
	AppKey string `json:"appKey"`
}

type JianDaoConfig struct {
	IOS     *JianDaoSDKConfig `json:"iOS"`
	Android *JianDaoSDKConfig `json:"android"`
}

func (c *JianDaoConfig) FileName() string {
	return "jiandaowenqing.json"
}

func (c *JianDaoConfig) Platform() types.SDKType {
	return types.SDKTypeJianDao
}

func (c *JianDaoConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *JianDaoConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}
