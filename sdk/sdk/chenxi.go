package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*ChenXiConfig)(nil))

}

type ChenXiSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type ChenXiConfig struct {
	IOS     *ChenXiSDKConfig `json:"iOS"`
	Android *ChenXiSDKConfig `json:"android"`
}

func (c *ChenXiConfig) FileName() string {
	return "chenxi.json"
}

func (c *ChenXiConfig) Platform() types.SDKType {
	return types.SDKTypeChenXi
}

func (c *ChenXiConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *ChenXiConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
