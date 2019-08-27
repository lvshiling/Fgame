package feifan_test

import (
	logintypes "fgame/fgame/account/login/types"
	. "fgame/fgame/charge_server/api/feifan"
	"testing"
)

var (
	devicePlatformType = logintypes.DevicePlatformTypeAndroid
	platformUserId     = "13282760"
	productId          = int32(1)
	money              = int32(6)
	roleId             = int64(1)
	roleName           = "test"
	serverId           = int32(1)
	orderId            = "test"
	appSecret          = "bbbeb1f3f0f168e1ccb9bdab9e1f1375"
	privateKey         = "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAIgn1wqjiVrcjYZD1tP5QUZMwRADo7mhv/qHYSjsPSa+Zx5yukiN3tGtVG9lkjSirwG9DzVW1I3OoIlkn64nedRs2W7Vo1oh0X0DSoGUOEnXe/Cman3KqiOfal5ZIDkulrGy7XPLGwxhdLGsemXaAD3NDllB5gWaIN2lA1d+7XInAgMBAAECgYArZc90E6YfMPdnGU5rKCJ3HtXWneJcs8K2PtpoKcxgAgZqPRVFNPsViBLGovBUGJqBilpDnRaI0Jh40nrXDrwU0adQbF6pfAR7hXJq3LF1MxMgsVMNWnU9RnRCMz/10VMHOJjsqwpd/lH/6UQJ9eotXLpFLKv2H6qTqsTOkc3kkQJBANxyZ08AnxLQch2Tyw3nBionmPx36ZL4/zLQGhO30kW5WhGBTFkDf7MhWaZH3WT1yOVbwuqek4O3ybyXtX2I5R8CQQCeHU0EX59qKLfuG/yTCNMq1N27FWNJz2qA5ywyESsTdBmJq4XS5VO8EFRLnJya+z6XSHuD0CGJz0ILcK2Ht4n5AkBV5GFqP8S7MOp1qcMhHJWjUSBjplkkwc21P64ZZrMQJaL5VRapTBqycdkbV77kenuXGgS9I6I4XSDGUZoOWotjAkBumcps+8KUTMVUXulPpMWp2Vr86doZIGi8oHhu0UmTgwv2HDAxNM9c5wNAHN4DHypKQp57ttQvBPaK8BfCrqVpAkEAkm1PjwufKbINGAxCvEhTqrOgOUSMQbNjK/807pL02ijB+TRDOydhdoPWkVUIqhMB3XH444vD3Qflfr2jmx3U0A=="
	notifyUrl          = "http://test.fhgamers.com/api/charge/feifan/android"
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
