package feiyang_test

import (
	"net/url"
	"testing"

	. "fgame/fgame/charge_server/api/feiyang"
)

func TestSign(t *testing.T) {
	newForm := &FeiYangRequest{}
	newForm.AppId = "1"
	newForm.CpOrderId = "20161028111"
	newForm.MemId = ""
	newForm.OrderId = "14794504894304304120001"
	newForm.OrderStatus = "2"
	newForm.PayTime = "1479450489"
	newForm.ProductId = "1"
	newForm.ProductName = "元宝"
	newForm.ProductPrice = "1"
	appKey := "f875364690581668449d4cf0aeb60560"
	sign := "3eaacb162b1f0fa12ad29dcd8e48ac1b"

	encodeName := "%E5%85%83%E5%AE%9D"
	if encodeName != url.QueryEscape(newForm.ProductName) {
		t.Fatalf("测试商品名加密失败,期望[%s],得到[%s]", encodeName, url.QueryEscape(newForm.ProductName))
	}

	newSign := GetFeiYangSign(newForm, appKey)
	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}
}
