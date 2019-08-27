package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*BoCaiConfig)(nil))

}

type BoCaiSDKConfig struct {
	CpId     string `json:"cpId"`
	ParamKey string `json:"paramKey"`
	GameId   string `json:"gameId"`
	PayKey   string `json:"payKey"`
}

type BoCaiConfig struct {
	IOS     *BoCaiSDKConfig `json:"iOS"`
	Android *BoCaiSDKConfig `json:"android"`
}

func (c *BoCaiConfig) FileName() string {
	return "bocai.json"
}

func (c *BoCaiConfig) Platform() types.SDKType {
	return types.SDKTypeBoCai
}

func (c *BoCaiConfig) GetCpId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.CpId
	case types.DevicePlatformTypeIOS:
		return c.IOS.CpId
	default:
		return ""
	}
}

func (c *BoCaiConfig) GetParamKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ParamKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ParamKey
	default:
		return ""
	}
}

func (c *BoCaiConfig) GetGameId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GameId
	case types.DevicePlatformTypeIOS:
		return c.IOS.GameId
	default:
		return ""
	}
}

func (c *BoCaiConfig) GetPayKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PayKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PayKey
	default:
		return ""
	}
}
