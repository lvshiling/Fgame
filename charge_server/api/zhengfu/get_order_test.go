package zhengfu_test

import (
	logintypes "fgame/fgame/account/login/types"
	. "fgame/fgame/charge_server/api/zhengfu"
	"testing"
)

var (
	devicePlatformType = logintypes.DevicePlatformTypeAndroid
	platformUserId     = "13287483"
	productId          = int32(1)
	money              = int32(6)
	roleId             = int64(1)
	roleName           = "test"
	serverId           = int32(1)
	orderId            = "test"
	appSecret          = "c2cbe80f4e52e78d7dc7b3598892f0e0"
	privateKey         = "MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBANKSxrU/82w8Nwbue11F8INRvCnqhDn/O2Kmjj0jfx52r1IIIQX6pVrUqa87ay1qNt8X6dJKBZmIGxwqicr5EVga1zdsuNPOPb8wnqIqhpDm+QrUDRHn8ZuZjyxX8SghHNOzH4ryv8zLZV1hOchRn6N+u3VfCrjwuqJdFgqiTnprAgMBAAECgYAR+8EGp7CFVNsqN2HHxHpW7LsSJVonjdmngivxor9vfZlZeyI+3XoTuMfJFF0B4ulOwj8Q24uA4jPWgveDoyPM57G+3FK4XSujkDnNB4L9anmMvejRDA5fxQLFJV1UEz4tOC+2hXvR1RQrG9Lug4cakUv8+xFRtQ3RTBvmr8zAcQJBAPoDSP5HWCjRNbHN115/y3zcnWxX+WOPTh0PWBLBHuSrFt/PBbz75nu6AgbMkY1bYyaF8I6k6+yBMgEz6lJgxBkCQQDXnbSdTIpVVdEsmzhrJpvREhSa6MhXxPgewYC+vVlMcPwQymHA4xQ80y/PbLIC/fJLviC8KtqVd32I+QFE7mMjAkB4Nowqd/OT7MR8shUUgy4843duWP65OHa+0lnu6p0IJpvhEZIYxKaWZ2ICEusJpR+Prmd0ryghmB2LJoNNCOpBAkBR3juaDlnoFPGbckR1yu8W7zqLpx+K0+syIl70DYk+kRfkeDOtvYsNnVJl++uLX0kEoWhkihD896XewE1PEwTpAkASPMpHHxe4Mqer6TxQXFGdaHTsESYSfnV+kWSlftYegr113ltAaownaVpoSc26OhPxNSwafj7mfEeCRz/DCtP5"
	notifyUrl          = "http://test.fhgamers.com/api/charge/zhengfu/android"
)

func TestGetOrder(t *testing.T) {
	sdkOrderId, _, err := GetSdkOrderId(appSecret, privateKey, notifyUrl, orderId, platformUserId, productId, money, roleId, roleName, serverId)
	if err != nil {
		t.Fatalf("TestGetOrder:error[%s]", err.Error())
	}

	if sdkOrderId == "" {
		t.Fatalf("TestGetOrder:下单失败")
	}
	t.Logf("TestGetOrder:get order [%d]", sdkOrderId)
}
