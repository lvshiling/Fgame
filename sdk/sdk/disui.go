package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*DiSuiConfig)(nil))

}

type DiSuiSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type DiSuiConfig struct {
	IOS     *DiSuiSDKConfig `json:"iOS"`
	Android *DiSuiSDKConfig `json:"android"`
}

func (c *DiSuiConfig) FileName() string {
	return "disui.json"
}

func (c *DiSuiConfig) Platform() types.SDKType {
	return types.SDKTypeDiSui
}

func (c *DiSuiConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *DiSuiConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
