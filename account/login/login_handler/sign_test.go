package login_handler_test

import (
	. "fgame/fgame/account/login/login_handler"
	"testing"
)

func TestFeiYangSign(t *testing.T) {
	appId := "1"
	appKey := "de933fdbede098c62cb309443c3cf251"
	userId := "23"
	userToken := "rkmi2huqu9dv6750g5os11ilv2"

	sign := "4753dce3ae736e7f894ebcc6cd3cff7a"
	newSign := GetFeiYangSign(appId, appKey, userId, userToken)

	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}

}

func TestXiXiYouSign(t *testing.T) {
	key := "0c260798f6dbbceda56a0a32c9aab7c3"
	msg := "766dca2a7a7aa85f73715947fc44c70d"
	userId := "10485195"
	timeStr := "1544105252"

	sign := "2a01592a91eecd0d147956407e9fdfa2"
	newSign := XiXiYouSign(key, msg, userId, timeStr)

	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}

}

func TestZuoWanSign(t *testing.T) {
	cpId := "90b6de55-ebc9-4418-9c81-1b1fa81ec041"
	sid := "b2951869-50b6-4c87-b34b-17aa8f2f9095"
	gameId := "37c55cb8-e82e-4dfe-b436-ac52e6bc38e8"
	loginToken := "40654D9D5D7747FEA2DF267FEF452152"
	paramKey := "88GFcDmncXfbMlvPjlNFaPBKf2fXkcOo9EKaGuyprJqljCGPXcnUTCSiwVpaon20"
	paramsign := "EC5D7517DB0BE78BC4984821C3FB35AF"
	newSign := GetZuoWanSign(cpId, gameId, sid, loginToken, paramKey)

	if paramsign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", paramsign, newSign)
	}

}

func TestQiLingSign(t *testing.T) {
	key := "5559dd9748f6b9cbe060561fb327a05a"
	msg := "9cd32239109283d9b6c3a822d56b23d0"
	userId := "26580_16410"
	timeStr := "1548229856"

	sign := "3d89768bd7299e9f9d1210985c2dbeb0"
	newSign := QiLingSign(key, msg, userId, timeStr)

	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}

}

func TestJiuLingSign(t *testing.T) {
	key := "e97f0c9486cb5286351f0ea32d9da29"
	useName := "wzy0010"
	loginTime := "1552376567"
	sign := "5ec2b18ee74fc3d5ecdfa4fb161304dd"

	newSign := JiuLingSign(useName, key, loginTime)

	if sign != newSign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, newSign)
	}

}
