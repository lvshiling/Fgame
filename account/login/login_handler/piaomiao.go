package login_handler

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fgame/fgame/account/login/login"
	"fgame/fgame/account/login/types"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/sdk"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLogin(types.SDKTypePiaoMiao, login.LoginHandlerFunc(handlePiaoMiao))
}

func handlePiaoMiao(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	piaoMiao := csAccountLogin.GetPiaoMiao()
	if piaoMiao == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆缥缈手游,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆缥缈手游,sdk配置为空")
		return
	}

	userId = piaoMiao.GetUserId()
	userToken := piaoMiao.GetToken()
	if len(userToken) == 0 {
		log.WithFields(
			log.Fields{}).Warn("login:登陆缥缈手游,userToken为空")
		return
	}

	//登录认证
	flag, err = piaoMiaoLogin(userId, userToken)
	if err != nil {
		return
	}
	if !flag {
		return
	}
	returnPlatform = int32(platform)
	log.WithFields(
		log.Fields{
			"userId":    userId,
			"userToken": userToken,
		}).Info("login:登陆缥缈手游,登陆成功")
	return
}

const (
	piaoMiaoApiPath = "http://sy.41419.com/sdk.php/LoginNotify/login_verify"
)

type piaoMiaoLoginCheckResult struct {
	Status      int32  `json:"status"`       //返回状态码
	UserId      string `json:"user_id"`      //返回信息
	UserAccount string `json:"user_account"` //返回信息
}

func piaoMiaoLogin(userId, userToken string) (flag bool, err error) {

	params := make(map[string]string)
	params["user_id"] = userId
	params["token"] = userToken

	bodyBytes, err := json.Marshal(params)
	if err != nil {
		return
	}
	request, err := http.NewRequest("POST", piaoMiaoApiPath, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
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
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆缥缈手游,回包数据读取错误")
		return false, nil
	}

	result := &piaoMiaoLoginCheckResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆缥缈手游,回包数据解析错误")
		return false, nil
	}

	if result.Status != 200 {
		log.WithFields(
			log.Fields{
				"status":    result.Status,
				"userId":    userId,
				"userToken": userToken,
			}).Warnf("login:登陆缥缈手游,登录验证失败")
		return
	}

	flag = true
	return
}

func GetPiaoMiaoSign(appId, appKey, userId, userToken string) (sign string) {
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
