package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*LuoLiWanConfig)(nil))

}

type LuoLiWanSDKConfig struct {
	ChargeKey string `json:"chargeKey"`
}

type LuoLiWanConfig struct {
	IOS     *LuoLiWanSDKConfig `json:"iOS"`
	Android *LuoLiWanSDKConfig `json:"android"`
}

func (c *LuoLiWanConfig) FileName() string {
	return "luoliwan.json"
}

func (c *LuoLiWanConfig) Platform() types.SDKType {
	return types.SDKTypeLuoLiWan
}

func (c *LuoLiWanConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
