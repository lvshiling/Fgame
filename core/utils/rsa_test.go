package utils_test

import "testing"

import . "fgame/fgame/core/utils"

var (
	originalData = "channelID=3413&currency=RMB&extension=b6389481043d46e7aab9679d20e176f3&gameID=599&money=100&orderID=1635776562840513334&productID=49&serverID=1&userID=13282760&bbbeb1f3f0f168e1ccb9bdab9e1f1375"
	publicKey    = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCIJ9cKo4la3I2GQ9bT+UFGTMEQA6O5ob/6h2Eo7D0mvmcecrpIjd7RrVRvZZI0oq8BvQ81VtSNzqCJZJ+uJ3nUbNlu1aNaIdF9A0qBlDhJ13vwpmp9yqojn2peWSA5Lpaxsu1zyxsMYXSxrHpl2gA9zQ5ZQeYFmiDdpQNXfu1yJwIDAQAB"
	sign         = "ArWgrio4RLDYcWaJF/6LNbg6GvtTAnMHSCtqRaAOJ6U9zYpXbQkDhK6TrLh9Wg7TiWaQRphpf2f0ue6GcUu+GAL+ndc0zH/dsCxj7G94WnO/0/zmZ7tnSBaBD6806rvqkPG0Fw1L10pEe9bJL+q+Sglh7cSSmoRCRKO23gNpDWs="
)

func TestCheckRsa(t *testing.T) {
	err := CheckRsaSign(publicKey, sign, originalData)
	if err != nil {
		t.Fatalf("获得错误[%s]", err.Error())
	}
}
