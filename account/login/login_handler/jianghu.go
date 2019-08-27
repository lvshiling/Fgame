package login_handler

import (
	"crypto/md5"
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeJiangHu, login.LoginHandlerFunc(handleJiangHu))
}

func handleJiangHu(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	jiangHu := csAccountLogin.GetJiangHu()
	if jiangHu == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆江湖手游,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆江湖手游,sdk配置为空")
		return
	}
	jiangHuConfig, ok := sdkConfig.(*sdksdk.JiangHuConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆江湖手游,sdk配置强制转换失败")
		return
	}

	appId := jiangHuConfig.GetAppId(devicePlatformType)
	gameId := jiangHuConfig.GetGameId(devicePlatformType)
	appKey := jiangHuConfig.GetAppKey(devicePlatformType)
	userName := jiangHu.GetUserName()
	loginTime := jiangHu.GetLoginTime()
	sign := jiangHu.GetSign()
	getSign := GetJiangHuSign(userName, appKey, loginTime)
	if getSign != sign {
		log.WithFields(
			log.Fields{
				"userName":  userName,
				"appId":     appId,
				"gameId":    gameId,
				"loginTime": loginTime,
				"sign":      sign,
				"getSign":   getSign,
			}).Warn("login:登陆江湖手游,签名错误")
		return
	}
	flag = true
	userId = userName
	returnPlatform = int32(platform)
	log.WithFields(
		log.Fields{
			"userName":  userName,
			"appId":     appId,
			"gameId":    gameId,
			"loginTime": loginTime,
			"sign":      sign,
		}).Info("login:登陆江湖手游,登陆成功")
	return
}

func GetJiangHuSign(userName, appKey, loginTime string) (sign string) {
	signMap := make(map[string]string)
	signMap["username"] = userName
	signMap["appkey"] = appKey
	signMap["logintime"] = loginTime
	keyList := []string{"username", "appkey", "logintime"}
	allStr := ""
	for _, key := range keyList {
		keyValue := fmt.Sprintf("%s=%s&", key, signMap[key])
		allStr += keyValue
	}
	if len(allStr) > 0 {
		allStr = allStr[:len(allStr)-1]
	}
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
