package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*TianJiConfig)(nil))

}

type TianJiSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type TianJiConfig struct {
	IOS     *TianJiSDKConfig `json:"iOS"`
	Android *TianJiSDKConfig `json:"android"`
}

func (c *TianJiConfig) FileName() string {
	return "tianji.json"
}

func (c *TianJiConfig) Platform() types.SDKType {
	return types.SDKTypeTianJi
}

func (c *TianJiConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *TianJiConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
