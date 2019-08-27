package taifeng_test

import (
	"fmt"
	"testing"

	. "fgame/fgame/charge_server/api/taifeng"
)

var (
	appId       = "10118"
	chlOrderNum = ""
	cpTradeId   = "75ab8c600186466c9ac0fa24f6eabd67"
	extra       = ""

	moneyType = "1"
	payResult = "1"
	payType   = "255"
	roleId    = "1134543915518435"
	serverId  = "2"
	sign      = "16c65b92905061d5a33c633c949a6abe"
	tfTradeNo = "20190221174405421020R4U66CF0"
	totalFee  = "1000"
	payKey    = "cde7ca411090eacb48e71681e44c4889"
)

func TestTaifengSign(t *testing.T) {
	getSign := TaifengSign(appId, tfTradeNo, chlOrderNum, extra, cpTradeId, roleId, serverId, moneyType, totalFee, payType, payResult, payKey)

	fmt.Println(getSign)
}
