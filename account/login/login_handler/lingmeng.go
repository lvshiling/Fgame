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
	login.RegisterLogin(types.SDKTypeLingMeng, login.LoginHandlerFunc(handleLingMeng))
}

func handleLingMeng(devicePlatformType types.DevicePlatformType, msg interface{}) (flag bool, returnPlatform int32, userId string, err error) {
	csAccountLogin := msg.(*uipb.CSAccountLogin)
	lingMeng := csAccountLogin.GetLingMeng()
	if lingMeng == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆灵梦仙界手游,数据是空的")
		return
	}
	platform := types.SDKType(csAccountLogin.GetPlatform())
	sdkConfig := sdk.GetSdkService().GetSdkConfig(platform)
	if sdkConfig == nil {
		log.WithFields(
			log.Fields{}).Warn("login:登陆灵梦仙界手游,sdk配置为空")
		return
	}
	lingMengConfig, ok := sdkConfig.(*sdksdk.LingMengConfig)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("login:登陆灵梦仙界手游,sdk配置强制转换失败")
		return
	}

	appId := lingMengConfig.GetAppId(devicePlatformType)
	appKey := lingMengConfig.GetAppKey(devicePlatformType)
	memId := lingMeng.GetMemId()
	userToken := lingMeng.GetUserToken()
	sign := GetLingMengSign(appId, appKey, memId, userToken)

	//登录认证
	flag, err = lingMengLogin(appId, memId, userToken, sign)
	if err != nil {
		return
	}
	if !flag {
		return
	}
	flag = true
	returnPlatform = int32(platform)
	userId = memId
	log.WithFields(
		log.Fields{
			"memId":     memId,
			"appId":     appId,
			"userToken": userToken,
			"sign":      sign,
		}).Info("login:登陆灵梦仙界手游,登陆成功")
	return
}

const (
	lingMengPath = "http://sdk.648game.cn/sdk/checkUsertoken.php"
)

type lingMengLoginResult struct {
	Status int32  `json:"status"` //返回状态码
	Msg    string `json:"msg"`    //返回信息
}

func lingMengLogin(appId, memId, userToken, sign string) (flag bool, err error) {

	params := make(map[string][]string)
	params["app_id"] = []string{appId}
	params["mem_id"] = []string{memId}
	params["user_token"] = []string{userToken}
	params["sign"] = []string{sign}
	data := url.Values(params)
	// TODO 请求超时
	resp, err := http.PostForm(lingMengPath, data)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆灵梦仙界手游,登录验证请求失败")
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
			}).Warnf("login:登陆灵梦仙界手游,回包数据读取错误")
		return false, nil
	}

	result := &lingMengLoginResult{}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warnf("login:登陆灵梦仙界手游,回包数据解析错误")
		return false, nil
	}

	if result.Status != 1 {
		log.WithFields(
			log.Fields{
				"status": result.Status,
				"msg":    result.Msg,
			}).Warnf("login:登陆灵梦仙界手游,登录验证失败")
		return
	}

	flag = true
	return
}

func GetLingMengSign(appId, appKey, memId, userToken string) (sign string) {
	signMap := make(map[string]string)
	signMap["app_id"] = appId
	signMap["app_key"] = appKey
	signMap["mem_id"] = memId
	signMap["user_token"] = userToken
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
