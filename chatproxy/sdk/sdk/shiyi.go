package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/chatproxy/sdk"
)

func init() {
	sdk.RegisterSDK((*ShiYiConfig)(nil))

}

type ShiYiSDKConfig struct {
	Key  string `json:"key"`
	Game int32  `json:"game"`
	Url  string `json:"url"`
}

type ShiYiConfig struct {
	Ios     *ShiYiSDKConfig `json:"ios"`
	Android *ShiYiSDKConfig `json:"android"`
}

func (c *ShiYiConfig) FileName() string {
	return "shiyi.json"
}

func (c *ShiYiConfig) Platform() types.SDKType {
	return types.SDKTypeQiLing
}

func (c *ShiYiConfig) GetKey(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Key
	}
	return c.Android.Key
}

func (c *ShiYiConfig) GetGame(deviceType types.DevicePlatformType) int32 {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Game
	}
	return c.Android.Game
}

func (c *ShiYiConfig) GetUrl(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Url
	}
	return c.Android.Url
}
