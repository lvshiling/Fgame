package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*QiYiYouZhangJianConfig)(nil))

}

type QiYiYouZhangJianSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type QiYiYouZhangJianConfig struct {
	IOS     *QiYiYouZhangJianSDKConfig `json:"iOS"`
	Android *QiYiYouZhangJianSDKConfig `json:"android"`
}

func (c *QiYiYouZhangJianConfig) FileName() string {
	return "qiyiyouzhangjian.json"
}

func (c *QiYiYouZhangJianConfig) Platform() types.SDKType {
	return types.SDKTypeQiYiYouZhangJian
}

func (c *QiYiYouZhangJianConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *QiYiYouZhangJianConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
