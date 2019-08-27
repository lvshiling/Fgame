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
	login.RegisterLogin(types.SDKTypeJiuLing, login.LoginHandlerFunc(handleJiuLing))
}

func handleJiuLing(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	jiuLing := csAccountLogin.GetJiuLing()
	if jiuLing == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆九零,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆九零,sdk配置为空")
		return
	}
	jiuLingConfig, ok := sdkConfig.(*sdksdk.JiuLingConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆九零,sdk配置强制转换失败")
		return
	}
	appKey := jiuLingConfig.GetAppKey(devicePlatformType)
	timeStr := jiuLing.GetLoginTime()
	userId = jiuLing.GetUserName()
	sign := jiuLing.GetSign()
	//验证签名正确性
	getSign := JiuLingSign(userId, appKey, timeStr)
	if getSign != sign {
		log.WithFields(
			log.Fields{
				"msg":     msg,
				"userId":  userId,
				"timeStr": timeStr,
				"sign":    sign,
				"getSign": getSign,
			}).Warn("login:登陆九零,签名错误")
		return
	}
	returnPlatform = int32(platform)
	flag = true
	log.WithFields(
		log.Fields{
			"timeStr": timeStr,
			"userId":  userId,
			"sign":    sign,
		}).Info("login:登陆九零,登陆成功")
	return
}

func JiuLingSign(userName string, appkey string, timeStr string) (sign string) {
	signMap := make(map[string]string)
	signMap["username"] = userName
	signMap["appkey"] = appkey
	signMap["logintime"] = timeStr
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
