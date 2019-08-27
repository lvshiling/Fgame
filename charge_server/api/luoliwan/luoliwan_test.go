package luoliwan_test

import (
	"testing"

	. "fgame/fgame/charge_server/api/luoliwan"
)

func TestSign(t *testing.T) {
	cpOrderIdStr := "1542281597174"
	orderIdStr := "sdk2_2018111510051991"
	appId := "110073"
	uid := "199900"
	amount := int32(10)
	receiveTime := int64(1542281613)
	key := "adadd0da96310fda8008510ec214d3b8"
	sign := "8eff6608205faee20cd34d50e4a8defe"
	// 8eff6608205faee20cd34d50e4a8defe
	newSign := GetLuoLiWanSign(cpOrderIdStr, orderIdStr, appId, uid, amount, receiveTime, key)
	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}
}
