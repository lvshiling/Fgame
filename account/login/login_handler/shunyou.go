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
	login.RegisterLogin(types.SDKTypeShunYou, login.LoginHandlerFunc(handleShunYou))
}

func handleShunYou(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	shunYou := csAccountLogin.GetShunYou()
	if shunYou == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆顺游手游,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆顺游手游,sdk配置为空")
		return
	}
	shunYouConfig, ok := sdkConfig.(*sdksdk.ShunYouConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆顺游手游,sdk配置强制转换失败")
		return
	}

	appId := shunYouConfig.GetAppId(devicePlatformType)
	appKey := shunYouConfig.GetAppKey(devicePlatformType)
	userId = shunYou.GetUserId()
	token := shunYou.GetToken()
	if len(userId) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆顺游手游,userId为空")
		return
	}
	if len(token) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆顺游手游,token为空")
		return
	}

	//登录认证
	flag, err = shunYouLogin(appKey, userId, token)
	if err != nil {
		log.WithFields(
			log.Fields{
				"appId":  appId,
				"appKey": appKey,
				"userId": userId,

				"userToken": token,
				"err":       err,
			}).Error("login:登陆顺游手游,错误")
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"appId":  appId,
				"appKey": appKey,
				"userId": userId,

				"userToken": token,
			}).Warn("login:登陆顺游手游,失败")
		return
	}
	returnPlatform = int32(platform)
	flag = true
	log.WithFields(
		log.Fields{
			"appId":     appId,
			"appKey":    appKey,
			"userId":    userId,
			"userToken": token,
		}).Info("login:登陆顺游手游,登陆成功")
	return
}

const (
	shunYouCheckUrl = "https://issue.23youxi.com/user/verifyAccount"
)

type ShunYouCheckData struct {
	UserId     int64  `json:"userID`
	UserName   string `json:"username"`
	ChannnelId int64  `json:"channelID"`
	SdkUserId  string `json:"sdkUserID"`
}

type ShunYouLoginCheckResult struct {
	State int32             `json:"state"` //返回状态码
	Data  *ShunYouCheckData `json:"data"`  //返回信息
}

func shunYouLogin(appKey, userId, token string) (flag bool, err error) {
	sign := getShunYouSign(appKey, userId, token)

	params := make(map[string][]string)
	params["userID"] = []string{userId}
	params["token"] = []string{token}
	params["sign"] = []string{sign}
	data := url.Values(params)

	resp, err := http.PostForm(shunYouCheckUrl, data)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code %d", resp.StatusCode)
		return
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	result := &ShunYouLoginCheckResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return
	}

	statusInt := result.State
	if statusInt != 1 {
		log.WithFields(
			log.Fields{
				"appKey":    appKey,
				"userId":    userId,
				"userToken": token,
				"sign":      sign,
				"status":    result.State,
				"data":      result.Data,
			}).Warnf("login:登陆顺游手游,登录验证失败")
		return
	}
	log.WithFields(
		log.Fields{
			"appKey":    appKey,
			"userId":    userId,
			"userToken": token,
			"sign":      sign,
			"data":      result.Data,
		}).Info("login:登陆顺游手游,登陆验证成功")

	flag = true
	return
}

func getShunYouSign(appKey, userId, token string) (sign string) {
	signMap := make(map[string]string)
	signMap["userID"] = userId
	signMap["token"] = token
	keyList := []string{"userID", "token"}
	allStr := ""
	for _, key := range keyList {
		keyValue := fmt.Sprintf("%s=%s", key, signMap[key])
		allStr += keyValue
	}
	allStr += appKey
	hw := md5.Sum([]byte(allStr))
	return fmt.Sprintf("%x", hw)
}
