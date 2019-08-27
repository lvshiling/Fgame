package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*PaiQuYouXiXianLuConfig)(nil))

}

type PaiQuYouXiXianLuSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type PaiQuYouXiXianLuConfig struct {
	IOS     *PaiQuYouXiXianLuSDKConfig `json:"iOS"`
	Android *PaiQuYouXiXianLuSDKConfig `json:"android"`
}

func (c *PaiQuYouXiXianLuConfig) FileName() string {
	return "paiquyouxixianlu.json"
}

func (c *PaiQuYouXiXianLuConfig) Platform() types.SDKType {
	return types.SDKTypePaiQuYouXiXianLu
}

func (c *PaiQuYouXiXianLuConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *PaiQuYouXiXianLuConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
