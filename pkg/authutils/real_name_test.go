package authutils_test

import "testing"
import . "fgame/fgame/pkg/authutils"

var (
	appKey = "bee25f245320743d5b5cec63e7070e87"
)
var (
	wrongName   = "张荣昌"
	wrongIdCard = "350582198811100513"
)
var (
	name   = "张荣昌"
	idCard = "350582198811100514"
)

func TestWrongJuHeIdCardQuery(t *testing.T) {
	flag, err := JuHeIdCardQuery(appKey, wrongName, wrongIdCard)
	if err != nil {
		t.Fatalf("id card query get error %s", err.Error())
	}
	if flag {
		t.Fatalf("id card query expect wrong but get right")
	}
}

func TestJuHeIdCardQuery(t *testing.T) {
	flag, err := JuHeIdCardQuery(appKey, name, idCard)
	if err != nil {
		t.Fatalf("id card query get error %s", err.Error())
	}
	if !flag {
		t.Fatalf("id card query expect right but get wrong")
	}
}
