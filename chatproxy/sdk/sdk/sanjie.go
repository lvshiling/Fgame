package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/chatproxy/sdk"
)

func init() {
	sdk.RegisterSDK((*SanJieConfig)(nil))

}

type SanJieSDKConfig struct {
	Key  string `json:"key"`
	Game int32  `json:"game"`
	Url  string `json:"url"`
}

type SanJieConfig struct {
	Ios     *SanJieSDKConfig `json:"ios"`
	Android *SanJieSDKConfig `json:"android"`
}

func (c *SanJieConfig) FileName() string {
	return "sanjie.json"
}

func (c *SanJieConfig) Platform() types.SDKType {
	return types.SDKTypeSanJie
}

func (c *SanJieConfig) GetKey(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Key
	}
	return c.Android.Key
}

func (c *SanJieConfig) GetGame(deviceType types.DevicePlatformType) int32 {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Game
	}
	return c.Android.Game
}

func (c *SanJieConfig) GetUrl(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Url
	}
	return c.Android.Url
}
