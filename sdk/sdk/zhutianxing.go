package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*ZhuTianXingConfig)(nil))

}

type ZhuTianXingSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type ZhuTianXingConfig struct {
	IOS     *ZhuTianXingSDKConfig `json:"iOS"`
	Android *ZhuTianXingSDKConfig `json:"android"`
}

func (c *ZhuTianXingConfig) FileName() string {
	return "zhutianxing.json"
}

func (c *ZhuTianXingConfig) Platform() types.SDKType {
	return types.SDKTypeZhuTianXing
}

func (c *ZhuTianXingConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *ZhuTianXingConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
