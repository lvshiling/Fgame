package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*WenJianConfig)(nil))

}

type WenJianSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type WenJianConfig struct {
	IOS     *WenJianSDKConfig `json:"iOS"`
	Android *WenJianSDKConfig `json:"android"`
}

func (c *WenJianConfig) FileName() string {
	return "wenjian.json"
}

func (c *WenJianConfig) Platform() types.SDKType {
	return types.SDKTypeWenJian
}

func (c *WenJianConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *WenJianConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
