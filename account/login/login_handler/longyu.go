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
	login.RegisterLogin(types.SDKTypeLongYu, login.LoginHandlerFunc(handleLongYu))
}

func handleLongYu(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)

	longYu := csAccountLogin.GetLongYu()
	if longYu == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆龙语大陆,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆龙语大陆,sdk配置为空")
		return
	}
	longYuConfig, ok := sdkConfig.(*sdksdk.LongYuConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆龙语大陆,sdk配置强制转换失败")
		return
	}
	key := longYuConfig.GetAppKey(devicePlatformType)
	timeStr := longYu.GetLogintime()
	userId = longYu.GetUsername()
	sign := longYu.GetSign()
	//验证签名正确性
	getSign := LongYuSign(key, userId, timeStr)
	if getSign != sign {
		log.WithFields(
			log.Fields{
				"key":       key,
				"username":  userId,
				"logintime": timeStr,
				"sign":      sign,
				"getSign":   getSign,
			}).Warn("login:登陆龙语大陆,签名错误")
		return
	}
	returnPlatform = int32(platform)
	flag = true
	log.WithFields(
		log.Fields{
			"key":       key,
			"username":  userId,
			"logintime": timeStr,
			"sign":      sign,
		}).Info("login:登陆龙语大陆,登陆成功")
	return
}

func LongYuSign(key string, userid string, timeStr string) (sign string) {
	signMap := make(map[string]string)
	signMap["username"] = userid
	signMap["logintime"] = timeStr
	signMap["appkey"] = key
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
