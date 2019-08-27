package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*ZuoWanConfig)(nil))

}

type ZuoWanSDKConfig struct {
	CpId       string `json:"cpId"`
	GameInstId string `json:"gameinstId"`
	Privatekey string `json:"privateKey"`
}

type ZuoWanConfig struct {
	IOS     *ZuoWanSDKConfig `json:"iOS"`
	Android *ZuoWanSDKConfig `json:"android"`
}

func (c *ZuoWanConfig) FileName() string {
	return "zuowan.json"
}

func (c *ZuoWanConfig) Platform() types.SDKType {
	return types.SDKTypeZuoWan
}

func (c *ZuoWanConfig) GetCpId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.CpId
	case types.DevicePlatformTypeIOS:
		return c.IOS.CpId
	default:
		return ""
	}
}

func (c *ZuoWanConfig) GetGameInstId(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.GameInstId
	case types.DevicePlatformTypeIOS:
		return c.IOS.GameInstId
	default:
		return ""
	}
}

func (c *ZuoWanConfig) GetPrivatekey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Privatekey
	case types.DevicePlatformTypeIOS:
		return c.IOS.Privatekey
	default:
		return ""
	}
}
