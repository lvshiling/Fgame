package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*MengHuanConfig)(nil))

}

type MengHuanSDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type MengHuanConfig struct {
	IOS     *MengHuanSDKConfig `json:"iOS"`
	Android *MengHuanSDKConfig `json:"android"`
}

func (c *MengHuanConfig) FileName() string {
	return "menghuan.json"
}

func (c *MengHuanConfig) Platform() types.SDKType {
	return types.SDKTypeMengHuan
}

func (c *MengHuanConfig) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *MengHuanConfig) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
