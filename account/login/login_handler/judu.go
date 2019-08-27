package login_handler

import (
	"crypto/md5"
	"encoding/json"
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/sdk"
	sdksdk "fgame/fgame/sdk/sdk"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypeJuDu, login.LoginHandlerFunc(handleJuDu))
}

func handleJuDu(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	judu := csAccountLogin.GetJuDu()
	if judu == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆巨都手游,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆巨都手游,sdk配置为空")
		return
	}
	juDuConfig, ok := sdkConfig.(*sdksdk.JuDuConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆巨都手游,sdk配置强制转换失败")
		return
	}

	appId := juDuConfig.GetAppId(devicePlatformType)
	gameId := juDuConfig.GetGameId(devicePlatformType)
	appKey := juDuConfig.GetAppKey(devicePlatformType)
	userName := judu.GetUserName()
	loginTime := judu.GetLoginTime()
	sign := judu.GetSign()
	getSign := GetJuDuSign(userName, appKey, loginTime)
	if getSign != sign {
		log.WithFields(
			log.Fields{
				"userName":  userName,
				"appId":     appId,
				"gameId":    gameId,
				"loginTime": loginTime,
				"sign":      sign,
				"getSign":   getSign,
			}).Warn("login:登陆享哥玩,签名错误")
		return
	}

	//登录认证
	flag, err = juduLogin(userName, appId, gameId, loginTime, sign)
	if err != nil {
		return
	}
	if !flag {
		return
	}
	flag = true
	returnPlatform = int32(platform)
	userId = userName
	log.WithFields(
		log.Fields{
			"userName":  userName,
			"appId":     appId,
			"gameId":    gameId,
			"loginTime": loginTime,
			"sign":      sign,
		}).Info("login:登陆巨都手游,登陆成功")
	return
}

const (
	juduPath = "http://sdk.juduwl.com/sdkapi/login/verification"
)

type juduLoginResult struct {
	Status int32  `json:"status"` //返回状态码
	Msg    string `json:"msg"`    //返回信息
}

func juduLogin(userName, appId, gameId, loginTime, sign string) (flag bool, err error) {

	params := make(map[string][]string)
	params["username"] = []string{userName}
	params["appid"] = []string{appId}
	params["gameid"] = []string{gameId}
	params["logintime"] = []string{loginTime}
	params["sign"] = []string{sign}
	data := url.Values(params)
	// TODO 请求超时
	resp, err := http.PostForm(juduPath, data)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆巨都手游,登录验证请求失败")
		return false, nil
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code %d", resp.StatusCode)
		return
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆巨都手游,回包数据读取错误")
		return false, nil
	}

	result := &juduLoginResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆巨都手游,回包数据解析错误")
		return false, nil
	}

	if result.Status != 1 {
		log.WithFields(
			log.Fields{
				"status": result.Status,
				"msg":    result.Msg,
			}).Warnf("login:登陆巨都手游,登录验证失败")
		return
	}

	flag = true
	return
}

func GetJuDuSign(userName, appKey, loginTime string) (sign string) {
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
