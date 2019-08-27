package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*MoFangYouXiFengMoConfig)(nil))

}

type MoFangYouXiFengMoSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type MoFangYouXiFengMoConfig struct {
	IOS     *MoFangYouXiFengMoSDKConfig `json:"iOS"`
	Android *MoFangYouXiFengMoSDKConfig `json:"android"`
}

func (c *MoFangYouXiFengMoConfig) FileName() string {
	return "mofangyouxifengmo.json"
}

func (c *MoFangYouXiFengMoConfig) Platform() types.SDKType {
	return types.SDKTypeMoFangYouXiFengMo
}

func (c *MoFangYouXiFengMoConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *MoFangYouXiFengMoConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
