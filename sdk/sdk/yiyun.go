package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*YiYunConfig)(nil))

}

type YiYunSDKConfig struct {
	Key string `json:"key"`
}

type YiYunConfig struct {
	IOS     *YiYunSDKConfig `json:"iOS"`
	Android *YiYunSDKConfig `json:"android"`
}

func (c *YiYunConfig) FileName() string {
	return "yiyun.json"
}

func (c *YiYunConfig) Platform() types.SDKType {
	return types.SDKTypeYiYun
}

func (c *YiYunConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}
