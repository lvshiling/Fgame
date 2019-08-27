package login_handler

import (
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeYouMeng, login.LoginHandlerFunc(handleYouMeng))
}

func handleYouMeng(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)

	youMeng := csAccountLogin.GetYouMeng()
	if youMeng == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆游梦江湖,数据是空的")
		return
	}
	flag = true
	platform := types.SDKType(csAccountLogin.GetPlatform())
	returnPlatform = int32(platform)
	userId = youMeng.GetUserid()

	log.WithFields(
		log.Fields{
			"userId": userId,
		}).Info("login:登陆游梦江湖,登陆成功")
	return
}
