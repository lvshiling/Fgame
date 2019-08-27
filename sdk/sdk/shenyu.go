package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*ShenYuConfig)(nil))

}

type ShenYuSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type ShenYuConfig struct {
	IOS     *ShenYuSDKConfig `json:"iOS"`
	Android *ShenYuSDKConfig `json:"android"`
}

func (c *ShenYuConfig) FileName() string {
	return "shenyu.json"
}

func (c *ShenYuConfig) Platform() types.SDKType {
	return types.SDKTypeShenYu
}

func (c *ShenYuConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *ShenYuConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
