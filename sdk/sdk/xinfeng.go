package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*XinFengConfig)(nil))

}

type XinFengSDKConfig struct {
	ChargeKey string `json:"chargeKey"`
}

type XinFengConfig struct {
	IOS     *XinFengSDKConfig `json:"iOS"`
	Android *XinFengSDKConfig `json:"android"`
}

func (c *XinFengConfig) FileName() string {
	return "xinfeng.json"
}

func (c *XinFengConfig) Platform() types.SDKType {
	return types.SDKTypeXinFeng
}

func (c *XinFengConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
