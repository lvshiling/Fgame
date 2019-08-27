package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*ZhangJianConfig)(nil))

}

type ZhangJianSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type ZhangJianConfig struct {
	IOS     *ZhangJianSDKConfig `json:"iOS"`
	Android *ZhangJianSDKConfig `json:"android"`
}

func (c *ZhangJianConfig) FileName() string {
	return "zhangjian.json"
}

func (c *ZhangJianConfig) Platform() types.SDKType {
	return types.SDKTypeZhangJian
}

func (c *ZhangJianConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *ZhangJianConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
