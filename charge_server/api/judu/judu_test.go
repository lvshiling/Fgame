package judu_test

import (
	"testing"

	. "fgame/fgame/charge_server/api/judu"
)

var (
	juDuRequest = &JuDuRequest{
		OrderId:  "15411478504456567",
		UserName: "13808520991",
		GameId:   4,
		RoleId:   "851840410215335",
		ServerId: 1,
		PayType:  "zfb",
		Amount:   10,
		PayTime:  1541147852,
		Attache:  "036c9a3ca3104c74af10363bf5f656c7",
		Sign:     "3824d5c77eb3a7e2862a5e7382a03ad0",
	}
	appKey = "f1ecc3af9346d2e8161821ac15a5c420"
)

func TestSign(t *testing.T) {

	newSign := GetJuDuSign(juDuRequest, appKey)
	if juDuRequest.Sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", juDuRequest.Sign, newSign)
	}
}
