package piaomiao_test

import (
	"testing"

	. "fgame/fgame/charge_server/api/piaomiao"
)

var (
	outTradeNo = "SP_20190516210157xB5W"
	price      = float32(98)
	payStatus  = int32(1)
	extend     = "6ad8cec522e346899b31b3b0d719985b"
	gameKey    = "26eb145c9ee33ba116f3413e1789b7ee"
	sign       = "60123536f71f19c160660f97d5ad1ad0"
)

func TestPiaoMiaoSign(t *testing.T) {
	newForm := &PiaoMiaoRequest{}
	newForm.OutTradeNo = outTradeNo
	newForm.Price = price
	newForm.PayStatus = payStatus
	newForm.Extend = extend

	newSign := GetPiaoMiaoSign(newForm, gameKey)
	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}
}
