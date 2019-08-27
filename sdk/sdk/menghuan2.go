package sdk

import (
	"fgame/fgame/account/login/types"
	"fgame/fgame/sdk"
)

func init() {
	sdk.RegisterSDK((*MengHuan2Config)(nil))

}

type MengHuan2SDKConfig struct {
	Key       string `json:"key"`
	ChargeKey string `json:"chargeKey"`
}

type MengHuan2Config struct {
	IOS     *MengHuan2SDKConfig `json:"iOS"`
	Android *MengHuan2SDKConfig `json:"android"`
}

func (c *MengHuan2Config) FileName() string {
	return "menghuan2.json"
}

func (c *MengHuan2Config) Platform() types.SDKType {
	return types.SDKTypeMengHuan2
}

func (c *MengHuan2Config) GetKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.Key
	case types.DevicePlatformTypeIOS:
		return c.IOS.Key
	default:
		return ""
	}
}

func (c *MengHuan2Config) GetChargeKey(devicePlatform types.DevicePlatformType) string {
	switch devicePlatform {
	case types.DevicePlatformTypeAndroid:
		return c.Android.ChargeKey
	case types.DevicePlatformTypeIOS:
		return c.IOS.ChargeKey
	default:
		return ""
	}
}
