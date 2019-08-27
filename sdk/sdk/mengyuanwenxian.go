package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*FeiYangMengYuanWenXianConfig)(nil))

}

type FeiYangMengYuanWenXianSDKConfig struct {
	AppId  string `json:"appId"`
	AppKey string `json:"appKey"`
}

type FeiYangMengYuanWenXianConfig struct {
	IOS     *FeiYangSDKConfig `json:"iOS"`
	Android *FeiYangSDKConfig `json:"android"`
}

func (c *FeiYangMengYuanWenXianConfig) FileName() string {
	return "mengyuanwenxian.json"
}

func (c *FeiYangMengYuanWenXianConfig) Platform() types.SDKType {
	return types.SDKTypeMengYuanWenXian
}

func (c *FeiYangMengYuanWenXianConfig) GetAppId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppId
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppId
	default:
		return ""
	}
}

func (c *FeiYangMengYuanWenXianConfig) GetAppKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.AppKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.AppKey
	default:
		return ""
	}
}
