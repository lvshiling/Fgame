package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/chatproxy/sdk"
)

func init() {
	sdk.RegisterSDK((*TianShenConfig)(nil))

}

type TianShenSDKConfig struct {
	Key  string `json:"key"`
	Game int32  `json:"game"`
	Url  string `json:"url"`
}

type TianShenConfig struct {
	Ios     *TianShenSDKConfig `json:"ios"`
	Android *TianShenSDKConfig `json:"android"`
}

func (c *TianShenConfig) FileName() string {
	return "tianshen.json"
}

func (c *TianShenConfig) Platform() types.SDKType {
	return types.SDKTypeTianShen
}

func (c *TianShenConfig) GetKey(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Key
	}
	return c.Android.Key
}

func (c *TianShenConfig) GetGame(deviceType types.DevicePlatformType) int32 {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Game
	}
	return c.Android.Game
}

func (c *TianShenConfig) GetUrl(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Url
	}
	return c.Android.Url
}
