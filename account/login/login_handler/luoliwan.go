package login_handler

import (
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeLuoLiWan, login.LoginHandlerFunc(handleLuoLiWan))
}

func handleLuoLiWan(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	luoLiWan := csAccountLogin.GetLuoLiWan()
	if luoLiWan == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆萝莉玩,数据是空的")
		return
	}

	flag = true
	userId = luoLiWan.GetUserId()
	returnPlatform = csAccountLogin.GetPlatform()
	log.WithFields(
		log.Fields{
			"userId": userId,
		}).Info("login:登陆萝莉玩,登陆成功")
	return
}
