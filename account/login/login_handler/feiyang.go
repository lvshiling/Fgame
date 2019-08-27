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
	login.RegisterLogin(types.SDKTypeFeiYang, login.LoginHandlerFunc(handleFeiYang))
}

func handleFeiYang(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	feiYang := csAccountLogin.GetFeiYang()
	if feiYang == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆飞扬手游,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆飞扬手游,sdk配置为空")
		return
	}
	feiYangConfig, ok := sdkConfig.(*sdksdk.FeiYangConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆飞扬手游,sdk配置强制转换失败")
		return
	}

	appId := feiYangConfig.GetAppId(devicePlatformType)
	appKey := feiYangConfig.GetAppKey(devicePlatformType)
	userId = feiYang.GetMemId()
	userToken := feiYang.GetUserToken()
	if len(userToken) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆飞扬手游,userToken为空")
		return
	}

	//登录认证
	flag, err = feiYangLogin(appId, appKey, userId, userToken)
	if err != nil {
		return
	}
	if !flag {
		return
	}
	returnPlatform = int32(platform)
	log.WithFields(
		log.Fields{
			"appId":     appId,
			"appKey":    appKey,
			"userId":    userId,
			"userToken": userToken,
		}).Info("login:登陆飞扬手游,登陆成功")
	return
}

const (
	apiPath = "https://api.519wan.cn/api/cp/user/check"
)

type feiYangLoginCheckResult struct {
	Status string `json:"status"` //返回状态码
	Msg    string `json:"msg"`    //返回信息
}

func feiYangLogin(appId, appKey, userId, userToken string) (flag bool, err error) {
	sign := GetFeiYangSign(appId, appKey, userId, userToken)

	params := make(map[string][]string)
	params["app_id"] = []string{appId}
	params["mem_id"] = []string{userId}
	params["user_token"] = []string{userToken}
	params["sign"] = []string{sign}
	data := url.Values(params)
	// TODO 请求超时
	resp, err := http.PostForm(apiPath, data)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆飞扬手游,登录验证请求失败")
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
			}).Warnf("login:登陆飞扬手游,回包数据读取错误")
		return false, nil
	}

	result := &feiYangLoginCheckResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆飞扬手游,回包数据解析错误")
		return false, nil
	}

	if result.Status != "1" {
		log.WithFields(
			log.Fields{
				"status": result.Status,
				"msg":    result.Msg,
			}).Warnf("login:登陆飞扬手游,登录验证失败")
		return
	}

	flag = true
	return
}

func GetFeiYangSign(appId, appKey, userId, userToken string) (sign string) {
	signMap := make(map[string]string)
	signMap["app_id"] = appId
	signMap["mem_id"] = userId
	signMap["user_token"] = userToken
	signMap["app_key"] = appKey
	keyList := []string{"app_id", "mem_id", "user_token", "app_key"}
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
