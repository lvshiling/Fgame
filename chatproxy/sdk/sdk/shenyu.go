package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/chatproxy/sdk"
)

func init() {
	sdk.RegisterSDK((*ShenYuConfig)(nil))

}

type ShenYuSDKConfig struct {
	Key  string `json:"key"`
	Game int32  `json:"game"`
	Url  string `json:"url"`
}

type ShenYuConfig struct {
	Ios     *ShenYuSDKConfig `json:"ios"`
	Android *ShenYuSDKConfig `json:"android"`
}

func (c *ShenYuConfig) FileName() string {
	return "shenyu.json"
}

func (c *ShenYuConfig) Platform() types.SDKType {
	return types.SDKTypeShenYu
}

func (c *ShenYuConfig) GetKey(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Key
	}
	return c.Android.Key
}

func (c *ShenYuConfig) GetGame(deviceType types.DevicePlatformType) int32 {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Game
	}
	return c.Android.Game
}

func (c *ShenYuConfig) GetUrl(deviceType types.DevicePlatformType) string {
	if deviceType == types.DevicePlatformTypeIOS {
		return c.Ios.Url
	}
	return c.Android.Url
}
