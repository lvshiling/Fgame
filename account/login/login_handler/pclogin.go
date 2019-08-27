package login_handler

import (
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypePC, login.LoginHandlerFunc(handlePCLogin))
}

func handlePCLogin(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	pcLoginData := csAccountLogin.GetPcLoginData()
	if pcLoginData == nil {
		log.WithFields(
			log.Fields{}).Warn("login:pc登陆,数据是空的")
		return
	}
	userId = pcLoginData.GetName()
	if len(userId) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:pc登陆,登录名为空")
		return
	}
	returnPlatform = csAccountLogin.GetPlatform()
	flag = true
	return
}
