package bocai_test

import (
	"testing"

	. "fgame/fgame/charge_server/api/bocai"
)

var (
	payKey = "noz3ursqlfzbu7e6ffhnurw015w35cur6oj20n2suhlasmoxh5"
)

func TestSign(t *testing.T) {
	newForm := &BoCaiForm{}
	newForm.AsyxOrderId = "bc1812060344579823"
	newForm.Subject = "元宝"
	newForm.SubjectDesc = "100元宝"
	newForm.TradeStatus = "1"
	newForm.Amount = "10.00"
	newForm.Channel = "聚合微信"
	newForm.OrderCreatdt = "2018-12-06 15:57:44"
	newForm.OrderPaydt = "2018-12-06 15:58:19"
	newForm.AsyxGameId = "fd2c3b1b-326b-4f55-96a9-32cdb2d17995"
	newForm.PayOrderId = "201001201812061557449304217"
	newForm.Sign = "78e4d1ade88070d74d0fd2dae7b9e985"
	newForm.GameZero = "1"
	newForm.Memo = "24561cf8be0a4fcc8594cdb2b89064d0"
	newForm.Uid = "25435"
	sign := newForm.Sign
	newSign := GetSign(newForm, payKey)
	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}
}
