package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*NiuChaYouFuTuConfig)(nil))

}

type NiuChaYouFuTuSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type NiuChaYouFuTuConfig struct {
	IOS     *NiuChaYouFuTuSDKConfig `json:"iOS"`
	Android *NiuChaYouFuTuSDKConfig `json:"android"`
}

func (c *NiuChaYouFuTuConfig) FileName() string {
	return "niuchayoufutu.json"
}

func (c *NiuChaYouFuTuConfig) Platform() types.SDKType {
	return types.SDKTypeNiuChaYouFuTu
}

func (c *NiuChaYouFuTuConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *NiuChaYouFuTuConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
