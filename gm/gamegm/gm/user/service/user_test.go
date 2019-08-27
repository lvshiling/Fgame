package service_test

import (
	"fmt"
	user "fgame/fgame/gm/gamegm/gm/user/service"
	testInst "fgame/fgame/gm/gamegm/test"
	"testing"
)

func TestGetUserList(t *testing.T) {
	db := testInst.NewQiPaiDB()
	service := user.NewGmUserService(db)
	if service == nil {
		t.Fail()
		return
	}
	rst, err := service.GetUserList("super", -1, 1)
	if err != nil {
		t.Fail()
	}
	fmt.Println("length:", len(rst))
	t.Fail()
}
