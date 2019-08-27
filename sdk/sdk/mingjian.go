package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*MingJianConfig)(nil))

}

type MingJianSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type MingJianConfig struct {
	IOS     *MingJianSDKConfig `json:"iOS"`
	Android *MingJianSDKConfig `json:"android"`
}

func (c *MingJianConfig) FileName() string {
	return "mingjian.json"
}

func (c *MingJianConfig) Platform() types.SDKType {
	return types.SDKTypeMingJian
}

func (c *MingJianConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *MingJianConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
