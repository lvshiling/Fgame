package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/chatproxy/sdk"
)

func init() {
	sdk.RegisterSDK((*WenJianConfig)(nil))

}

type WenJianSDKConfig struct {
	Key  string `json:"key"`
	Game int32  `json:"game"`
	Url  string `json:"url"`
}

type WenJianConfig struct {
	Ios     *WenJianSDKConfig `json:"ios"`
	Android *WenJianSDKConfig `json:"android"`
}

func (c *WenJianConfig) FileName() string {
	return "wenjian.json"
}

func (c *WenJianConfig) Platform() types.SDKType {
	return types.SDKTypeWenJian
}

func (c *WenJianConfig) GetKey(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Key
	}
	return c.Android.Key
}

func (c *WenJianConfig) GetGame(deviceType types.DevicePlatformType) int32 {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Game
	}
	return c.Android.Game
}

func (c *WenJianConfig) GetUrl(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Url
	}
	return c.Android.Url
}
