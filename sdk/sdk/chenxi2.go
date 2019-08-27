package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*ChenXi2Config)(nil))

}

type ChenXi2SDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type ChenXi2Config struct {
	IOS     *ChenXi2SDKConfig `json:"iOS"`
	Android *ChenXi2SDKConfig `json:"android"`
}

func (c *ChenXi2Config) FileName() string {
	return "chenxi2.json"
}

func (c *ChenXi2Config) Platform() types.SDKType {
	return types.SDKTypeChenXi2
}

func (c *ChenXi2Config) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *ChenXi2Config) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
