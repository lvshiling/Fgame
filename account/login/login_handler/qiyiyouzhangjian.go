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
	login.RegisterLogin(types.SDKTypeQiYiYouZhangJian, login.LoginHandlerFunc(handleQiYiYouZhangJian))
}

func handleQiYiYouZhangJian(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)

	qiYiYouZhangJian := csAccountLogin.GetQiYiYouZhangJian()
	if qiYiYouZhangJian == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆7亿游-仗剑凡尘,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆7亿游-仗剑凡尘,sdk配置为空")
		return
	}
	qiYiYouZhangJianConfig, ok := sdkConfig.(*sdksdk.QiYiYouZhangJianConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆7亿游-仗剑凡尘,sdk配置强制转换失败")
		return
	}
	key := qiYiYouZhangJianConfig.GetKey(devicePlatformType)
	timeStr := qiYiYouZhangJian.GetTime()
	userId = qiYiYouZhangJian.GetUserid()
	sign := qiYiYouZhangJian.GetSign()
	tmsg := qiYiYouZhangJian.GetMsg()
	//验证签名正确性
	getSign := QiYiYouZhangJianSign(key, tmsg, userId, timeStr)
	if getSign != sign {
		log.WithFields(
			log.Fields{
				"key":     key,
				"msg":     msg,
				"userId":  userId,
				"timeStr": timeStr,
				"sign":    sign,
				"getSign": getSign,
			}).Warn("login:登陆7亿游-仗剑凡尘,签名错误")
		return
	}
	returnPlatform = int32(platform)
	flag = true
	log.WithFields(
		log.Fields{
			"timeStr": timeStr,
			"userId":  userId,
			"sign":    sign,
			"tmsg":    tmsg,
		}).Info("login:登陆7亿游-仗剑凡尘,登陆成功")
	return
}

func QiYiYouZhangJianSign(key string, msg string, userid string, timeStr string) (sign string) {
	signMap := make(map[string]string)
	signMap["code"] = "success"
	signMap["msg"] = msg
	signMap["userid"] = userid
	signMap["time"] = timeStr
	signMap["key"] = key
	keyList := []string{"code", "msg", "userid", "time", "key"}
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
