package shunyou_test

import (
	logintypes "fgame/fgame/account/login/types"
	. "fgame/fgame/charge_server/api/shunyou"
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
	appSecret          = "0cb270853ff271615b2a90c31b95fe09"
	privateKey         = "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBALJj7Kl/HEbs1S/ApC3qJ7pZx46APgyeKrZ938F/OgT/7bLbZZZcorGAtwJbCSWquX3h4wavFEnipo5eiJEnJUp+fuhC1DuOo3pi6r3JzVUhQA1Jzwoos+GRNBLYRpZ8KSWIYQiF+IbUDUFB3y3xWjFtQze3V0kqdWK6da3jBUmvAgMBAAECgYACwTBCXcgeAEI6fosKencqlYBTXv+WSkr2jnMKFeDbeug8vs6Ox9drTkWFL8qwXjaHDxnmXIW/rlRMFoGdXDFjMm4C+hOGz0bTinKhwxvk5Pkh1F1+p4cXyKzH8Ojybv3STA8cz2Stq2f0JtjKRCbbVmH+h0yPlC6GtH+HfAN3wQJBAPaFxwSIsMeltRX9qQ2tEpmLcxW8vZEHvsUVZTmT1sly/kKermlOqMd4CilOa7QYjayEOJ2jTnWfeF6D0a9f74kCQQC5P5qk3wyZCzlXgcVenbq4xhtXqE1YfsbG6urqXOgwHgx9DgV+b+wxnGj6+7vuDCPx/YBKgKiO3UeyZjprkCl3AkEArYpZIpjzEWhWhQePVWBL4qknN9so+4qfQfAg1Rp8rk10LgO0tc84w0p+pLte2GYcfaCKlnYaynSbcLWNC88WOQJAb09DibuoozE2VFlakd6uuqX2+fXb+8e5gv7XBtmqfncfw+iv7mgsASddgSnPo1rSIm7TLnEeVzGpCg4ZHlayQwJBAJl/i1A19zUkcMiEw0xS8Gr66IC5d4G9jKNyAQZc8FYRIUkb+OxHFi45miuqsnmRVKJ8Bi9B9bjgmBJcoAz7glE="
	notifyUrl          = "http://test.fhgamers.com/api/charge/shunyou/android"
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
