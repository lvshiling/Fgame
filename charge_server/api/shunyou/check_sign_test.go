package shunyou_test

import (
	"fmt"
	"testing"

	. "fgame/fgame/charge_server/api/shunyou"
)

var (
	dataMap = map[string]string{
		"productID": "49",
		"orderID":   "1635776562840513334",
		"userID":    "13282760",
		"channelID": "3413",
		"gameID":    "608",
		"serverID":  "1",
		"money":     "100",
		"currency":  "RMB",
		"extension": "b6389481043d46e7aab9679d20e176f3",
		"signType":  "rsa",
		"sign":      "ArWgrio4RLDYcWaJF/6LNbg6GvtTAnMHSCtqRaAOJ6U9zYpXbQkDhK6TrLh9Wg7TiWaQRphpf2f0ue6GcUu+GAL+ndc0zH/dsCxj7G94WnO/0/zmZ7tnSBaBD6806rvqkPG0Fw1L10pEe9bJL+q+Sglh7cSSmoRCRKO23gNpDWs=",
	}
	secretKey = "0cb270853ff271615b2a90c31b95fe09"
)

func TestGetShunYouOriginalData(t *testing.T) {
	originData := GetShunYouOriginalData(dataMap, secretKey)
	fmt.Println(originData)
}
