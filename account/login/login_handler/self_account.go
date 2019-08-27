package login_handler

import (
	"context"
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"
	"time"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeSelf, login.LoginHandlerFunc(handleSelfAccount))
}

func handleSelfAccount(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	selfAccountData := csAccountLogin.GetSelfAccount()
	if selfAccountData == nil {
		log.WithFields(
			log.Fields{}).Warn("login:自己账户登陆,数据是空的")
		return
	}
	name := selfAccountData.GetName()
	if len(name) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:自己账户登陆登陆,登录名为空")
		return
	}
	password := selfAccountData.GetPassword()
	if len(password) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:自己账户登陆登陆,登录名为空")
		return
	}
	timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	//发送登陆
	returnPlatform, userId, err = login.GetLoginService().SelfLogin(timeoutCtx, name, password)
	if err != nil {
		return
	}
	if len(userId) <= 0 {
		return
	}
	cancel()
	flag = true
	return
}

var (
	rpcTimeout = time.Second * 10
)
