package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*PiaoMiaoConfig)(nil))

}

type PiaoMiaoSDKConfig struct {
	GamgeId string `json:"gamgeId"`
	GameKey string `json:"gameKey"`
}

type PiaoMiaoConfig struct {
	IOS     *PiaoMiaoSDKConfig `json:"iOS"`
	Android *PiaoMiaoSDKConfig `json:"android"`
}

func (c *PiaoMiaoConfig) FileName() string {
	return "piaomiao.json"
}

func (c *PiaoMiaoConfig) Platform() types.SDKType {
	return types.SDKTypePiaoMiao
}

func (c *PiaoMiaoConfig) GetGamgeId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GamgeId
	case types.DevicePlatformTypeIOS:
		return c.IOS.GamgeId
	default:
		return ""
	}
}

func (c *PiaoMiaoConfig) GetGameKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GameKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.GameKey
	default:
		return ""
	}
}
