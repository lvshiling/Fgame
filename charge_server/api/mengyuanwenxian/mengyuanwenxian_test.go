package mengyuanwenxian_test

import (
	"testing"

	. "fgame/fgame/charge_server/api/mengyuanwenxian"
)

func TestSign(t *testing.T) {
	newForm := &MengYuanWenXianRequest{}
	newForm.AppId = "6046"
	newForm.CpOrderId = "52a233016e1d4335888347ba37638b05"
	newForm.MemId = "3"
	newForm.OrderId = "15415759405791800030001"
	newForm.OrderStatus = "2"
	newForm.PayTime = "1541575940"
	newForm.ProductId = "85"
	newForm.ProductName = "100元宝"
	newForm.ProductPrice = "10.00"
	appKey := "c2fc43e5342e216a2246ae4e849fb3c3"
	sign := "e1c0a7a6e9f7c0c03b6222577640de59"

	// encodeName := "%E5%85%83%E5%AE%9D"
	// if encodeName != url.QueryEscape(newForm.ProductName) {
	// 	t.Fatalf("测试商品名加密失败,期望[%s],得到[%s]", encodeName, url.QueryEscape(newForm.ProductName))
	// }

	newSign := GetMengYuanWenXianSign(newForm, appKey)
	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}
}

func TestAndroidSign(t *testing.T) {
	newForm := &MengYuanWenXianRequest{}
	newForm.AppId = "6045"
	newForm.CpOrderId = "cac7eeaa66964d78bb94ad725948efd4"
	newForm.MemId = "2"
	newForm.OrderId = "15415758647923600020001"
	newForm.OrderStatus = "2"
	newForm.PayTime = "1541575864"
	newForm.ProductId = "85"
	newForm.ProductName = "100元宝"
	newForm.ProductPrice = "10.0"
	appKey := "a33fe56005f5db6bbaf885f8c658b1db"
	sign := "35d645e7f2f3327830fa8b5e830b166b"

	// encodeName := "%E5%85%83%E5%AE%9D"
	// if encodeName != url.QueryEscape(newForm.ProductName) {
	// 	t.Fatalf("测试商品名加密失败,期望[%s],得到[%s]", encodeName, url.QueryEscape(newForm.ProductName))
	// }

	newSign := GetMengYuanWenXianSign(newForm, appKey)
	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}
}
