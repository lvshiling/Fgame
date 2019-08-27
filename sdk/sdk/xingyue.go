package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*XingYueConfig)(nil))

}

type XingYueSDKConfig struct {
	CpId     string `json:"cpId"`
	ParamKey string `json:"paramKey"`
	GameId   string `json:"gameId"`
	PayKey   string `json:"payKey"`
}

type XingYueConfig struct {
	IOS     *XingYueSDKConfig `json:"iOS"`
	Android *XingYueSDKConfig `json:"android"`
}

func (c *XingYueConfig) FileName() string {
	return "xingyue.json"
}

func (c *XingYueConfig) Platform() types.SDKType {
	return types.SDKTypeXingYue
}

func (c *XingYueConfig) GetCpId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.CpId
	case types.DevicePlatformTypeIOS:
		return c.IOS.CpId
	default:
		return ""
	}
}

func (c *XingYueConfig) GetParamKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ParamKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ParamKey
	default:
		return ""
	}
}

func (c *XingYueConfig) GetGameId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GameId
	case types.DevicePlatformTypeIOS:
		return c.IOS.GameId
	default:
		return ""
	}
}

func (c *XingYueConfig) GetPayKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.PayKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.PayKey
	default:
		return ""
	}
}
