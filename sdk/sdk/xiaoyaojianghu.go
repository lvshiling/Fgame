package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*XiaoYaoJiangHuConfig)(nil))

}

type XiaoYaoJiangHuSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type XiaoYaoJiangHuConfig struct {
	IOS     *XiaoYaoJiangHuSDKConfig `json:"iOS"`
	Android *XiaoYaoJiangHuSDKConfig `json:"android"`
}

func (c *XiaoYaoJiangHuConfig) FileName() string {
	return "xiaoyaojianghu.json"
}

func (c *XiaoYaoJiangHuConfig) Platform() types.SDKType {
	return types.SDKTypeXiaoYaoJiangHu
}

func (c *XiaoYaoJiangHuConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *XiaoYaoJiangHuConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
