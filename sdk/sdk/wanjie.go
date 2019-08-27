package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*WanJieConfig)(nil))

}

type WanJieSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type WanJieConfig struct {
	IOS     *WanJieSDKConfig `json:"iOS"`
	Android *WanJieSDKConfig `json:"android"`
}

func (c *WanJieConfig) FileName() string {
	return "wanjie.json"
}

func (c *WanJieConfig) Platform() types.SDKType {
	return types.SDKTypeWanJie
}

func (c *WanJieConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *WanJieConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
