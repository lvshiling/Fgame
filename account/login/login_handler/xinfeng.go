package login_handler

import (
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeXinFeng, login.LoginHandlerFunc(handleXinFeng))
}

func handleXinFeng(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	xinFeng := csAccountLogin.GetXinFeng()
	if xinFeng == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆新蜂,数据是空的")
		return
	}
	returnPlatform = csAccountLogin.GetPlatform()
	flag = true
	userId = xinFeng.GetUserId()
	log.WithFields(
		log.Fields{
			"userId": userId,
		}).Info("login:登陆新蜂,登陆成功")
	return
}
