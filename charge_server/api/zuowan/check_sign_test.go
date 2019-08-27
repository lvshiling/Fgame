package zuowan_test

import (
	"testing"

	. "fgame/fgame/charge_server/api/zuowan"
)

var (
	dataMap = map[string]string{
		"amount":                    "10.00",
		"gamecallbackdeliveryvalue": "183086c772354dc8980402ec8252f5c6",
		"gameid":                    "b2951869-50b6-4c87-b34b-17aa8f2f9095",
		"ordgoods":                  "100元宝",
		"ordgoodsdescription":       "100元宝",
		"ordno":                     "3004404012201812260932378111266",
		"ordtime":                   "2018-12-26 21:37:32",
		"paytime":                   "2018-12-26 21:37:43",
		"paytradeid":                "H1812263972247AM",
		"paymenttype":               "4",
		"tradestatus":               "1",
		"sid":                       "37c55cb8-e82e-4dfe-b436-ac52e6bc38e8",
	}
	sign = "0A22BE004357E0B63D5CD8BEB8E8C928"

	secretKey = "88GFcDmncXfbMlvPjlNFaPBKf2fXkcOo9EKaGuyprJqljCGPXcnUTCSiwVpaon20"
)

func TestZuoWanSign(t *testing.T) {
	getSign := GetZuoWanSign(dataMap, secretKey)
	if getSign != sign {
		t.Fatalf("测试失败,期望[%s],得到[%s]", sign, getSign)
	}
}
